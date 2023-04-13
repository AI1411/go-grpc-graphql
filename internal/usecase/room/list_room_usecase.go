package room

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type ListRoomUsecaseImpl interface {
	Exec(context.Context, string) (*grpc.ListRoomResponse, error)
}

type listRoomUsecaseImpl struct {
	userRepository user.Repository
	roomRepository room.Repository
}

func NewListRoomUsecaseImpl(userRepository user.Repository, roomRepository room.Repository) ListRoomUsecaseImpl {
	return &listRoomUsecaseImpl{
		userRepository: userRepository,
		roomRepository: roomRepository,
	}
}

func (c listRoomUsecaseImpl) Exec(ctx context.Context, userID string) (*grpc.ListRoomResponse, error) {
	rooms, err := c.roomRepository.ListRoom(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*grpc.Room, len(rooms))
	for i, Room := range rooms {
		res[i] = &grpc.Room{
			Id:        util.NullUUIDToString(Room.ID),
			UserId:    util.NullUUIDToString(Room.UserID),
			CreatedAt: timestamppb.New(Room.CreatedAt),
			UpdatedAt: timestamppb.New(Room.UpdatedAt),
		}
	}
	return &grpc.ListRoomResponse{Rooms: res}, err
}
