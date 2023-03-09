package user

import (
	"context"

	"github.com/AI1411/go-grpc-praphql/grpc"
	userRepo "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type GetUserUsecase interface {
	Exec(ctx context.Context, userId string) (*grpc.GetUserResponse, error)
}

type getUserUsecaseImpl struct {
	userRepository userRepo.UserRepository
}

func NewGetUserUsecaseImpl(userRepository userRepo.UserRepository) GetUserUsecase {
	return &getUserUsecaseImpl{
		userRepository: userRepository,
	}
}

func (usecase *getUserUsecaseImpl) Exec(ctx context.Context, userID string) (*grpc.GetUserResponse, error) {
	user, err := usecase.userRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &grpc.GetUserResponse{
		User: &grpc.User{
			Id:           util.NullUUIDToString(user.ID),
			Email:        user.Email,
			Username:     user.Username,
			Password:     string(user.Password),
			Status:       grpc.Status_ACTIVE,
			Introduction: user.Introduction,
		},
	}, nil
}
