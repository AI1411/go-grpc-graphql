package chat_test

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

	"github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
)

const (
	testChatID     = "da0b1f2b-276a-417d-b4c2-77b81c8ad3c3"
	testChatID2    = "d183dc26-2083-4375-8abe-5d292d84c0ce"
	testRoomID     = "18bb7429-e891-4f41-b045-a52aaf53ea93"
	testFromUserID = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	notExistUserID = "a35f1d56-068d-4ec5-a892-a83e479393d7"
	testToUserID   = "3975482e-0133-4b4e-8d91-b8c983fbc9e6"
	testToUserID2  = "e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b"
)

func TestChatRepository_ListChat(t *testing.T) {
	ctx := context.Background()

	testcasesListChat := []struct {
		id        int
		name      string
		request   *entity.Chat
		want      []*entity.Chat
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系/Chat一覧が取得できること",
			request: &entity.Chat{
				RoomID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				FromUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testFromUserID),
					Valid: true,
				},
			},
			want: []*entity.Chat{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testChatID),
						Valid: true,
					},
					RoomID: uuid.NullUUID{
						UUID:  uuid.MustParse(testRoomID),
						Valid: true,
					},
					FromUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testFromUserID),
						Valid: true,
					},
					ToUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testToUserID),
						Valid: true,
					},
					Body:      "body",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testChatID2),
						Valid: true,
					},
					RoomID: uuid.NullUUID{
						UUID:  uuid.MustParse(testRoomID),
						Valid: true,
					},
					FromUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testFromUserID),
						Valid: true,
					},
					ToUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testToUserID2),
						Valid: true,
					},
					Body:      "test",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b','test',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesListChat {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				chatRepo := repository.NewChatRepository(testClient)

				got, err := chatRepo.ListChat(ctx, tt.request)
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

func TestChatRepository_CreateChat(t *testing.T) {
	ctx := context.Background()

	testcasesCreateChat := []struct {
		id        int
		name      string
		request   *entity.Chat
		want      []*entity.Chat
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:   1,
			name: "正常系/Chatが作成できること",
			request: &entity.Chat{
				RoomID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				FromUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testFromUserID),
					Valid: true,
				},
				ToUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testToUserID),
					Valid: true,
				},
			},
			want: []*entity.Chat{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testChatID),
						Valid: true,
					},
					RoomID: uuid.NullUUID{
						UUID:  uuid.MustParse(testRoomID),
						Valid: true,
					},
					FromUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testFromUserID),
						Valid: true,
					},
					ToUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testToUserID),
						Valid: true,
					},
					Body:      "body",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testChatID2),
						Valid: true,
					},
					RoomID: uuid.NullUUID{
						UUID:  uuid.MustParse(testRoomID),
						Valid: true,
					},
					FromUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testFromUserID),
						Valid: true,
					},
					ToUserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testToUserID2),
						Valid: true,
					},
					Body:      "test",
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   2,
			name: "異常系/送り元（FromUserID）が存在しない場合、NotFoundエラーが返ること",
			request: &entity.Chat{
				RoomID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				FromUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(notExistUserID),
					Valid: true,
				},
				ToUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testToUserID),
					Valid: true,
				},
			},
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   3,
			name: "異常系/送り先（TOUserID）が存在しない場合、NotFoundエラーが返ること",
			request: &entity.Chat{
				RoomID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				FromUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testFromUserID),
					Valid: true,
				},
				ToUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(notExistUserID),
					Valid: true,
				},
			},
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   4,
			name: "異常系/送り先（TOUserID）のステータスが異常な場合、FailedPreconditionエラーが返ること",
			request: &entity.Chat{
				RoomID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				FromUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testFromUserID),
					Valid: true,
				},
				ToUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testToUserID),
					Valid: true,
				},
			},
			wantError: codes.FailedPrecondition,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','退会済','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:   5,
			name: "異常系/送り先（TOUserID）のステータスが異常な場合、FailedPreconditionエラーが返ること",
			request: &entity.Chat{
				RoomID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				FromUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testFromUserID),
					Valid: true,
				},
				ToUserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testToUserID),
					Valid: true,
				},
			},
			wantError: codes.FailedPrecondition,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','アカウント停止','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','通常会員','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesCreateChat {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				chatRepo := repository.NewChatRepository(testClient)

				_, err = chatRepo.CreateChat(ctx, tt.request)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}
			},
		)
	}
}

func TestChatRepository_MarkChatAsRead(t *testing.T) {
	ctx := context.Background()

	testcasesMarkChatAsRead := []struct {
		id        int
		name      string
		chatID    string
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:     1,
			name:   "正常系/Chatが既読にできること",
			chatID: testChatID,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        2,
			name:      "異常系/Chatが見つからない場合、NotFoundエラーが返ること",
			chatID:    notExistUserID,
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        3,
			name:      "異常系/Chatが既読の場合はInvalidArgumentが返ること",
			chatID:    testChatID,
			wantError: codes.InvalidArgument,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','username','test@gmail.com','$2a$10$Ig2ubFhcRtxTswDOZ95ymOfpnhRjm4DhmTPwlp1VtC.3NoCO4y2aC','通常会員','岡山県','introduction','A型','2017-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO users (id,username,email,password,status,prefecture,introduction,blood_type,created_at,updated_at) VALUES ('3975482e-0133-4b4e-8d91-b8c983fbc9e6','tetuser','usr@gmail.com','$2a$10$.0GNxvJhIqEuE4riZhpvAe/H83bbmstg2PGtlsPBidyd/R51ooW9y','プレミアム','岩手県','自己紹介','B型','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',true,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesMarkChatAsRead {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				chatRepo := repository.NewChatRepository(testClient)

				err = chatRepo.MarkChatAsRead(ctx, tt.chatID)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}
			},
		)
	}
}
