package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	hobbyEntity "github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
	mockCategory "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/category"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func TestGetCategoryUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepo := mockCategory.NewMockRepository(ctrl)

	getCategoryUsecase := category.NewGetCategoryUsecaseImpl(mockCategoryRepo)

	ctx := context.Background()
	categoryID := "123e4567-e89b-12d3-a456-426614174000"
	id := util.StringToNullUUID(categoryID)
	category := &entity.Category{
		ID:          id,
		Name:        "Sports",
		Description: "All about sports",
		Hobbies: []*hobbyEntity.Hobby{
			{
				ID:          util.StringToNullUUID("123e4567-e89b-12d3-a456-426614174001"),
				Name:        "Football",
				Description: "All about football",
				CategoryID:  id,
			},
			{
				ID:          util.StringToNullUUID("123e4567-e89b-12d3-a456-426614174002"),
				Name:        "Basketball",
				Description: "All about basketball",
				CategoryID:  id,
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.EXPECT().GetCategory(ctx, categoryID).Return(category, nil)

		res, err := getCategoryUsecase.Exec(ctx, categoryID)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, categoryID, res.Category.Id)
		assert.Equal(t, category.Name, res.Category.Name)
		assert.Equal(t, category.Description, res.Category.Description)
		assert.Equal(t, len(category.Hobbies), len(res.Hobbies))
	})

	t.Run("error", func(t *testing.T) {
		mockCategoryRepo.EXPECT().GetCategory(ctx, categoryID).Return(nil, errors.New("category not found"))

		res, err := getCategoryUsecase.Exec(ctx, categoryID)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
