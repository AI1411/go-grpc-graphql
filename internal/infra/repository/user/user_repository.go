package user

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UserRepository interface {
	GetUser(context.Context, string) (*entity.User, error)
	CreateUser(context.Context, *entity.User) (string, error)
	UpdateUserProfile(context.Context, *entity.User) error
	UpdateUserStatus(context.Context, *entity.User) error
	UpdateUserPassword(context.Context, *entity.UserPassword) error
	Login(context.Context, string, string) (string, error)
	UploadUserImage(context.Context, *entity.User) error
}

var awardPoint = os.Getenv("STAR_AWARD_POINT")

type userRepository struct {
	dbClient   *db.Client
	awsSession *session.Session
}

func NewUserRepository(dbClient *db.Client, awsSession *session.Session) UserRepository {
	return &userRepository{
		dbClient:   dbClient,
		awsSession: awsSession,
	}
}

func (u *userRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	if err := u.dbClient.Conn(ctx).Where("id", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	return &user, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (string, error) {
	if err := u.dbClient.Conn(ctx).Create(&user).Error; err != nil {
		return "", status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return util.NullUUIDToString(user.ID), nil
}

func (u *userRepository) UpdateUserProfile(ctx context.Context, user *entity.User) error {
	var userEntity entity.User
	if err := u.dbClient.Conn(ctx).Where("id", util.NullUUIDToString(user.ID)).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if err := u.dbClient.Conn(ctx).
		Where("id", user.ID).
		Select("Username", "Prefecture", "Introduction", "BloodType").
		Updates(&user).Error; err != nil {
		return status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return nil
}

func (u *userRepository) UpdateUserStatus(ctx context.Context, user *entity.User) error {
	var userEntity entity.User
	if err := u.dbClient.Conn(ctx).Where("id", util.NullUUIDToString(user.ID)).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if err := u.dbClient.Conn(ctx).
		Model(&entity.User{}).
		Where("id", user.ID).
		Update("status", user.Status).Error; err != nil {
		return status.Errorf(codes.Internal, "failed to update user status: %v", err)
	}
	return nil
}

func (u *userRepository) UpdateUserPassword(ctx context.Context, user *entity.UserPassword) error {
	userEntity := &entity.User{}
	// get user
	if err := u.dbClient.Conn(ctx).Where("id", user.ID).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	// set new password
	userEntity.Password = user.Password
	if err := util.SetPassword(userEntity); err != nil {
		return status.Errorf(codes.Internal, "failed to set password: %v", err)
	}
	user.Password = userEntity.Password

	if err := u.dbClient.Conn(ctx).
		Model(&entity.User{}).
		Where("id", user.ID).
		Update("password", user.Password).Error; err != nil {
		return status.Errorf(codes.Internal, "failed to update password: %v", err)
	}
	return nil
}

func (u *userRepository) Login(ctx context.Context, email, password string) (string, error) {
	// TODO とりあえずTokenだけ返す。保存処理は後で実装
	var user *entity.User
	if err := u.dbClient.Conn(ctx).Where("email", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", status.Errorf(codes.InvalidArgument, "invalid password: %v", err)
	}

	token, err := util.GenerateToken(util.NullUUIDToString(user.ID))
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	// user_loginsに登録もしくは更新
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	var userLogin *entity.UserLogin
	err = u.dbClient.Conn(ctx).
		Where("login_date = ? AND user_id = ?", today, user.ID).First(&userLogin).Error
	if err != nil {
		// 今日ログインしていない場合
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := u.dbClient.Conn(ctx).
				Table("user_logins").
				Create(&entity.UserLogin{
					UserID:    user.ID,
					LoginDate: today,
				}).Error; err != nil {
				return "", status.Errorf(codes.Internal, "failed to create user_login: %v", err)
			}

			if err := u.awardLoginPoint(ctx, util.NullUUIDToString(user.ID)); err != nil {
				return "", status.Errorf(codes.Internal, "failed to award login point: %v", err)
			}
		} else {
			return "", status.Errorf(codes.Internal, "failed to get user_login: %v", err)
		}
	} else {
		// 今日ログインしている場合
		err = u.dbClient.Conn(ctx).
			Table("user_logins").
			Where("login_date", today).Update("updated_at", time.Now()).Error
		if err != nil {
			return "", status.Errorf(codes.Internal, "failed to update user_login: %v", err)
		}
	}

	return token, nil
}

// awardLoginPoint ログインボーナスを付与する
// ログイン時に50ポイント付与
// 今日ログインしていない場合はログインボーナスを付与する
// 今日ログインしている場合はログインボーナスを付与しない
// 今日ログインしている場合はログイン日時を更新する
func (u *userRepository) awardLoginPoint(ctx context.Context, userID string) error {
	if awardPoint == "" {
		awardPoint = "50"
	}

	// pointを数値に変換
	point, err := strconv.Atoi(awardPoint)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to convert string to int: %v", err)
	}

	if err = u.dbClient.Conn(ctx).Transaction(func(tx *gorm.DB) error {
		var user *entity.User
		if err = u.dbClient.Conn(ctx).
			Where("id = ?", userID).
			First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "user not found: %v", err)
			}
			return status.Errorf(codes.Internal, "failed to get user: %v", err)
		}

		var userPoint *entity.UserPoint
		err = u.dbClient.Conn(ctx).
			Where("user_id = ?", userID).First(&userPoint).Error

		if err != nil {
			// user_pointsに登録
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err = u.dbClient.Conn(ctx).
					Create(&entity.UserPoint{
						UserID: util.StringToNullUUID(userID),
						Point:  point,
					}).Error; err != nil {
					return status.Errorf(codes.Internal, "failed to create user_point: %v", err)
				}

				// user_point_historyに登録
				if err = u.dbClient.Conn(ctx).
					Create(&entity.UserPointHistory{
						UserID:        util.StringToNullUUID(userID),
						Point:         point,
						OperationType: "LOGIN",
					}).Error; err != nil {
					return status.Errorf(codes.Internal, "failed to create user_point: %v", err)
				}
			} else {
				return status.Errorf(codes.Internal, "failed to get user_point: %v", err)
			}
		} else {
			// user_pointsを更新
			err = u.dbClient.Conn(ctx).
				Table("user_points").
				Where("user_id", userID).Update("point", gorm.Expr("point + ?", point)).Error
			if err != nil {
				return status.Errorf(codes.Internal, "failed to update user_point: %v", err)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UploadUserImage(ctx context.Context, user *entity.User) error {
	var userEntity entity.User
	if err := u.dbClient.Conn(ctx).Where("id", util.NullUUIDToString(user.ID)).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	// base64データをデコード
	decodedImageBuffer, err := decodeBase64Image(user.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	uploadedImagePath, err := uploadToS3(u.awsSession, decodedImageBuffer, "star-user-image", util.NullUUIDToString(user.ID))

	if err := u.dbClient.Conn(ctx).
		Model(&entity.User{}).
		Where("id", user.ID).
		Update("ImagePath", uploadedImagePath).Error; err != nil {
		return status.Errorf(codes.Internal, "failed to update user image: %v", err)
	}
	return nil
}

func decodeBase64Image(image string) (*bytes.Buffer, error) {
	// base64データをデコード
	decodedImageData, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		log.Fatal(err)
	}

	// 画像データをバッファに格納
	decodedImageBuffer := bytes.NewBuffer(decodedImageData)

	return decodedImageBuffer, nil
}

// 画像データをS3にアップロードする関数
func uploadToS3(awsSession *session.Session, imageBuffer *bytes.Buffer, bucketName, imageKey string) (string, error) {
	uploader := s3manager.NewUploader(awsSession)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imageKey),
		Body:   imageBuffer,
	})
	if err != nil {
		return "", err

	}

	return result.Location, nil
}
