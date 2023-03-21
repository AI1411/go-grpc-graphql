package room_test

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

	chatEntity "github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/room/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
)

const (
	testUserID     = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	testRoomID     = "18bb7429-e891-4f41-b045-a52aaf53ea93"
	testRoomID2    = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	testChatID     = "da0b1f2b-276a-417d-b4c2-77b81c8ad3c3"
	testChatID2    = "d183dc26-2083-4375-8abe-5d292d84c0ce"
	testFromUserID = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	notExistUserID = "a35f1d56-068d-4ec5-a892-a83e479393d7"
	testToUserID   = "3975482e-0133-4b4e-8d91-b8c983fbc9e6"
	testToUserID2  = "e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b"
)

func TestChatRepository_ListRoom(t *testing.T) {
	ctx := context.Background()

	testcasesListRoom := []struct {
		id        int
		name      string
		userID    string
		want      []*entity.Room
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:     1,
			name:   "正常系/Room一覧が取得できること",
			userID: testUserID,
			want: []*entity.Room{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testRoomID),
						Valid: true,
					},
					UserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testUserID),
						Valid: true,
					},
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(testRoomID2),
						Valid: true,
					},
					UserID: uuid.NullUUID{
						UUID:  uuid.MustParse(testUserID),
						Valid: true,
					},
					CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesListRoom {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				roomRepo := repository.NewRoomRepository(testClient)

				got, err := roomRepo.ListRoom(ctx, tt.userID)
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

func TestChatRepository_GetRoom(t *testing.T) {
	ctx := context.Background()

	testcasesGetRoom := []struct {
		id        int
		name      string
		roomID    string
		want      *entity.Room
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:     1,
			name:   "正常系/Roomが取得できること",
			roomID: testRoomID,
			want: &entity.Room{
				ID: uuid.NullUUID{
					UUID:  uuid.MustParse(testRoomID),
					Valid: true,
				},
				UserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testUserID),
					Valid: true,
				},
				Chats: []*chatEntity.Chat{
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
				CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b','test',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
		{
			id:        2,
			name:      "正常系/Roomが取得できること",
			roomID:    notExistUserID,
			wantError: codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO rooms (id,user_id,created_at,updated_at) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc','cc293e0a-7342-4aac-b49b-a851e8af9dfc','2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','3975482e-0133-4b4e-8d91-b8c983fbc9e6','body',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO chats (id,room_id,from_user_id,to_user_id,body,is_read,created_at,updated_at) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce','18bb7429-e891-4f41-b045-a52aaf53ea93','cc293e0a-7342-4aac-b49b-a851e8af9dfc','e5b5b2a1-0b1f-4b0e-8c1f-8b5b0b2b2b2b','test',false,'2018-01-01T00:00:00+00:00','2018-01-01T00:00:00+00:00');`).Error)
			},
		},
	}

	for _, tt := range testcasesGetRoom {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				roomRepo := repository.NewRoomRepository(testClient)

				got, err := roomRepo.GetRoom(ctx, tt.roomID)
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

func TestChatRepository_CreateRoom(t *testing.T) {
	ctx := context.Background()

	testcasesCreateRoom := []struct {
		id        int
		name      string
		request   *entity.Room
		want      *entity.Room
		wantError codes.Code
	}{
		{
			id:   1,
			name: "正常系/Roomが取得できること",
			request: &entity.Room{
				UserID: uuid.NullUUID{
					UUID:  uuid.MustParse(testUserID),
					Valid: true,
				},
			},
		},
	}

	for _, tt := range testcasesCreateRoom {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})

				roomRepo := repository.NewRoomRepository(testClient)

				_, err = roomRepo.CreateRoom(ctx, tt.request)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}
			},
		)
	}
}

func TestChatRepository_DeleteRoom(t *testing.T) {
	ctx := context.Background()

	testcasesDeleteRoom := []struct {
		id        int
		name      string
		roomID    string
		wantError codes.Code
	}{
		{
			id:     1,
			name:   "正常系/Roomが削除できること",
			roomID: testRoomID,
		},
	}

	for _, tt := range testcasesDeleteRoom {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				testClient.TruncateTable(ctx, t, []string{"users", "chats", "rooms"})

				roomRepo := repository.NewRoomRepository(testClient)

				err = roomRepo.DeleteRoom(ctx, tt.roomID)
				if tt.wantError != 0 {
					a.Equal(status.Code(err), tt.wantError)
				} else {
					a.NoError(err)
				}
			},
		)
	}
}
