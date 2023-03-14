package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/env"
)

type Client struct {
	db *gorm.DB
}

func NewClient(e *env.DB, zapLogger *zap.Logger) (*Client, error) {
	gormLogger := initGormLogger(zapLogger)
	db, err := open(e)
	if err != nil {
		return nil, fmt.Errorf("failed to connect master DB: %v", err)
	}

	db.Logger = db.Logger.LogMode(gormLogger.LogLevel)

	return &Client{
		db: db,
	}, nil
}

func open(env *env.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		env.PostgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDatabase,
		env.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *Client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}

func (c *Client) TruncateTable(ctx context.Context, t *testing.T, tables []string) {
	for _, table := range tables {
		require.NoError(t, c.Conn(ctx).Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY;", table)).Error)
	}
}
