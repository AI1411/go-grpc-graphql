package hobby

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetHobbyUsecaseImpl interface {
	Exec(ctx context.Context, id string) (*grpc.GetHobbyResponse, error)
}

type getHobbyUsecaseImpl struct {
	hobbyRepo hobby.HobbyRepository
}

func NewGetHobbyUsecaseImpl(hobbyRepo hobby.HobbyRepository) GetHobbyUsecaseImpl {
	return &getHobbyUsecaseImpl{
		hobbyRepo: hobbyRepo,
	}
}

func (u *getHobbyUsecaseImpl) Exec(ctx context.Context, id string) (*grpc.GetHobbyResponse, error) {
	res, err := u.hobbyRepo.GetHobby(ctx, id)
	if err != nil {
		return nil, err
	}

	hobby := &grpc.Hobby{
		Id:          util.NullUUIDToString(res.ID),
		Name:        res.Name,
		Description: res.Description,
		CategoryId:  util.NullUUIDToString(res.CategoryID),
	}

	return &grpc.GetHobbyResponse{Hobby: hobby}, nil
}
