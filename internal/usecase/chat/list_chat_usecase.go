package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type ListChatUsecaseImpl interface {
	Exec(ctx context.Context, userID string) (*grpc.ListChatResponse, error)
}

type listChatUsecaseImpl struct {
	userRepository userRepo.UserRepository
	chatRepo       chatRepo.ChatRepository
}

func NewListChatUsecaseImpl(userRepository userRepo.UserRepository, chatRepo chatRepo.ChatRepository) ListChatUsecaseImpl {
	return &listChatUsecaseImpl{
		userRepository: userRepository,
		chatRepo:       chatRepo,
	}
}

func (u *listChatUsecaseImpl) Exec(ctx context.Context, userID string) (*grpc.ListChatResponse, error) {
	chats, err := u.chatRepo.ListChat(ctx, userID)
	if err != nil {
		return nil, err
	}

	chatResponses := make([]*grpc.Chat, len(chats))
	for i, chat := range chats {
		chatResponses[i] = &grpc.Chat{
			Id:         util.NullUUIDToString(chat.ID),
			FromUserId: util.NullUUIDToString(chat.FromUserID),
			ToUserId:   util.NullUUIDToString(chat.ToUserID),
			Body:       chat.Body,
			CreatedAt:  timestamppb.New(chat.CreatedAt),
			UpdatedAt:  timestamppb.New(chat.UpdatedAt),
		}
	}

	return &grpc.ListChatResponse{Chats: chatResponses}, nil
}
