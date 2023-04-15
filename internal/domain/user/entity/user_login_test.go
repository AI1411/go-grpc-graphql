package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestNewUserLogin(t *testing.T) {
	id := uuid.New()
	userID := uuid.New()
	loginDate := time.Now().UTC()
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()

	userLogin := entity.NewUserLogin(uuid.NullUUID{UUID: id, Valid: true}, uuid.NullUUID{UUID: userID, Valid: true}, loginDate, createdAt, updatedAt)

	assert.Equal(t, id, userLogin.ID.UUID)
	assert.Equal(t, userID, userLogin.UserID.UUID)
	assert.Equal(t, loginDate, userLogin.LoginDate)
	assert.Equal(t, createdAt, userLogin.CreatedAt)
	assert.Equal(t, updatedAt, userLogin.UpdatedAt)
}
