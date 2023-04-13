package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	roomEntity "github.com/AI1411/go-grpc-graphql/internal/domain/room/entity"
)

func TestNewRoom(t *testing.T) {
	roomID := uuid.New()
	userID := uuid.New()
	createdAt := time.Now().UTC().Truncate(time.Millisecond)
	updatedAt := createdAt.Add(10 * time.Minute)

	chats := []*entity.Chat{
		{
			ID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			RoomID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			FromUserID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			ToUserID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			Body:      "Hello",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		{
			ID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			RoomID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			FromUserID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			ToUserID: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
			Body:      "Hello",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	room := roomEntity.NewRoom(roomID, userID, chats, createdAt, updatedAt)

	assert.NotNil(t, room)
	assert.Equal(t, roomID, room.ID.UUID)
	assert.True(t, room.ID.Valid)
	assert.Equal(t, userID, room.UserID.UUID)
	assert.True(t, room.UserID.Valid)
	assert.Equal(t, chats, room.Chats)
	assert.Equal(t, createdAt, room.CreatedAt)
	assert.Equal(t, updatedAt, room.UpdatedAt)
}
