package entity

import "github.com/google/uuid"

type Category struct {
	ID          uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string
}

type CategoryCondition struct {
	Name   string
	Order  string
	Limit  int32
	Offset int32
}
