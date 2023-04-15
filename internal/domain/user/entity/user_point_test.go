package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestNewUserPoint(t *testing.T) {
	id := uuid.New()
	userID := uuid.New()
	point := 100
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()

	userPoint := entity.NewUserPoint(uuid.NullUUID{UUID: id, Valid: true}, uuid.NullUUID{UUID: userID, Valid: true}, point, createdAt, updatedAt)

	assert.Equal(t, id, userPoint.ID.UUID)
	assert.Equal(t, userID, userPoint.UserID.UUID)
	assert.Equal(t, point, userPoint.Point)
	assert.Equal(t, createdAt, userPoint.CreatedAt)
	assert.Equal(t, updatedAt, userPoint.UpdatedAt)
}
