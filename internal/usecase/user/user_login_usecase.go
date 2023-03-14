package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type LoginUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.LoginRequest) (*grpc.LoginResponse, error)
}

type loginUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewLoginUsecaseImpl(userRepository repository.UserRepository) LoginUsecaseImpl {
	return &loginUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *loginUsecaseImpl) Exec(ctx context.Context, in *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	token, err := u.userRepository.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, err
	}

	return &grpc.LoginResponse{
		Token: token,
	}, nil
}
