package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
)

func Test_distributePointAllUsers(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		id        int
		name      string
		want      *entity.UserPoint
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系",
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO user_points (id,user_id,point,created_at,updated_at) VALUES ('178278fb-08ec-4fd4-85b8-06eeec1f505e','27220eac-e75d-40cf-8163-e252c78bf2fe',50,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO user_points (id,user_id,point,created_at,updated_at) VALUES ('d3c4e037-1f90-4100-bc2f-15709740bb8f','242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a',10,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        2,
			name:      "異常系/ユーザが存在しない場合、NotFoundエラーが返ること",
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			testClient, err := testutil.TestConnection(t)
			a.NoError(err)
			testClient.TruncateTable(ctx, t, []string{"users", "user_points"})
			if tt.setup != nil {
				tt.setup(ctx, t, testClient)
			}

			err = distributePointAllUsers()
			if tt.wantError != codes.OK {
				a.Error(err)
				a.Equal(tt.wantError, status.Code(err))
			} else {
				a.NoError(err)
				var userPoints []*entity.UserPoint
				a.NoError(testClient.Conn(ctx).Find(&userPoints).Error)
				a.Equal(2, len(userPoints))
				a.Equal(100, userPoints[0].Point)
				a.Equal(60, userPoints[1].Point)
			}
		})
	}
}
