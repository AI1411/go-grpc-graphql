package category

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
)

type DeleteCategoryUsecaseImpl interface {
	Exec(ctx context.Context, id string) error
}

type deleteCategoryUsecaseImpl struct {
	categoryRepo category.Repository
}

func NewDeleteCategoryUsecaseImpl(categoryRepo category.Repository) DeleteCategoryUsecaseImpl {
	return &deleteCategoryUsecaseImpl{
		categoryRepo: categoryRepo,
	}
}

func (u *deleteCategoryUsecaseImpl) Exec(ctx context.Context, id string) error {
	return u.categoryRepo.DeleteCategory(ctx, id)
}
