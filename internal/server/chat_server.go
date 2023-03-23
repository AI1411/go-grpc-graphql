package server

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	chatForm "github.com/AI1411/go-grpc-graphql/internal/server/form/chat"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/chat"
)

type ChatServer struct {
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

func (s *ChatServer) ListChat(ctx context.Context, in *connect.Request[grpc.ListChatRequest]) (*connect.Response[grpc.ListChatResponse], error) {
	validator := form.NewFormValidator(chatForm.NewListChatForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := chat.NewListChatUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.ListChatResponse{
		Chats: res.Chats,
	})

	return resp, nil
}

func (s *ChatServer) CreateChat(ctx context.Context, in *connect.Request[grpc.CreateChatRequest]) (*connect.Response[grpc.CreateChatResponse], error) {
	validator := form.NewFormValidator(chatForm.NewCreateChatForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := chat.NewCreateChatUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.CreateChatResponse{
		Id: res.Id,
	})

	return resp, nil
}

func (s *ChatServer) MarkChatAsRead(ctx context.Context, in *connect.Request[grpc.MarkChatAsReadRequest]) (*connect.Response[emptypb.Empty], error) {
	validator := form.NewFormValidator(chatForm.NewMarkChatAsReadForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := chat.NewMarkChatAsReadUsecaseImpl(s.userRepo, s.chatRepo)
	err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&emptypb.Empty{})

	return resp, nil
}
