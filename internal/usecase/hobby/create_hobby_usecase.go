package hobby

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type CreateHobbyUsecaseImpl interface {
	Exec(ctx context.Context, req *grpc.CreateHobbyRequest) (*grpc.CreateHobbyResponse, error)
}

type createHobbyUsecaseImpl struct {
	hobbyRepo hobby.HobbyRepository
}

func NewCreateHobbyUsecaseImpl(hobbyRepo hobby.HobbyRepository) CreateHobbyUsecaseImpl {
	return &createHobbyUsecaseImpl{
		hobbyRepo: hobbyRepo,
	}
}

func (u *createHobbyUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateHobbyRequest) (*grpc.CreateHobbyResponse, error) {
	res, err := u.hobbyRepo.CreateHobby(ctx, &entity.Hobby{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		CategoryID:  util.StringToNullUUID(in.GetCategoryId()),
	})
	if err != nil {
		return nil, err
	}

	return &grpc.CreateHobbyResponse{Id: res}, nil
}
