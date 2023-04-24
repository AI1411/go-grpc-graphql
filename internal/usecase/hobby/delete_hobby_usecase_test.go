package hobby_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/grpc"
	mockHobby "github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/hobby"
)

func TestDeleteHobbyUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHobbyRepo := mockHobby.NewMockRepository(ctrl)

	deleteHobbyUsecase := hobby.NewDeleteHobbyUsecaseImpl(mockHobbyRepo)

	ctx := context.Background()
	newHobby := &grpc.DeleteHobbyRequest{
		Id: hobbyID,
	}

	t.Run("success", func(t *testing.T) {
		mockHobbyRepo.EXPECT().DeleteHobby(ctx, hobbyID).Return(nil).Times(1)

		err := deleteHobbyUsecase.Exec(ctx, newHobby.GetId())

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockHobbyRepo.
			EXPECT().
			DeleteHobby(ctx, hobbyID).
			Return(errors.New("error creating hobby")).
			Times(1)

		err := deleteHobbyUsecase.Exec(ctx, newHobby.GetId())

		assert.Error(t, err)
	})
}
