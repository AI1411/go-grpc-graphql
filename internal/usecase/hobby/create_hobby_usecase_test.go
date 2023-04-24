package hobby_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
	mockHobby "github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/hobby"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

const (
	hobbyID    = "123e4567-e89b-12d3-a456-426614174000"
	categoryID = "123e4567-e89b-12d3-a456-426614174001"
)

func TestCreateHobbyUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHobbyRepo := mockHobby.NewMockRepository(ctrl)

	createHobbyUsecase := hobby.NewCreateHobbyUsecaseImpl(mockHobbyRepo)

	ctx := context.Background()
	newHobby := &grpc.CreateHobbyRequest{
		Name:        "Sports",
		Description: "All about sports",
		CategoryId:  categoryID,
	}

	t.Run("success", func(t *testing.T) {
		mockHobbyRepo.EXPECT().CreateHobby(ctx, &entity.Hobby{
			Name:        newHobby.GetName(),
			Description: newHobby.GetDescription(),
			CategoryID:  util.StringToNullUUID(newHobby.GetCategoryId()),
		}).Return(hobbyID, nil).Times(1)

		res, err := createHobbyUsecase.Exec(ctx, newHobby)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, hobbyID, res.Id)
	})

	t.Run("error", func(t *testing.T) {
		mockHobbyRepo.EXPECT().CreateHobby(ctx, &entity.Hobby{
			Name:        newHobby.GetName(),
			Description: newHobby.GetDescription(),
			CategoryID:  util.StringToNullUUID(categoryID),
		}).Return("", errors.New("error creating hobby"))

		res, err := createHobbyUsecase.Exec(ctx, newHobby)

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
