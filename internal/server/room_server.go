package server

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	roomRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/room"
)

type RoomServer struct {
	dbClient  *db.Client
	zapLogger *zap.Logger
	userRepo  userRepo.UserRepository
	chatRepo  roomRepo.RoomRepository
}

func NewRoomServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	userRepo userRepo.UserRepository,
	chatRepo roomRepo.RoomRepository,
) *RoomServer {
	return &RoomServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		userRepo:  userRepo,
		chatRepo:  chatRepo,
	}
}

func (s *RoomServer) ListRoom(ctx context.Context, in *connect.Request[grpc.ListRoomRequest]) (*connect.Response[grpc.ListRoomResponse], error) {
	usecase := room.NewListRoomUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.Msg.GetUserId())
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.ListRoomResponse{
		Rooms: res.Rooms,
	})

	return resp, nil
}

func (s *RoomServer) GetRoom(ctx context.Context, in *connect.Request[grpc.GetRoomRequest]) (*connect.Response[grpc.GetRoomResponse], error) {
	usecase := room.NewGetRoomUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.Msg.GetId())
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.GetRoomResponse{
		Room: res.Room,
	})

	return resp, nil
}

func (s *RoomServer) CreateRoom(ctx context.Context, in *connect.Request[grpc.CreateRoomRequest]) (*connect.Response[grpc.CreateRoomResponse], error) {
	usecase := room.NewCreateRoomUsecaseImpl(s.userRepo, s.chatRepo)
	res, err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.CreateRoomResponse{
		Id: res.Id,
	})

	return resp, nil
}

func (s *RoomServer) DeleteRoom(ctx context.Context, in *connect.Request[grpc.DeleteRoomRequest]) (*connect.Response[emptypb.Empty], error) {
	usecase := room.NewDeleteRoomUsecaseImpl(s.userRepo, s.chatRepo)
	err := usecase.Exec(ctx, in.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
