package chat

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type MarkChatAsReadUsecaseImpl interface {
	Exec(ctx context.Context, chat *grpc.MarkChatAsReadRequest) error
}

type markChatAsReadUsecaseImpl struct {
	userRepository userRepo.Repository
	chatRepo       chatRepo.Repository
}

func NewMarkChatAsReadUsecaseImpl(userRepository userRepo.Repository, chatRepo chatRepo.Repository) MarkChatAsReadUsecaseImpl {
	return &markChatAsReadUsecaseImpl{
		userRepository: userRepository,
		chatRepo:       chatRepo,
	}
}

func (u *markChatAsReadUsecaseImpl) Exec(ctx context.Context, in *grpc.MarkChatAsReadRequest) error {
	err := u.chatRepo.MarkChatAsRead(ctx, in.GetId())
	if err != nil {
		return err
	}

	return nil
}
