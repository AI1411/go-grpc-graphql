package tweet

import (
	"context"

	"github.com/AI1411/go-grpc-praphql/internal/domain/tweet/entity"
	userEntity "github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type TweetRepository interface {
	ListTweet(context.Context) ([]*entity.Tweet, error)
	CreateTweet(ctx context.Context, tweet *entity.Tweet) (string, error)
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
	var users []*userEntity.User
	// 退会、バンされているユーザは除く
	userStatusActive := []string{userEntity.UserStatusActive.String(), userEntity.UserStatusPremium.String()}
	if err := r.dbClient.Conn(ctx).Where("status", userStatusActive).Find(&users).Error; err != nil {
		return nil, err
	}

	userIDs := make([]string, len(users))
	for i, user := range users {
		userIDs[i] = util.NullUUIDToString(user.ID)
	}

	var tweets []*entity.Tweet
	if err := r.dbClient.Conn(ctx).Where("user_id", userIDs).Order("created_at DESC").Find(&tweets).Error; err != nil {
		return nil, err
	}

	return tweets, nil
}

func (r *tweetRepository) CreateTweet(ctx context.Context, tweet *entity.Tweet) (string, error) {
	if err := r.dbClient.Conn(ctx).Create(tweet).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(tweet.ID), nil
}
