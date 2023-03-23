package server

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/user"
)

type UserPointServer struct {
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

func (s *UserPointServer) GetUserPoint(ctx context.Context, in *connect.Request[grpc.GetUserPointRequest]) (*connect.Response[grpc.GetUserPointResponse], error) {
	usecase := user.NewGetUserPointUsecaseImpl(s.userRepo, s.userPointRepo)
	res, err := usecase.Exec(ctx, in.Msg.GetUserId())
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.GetUserPointResponse{
		Point: res.Point,
	})

	return resp, nil
}

func (s *UserPointServer) UpdateUserPoint(ctx context.Context, in *connect.Request[grpc.UpdateUserPointRequest]) (*connect.Response[emptypb.Empty], error) {
	usecase := user.NewUpdateUserPointUsecaseImpl(s.userRepo, s.userPointRepo)
	err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
