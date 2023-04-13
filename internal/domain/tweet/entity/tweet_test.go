package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/tweet/entity"
)

func TestNewTweet(t *testing.T) {
	id := uuid.NullUUID{UUID: uuid.New(), Valid: true}
	userID := uuid.NullUUID{UUID: uuid.New(), Valid: true}
	body := "This is a test tweet"
	createdAt := time.Now()
	updatedAt := time.Now()

	tweet := entity.NewTweet(id, userID, body, createdAt, updatedAt)

	assert.Equal(t, id, tweet.ID)
	assert.Equal(t, userID, tweet.UserID)
	assert.Equal(t, body, tweet.Body)
	assert.Equal(t, createdAt, tweet.CreatedAt)
	assert.Equal(t, updatedAt, tweet.UpdatedAt)
}
