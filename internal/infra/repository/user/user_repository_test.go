package user_test

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

	commonEntity "github.com/AI1411/go-grpc-praphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/testutil"
)

const (
	testUUID         = "27220eac-e75d-40cf-8163-e252c78bf2fe"
	notExistTestUUID = "27220eac-e75d-40cf-8163-e252c78bf2ff"
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	testcasesGetUser := []struct {
		id        int
		name      string
		userID    string
		want      *entity.User
		wantError error
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:     1,
			name:   "正常系/Userが取得できること",
			userID: testUUID,
			want: &entity.User{
				ID: uuid.NullUUID{
					UUID:  uuid.MustParse(testUUID),
					Valid: true,
				},
				Username:     "username",
				Email:        "test@gmail.com",
				Password:     "$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC",
				Status:       entity.UserStatusActive,
				Prefecture:   commonEntity.PrefectureOkayama,
				Introduction: "introduction",
				BloodType:    commonEntity.BloodTypeA,
				CreatedAt:    time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        2,
			name:      "異常系/Userが取得できない場合、NotFoundエラーを返すこと",
			userID:    notExistTestUUID,
			wantError: status.Errorf(codes.NotFound, "user not found"),
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesGetUser {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				userRepo := repository.NewUserRepository(testClient)

				got, err := userRepo.GetUser(ctx, tt.userID)
				if tt.wantError != nil {
					a.Error(err)
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
