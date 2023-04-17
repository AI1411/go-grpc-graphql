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

func TestGetHobbyUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHobbyRepo := mockHobby.NewMockRepository(ctrl)

	getHobbyUsecase := hobby.NewGetHobbyUsecaseImpl(mockHobbyRepo)

	ctx := context.Background()
	newHobby := &grpc.GetHobbyRequest{
		Id: hobbyID,
	}

	expected := &entity.Hobby{
		ID:          util.StringToNullUUID(hobbyID),
		Name:        "hobby",
		Description: "hobby description",
		CategoryID:  util.StringToNullUUID(categoryID),
	}

	t.Run("success", func(t *testing.T) {
		mockHobbyRepo.EXPECT().GetHobby(ctx, newHobby.GetId()).Return(expected, nil).Times(1)

		actual, err := getHobbyUsecase.Exec(ctx, newHobby.GetId())

		assert.NoError(t, err)
		assert.Equal(t, util.NullUUIDToString(expected.ID), actual.Hobby.Id)
		assert.Equal(t, expected.Name, actual.Hobby.Name)
		assert.Equal(t, expected.Description, actual.Hobby.Description)
		assert.Equal(t, util.NullUUIDToString(expected.CategoryID), actual.Hobby.CategoryId)
	})

	t.Run("error", func(t *testing.T) {
		mockHobbyRepo.
			EXPECT().
			GetHobby(ctx, hobbyID).
			Return(nil, errors.New("error creating hobby")).
			Times(1)

		actual, err := getHobbyUsecase.Exec(ctx, newHobby.GetId())

		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
