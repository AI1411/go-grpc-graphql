package entity

import "github.com/google/uuid"

type Hobby struct {
	ID          uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string
	CategoryID  uuid.NullUUID
}

func NewHobby(id uuid.NullUUID, name string, categoryID uuid.NullUUID) *Hobby {
	return &Hobby{
		ID:         id,
		Name:       name,
		CategoryID: categoryID,
	}
}
