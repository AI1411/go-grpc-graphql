package category

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	categoryRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetCategoryUsecaseImpl interface {
	Exec(ctx context.Context, id string) (*grpc.GetCategoryResponse, error)
}

type getCategoryUsecaseImpl struct {
	categoryRepository categoryRepo.CategoryRepository
}

func NewGetCategoryUsecaseImpl(categoryRepository categoryRepo.CategoryRepository) GetCategoryUsecaseImpl {
	return &getCategoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (g getCategoryUsecaseImpl) Exec(ctx context.Context, id string) (*grpc.GetCategoryResponse, error) {
	res, err := g.categoryRepository.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	category := &grpc.Category{
		Id:          util.NullUUIDToString(res.ID),
		Name:        res.Name,
		Description: res.Description,
	}

	return &grpc.GetCategoryResponse{Category: category}, nil
}
