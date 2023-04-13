package hobby

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type Repository interface {
	GetHobby(ctx context.Context, id string) (*entity.Hobby, error)
	CreateHobby(ctx context.Context, Hobby *entity.Hobby) (string, error)
	DeleteHobby(ctx context.Context, id string) error
}

type hobbyRepository struct {
	dbClient *db.Client
}

func NewHobbyRepository(dbClient *db.Client) Repository {
	return &hobbyRepository{
		dbClient: dbClient,
	}
}

func (h hobbyRepository) GetHobby(ctx context.Context, id string) (*entity.Hobby, error) {
	var hobby *entity.Hobby
	if err := h.dbClient.Conn(ctx).Where("id", id).First(&hobby).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "hobby not found: %v", err)
		}
		return nil, err
	}

	return hobby, nil
}

func (h hobbyRepository) CreateHobby(ctx context.Context, hobby *entity.Hobby) (string, error) {
	if err := h.dbClient.Conn(ctx).Create(hobby).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(hobby.ID), nil
}

func (h hobbyRepository) DeleteHobby(ctx context.Context, id string) error {
	if err := h.dbClient.Conn(ctx).Where("id", id).Delete(&entity.Hobby{}).Error; err != nil {
		return err
	}
	return nil
}
