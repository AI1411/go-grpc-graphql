package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserPointHistory struct {
	ID            uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        uuid.NullUUID
	Point         int
	OperationType string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUserPointHistory(id, userID uuid.NullUUID, point int, operationType string, createdAt, updatedAt time.Time) *UserPointHistory {
	return &UserPointHistory{
		ID:            id,
		UserID:        userID,
		Point:         point,
		OperationType: operationType,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
