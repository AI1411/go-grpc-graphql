package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-praphql/grpc"
)

type UserServer struct {
	grpc.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, in *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	return &grpc.GetUserResponse{}, nil
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
