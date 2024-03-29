package tweet

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	tweetEntity "github.com/AI1411/go-grpc-graphql/internal/domain/tweet/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type CreateTweetUsecaseImpl interface {
	Exec(ctx context.Context, req *grpc.CreateTweetRequest) (*grpc.CreateTweetResponse, error)
}

type createTweetUsecaseImpl struct {
	tweetRepository tweet.Repository
}

func NewCreateTweetUsecaseImpl(tweetRepository tweet.Repository) CreateTweetUsecaseImpl {
	return &createTweetUsecaseImpl{
		tweetRepository: tweetRepository,
	}
}

func (c createTweetUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateTweetRequest) (*grpc.CreateTweetResponse, error) {
	tweetID, err := c.tweetRepository.CreateTweet(ctx, &tweetEntity.Tweet{
		UserID: util.StringToNullUUID(in.UserId),
		Body:   in.Body,
	})
	if err != nil {
		return nil, err
	}

	return &grpc.CreateTweetResponse{Id: tweetID}, nil
}
