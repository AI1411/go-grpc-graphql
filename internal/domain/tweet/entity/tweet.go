package entity

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	ID        uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTweet(id, userID uuid.NullUUID, body string, createdAt, updatedAt time.Time) *Tweet {
	return &Tweet{
		ID:        id,
		UserID:    userID,
		Body:      body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
