package repository

import (
	"context"
	"errors"

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

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteUser(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}
