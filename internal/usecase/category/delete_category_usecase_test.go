package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockCategory "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/category"
)

func TestDeleteCategoryUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepo := mockCategory.NewMockRepository(ctrl)

	deleteCategoryUsecase := category.NewDeleteCategoryUsecaseImpl(mockCategoryRepo)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.EXPECT().DeleteCategory(ctx, categoryID).Return(nil)

		err := deleteCategoryUsecase.Exec(ctx, categoryID)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockCategoryRepo.EXPECT().DeleteCategory(ctx, categoryID).Return(errors.New("category not found"))

		err := deleteCategoryUsecase.Exec(ctx, categoryID)

		assert.Error(t, err)
	})
}
