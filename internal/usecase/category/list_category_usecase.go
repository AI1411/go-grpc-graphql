package category

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type ListCategoryUsecaseImpl interface {
	Exec(ctx context.Context, request *grpc.ListCategoryRequest) (*grpc.ListCategoryResponse, error)
}

type listCategoryUsecaseImpl struct {
	categoryRepository category.CategoryRepository
}

func NewListCategoryUsecaseImpl(categoryRepository category.CategoryRepository) ListCategoryUsecaseImpl {
	return &listCategoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (l listCategoryUsecaseImpl) Exec(ctx context.Context, in *grpc.ListCategoryRequest) (*grpc.ListCategoryResponse, error) {
	res, err := l.categoryRepository.ListCategory(ctx, &entity.CategoryCondition{
		Name:   in.GetName(),
		Order:  in.GetOrder(),
		Offset: in.GetOffset(),
		Limit:  in.GetLimit(),
	})
	if err != nil {
		return nil, err
	}

	categories := make([]*grpc.Category, len(res))
	for i, v := range res {
		categories[i] = &grpc.Category{
			Id:          util.NullUUIDToString(v.ID),
			Name:        v.Name,
			Description: v.Description,
		}
	}

	return &grpc.ListCategoryResponse{Categories: categories}, nil
}
