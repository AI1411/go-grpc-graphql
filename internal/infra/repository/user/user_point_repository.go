package user

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
)

type UserPointRepository interface {
	GetPoint(ctx context.Context, userID string) (int, error)
	UpdateUserPoint(ctx context.Context, point *entity.UserPoint) error
}

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
			return 0, nil
		}
		return 0, err
	}

	return point.Point, nil
}

func (u userPointRepository) UpdateUserPoint(ctx context.Context, point *entity.UserPoint) error {
	//TODO implement me
	panic("implement me")
}
