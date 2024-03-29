package entity

import (
	"github.com/google/uuid"

	"github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
)

type Category struct {
	ID          uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string

	Hobbies []*entity.Hobby `gorm:"foreignKey:CategoryID;references:ID"`
}

type CategoryCondition struct {
	Name   string
	Order  string
	Limit  int32
	Offset int32
}

func NewCategory(name, description string) *Category {
	return &Category{
		Name:        name,
		Description: description,
	}
}

func NewCategoryCondition(name, order string, limit, offset int32) *CategoryCondition {
	return &CategoryCondition{
		Name:   name,
		Order:  order,
		Limit:  limit,
		Offset: offset,
	}
}
