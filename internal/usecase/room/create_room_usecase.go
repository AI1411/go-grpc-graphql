package room

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/room/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type CreateRoomUsecaseImpl interface {
	Exec(context.Context, *grpc.CreateRoomRequest) (*grpc.CreateRoomResponse, error)
}

type createRoomUsecaseImpl struct {
	userRepository user.UserRepository
	roomRepository room.RoomRepository
}

func NewCreateRoomUsecaseImpl(userRepository user.UserRepository, roomRepository room.RoomRepository) CreateRoomUsecaseImpl {
	return &createRoomUsecaseImpl{
		userRepository: userRepository,
		roomRepository: roomRepository,
	}
}

func (c *createRoomUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateRoomRequest) (*grpc.CreateRoomResponse, error) {
	res, err := c.roomRepository.CreateRoom(ctx, &entity.Room{
		UserID: util.StringToNullUUID(in.UserId),
	})
	if err != nil {
		return nil, err
	}

	return &grpc.CreateRoomResponse{Id: res}, err
}
