package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestNewUserHobby(t *testing.T) {
	id := uuid.New()
	userID := uuid.New()
	hobbyID := uuid.New()

	userPoint := entity.NewUserHobby(
		uuid.NullUUID{UUID: id, Valid: true},
		uuid.NullUUID{UUID: userID, Valid: true},
		uuid.NullUUID{UUID: hobbyID, Valid: true},
	)

	assert.Equal(t, id, userPoint.ID.UUID)
	assert.Equal(t, userID, userPoint.UserID.UUID)
	assert.Equal(t, hobbyID, userPoint.HobbyID.UUID)
}
