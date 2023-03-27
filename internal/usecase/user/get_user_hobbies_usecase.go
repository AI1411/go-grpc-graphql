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
	userHobbyRepo user.UserHobbyRepository
}

func NewGetUserHobbiesUsecaseImpl(userHobbyRepo user.UserHobbyRepository) GetUserHobbiesUsecaseImpl {
	return &getUserHobbiesUsecaseImpl{
		userHobbyRepo: userHobbyRepo,
	}
}

func (u getUserHobbiesUsecaseImpl) Exec(ctx context.Context, in *grpc.GetUserHobbiesRequest) (*grpc.GetUserHobbiesResponse, error) {
	userHobbies, err := u.userHobbyRepo.GetUserHobbies(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	uh := make([]*grpc.UserHobby, len(userHobbies))
	for i, hobby := range userHobbies {
		uh[i] = &grpc.UserHobby{
			Id:      util.NullUUIDToString(hobby.ID),
			UserId:  util.NullUUIDToString(hobby.UserID),
			HobbyId: util.NullUUIDToString(hobby.HobbyID),
		}
	}

	return &grpc.GetUserHobbiesResponse{
		UserHobbies: uh,
	}, nil
}
