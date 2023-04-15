package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserPointHistory(t *testing.T) {
	id := uuid.New()
	userID := uuid.New()
	point := 100
	operationType := "test"
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()

	userPoint := NewUserPointHistory(
		uuid.NullUUID{UUID: id, Valid: true},
		uuid.NullUUID{UUID: userID, Valid: true},
		point,
		operationType,
		createdAt,
		updatedAt,
	)

	assert.Equal(t, id, userPoint.ID.UUID)
	assert.Equal(t, userID, userPoint.UserID.UUID)
	assert.Equal(t, point, userPoint.Point)
	assert.Equal(t, operationType, userPoint.OperationType)
	assert.Equal(t, createdAt, userPoint.CreatedAt)
	assert.Equal(t, updatedAt, userPoint.UpdatedAt)
}
