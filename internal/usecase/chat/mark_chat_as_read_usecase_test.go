package chat_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/grpc"
	chatEntity "github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	mockChat "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat/mock"
	mockUser "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/chat"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func Test_markChatAsReadUsecaseImpl_Exec(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mockChat.NewMockChatRepository(ctrl)
	mockUserRepo := mockUser.NewMockUserRepository(ctrl)

	// Test data
	roomID := uuid.New().String()
	userID := uuid.New().String()

	testChat := &chatEntity.Chat{
		ID:         util.StringToNullUUID(uuid.New().String()),
		RoomID:     util.StringToNullUUID(roomID),
		FromUserID: util.StringToNullUUID(userID),
		ToUserID:   util.StringToNullUUID(uuid.New().String()),
		Body:       "Hello!",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	t.Run("正常系/Test markChatAsReadUsecaseImpl_Exec", func(t *testing.T) {
		mockChatRepo.
			EXPECT().
			MarkChatAsRead(ctx, gomock.Any()).
			Return(nil).
			Times(1)

		usecase := chat.NewMarkChatAsReadUsecaseImpl(mockUserRepo, mockChatRepo)

		in := &grpc.MarkChatAsReadRequest{
			Id: util.NullUUIDToString(testChat.ID),
		}
		err := usecase.Exec(ctx, in)
		a.NoError(err)
	})

	t.Run("異常系/Test markChatAsReadUsecaseImpl_Exec", func(t *testing.T) {
		mockChatRepo.
			EXPECT().
			MarkChatAsRead(ctx, gomock.Any()).
			Return(status.Errorf(codes.NotFound, "chat not found")).
			Times(1)

		usecase := chat.NewMarkChatAsReadUsecaseImpl(mockUserRepo, mockChatRepo)

		in := &grpc.MarkChatAsReadRequest{
			Id: util.NullUUIDToString(testChat.ID),
		}
		err := usecase.Exec(ctx, in)
		a.Error(err)
	})
}
