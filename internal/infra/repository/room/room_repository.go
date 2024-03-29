package room

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/room/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type Repository interface {
	ListRoom(ctx context.Context, userID string) ([]*entity.Room, error)
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
	CreateRoom(ctx context.Context, Room *entity.Room) (string, error)
	DeleteRoom(ctx context.Context, id string) error
}

type roomRepository struct {
	dbClient db.Client
}

func NewRoomRepository(dbClient db.Client) Repository {
	return &roomRepository{
		dbClient: dbClient,
	}
}

func (c *roomRepository) ListRoom(ctx context.Context, userID string) ([]*entity.Room, error) {
	var rooms []*entity.Room
	if err := c.dbClient.Conn(ctx).
		Where("user_id", userID).
		Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (c *roomRepository) GetRoom(ctx context.Context, id string) (*entity.Room, error) {
	var room *entity.Room
	if err := c.dbClient.Conn(ctx).
		Where("id", id).
		Preload("Chats").
		First(&room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "room not found: %v", err)
		}
		return nil, err
	}

	return room, nil
}

func (c *roomRepository) CreateRoom(ctx context.Context, Room *entity.Room) (string, error) {
	if err := c.dbClient.Conn(ctx).Create(Room).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(Room.ID), nil
}

func (c *roomRepository) DeleteRoom(ctx context.Context, id string) error {
	var room *entity.Room
	if err := c.dbClient.Conn(ctx).Where("id", id).Delete(&room).Error; err != nil {
		return err
	}
	return nil
}
