package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID         uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoomID     uuid.NullUUID
	FromUserID uuid.NullUUID
	ToUserID   uuid.NullUUID
	Body       string
	IsRead     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewChat(id, roomID, fromUserID, toUserID uuid.UUID, body string, isRead bool, createdAt, updatedAt time.Time) *Chat {
	return &Chat{
		ID: uuid.NullUUID{
			UUID:  id,
			Valid: true,
		},
		RoomID: uuid.NullUUID{
			UUID:  roomID,
			Valid: true,
		},
		FromUserID: uuid.NullUUID{
			UUID:  fromUserID,
			Valid: true,
		},
		ToUserID: uuid.NullUUID{
			UUID:  toUserID,
			Valid: true,
		},
		Body:      body,
		IsRead:    isRead,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
