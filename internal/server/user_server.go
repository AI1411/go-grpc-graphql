package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/usecase/user"
)

type UserServer struct {
	grpc.UnimplementedUserServiceServer
	dbClient *db.Client
}

func NewUserServer(dbClient *db.Client) *UserServer {
	return &UserServer{
		dbClient: dbClient,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	userRepo := repository.NewUserRepository(s.dbClient)
	usecase := user.NewGetUserUsecaseImpl(userRepo)
	return usecase.Exec(ctx, in.GetId())
}

func (s *UserServer) CreateUser(ctx context.Context, in *grpc.CreateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, in *grpc.UpdateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, in *grpc.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
