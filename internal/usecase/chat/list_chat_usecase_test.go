package chat

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/grpc"
	chatEntity "github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	mockChat "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat/mock"
	mockUser "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user/mock"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func Test_listChatUsecaseImpl_Exec(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mockChat.NewMockChatRepository(ctrl)
	mockUserRepo := mockUser.NewMockUserRepository(ctrl)

	// Test data
	roomID := uuid.New().String()
	roomID2 := uuid.New().String()
	userID := uuid.New().String()
	userID2 := uuid.New().String()

	testChats := []*chatEntity.Chat{
		{
			ID:         util.StringToNullUUID(uuid.New().String()),
			RoomID:     util.StringToNullUUID(roomID),
			FromUserID: util.StringToNullUUID(userID),
			ToUserID:   util.StringToNullUUID(uuid.New().String()),
			Body:       "Hello!",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         util.StringToNullUUID(uuid.New().String()),
			RoomID:     util.StringToNullUUID(roomID2),
			FromUserID: util.StringToNullUUID(userID2),
			ToUserID:   util.StringToNullUUID(uuid.New().String()),
			Body:       "Hello!",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	mockChatRepo.
		EXPECT().
		ListChat(ctx, gomock.Any()).
		Return(testChats, nil).
		Times(1)

	// Create listChatUsecaseImpl
	usecase := NewListChatUsecaseImpl(mockUserRepo, mockChatRepo)

	// Call Exec() method
	in := &grpc.ListChatRequest{
		RoomId: roomID,
		UserId: userID,
	}
	resp, err := usecase.Exec(context.Background(), in)

	t.Run("正常系/Test listChatUsecaseImpl_Exec", func(t *testing.T) {
		// Assertions
		a.NoError(err, "Exec should not return an error")
		a.NotNil(resp, "Exec should return a non-nil response")
		a.Equal(len(testChats), len(resp.Chats), "The number of chats in the response should match the test data")

		for i, chat := range testChats {
			assert.Equal(t, chat.ID.UUID.String(), resp.Chats[i].GetId())
			assert.Equal(t, chat.RoomID.UUID.String(), resp.Chats[i].GetRoomId())
			assert.Equal(t, chat.FromUserID.UUID.String(), resp.Chats[i].GetFromUserId())
			assert.Equal(t, chat.ToUserID.UUID.String(), resp.Chats[i].GetToUserId())
			assert.Equal(t, chat.Body, resp.Chats[i].GetBody())
			assert.Equal(t, chat.CreatedAt.Unix(), resp.Chats[i].GetCreatedAt().AsTime().Unix())
			assert.Equal(t, chat.UpdatedAt.Unix(), resp.Chats[i].GetUpdatedAt().AsTime().Unix())
		}
	})
}
