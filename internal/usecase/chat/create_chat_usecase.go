package chat

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type CreateChatUsecaseImpl interface {
	Exec(ctx context.Context, chat *grpc.CreateChatRequest) (*grpc.CreateChatResponse, error)
}

type createChatUsecaseImpl struct {
	userRepository userRepo.UserRepository
	chatRepo       chatRepo.ChatRepository
}

func NewCreateChatUsecaseImpl(userRepository userRepo.UserRepository, chatRepo chatRepo.ChatRepository) CreateChatUsecaseImpl {
	return &createChatUsecaseImpl{
		userRepository: userRepository,
		chatRepo:       chatRepo,
	}
}

func (u *createChatUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateChatRequest) (*grpc.CreateChatResponse, error) {
	chatID, err := u.chatRepo.CreateChat(ctx, &entity.Chat{
		RoomID:     util.StringToNullUUID(in.GetRoomId()),
		FromUserID: util.StringToNullUUID(in.GetFromUserId()),
		ToUserID:   util.StringToNullUUID(in.GetToUserId()),
		Body:       in.GetBody(),
	})

	if err != nil {
		return nil, err
	}

	return &grpc.CreateChatResponse{Id: chatID}, nil
}
