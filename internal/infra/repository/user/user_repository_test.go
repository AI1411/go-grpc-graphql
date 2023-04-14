package user_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	commonEntity "github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
	"github.com/AI1411/go-grpc-graphql/internal/util"
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
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
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
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        2,
			name:      "異常系/Userが取得できない場合、NotFoundエラーを返すこと",
			userID:    notExistTestUUID,
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('242a9ea8-a9c7-40ba-b44a-deb3bae8ac6a','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        3,
			name:      "異常系/エラーが起きた場合はInternalエラーを返すこと",
			userID:    notExistTestUUID,
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
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
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				userRepo := repository.NewUserRepository(testClient, nil)

				got, err := userRepo.GetUser(ctx, tt.userID)
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

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	testcasesCreateUser := []struct {
		id        int
		name      string
		request   *entity.User
		want      *entity.User
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/Userが作成できること",
			request: &entity.User{
				Username:     "username",
				Email:        "test@gmail.com",
				Password:     "password",
				Status:       entity.UserStatusActive,
				Prefecture:   "岡山県",
				Introduction: "自己紹介",
				BloodType:    "A型",
			},
			want: &entity.User{
				Username:     "username",
				Email:        "test@gmail.com",
				Status:       entity.UserStatusActive,
				Prefecture:   commonEntity.PrefectureOkayama,
				Introduction: "自己紹介",
				BloodType:    commonEntity.BloodTypeA,
			},
		},
		{
			id:   2,
			name: "異常系/Userが作成できない場合、Internalエラーが返ること",
			request: &entity.User{
				Username:     strings.Repeat("a", 101),
				Email:        "test@gmail.com",
				Password:     "password",
				Status:       entity.UserStatusActive,
				Prefecture:   "岡山県",
				Introduction: "自己紹介",
				BloodType:    "A型",
			},
			wantError: codes.Internal,
		},
	}

	for _, tt := range testcasesCreateUser {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				userRepo := repository.NewUserRepository(testClient, nil)

				_, err = userRepo.CreateUser(ctx, tt.request)

				var got *entity.User

				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(testClient.Conn(ctx).First(&got).Error)
					a.NoError(err)
				}

				if tt.want != nil {
					opt := cmpopts.IgnoreFields(entity.User{}, "ID", "Password", "CreatedAt", "UpdatedAt")
					if !cmp.Equal(got, tt.want, opt) {
						t.Errorf("diff %s", cmp.Diff(got, tt.want))
					}
				}
			},
		)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	ctx := context.Background()

	testcasesUpdateUserProfile := []struct {
		id        int
		name      string
		request   *entity.User
		want      *entity.User
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/UserProfileが更新できること",
			request: &entity.User{
				ID:           util.StringToNullUUID(testUUID),
				Username:     "update",
				Prefecture:   commonEntity.PrefectureTokyo,
				Introduction: "更新",
				BloodType:    commonEntity.BloodTypeB,
			},
			want: &entity.User{
				ID:           util.StringToNullUUID(testUUID),
				Username:     "update",
				Email:        "test@gmail.com",
				Status:       entity.UserStatusActive,
				Prefecture:   commonEntity.PrefectureTokyo,
				Introduction: "更新",
				BloodType:    commonEntity.BloodTypeB,
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   2,
			name: "異常系/Userが見つからない場合、NotFoundエラーが返ること",
			request: &entity.User{
				ID:           util.StringToNullUUID(notExistTestUUID),
				Username:     "update",
				Prefecture:   commonEntity.PrefectureTokyo,
				Introduction: "更新",
				BloodType:    commonEntity.BloodTypeB,
			},
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   3,
			name: "異常系/UserProfileの更新に失敗した場合、Internalエラーが返ること",
			request: &entity.User{
				ID:           util.StringToNullUUID(testUUID),
				Username:     strings.Repeat("a", 101),
				Prefecture:   commonEntity.PrefectureTokyo,
				Introduction: "更新",
				BloodType:    commonEntity.BloodTypeB,
			},
			wantError: codes.Internal,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesUpdateUserProfile {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				userRepo := repository.NewUserRepository(testClient, nil)

				err = userRepo.UpdateUserProfile(ctx, tt.request)

				var got *entity.User

				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(testClient.Conn(ctx).Where("id = ?", testUUID).First(&got).Error)
					a.NoError(err)
				}

				if tt.want != nil {
					opt := cmpopts.IgnoreFields(entity.User{}, "ID", "Password", "CreatedAt", "UpdatedAt")
					if !cmp.Equal(got, tt.want, opt) {
						t.Errorf("diff %s", cmp.Diff(got, tt.want))
					}
				}
			},
		)
	}
}

func TestUpdateUserStatus(t *testing.T) {
	ctx := context.Background()

	testcasesUpdateUserStaus := []struct {
		id        int
		name      string
		request   *entity.User
		want      *entity.User
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/UserStatusが更新できること",
			request: &entity.User{
				ID:     util.StringToNullUUID(testUUID),
				Status: entity.UserStatusPremium,
			},
			want: &entity.User{
				ID:           util.StringToNullUUID(testUUID),
				Username:     "username",
				Email:        "test@gmail.com",
				Status:       entity.UserStatusPremium,
				Prefecture:   commonEntity.PrefectureOkayama,
				Introduction: "introduction",
				BloodType:    commonEntity.BloodTypeA,
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   2,
			name: "異常系/Userが見つからない場合、NotFoundエラーが返ること",
			request: &entity.User{
				ID:     util.StringToNullUUID(notExistTestUUID),
				Status: entity.UserStatusPremium,
			},
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   3,
			name: "異常系/UserStatusの更新に失敗した場合、Internalエラーが返ること",
			request: &entity.User{
				ID:     util.StringToNullUUID(testUUID),
				Status: "invalid status",
			},
			wantError: codes.Internal,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesUpdateUserStaus {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				userRepo := repository.NewUserRepository(testClient, nil)

				err = userRepo.UpdateUserStatus(ctx, tt.request)

				var got *entity.User

				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(testClient.Conn(ctx).Where("id = ?", testUUID).First(&got).Error)
					a.NoError(err)
				}

				if tt.want != nil {
					opt := cmpopts.IgnoreFields(entity.User{}, "ID", "Password", "CreatedAt", "UpdatedAt")
					if !cmp.Equal(got, tt.want, opt) {
						t.Errorf("diff %s", cmp.Diff(got, tt.want))
					}
				}
			},
		)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	ctx := context.Background()

	testcasesUpdateUserPassword := []struct {
		id        int
		name      string
		request   *entity.UserPassword
		want      *entity.User
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/UserPasswordが更新できること",
			request: &entity.UserPassword{
				ID:                   util.StringToNullUUID(testUUID),
				Password:             "password!",
				PasswordConfirmation: "password!",
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   2,
			name: "異常系/Userが見つからない場合、NotFoundエラーが返ること",
			request: &entity.UserPassword{
				ID:                   util.StringToNullUUID(notExistTestUUID),
				Password:             "password!",
				PasswordConfirmation: "password!",
			},
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('27220eac-e75d-40cf-8163-e252c78bf2fe','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesUpdateUserPassword {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"users"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				userRepo := repository.NewUserRepository(testClient, nil)

				err = userRepo.UpdateUserPassword(ctx, tt.request)

				var got *entity.User

				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(testClient.Conn(ctx).Where("id = ?", testUUID).First(&got).Error)

					a.NoError(bcrypt.CompareHashAndPassword([]byte(got.Password), []byte("password!")))
				}
			},
		)
	}
}
