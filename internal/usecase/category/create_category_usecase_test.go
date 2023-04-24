package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	mockCategory "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/category"
)

const categoryID = "123e4567-e89b-12d3-a456-426614174000"

func TestCreateCategoryUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepo := mockCategory.NewMockRepository(ctrl)

	createCategoryUsecase := category.NewCreateCategoryUsecaseImpl(mockCategoryRepo)

	ctx := context.Background()
	newCategory := &grpc.CreateCategoryRequest{
		Name:        "Sports",
		Description: "All about sports",
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.EXPECT().CreateCategory(ctx, &entity.Category{
			Name:        newCategory.GetName(),
			Description: newCategory.GetDescription(),
		}).Return(categoryID, nil)

		res, err := createCategoryUsecase.Exec(ctx, newCategory)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, categoryID, res.Id)
	})

	t.Run("error", func(t *testing.T) {
		mockCategoryRepo.EXPECT().CreateCategory(ctx, &entity.Category{
			Name:        newCategory.GetName(),
			Description: newCategory.GetDescription(),
		}).Return("", errors.New("error creating category"))

		res, err := createCategoryUsecase.Exec(ctx, newCategory)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
