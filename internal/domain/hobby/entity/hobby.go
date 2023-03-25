package entity

import "github.com/google/uuid"

type Hobby struct {
	ID         uuid.NullUUID
	Name       string
	CategoryID uuid.NullUUID
}

func NewHobby(id uuid.NullUUID, name string, categoryID uuid.NullUUID) *Hobby {
	return &Hobby{
		ID:         id,
		Name:       name,
		CategoryID: categoryID,
	}
}
