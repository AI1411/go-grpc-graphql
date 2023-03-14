package tweet

import (
	"context"

	"github.com/AI1411/go-grpc-praphql/internal/domain/tweet/entity"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
)

type TweetRepository interface {
	ListTweet(context.Context) ([]*entity.Tweet, error)
}

type tweetRepository struct {
	dbClient *db.Client
}

func NewTweetRepository(dbClient *db.Client) TweetRepository {
	return &tweetRepository{
		dbClient: dbClient,
	}
}

func (r *tweetRepository) ListTweet(ctx context.Context) ([]*entity.Tweet, error) {
	var tweets []*entity.Tweet
	if err := r.dbClient.Conn(ctx).Find(&tweets).Error; err != nil {
		return nil, err
	}

	return tweets, nil
}
