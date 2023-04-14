package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	roomRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/room"
)

type RoomServer struct {
	grpc.UnimplementedRoomServiceServer
	dbClient  db.Client
	zapLogger *zap.Logger
	userRepo  userRepo.Repository
	chatRepo  roomRepo.Repository
}

func NewRoomServer(
	dbClient db.Client,
	zapLogger *zap.Logger,
	userRepo userRepo.Repository,
	chatRepo roomRepo.Repository,
) *RoomServer {
	return &RoomServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		userRepo:  userRepo,
		chatRepo:  chatRepo,
	}
}

func (s *RoomServer) ListRoom(ctx context.Context, in *grpc.ListRoomRequest) (*grpc.ListRoomResponse, error) {
	usecase := room.NewListRoomUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RoomServer) GetRoom(ctx context.Context, in *grpc.GetRoomRequest) (*grpc.GetRoomResponse, error) {
	usecase := room.NewGetRoomUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RoomServer) CreateRoom(ctx context.Context, in *grpc.CreateRoomRequest) (*grpc.CreateRoomResponse, error) {
	usecase := room.NewCreateRoomUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RoomServer) DeleteRoom(ctx context.Context, in *grpc.DeleteRoomRequest) (*emptypb.Empty, error) {
	usecase := room.NewDeleteRoomUsecaseImpl(s.userRepo, s.chatRepo)
	err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
