package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UpdateUserPointUsecase interface {
	Exec(ctx context.Context, request *grpc.UpdateUserPointRequest) error
}

type updateUserPointUsecaseImpl struct {
	userRepository      repository.UserRepository
	userPointRepository repository.UserPointRepository
}

func NewUpdateUserPointUsecaseImpl(userRepository repository.UserRepository, userPointRepository repository.UserPointRepository) UpdateUserPointUsecase {
	return &updateUserPointUsecaseImpl{
		userRepository:      userRepository,
		userPointRepository: userPointRepository,
	}
}

func (usecase *updateUserPointUsecaseImpl) Exec(ctx context.Context, in *grpc.UpdateUserPointRequest) error {
	err := usecase.userPointRepository.UpdateUserPoint(ctx, &entity.UserPoint{
		UserID: util.StringToNullUUID(in.GetUserId()),
		Point:  int(in.GetPoint()),
	})
	if err != nil {
		return err
	}

	return nil
}
