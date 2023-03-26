package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	redisRepository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/redis"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type LoginUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.LoginRequest) (*grpc.LoginResponse, error)
}

type loginUsecaseImpl struct {
	userRepository  repository.UserRepository
	redisRepository redisRepository.RedisRepository
}

func NewLoginUsecaseImpl(userRepository repository.UserRepository, redisRepository redisRepository.RedisRepository) LoginUsecaseImpl {
	return &loginUsecaseImpl{
		userRepository:  userRepository,
		redisRepository: redisRepository,
	}
}

func (u *loginUsecaseImpl) Exec(ctx context.Context, in *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	token, err := u.userRepository.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, err
	}

	userID, err := util.GetUserIDFromJWT(token)
	if err != nil {
		return nil, err
	}

	if err := u.redisRepository.Set(ctx, token, userID); err != nil {
		return nil, err
	}

	return &grpc.LoginResponse{
		Token: token,
	}, nil
}
