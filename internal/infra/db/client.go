package db

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-praphql/internal/env"
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
