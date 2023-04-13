package user

import (
	"context"

	"github.com/pkg/errors"

	"github.com/AI1411/go-grpc-graphql/grpc"
	redisRepository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/redis"
)

type LogoutUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.LogoutRequest) error
}

type logoutUsecaseImpl struct {
	redisRepository redisRepository.Repository
}

func NewLogoutUsecaseImpl(redisRepository redisRepository.Repository) LogoutUsecaseImpl {
	return &logoutUsecaseImpl{
		redisRepository: redisRepository,
	}
}

func (u *logoutUsecaseImpl) Exec(ctx context.Context, in *grpc.LogoutRequest) error {
	userID, err := u.redisRepository.Get(ctx, in.GetToken())
	if err != nil {
		return err
	}

	if userID != in.GetUserId() {
		return errors.New("invalid userID")
	}

	return u.redisRepository.Delete(ctx, in.GetToken())
}
