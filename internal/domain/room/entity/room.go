package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
)

type Room struct {
	ID        uuid.NullUUID  `gorm:"type:uuid;default;uuid_generate_v4();primaryKey"`
	UserID    uuid.NullUUID  `gorm:"type:uuid;not null"`
	Chats     []*entity.Chat `gorm:"foreignKey:RoomID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRoom(id, userID uuid.UUID, chats []*entity.Chat, createdAt, updatedAt time.Time) *Room {
	return &Room{
		ID: uuid.NullUUID{
			UUID:  id,
			Valid: true,
		},
		UserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
		Chats:     chats,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
