package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type RegisterUserHobbyUsecaseImpl interface {
	Exec(ctx context.Context, in *grpc.RegisterUserHobbyRequest) (*grpc.RegisterUserHobbyResponse, error)
}

type registerUserHobbyUsecaseImpl struct {
	userHobbyRepo userRepo.UserHobbyRepository
}

func NewRegisterUserHobbyUsecaseImpl(userHobbyRepo userRepo.UserHobbyRepository) RegisterUserHobbyUsecaseImpl {
	return &registerUserHobbyUsecaseImpl{
		userHobbyRepo: userHobbyRepo,
	}
}

func (u registerUserHobbyUsecaseImpl) Exec(ctx context.Context, in *grpc.RegisterUserHobbyRequest) (*grpc.RegisterUserHobbyResponse, error) {
	id, err := u.userHobbyRepo.RegisterUserHobby(ctx, &entity.UserHobby{
		UserID:  util.StringToNullUUID(in.GetUserId()),
		HobbyID: util.StringToNullUUID(in.GetHobbyId()),
	})
	if err != nil {
		return nil, err
	}

	return &grpc.RegisterUserHobbyResponse{
		Id: id,
	}, nil
}
