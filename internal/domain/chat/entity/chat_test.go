package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewChat(t *testing.T) {
	chatID := uuid.New()
	roomID := uuid.New()
	fromUserID := uuid.New()
	toUserID := uuid.New()
	body := "Hello, how are you?"
	isRead := false
	createdAt := time.Now().UTC().Truncate(time.Millisecond)
	updatedAt := createdAt.Add(10 * time.Minute)

	chat := NewChat(chatID, roomID, fromUserID, toUserID, body, isRead, createdAt, updatedAt)

	assert.NotNil(t, chat)
	assert.Equal(t, chatID, chat.ID.UUID)
	assert.True(t, chat.ID.Valid)
	assert.Equal(t, roomID, chat.RoomID.UUID)
	assert.True(t, chat.RoomID.Valid)
	assert.Equal(t, fromUserID, chat.FromUserID.UUID)
	assert.True(t, chat.FromUserID.Valid)
	assert.Equal(t, toUserID, chat.ToUserID.UUID)
	assert.True(t, chat.ToUserID.Valid)
	assert.Equal(t, body, chat.Body)
	assert.Equal(t, isRead, chat.IsRead)
	assert.Equal(t, createdAt, chat.CreatedAt)
	assert.Equal(t, updatedAt, chat.UpdatedAt)
}
