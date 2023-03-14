package tweet

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/infra/repository/tweet"
	"github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type TweetUsecaseImpl interface {
	Exec(ctx context.Context) (*grpc.ListTweetResponse, error)
}

type listTweetUsecaseImpl struct {
	userRepository  user.UserRepository
	tweetRepository tweet.TweetRepository
}

func NewListTweetUsecaseImpl(userRepository user.UserRepository, tweetRepository tweet.TweetRepository) TweetUsecaseImpl {
	return &listTweetUsecaseImpl{
		userRepository:  userRepository,
		tweetRepository: tweetRepository,
	}
}

func (l *listTweetUsecaseImpl) Exec(ctx context.Context) (*grpc.ListTweetResponse, error) {
	tweets, err := l.tweetRepository.ListTweet(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*grpc.Tweet, len(tweets))
	for i, tweet := range tweets {
		res[i] = &grpc.Tweet{
			Id:        util.NullUUIDToString(tweet.ID),
			UserId:    util.NullUUIDToString(tweet.UserID),
			Body:      tweet.Body,
			CreatedAt: timestamppb.New(tweet.CreatedAt),
			UpdatedAt: timestamppb.New(tweet.UpdatedAt),
		}
	}

	return &grpc.ListTweetResponse{Tweets: res}, nil
}
