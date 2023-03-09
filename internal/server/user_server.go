package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/server/form"
	userForm "github.com/AI1411/go-grpc-praphql/internal/server/form/user"
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
	validator := form.NewFormValidator(userForm.NewGetUserForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(s.dbClient)
	usecase := user.NewGetUserUsecaseImpl(userRepo)
	return usecase.Exec(ctx, in.GetId())
}

func (s *UserServer) CreateUser(ctx context.Context, in *grpc.CreateUserRequest) (*emptypb.Empty, error) {
	userRepo := repository.NewUserRepository(s.dbClient)
	usecase := user.NewCreateUserUsecaseImpl(userRepo)
	log.Printf("CreateUser: %v", in)
	if err := usecase.Exec(ctx, in); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, in *grpc.UpdateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, in *grpc.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
