package user

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
)

type UserPointRepository interface {
	GetPoint(ctx context.Context, userID string) (int, error)
	UpdateUserPoint(ctx context.Context, point *entity.UserPoint) error
	DistributePointAllUsers(ctx context.Context, reason string) error
}

const defaultDistributePoint = 50

type userPointRepository struct {
	dbClient *db.Client
}

func NewUserPointRepository(dbClient *db.Client) UserPointRepository {
	return &userPointRepository{
		dbClient: dbClient,
	}
}

func (u userPointRepository) GetPoint(ctx context.Context, userID string) (int, error) {
	var point entity.UserPoint
	if err := u.dbClient.Conn(ctx).Where("user_id", userID).First(&point).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, status.Errorf(codes.NotFound, "user point not found: %v", err)
		}
		return 0, err
	}

	return point.Point, nil
}

func (u userPointRepository) UpdateUserPoint(ctx context.Context, point *entity.UserPoint) error {
	var user entity.User
	if err := u.dbClient.Conn(ctx).Where("id", point.UserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "user not found: %v", err)
		}
		return err
	}

	var userPoint entity.UserPoint
	if err := u.dbClient.Conn(ctx).Where("user_id", point.UserID).First(&userPoint).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "user point not found: %v", err)
		}
		return err
	}

	if err := u.dbClient.Conn(ctx).Model(&userPoint).Select("UserID", "Point").Updates(point).Error; err != nil {
		return err
	}

	return nil
}

func (u userPointRepository) DistributePointAllUsers(ctx context.Context, reason string) error {
	var users []entity.User
	if err := u.dbClient.Conn(ctx).Where("status IN ?", entity.ActiveUser).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		var userPoint entity.UserPoint
		if err := u.dbClient.Conn(ctx).Where("user_id", user.ID).First(&userPoint).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "user point not found: %v", err)
			}
			return err
		}

		if err := u.dbClient.Conn(ctx).Model(&userPoint).Select("UserID", "Point").Updates(&entity.UserPoint{
			UserID: user.ID,
			Point:  userPoint.Point + defaultDistributePoint,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
