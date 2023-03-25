package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserLogin struct {
	ID        uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.NullUUID
	LoginDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserLogin(id, userID uuid.NullUUID, loginDate, createdAt, updatedAt time.Time) *UserLogin {
	return &UserLogin{
		ID:        id,
		UserID:    userID,
		LoginDate: loginDate,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
