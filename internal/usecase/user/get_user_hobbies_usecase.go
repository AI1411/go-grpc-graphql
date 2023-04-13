package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetUserHobbiesUsecaseImpl interface {
	Exec(context.Context, *grpc.GetUserHobbiesRequest) (*grpc.GetUserHobbiesResponse, error)
}

type getUserHobbiesUsecaseImpl struct {
	userHobbyRepo user.HobbyRepository
}

func NewGetUserHobbiesUsecaseImpl(userHobbyRepo user.HobbyRepository) GetUserHobbiesUsecaseImpl {
	return &getUserHobbiesUsecaseImpl{
		userHobbyRepo: userHobbyRepo,
	}
}

func (u getUserHobbiesUsecaseImpl) Exec(ctx context.Context, in *grpc.GetUserHobbiesRequest) (*grpc.GetUserHobbiesResponse, error) {
	hobbies, err := u.userHobbyRepo.GetUserHobbies(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	res := make([]*grpc.Hobby, len(hobbies))
	for i, hobby := range hobbies {
		res[i] = &grpc.Hobby{
			Id:          util.NullUUIDToString(hobby.ID),
			Name:        hobby.Name,
			Description: hobby.Description,
			CategoryId:  util.NullUUIDToString(hobby.CategoryID),
		}
	}

	return &grpc.GetUserHobbiesResponse{
		Hobbies: res,
	}, nil
}
