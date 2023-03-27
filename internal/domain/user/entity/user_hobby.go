package entity

import "github.com/google/uuid"

type UserHobby struct {
	ID      uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID  uuid.NullUUID
	HobbyID uuid.NullUUID
}
