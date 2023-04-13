package category

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
)

type CreateCategoryUsecaseImpl interface {
	Exec(ctx context.Context, category *grpc.CreateCategoryRequest) (*grpc.CreateCategoryResponse, error)
}

type createCategoryUsecaseImpl struct {
	categoryRepo category.CategoryRepository
}

func NewCreateCategoryUsecaseImpl(categoryRepo category.CategoryRepository) CreateCategoryUsecaseImpl {
	return &createCategoryUsecaseImpl{
		categoryRepo: categoryRepo,
	}
}

func (u *createCategoryUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateCategoryRequest) (*grpc.CreateCategoryResponse, error) {
	categoryID, err := u.categoryRepo.CreateCategory(ctx, &entity.Category{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	return &grpc.CreateCategoryResponse{Id: categoryID}, nil
}
