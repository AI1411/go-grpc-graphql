package report_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
)

const (
	testUserID    = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	testReportID  = "18bb7429-e891-4f41-b045-a52aaf53ea93"
	testReportID2 = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	testReportID3 = "da0b1f2b-276a-417d-b4c2-77b81c8ad3c3"
	testReportID4 = "d183dc26-2083-4375-8abe-5d292d84c0ce"
	testReportID5 = "d183dc26-2083-4375-8abe-5d292d84c0ce"
	testReportID6 = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	testChatID    = "a35f1d56-068d-4ec5-a892-a83e479393d7"
	testChatID2   = "3975482e-0133-4b4e-8d91-b8c983fbc9e6"
	testToUserID2 = "e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b"
)

func Test_reportRepository_ListReport(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		id        int
		name      string
		userID    string
		want      []*entity.Report
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:     1,
			name:   "正常系/Report一覧が取得できること",
			userID: testReportID3,
			want: []*entity.Report{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testReportID),
						Valid: true,
					},
					ReporterUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testReportID2),
						Valid: true,
					},
					ReportedUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testReportID3),
						Valid: true,
					},
					ReportedChatID: uuid.NullUUID{
						UUID:  uuid.MustParse(testChatID),
						Valid: true,
					},
					Status:    entity.ReportStatusPending,
					Reason:    "暴言",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testReportID5),
						Valid: true,
					},
					ReporterUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testReportID6),
						Valid: true,
					},
					ReportedUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testReportID3),
						Valid: true,
					},
					ReportedChatID: uuid.NullUUID{
						UUID:  uuid.MustParse(testChatID2),
						Valid: true,
					},
					Status:    entity.ReportStatusPending,
					Reason:    "規約違反",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantError: 0,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b','test',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO reports (id,reporter_user_id, reported_user_id, reported_chat_id,reason,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','a35f1d56-068d-4ec5-a892-a83e479393d7','暴言','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO reports (id,reporter_user_id, reported_user_id, reported_chat_id,reason,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','cc293e0a-7342-4aac-b49b-a851e8af9dfc','da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','3975482e-0133-4b4e-8d91-b8c983fbc9e6','規約違反','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms", "reports"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				reportRepo := repository.NewReportRepository(testClient)

				got, err := reportRepo.ListReport(ctx, tt.userID)
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

func Test_reportRepository_GetUserReportCount(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		id        int
		name      string
		userID    string
		want      []*entity.ReportCount
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系/ユーザごとのReport一覧と回数が取得できること",
			want: []*entity.ReportCount{
				{
					ReportedUserID: uuid.NullUUID{
						UUID:  uuid.MustParse("da0b1f2b-276a-417d-b4c2-77b81c8ad3c3"),
						Valid: true,
					},
					ReportCount: 2,
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b','test',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO reports (id,reporter_user_id, reported_user_id, reported_chat_id,reason,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','a35f1d56-068d-4ec5-a892-a83e479393d7','暴言','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO reports (id,reporter_user_id, reported_user_id, reported_chat_id,reason,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','cc293e0a-7342-4aac-b49b-a851e8af9dfc','da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','3975482e-0133-4b4e-8d91-b8c983fbc9e6','規約違反','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms", "reports"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				reportRepo := repository.NewReportRepository(testClient)

				got, err := reportRepo.GetUserReportCount(ctx)
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

func Test_reportRepository_CreateReportCount(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		id        int
		name      string
		request   *entity.Report
		want      *entity.Report
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系/ユーザごとのReport一覧と回数が取得できること",
			request: &entity.Report{
				ReporterUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testReportID2),
					Valid: true,
				},
				ReportedUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testReportID3),
					Valid: true,
				},
				ReportedChatID: uuid.NullUUID{
					UUID:  uuid.MustParse(testChatID),
					Valid: true,
				},
				Status: entity.ReportStatusPending,
				Reason: "暴言",
			},
			want: &entity.Report{
				ID: uuid.NullUUID{
					UUID:  uuid.MustParse(testReportID),
					Valid: true,
				},
				ReporterUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testReportID2),
					Valid: true,
				},
				ReportedUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testReportID3),
					Valid: true,
				},
				ReportedChatID: uuid.NullUUID{
					UUID:  uuid.MustParse(testChatID),
					Valid: true,
				},
				Status:    entity.ReportStatusPending,
				Reason:    "暴言",
				CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms", "reports"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				reportRepo := repository.NewReportRepository(testClient)

				id, err := reportRepo.CreateReport(ctx, tt.request)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}

				var got *entity.Report
				err = testClient.Conn(ctx).Raw(`SELECT * FROM reports WHERE id = ?`, id).Scan(&got).Error

				if tt.want != nil {
					opt := cmpopts.IgnoreFields(entity.Report{}, "ID", "CreatedAt", "UpdatedAt")
					if !cmp.Equal(got, tt.want, opt) {
						t.Errorf("diff %s", cmp.Diff(got, tt.want))
					}
				}
			},
		)
	}
}
