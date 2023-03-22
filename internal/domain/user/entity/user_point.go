package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserPoint struct {
	ID        uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.NullUUID
	Point     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserPoint(id, userID uuid.NullUUID, point int, createdAt, updatedAt time.Time) *UserPoint {
	return &UserPoint{
		ID:        id,
		UserID:    userID,
		Point:     point,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
