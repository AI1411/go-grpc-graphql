package category

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	categoryRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
	hobbyRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetCategoryUsecaseImpl interface {
	Exec(ctx context.Context, id string) (*grpc.GetCategoryResponse, error)
}

type getCategoryUsecaseImpl struct {
	categoryRepository categoryRepo.CategoryRepository
	hobbyRepo          hobbyRepo.HobbyRepository
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

	hobbies := make([]*grpc.Hobby, len(res.Hobbies))
	for i, hobby := range res.Hobbies {
		hobbies[i] = &grpc.Hobby{
			Id:          util.NullUUIDToString(hobby.ID),
			Name:        hobby.Name,
			Description: hobby.Description,
			CategoryId:  util.NullUUIDToString(hobby.CategoryID),
		}
	}

	return &grpc.GetCategoryResponse{Category: category, Hobbies: hobbies}, nil
}
