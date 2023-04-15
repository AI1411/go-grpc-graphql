package tweet_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/grpc"
	tweetEntity "github.com/AI1411/go-grpc-graphql/internal/domain/tweet/entity"
	mockTweet "github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/tweet"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func TestCreateTweetUsecaseImpl_Exec(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTweetRepo := mockTweet.NewMockTweetRepository(ctrl)

	testTweet := &tweetEntity.Tweet{
		UserID:    util.StringToNullUUID(uuid.New().String()),
		Body:      "Hello!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("正常系", func(t *testing.T) {
		mockTweetRepo.
			EXPECT().
			CreateTweet(ctx, gomock.Any()).
			Return(util.NullUUIDToString(testTweet.ID), nil)

		uc := tweet.NewCreateTweetUsecaseImpl(mockTweetRepo)

		in := &grpc.CreateTweetRequest{
			UserId: util.NullUUIDToString(testTweet.UserID),
			Body:   testTweet.Body,
		}

		resp, err := uc.Exec(ctx, in)
		a.NoError(err)
		a.Equal(util.NullUUIDToString(testTweet.ID), resp.Id)
	})

	t.Run("異常系", func(t *testing.T) {
		mockTweetRepo.
			EXPECT().
			CreateTweet(ctx, gomock.Any()).
			Return("", status.Errorf(codes.Internal, "Internal Server Error"))

		uc := tweet.NewCreateTweetUsecaseImpl(mockTweetRepo)

		in := &grpc.CreateTweetRequest{
			UserId: util.NullUUIDToString(testTweet.UserID),
			Body:   testTweet.Body,
		}

		resp, err := uc.Exec(ctx, in)
		a.Error(err)
		a.Nil(resp)
	})
}
