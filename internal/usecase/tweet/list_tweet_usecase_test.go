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

	tweetEntity "github.com/AI1411/go-grpc-graphql/internal/domain/tweet/entity"
	mockTweet "github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet/mock"
	mockUser "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/tweet"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func TestLIstTweetUsecaseImpl_Exec(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTweetRepo := mockTweet.NewMockTweetRepository(ctrl)
	mockUserRepo := mockUser.NewMockUserRepository(ctrl)

	testTweet := []*tweetEntity.Tweet{
		{
			UserID:    util.StringToNullUUID(uuid.New().String()),
			Body:      "Hello!",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    util.StringToNullUUID(uuid.New().String()),
			Body:      "Hello World!",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("正常系", func(t *testing.T) {
		mockTweetRepo.
			EXPECT().
			ListTweet(ctx).
			Return(testTweet, nil)

		uc := tweet.NewListTweetUsecaseImpl(mockUserRepo, mockTweetRepo)

		resp, err := uc.Exec(ctx)
		a.NoError(err)
		a.Equal(len(testTweet), len(resp.Tweets))
	})

	t.Run("異常系", func(t *testing.T) {
		mockTweetRepo.
			EXPECT().
			ListTweet(ctx).
			Return(nil, status.Errorf(codes.Internal, "Internal Server Error"))

		uc := tweet.NewListTweetUsecaseImpl(mockUserRepo, mockTweetRepo)

		resp, err := uc.Exec(ctx)
		a.Error(err)
		a.Nil(resp)
	})
}
