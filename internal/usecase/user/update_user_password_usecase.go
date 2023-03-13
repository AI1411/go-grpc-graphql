package user

import (
	"context"

	"github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type UpdateUserPasswordUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.UpdateUserPasswordRequest) error
}

type updateUserPasswordUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUpdateUserPasswordUsecaseImpl(userRepository repository.UserRepository) UpdateUserPasswordUsecaseImpl {
	return &updateUserPasswordUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *updateUserPasswordUsecaseImpl) Exec(ctx context.Context, in *grpc.UpdateUserPasswordRequest) error {
	return u.userRepository.UpdateUserPassword(ctx, &entity.UserPassword{
		ID:                   util.StringToNullUUID(in.GetId()),
		Password:             in.GetPassword(),
		PasswordConfirmation: in.GetPasswordConfirmation(),
	})
}
