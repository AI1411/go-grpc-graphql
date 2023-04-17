package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	mockCategory "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/category"
	"github.com/AI1411/go-grpc-graphql/internal/util"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListCategoryUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepo := mockCategory.NewMockRepository(ctrl)

	listCategoryUsecase := category.NewListCategoryUsecaseImpl(mockCategoryRepo)

	ctx := context.Background()
	categoryID := "123e4567-e89b-12d3-a456-426614174000"
	categoryID2 := "123e4567-e89b-12d3-a456-426614174001"

	listCategoryRequest := &grpc.ListCategoryRequest{
		Name:   "Sports",
		Order:  "ASC",
		Offset: 0,
		Limit:  10,
	}

	categories := []*entity.Category{
		{
			ID:          util.StringToNullUUID(categoryID),
			Name:        "Sports",
			Description: "All about sports",
		},
		{
			ID:          util.StringToNullUUID(categoryID2),
			Name:        "Music",
			Description: "All about music",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.EXPECT().ListCategory(ctx, &entity.CategoryCondition{
			Name:   listCategoryRequest.GetName(),
			Order:  listCategoryRequest.GetOrder(),
			Offset: listCategoryRequest.GetOffset(),
			Limit:  listCategoryRequest.GetLimit(),
		}).Return(categories, nil)

		res, err := listCategoryUsecase.Exec(ctx, listCategoryRequest)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, len(categories), len(res.Categories))
		for i, v := range categories {
			assert.Equal(t, util.NullUUIDToString(v.ID), res.Categories[i].Id)
			assert.Equal(t, v.Name, res.Categories[i].Name)
			assert.Equal(t, v.Description, res.Categories[i].Description)
		}
	})

	t.Run("error", func(t *testing.T) {
		mockCategoryRepo.EXPECT().ListCategory(ctx, &entity.CategoryCondition{
			Name:   listCategoryRequest.GetName(),
			Order:  listCategoryRequest.GetOrder(),
			Offset: listCategoryRequest.GetOffset(),
			Limit:  listCategoryRequest.GetLimit(),
		}).Return(nil, errors.New("error listing categories"))

		res, err := listCategoryUsecase.Exec(ctx, listCategoryRequest)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
