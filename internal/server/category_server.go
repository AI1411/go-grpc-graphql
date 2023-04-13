package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	categoryRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
	"github.com/AI1411/go-grpc-graphql/internal/server/form"
	categoryForm "github.com/AI1411/go-grpc-graphql/internal/server/form/category"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/category"
)

type CategoryServer struct {
	grpc.UnimplementedCategoryServiceServer
	dbClient     *db.Client
	zapLogger    *zap.Logger
	categoryRepo categoryRepo.Repository
}

func NewCategoryServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	categoryRepo categoryRepo.Repository,
) *CategoryServer {
	return &CategoryServer{
		dbClient:     dbClient,
		zapLogger:    zapLogger,
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryServer) GetCategory(ctx context.Context, in *grpc.GetCategoryRequest) (*grpc.GetCategoryResponse, error) {
	validator := form.NewFormValidator(categoryForm.NewGetCategoryForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := category.NewGetCategoryUsecaseImpl(s.categoryRepo)
	res, err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryServer) ListCategory(ctx context.Context, in *grpc.ListCategoryRequest) (*grpc.ListCategoryResponse, error) {
	validator := form.NewFormValidator(categoryForm.NewListCategoryForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := category.NewListCategoryUsecaseImpl(s.categoryRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryServer) CreateCategory(ctx context.Context, in *grpc.CreateCategoryRequest) (*grpc.CreateCategoryResponse, error) {
	validator := form.NewFormValidator(categoryForm.NewCreateCategoryForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := category.NewCreateCategoryUsecaseImpl(s.categoryRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryServer) DeleteCategory(ctx context.Context, in *grpc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	validator := form.NewFormValidator(categoryForm.NewDeleteCategoryForm(in))
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	usecase := category.NewDeleteCategoryUsecaseImpl(s.categoryRepo)
	err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
