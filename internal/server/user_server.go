package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	redisRepository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/redis"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	userForm "github.com/AI1411/go-grpc-graphql/internal/server/form/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/user"
)

type UserServer struct {
	grpc.UnimplementedUserServiceServer
	dbClient  *db.Client
	zapLogger *zap.Logger
	userRepo  repository.UserRepository
	redisRepo redisRepository.RedisRepository
}

func NewUserServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	userRepo repository.UserRepository,
	redisRepo redisRepository.RedisRepository,
) *UserServer {
	return &UserServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		userRepo:  userRepo,
		redisRepo: redisRepo,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	validator := form.NewFormValidator(userForm.NewGetUserForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewGetUserUsecaseImpl(s.userRepo)
	return usecase.Exec(ctx, in.GetId())
}

func (s *UserServer) CreateUser(ctx context.Context, in *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {
	validator := form.NewFormValidator(userForm.NewCreateUserForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewCreateUserUsecaseImpl(s.userRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}
	return &grpc.CreateUserResponse{Id: res.Id}, nil
}

func (s *UserServer) UpdateUserProfile(ctx context.Context, in *grpc.UpdateUserProfileRequest) (*emptypb.Empty, error) {
	validator := form.NewFormValidator(userForm.NewUpdateUserProfileForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewUpdateUserProfileUsecaseImpl(s.userRepo)
	if err := usecase.Exec(ctx, in); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) UpdateUserStatus(ctx context.Context, in *grpc.UpdateUserStatusRequest) (*emptypb.Empty, error) {
	validator := form.NewFormValidator(userForm.NewUpdateUserStatusForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewUpdateUserStatusUsecaseImpl(s.userRepo)
	if err := usecase.Exec(ctx, in); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) UpdateUserPassword(ctx context.Context, in *grpc.UpdateUserPasswordRequest) (*emptypb.Empty, error) {
	validator := form.NewFormValidator(userForm.NewUpdateUserPasswordForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewUpdateUserPasswordUsecaseImpl(s.userRepo)
	if err := usecase.Exec(ctx, in); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) Login(ctx context.Context, in *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	validator := form.NewFormValidator(userForm.NewLoginForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := user.NewLoginUsecaseImpl(s.userRepo, s.redisRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
