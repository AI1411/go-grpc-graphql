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

type Client interface {
	Conn(ctx context.Context) *gorm.DB
	TruncateTable(ctx context.Context, t *testing.T, tables []string)
}

type client struct {
	db *gorm.DB
}

func NewClient(e *env.DB, zapLogger *zap.Logger) Client {
	gormLogger := initGormLogger(zapLogger)
	db, _ := open(e)

	db.Logger = db.Logger.LogMode(gormLogger.LogLevel)

	return &client{
		db: db,
	}
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

func (c *client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}

func (c *client) TruncateTable(ctx context.Context, t *testing.T, tables []string) {
	for _, table := range tables {
		require.NoError(t, c.Conn(ctx).Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table)).Error)
	}
}
