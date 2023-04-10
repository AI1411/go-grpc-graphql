package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
)

func TestNewCategory(t *testing.T) {
	name := "Sports"
	description := "All about sports"

	category := entity.NewCategory(name, description)

	assert.Equal(t, name, category.Name)
	assert.Equal(t, description, category.Description)
}

func TestNewCategoryCondition(t *testing.T) {
	name := "Sports"
	order := "ASC"
	limit := int32(10)
	offset := int32(0)

	condition := entity.NewCategoryCondition(name, order, limit, offset)

	assert.Equal(t, name, condition.Name)
	assert.Equal(t, order, condition.Order)
	assert.Equal(t, limit, condition.Limit)
	assert.Equal(t, offset, condition.Offset)
}
