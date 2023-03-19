package server

import (
	"context"

	"go.uber.org/zap"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	chatForm "github.com/AI1411/go-grpc-graphql/internal/server/form/chat"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/chat"
)

type ChatServer struct {
	grpc.UnimplementedChatServiceServer
	dbClient  *db.Client
	zapLogger *zap.Logger
	userRepo  userRepo.UserRepository
	chatRepo  chatRepo.ChatRepository
}

func NewChatServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	userRepo userRepo.UserRepository,
	chatRepo chatRepo.ChatRepository,
) *ChatServer {
	return &ChatServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		userRepo:  userRepo,
		chatRepo:  chatRepo,
	}
}

func (s *ChatServer) ListChat(ctx context.Context, in *grpc.ListChatRequest) (*grpc.ListChatResponse, error) {
	validator := form.NewFormValidator(chatForm.NewListChatForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := chat.NewListChatUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ChatServer) CreateChat(ctx context.Context, in *grpc.CreateChatRequest) (*grpc.CreateChatResponse, error) {
	usecase := chat.NewCreateChatUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}
