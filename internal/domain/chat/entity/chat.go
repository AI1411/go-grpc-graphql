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
