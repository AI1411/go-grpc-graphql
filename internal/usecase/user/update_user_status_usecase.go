package user

import (
	"context"

	"github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type UpdateUserStatusUsecase interface {
	Exec(ctx context.Context, in *grpc.UpdateUserStatusRequest) error
}

type updateUserStatusUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUpdateUserStatusUsecaseImpl(userRepository repository.UserRepository) UpdateUserStatusUsecase {
	return &updateUserStatusUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *updateUserStatusUsecaseImpl) Exec(ctx context.Context, in *grpc.UpdateUserStatusRequest) error {
	return u.userRepository.UpdateUserStatus(ctx, &entity.User{
		ID:     util.StringToNullUUID(in.GetId()),
		Status: entity.UserStatusName[in.GetStatus().String()],
	})
}
