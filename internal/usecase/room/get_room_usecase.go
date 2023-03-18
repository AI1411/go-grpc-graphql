package room

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	roomRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetRoomUsecaseImpl interface {
	Exec(context.Context, string) (*grpc.GetRoomResponse, error)
}

type getRoomUsecaseImpl struct {
	userRepository user.UserRepository
	roomRepository roomRepo.RoomRepository
}

func NewGetRoomUsecaseImpl(userRepository user.UserRepository, roomRepository roomRepo.RoomRepository) GetRoomUsecaseImpl {
	return &getRoomUsecaseImpl{
		userRepository: userRepository,
		roomRepository: roomRepository,
	}
}

func (c *getRoomUsecaseImpl) Exec(ctx context.Context, userID string) (*grpc.GetRoomResponse, error) {
	room, err := c.roomRepository.GetRoom(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := &grpc.Room{
		Id:        util.NullUUIDToString(room.ID),
		UserId:    util.NullUUIDToString(room.UserID),
		CreatedAt: timestamppb.New(room.CreatedAt),
		UpdatedAt: timestamppb.New(room.UpdatedAt),
	}

	chats := make([]*grpc.Chat, len(room.Chats))
	for i, chat := range room.Chats {
		chats[i] = &grpc.Chat{
			Id:         util.NullUUIDToString(chat.ID),
			RoomId:     util.NullUUIDToString(chat.RoomID),
			FromUserId: util.NullUUIDToString(chat.FromUserID),
			ToUserId:   util.NullUUIDToString(chat.ToUserID),
			Body:       chat.Body,
			CreatedAt:  timestamppb.New(chat.CreatedAt),
			UpdatedAt:  timestamppb.New(chat.UpdatedAt),
		}
	}

	return &grpc.GetRoomResponse{Room: res, Chats: chats}, err
}
