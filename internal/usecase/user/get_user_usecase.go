package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/user/converter"
)

type GetUserUsecase interface {
	Exec(ctx context.Context, userID string) (*grpc.GetUserResponse, error)
}

type getUserUsecaseImpl struct {
	userRepository userRepo.Repository
}

func NewGetUserUsecaseImpl(userRepository userRepo.Repository) GetUserUsecase {
	return &getUserUsecaseImpl{
		userRepository: userRepository,
	}
}

func (usecase *getUserUsecaseImpl) Exec(ctx context.Context, userID string) (*grpc.GetUserResponse, error) {
	user, err := usecase.userRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return converter.UserEntityToGRPC(user), nil
}
