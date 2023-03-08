package repository

import (
	"context"
	"time"

	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
)

type UserRepository interface {
	GetUser(context.Context, string) (*entity.User, error)
	CreateUser(context.Context, *entity.User) error
	UpdateUser(context.Context, *entity.User) error
	DeleteUser(context.Context, string) error
}

type userRepository struct {
	dbClient *db.Client
}

func NewUserRepository(dbClient *db.Client) UserRepository {
	return &userRepository{dbClient: dbClient}
}

func (u *userRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	return &entity.User{
		ID:           "tes",
		Username:     "tes",
		Email:        "e",
		Password:     "e",
		Status:       "sss",
		Prefecture:   "sss",
		Introduction: "sss",
		BloodType:    "sss",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteUser(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}
