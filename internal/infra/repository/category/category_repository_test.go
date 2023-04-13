package category_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	hobbyEntity "github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	categoryRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/category"
	"github.com/AI1411/go-grpc-graphql/internal/testutil"
)

const (
	categoryID     = "cc293e0a-7342-4aac-b49b-a851e8af9dfc"
	categoryID2    = "18bb7429-e891-4f41-b045-a52aaf53ea93"
	hobbyID        = "da0b1f2b-276a-417d-b4c2-77b81c8ad3c3"
	hobbyID2       = "d183dc26-2083-4375-8abe-5d292d84c0ce"
	notExistUserID = "a35f1d56-068d-4ec5-a892-a83e479393d7"
)

func TestChatRepository_ListCategory(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		id        int
		name      string
		request   *entity.CategoryCondition
		want      []*entity.Category
		wantError codes.Code
		setup     func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:      1,
			name:    "正常系/Category一覧が取得できること",
			request: &entity.CategoryCondition{},
			want: []*entity.Category{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(categoryID),
						Valid: true,
					},
					Name:        "category",
					Description: "description",
				},
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(categoryID2),
						Valid: true,
					},
					Name:        "name",
					Description: "remarks",
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93', 'name', 'remarks');`).Error)
			},
		},
		{
			id:   2,
			name: "正常系/Category一覧が取得できること/Nameで絞り込み",
			request: &entity.CategoryCondition{
				Name: "category",
			},
			want: []*entity.Category{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(categoryID),
						Valid: true,
					},
					Name:        "category",
					Description: "description",
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93', 'name', 'remarks');`).Error)
			},
		},
		{
			id:   3,
			name: "正常系/Category一覧が取得できること/Limitで絞り込み",
			request: &entity.CategoryCondition{
				Limit: 1,
			},
			want: []*entity.Category{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(categoryID),
						Valid: true,
					},
					Name:        "category",
					Description: "description",
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93', 'name', 'remarks');`).Error)
			},
		},
		{
			id:   4,
			name: "正常系/Category一覧が取得できること/Offsetで絞り込み",
			request: &entity.CategoryCondition{
				Offset: 1,
			},
			want: []*entity.Category{
				{
					ID: uuid.NullUUID{
						UUID:  uuid.MustParse(categoryID2),
						Valid: true,
					},
					Name:        "name",
					Description: "remarks",
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93', 'name', 'remarks');`).Error)
			},
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"categories"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				roomRepo := categoryRepo.NewCategoryRepository(testClient)

				got, err := roomRepo.ListCategory(ctx, tt.request)
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

func TestChatRepository_GetCategory(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		id         int
		name       string
		categoryID string
		want       *entity.Category
		wantError  codes.Code
		setup      func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:         1,
			name:       "正常系/Categoryが取得できること",
			categoryID: categoryID,
			want: &entity.Category{
				ID: uuid.NullUUID{
					UUID:  uuid.MustParse(categoryID),
					Valid: true,
				},
				Name:        "category",
				Description: "description",
				Hobbies: []*hobbyEntity.Hobby{
					{
						ID: uuid.NullUUID{
							UUID:  uuid.MustParse(hobbyID),
							Valid: true,
						},
						Name:        "スポーツ",
						Description: "大谷翔平",
						CategoryID: uuid.NullUUID{
							UUID:  uuid.MustParse(categoryID),
							Valid: true,
						},
					},
					{
						ID: uuid.NullUUID{
							UUID:  uuid.MustParse(hobbyID2),
							Valid: true,
						},
						Name:        "ゲーム",
						Description: "フォートナイト",
						CategoryID: uuid.NullUUID{
							UUID:  uuid.MustParse(categoryID),
							Valid: true,
						},
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93', 'name', 'remarks');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO hobbies (id,name,description, category_id) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3', 'スポーツ', '大谷翔平', 'cc293e0a-7342-4aac-b49b-a851e8af9dfc');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO hobbies (id,name,description, category_id) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce', 'ゲーム', 'フォートナイト', 'cc293e0a-7342-4aac-b49b-a851e8af9dfc');`).Error)
			},
		},
		{
			id:         2,
			name:       "異常系/Categoryが存在しない場合はNotFoundエラー",
			categoryID: notExistUserID,
			wantError:  codes.NotFound,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('18bb7429-e891-4f41-b045-a52aaf53ea93', 'name', 'remarks');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO hobbies (id,name,description, category_id) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3', 'スポーツ', '大谷翔平', 'cc293e0a-7342-4aac-b49b-a851e8af9dfc');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO hobbies (id,name,description, category_id) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce', 'ゲーム', 'フォートナイト', 'cc293e0a-7342-4aac-b49b-a851e8af9dfc');`).Error)
			},
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"categories"})
				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				roomRepo := categoryRepo.NewCategoryRepository(testClient)

				got, err := roomRepo.GetCategory(ctx, tt.categoryID)
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

func TestChatRepository_CreateCategory(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		id      int
		name    string
		request *entity.Category
	}{
		{
			id:   1,
			name: "正常系/Categoryが作成できること",
			request: &entity.Category{
				Name:        "category",
				Description: "description",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"categories"})

				roomRepo := categoryRepo.NewCategoryRepository(testClient)

				_, err = roomRepo.CreateCategory(ctx, tt.request)
				a.NoError(err)
			},
		)
	}
}

func TestChatRepository_DeleteCategory(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		id         int
		name       string
		categoryID string
		setup      func(ctx context.Context, t *testing.T, dbClient *db.Client)
	}{
		{
			id:         1,
			name:       "正常系/Categoryが削除できること",
			categoryID: categoryID,
			setup: func(ctx context.Context, t *testing.T, dbClient *db.Client) {
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO categories (id,name,description) VALUES ('cc293e0a-7342-4aac-b49b-a851e8af9dfc', 'category', 'description');`).Error)

				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO hobbies (id,name,description, category_id) VALUES ('da0b1f2b-276a-417d-b4c2-77b81c8ad3c3', 'スポーツ', '大谷翔平', 'cc293e0a-7342-4aac-b49b-a851e8af9dfc');`).Error)
				require.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO hobbies (id,name,description, category_id) VALUES ('d183dc26-2083-4375-8abe-5d292d84c0ce', 'ゲーム', 'フォートナイト', 'cc293e0a-7342-4aac-b49b-a851e8af9dfc');`).Error)
			},
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)

				testClient, err := testutil.TestConnection(t)
				a.NoError(err)
				testClient.TruncateTable(ctx, t, []string{"categories"})

				if tt.setup != nil {
					tt.setup(ctx, t, testClient)
				}

				roomRepo := categoryRepo.NewCategoryRepository(testClient)

				err = roomRepo.DeleteCategory(ctx, tt.categoryID)
				a.NoError(err)
			},
		)
	}
}
