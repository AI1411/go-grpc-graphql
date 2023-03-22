package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/user"
)

type UserPointServer struct {
	grpc.UnimplementedUserPointServiceServer
	dbClient      *db.Client
	zapLogger     *zap.Logger
	userRepo      repository.UserRepository
	userPointRepo repository.UserPointRepository
}

func NewUserPointServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	userRepo repository.UserRepository,
	userPointRepo repository.UserPointRepository,
) *UserPointServer {
	return &UserPointServer{
		dbClient:      dbClient,
		zapLogger:     zapLogger,
		userRepo:      userRepo,
		userPointRepo: userPointRepo,
	}
}

func (s *UserPointServer) GetUserPoint(ctx context.Context, in *grpc.GetUserPointRequest) (*grpc.GetUserPointResponse, error) {
	usecase := user.NewGetUserPointUsecaseImpl(s.userRepo, s.userPointRepo)
	return usecase.Exec(ctx, in.GetUserId())
}

func (s *UserPointServer) UpdateUserPoint(ctx context.Context, in *grpc.UpdateUserPointRequest) (*emptypb.Empty, error) {
	usecase := user.NewUpdateUserPointUsecaseImpl(s.userRepo, s.userPointRepo)
	err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
