package server

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	userForm "github.com/AI1411/go-grpc-graphql/internal/server/form/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/user"
)

type UserServer struct {
	dbClient  *db.Client
	zapLogger *zap.Logger
	userRepo  repository.UserRepository
}

func NewUserServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	userRepo repository.UserRepository,
) *UserServer {
	return &UserServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		userRepo:  userRepo,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *connect.Request[grpc.GetUserRequest]) (*connect.Response[grpc.GetUserResponse], error) {
	validator := form.NewFormValidator(userForm.NewGetUserForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewGetUserUsecaseImpl(s.userRepo)
	res, err := usecase.Exec(ctx, in.Msg.GetId())
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.GetUserResponse{
		User: res.User,
	})

	return resp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, in *connect.Request[grpc.CreateUserRequest]) (*connect.Response[grpc.CreateUserResponse], error) {
	validator := form.NewFormValidator(userForm.NewCreateUserForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewCreateUserUsecaseImpl(s.userRepo)
	res, err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&grpc.CreateUserResponse{
		Id: res.Id,
	})
	return resp, nil
}

func (s *UserServer) UpdateUserProfile(ctx context.Context, in *connect.Request[grpc.UpdateUserProfileRequest]) (*connect.Response[emptypb.Empty], error) {
	validator := form.NewFormValidator(userForm.NewUpdateUserProfileForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewUpdateUserProfileUsecaseImpl(s.userRepo)
	if err := usecase.Exec(ctx, in.Msg); err != nil {
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *UserServer) UpdateUserStatus(ctx context.Context, in *connect.Request[grpc.UpdateUserStatusRequest]) (*connect.Response[emptypb.Empty], error) {
	validator := form.NewFormValidator(userForm.NewUpdateUserStatusForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewUpdateUserStatusUsecaseImpl(s.userRepo)
	if err := usecase.Exec(ctx, in.Msg); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *UserServer) UpdateUserPassword(ctx context.Context, in *connect.Request[grpc.UpdateUserPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	validator := form.NewFormValidator(userForm.NewUpdateUserPasswordForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewUpdateUserPasswordUsecaseImpl(s.userRepo)
	if err := usecase.Exec(ctx, in.Msg); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *UserServer) Login(ctx context.Context, in *connect.Request[grpc.LoginRequest]) (*connect.Response[grpc.LoginResponse], error) {
	validator := form.NewFormValidator(userForm.NewLoginForm(in.Msg))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewLoginUsecaseImpl(s.userRepo)
	res, err := usecase.Exec(ctx, in.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&grpc.LoginResponse{Token: res.Token}), nil
}
