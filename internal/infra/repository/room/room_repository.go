package room

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/internal/domain/room/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type RoomRepository interface {
	ListRoom(ctx context.Context, userID string) ([]*entity.Room, error)
	CreateRoom(ctx context.Context, Room *entity.Room) (string, error)
}

type roomRepository struct {
	dbClient *db.Client
}

func NewRoomRepository(dbClient *db.Client) RoomRepository {
	return &roomRepository{
		dbClient: dbClient,
	}
}

func (c *roomRepository) ListRoom(ctx context.Context, userID string) ([]*entity.Room, error) {
	var Rooms []*entity.Room
	if err := c.dbClient.Conn(ctx).
		Where("user_id", userID).
		Find(&Rooms).Error; err != nil {
		return nil, err
	}
	return Rooms, nil
}

func (c *roomRepository) CreateRoom(ctx context.Context, Room *entity.Room) (string, error) {
	if err := c.dbClient.Conn(ctx).Create(Room).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(Room.ID), nil
}
