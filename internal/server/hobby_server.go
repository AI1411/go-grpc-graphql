package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	hobbyRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	hobbyForm "github.com/AI1411/go-grpc-graphql/internal/server/form/hobby"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/hobby"
)

type HobbyServer struct {
	grpc.UnimplementedHobbyServiceServer
	dbClient  *db.Client
	zapLogger *zap.Logger
	hobbyRepo hobbyRepo.Repository
}

func NewHobbyServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	hobbyRepo hobbyRepo.Repository,
) *HobbyServer {
	return &HobbyServer{
		dbClient:  dbClient,
		zapLogger: zapLogger,
		hobbyRepo: hobbyRepo,
	}
}

func (s *HobbyServer) GetHobby(ctx context.Context, in *grpc.GetHobbyRequest) (*grpc.GetHobbyResponse, error) {
	usecase := hobby.NewGetHobbyUsecaseImpl(s.hobbyRepo)
	res, err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *HobbyServer) CreateHobby(ctx context.Context, in *grpc.CreateHobbyRequest) (*grpc.CreateHobbyResponse, error) {
	validator := form.NewFormValidator(hobbyForm.NewCreateHobbyForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := hobby.NewCreateHobbyUsecaseImpl(s.hobbyRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *HobbyServer) DeleteHobby(ctx context.Context, in *grpc.DeleteHobbyRequest) (*emptypb.Empty, error) {
	validator := form.NewFormValidator(hobbyForm.NewDeleteHobbyForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := hobby.NewDeleteHobbyUsecaseImpl(s.hobbyRepo)
	err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
