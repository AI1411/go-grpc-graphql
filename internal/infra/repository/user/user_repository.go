package repository

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type UserRepository interface {
	GetUser(context.Context, string) (*entity.User, error)
	CreateUser(context.Context, *entity.User) error
	UpdateUserProfile(context.Context, *entity.User) error
	UpdateUserStatus(context.Context, *entity.User) error
	UpdateUserPassword(context.Context, *entity.UserPassword) error
	Login(context.Context, string, string) (string, error)
}

type userRepository struct {
	dbClient *db.Client
}

func NewUserRepository(dbClient *db.Client) UserRepository {
	return &userRepository{
		dbClient: dbClient,
	}
}

func (u *userRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	if err := u.dbClient.Conn(ctx).Where("id", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := util.SetPassword(user); err != nil {
		return err
	}
	if err := u.dbClient.Conn(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUserProfile(ctx context.Context, user *entity.User) error {
	if err := u.dbClient.Conn(ctx).
		Where("id", user.ID).
		Select("Username", "Prefecture", "Introduction", "BloodType").
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUserStatus(ctx context.Context, user *entity.User) error {
	if err := u.dbClient.Conn(ctx).
		Model(&entity.User{}).
		Where("id", user.ID).
		Update("status", user.Status).Error; err != nil {
		return err
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

	return token, nil
}
