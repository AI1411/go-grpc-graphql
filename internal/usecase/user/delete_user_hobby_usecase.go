package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type DeleteUserHobbyUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.DeleteUserHobbyRequest) (*emptypb.Empty, error)
}

type deleteUserHobbyUsecaseImpl struct {
	userHobbyRepo userRepo.UserHobbyRepository
}

func NewDeleteUserHobbyUsecaseImpl(userHobbyRepo userRepo.UserHobbyRepository) DeleteUserHobbyUsecaseImpl {
	return &deleteUserHobbyUsecaseImpl{
		userHobbyRepo: userHobbyRepo,
	}
}

func (u deleteUserHobbyUsecaseImpl) Exec(ctx context.Context, in *grpc.DeleteUserHobbyRequest) (*emptypb.Empty, error) {
	err := u.userHobbyRepo.DeleteUserHobby(ctx, &entity.UserHobby{
		UserID:  util.StringToNullUUID(in.GetUserId()),
		HobbyID: util.StringToNullUUID(in.GetHobbyId()),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
