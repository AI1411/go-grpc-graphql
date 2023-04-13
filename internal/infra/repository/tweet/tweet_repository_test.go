package tweet_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/internal/domain/tweet/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
)

const (
	testTweetID  = "da0b1f2b-276a-417d-b4c2-77b81c8ad3c3"
	testTweetID2 = "d183dc26-2083-4375-8abe-5d292d84c0ce"
	testUserID   = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	testUserID2  = "3975482e-0133-4b4e-8d91-b8c983fbc9e6"
)

func TestTweetRepository_ListTweet(t *testing.T) {
	ctx := context.Background()

	testcasesListTweet := []struct {
		id        int
		name      string
		want      []*entity.Tweet
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系/アクティブなユーザのTweet一覧が取得できること",
			want: []*entity.Tweet{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testTweetID2),
						Valid: true,
					},
					UserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testUserID2),
						Valid: true,
					},
					Body:      "test",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testTweetID),
						Valid: true,
					},
					UserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testUserID),
						Valid: true,
					},
					Body:      "tweet",
					CreatedAt: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO tweets (id,user_id,body,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','cc293e0a-7342-4aac-b49b-a851e8af9dfc','tweet','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO tweets (id,user_id,body,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','3975482e-0133-4b4e-8d91-b8c983fbc9e6','test','2018-01-01T00:00:00+00:00','2019-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   2,
			name: "正常系/バンされた、退会済みのユーザのTweetは除いたtweet一覧が取得できること",
			want: []*entity.Tweet{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testTweetID2),
						Valid: true,
					},
					UserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testUserID2),
						Valid: true,
					},
					Body:      "test",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','アカウント停止','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO tweets (id,user_id,body,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','cc293e0a-7342-4aac-b49b-a851e8af9dfc','tweet','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO tweets (id,user_id,body,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','3975482e-0133-4b4e-8d91-b8c983fbc9e6','test','2018-01-01T00:00:00+00:00','2019-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesListTweet {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users", "tweets"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				tweetRepo := repository.NewTweetRepository(testClient)

				got, err := tweetRepo.ListTweet(ctx)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}

				if tt.want != nil {
					if !cmp.Equal(got, tt.want) {
						t.Errorf("diff %s", cmp.Diff(got, tt.want))
					}
				}
			},
		)
	}
}

func TestTweetRepository_CreateTweet(t *testing.T) {
	ctx := context.Background()

	testcasesCreateTweet := []struct {
		id        int
		name      string
		request   *entity.Tweet
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系/アクティブなユーザのTweet一覧が取得できること",
			request: &entity.Tweet{
				UserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testUserID),
					Valid: true,
				},
				Body:      "body",
				CreatedAt: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesCreateTweet {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users", "tweets"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				tweetRepo := repository.NewTweetRepository(testClient)

				_, err = tweetRepo.CreateTweet(ctx, tt.request)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}
			},
		)
	}
}
