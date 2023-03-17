package room

import (
	"context"

	roomRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type DeleteRoomUsecaseImpl interface {
	Exec(context.Context, string) error
}

type deleteRoomUsecaseImpl struct {
	userRepository user.UserRepository
	roomRepository roomRepo.RoomRepository
}

func NewDeleteRoomUsecaseImpl(userRepository user.UserRepository, roomRepository roomRepo.RoomRepository) DeleteRoomUsecaseImpl {
	return &deleteRoomUsecaseImpl{
		userRepository: userRepository,
		roomRepository: roomRepository,
	}
}

func (c *deleteRoomUsecaseImpl) Exec(ctx context.Context, userID string) error {
	err := c.roomRepository.DeleteRoom(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
