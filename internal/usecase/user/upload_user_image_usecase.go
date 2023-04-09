package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UploadUserImageUsecaseImpl interface {
	Exec(ctx context.Context, request *grpc.UploadUserImageRequest) error
}

type uploadUserImageUsecaseImpl struct {
	userRepository userRepo.UserRepository
}

func NewUploadUserImageUsecaseImpl(userRepository userRepo.UserRepository) UploadUserImageUsecaseImpl {
	return &uploadUserImageUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *uploadUserImageUsecaseImpl) Exec(ctx context.Context, in *grpc.UploadUserImageRequest) error {
	err := u.userRepository.UploadUserImage(ctx, &entity.User{
		ID:        util.StringToNullUUID(in.UserId),
		ImagePath: in.Image,
	})
	if err != nil {
		return err
	}

	return nil
}
