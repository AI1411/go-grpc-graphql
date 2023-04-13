package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UpdateUserPasswordUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.UpdateUserPasswordRequest) error
}

type updateUserPasswordUsecaseImpl struct {
	userRepository repository.Repository
}

func NewUpdateUserPasswordUsecaseImpl(userRepository repository.Repository) UpdateUserPasswordUsecaseImpl {
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
