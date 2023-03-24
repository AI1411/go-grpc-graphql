package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTweet(t *testing.T) {
	id := uuid.NullUUID{UUID: uuid.New(), Valid: true}
	userID := uuid.NullUUID{UUID: uuid.New(), Valid: true}
	body := "This is a test tweet"
	createdAt := time.Now()
	updatedAt := time.Now()

	tweet := NewTweet(id, userID, body, createdAt, updatedAt)

	assert.Equal(t, id, tweet.ID)
	assert.Equal(t, userID, tweet.UserID)
	assert.Equal(t, body, tweet.Body)
	assert.Equal(t, createdAt, tweet.CreatedAt)
	assert.Equal(t, updatedAt, tweet.UpdatedAt)
}
