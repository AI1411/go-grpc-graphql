package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AI1411/go-graphql-grpc/internal/env"
)

type Client struct {
	db *gorm.DB
}

func NewClient(e *env.DB) (*Client, error) {
	//gormLogger := initGormLogger(logger)
	db, err := open(e)
	if err != nil {
		return nil, fmt.Errorf("failed to connect master DB: %v", err)
	}

	//db.Logger = db.Logger.LogMode(gormLogger.LogLevel)

	return &Client{
		db: db,
	}, nil
}

func open(env *env.DB) (*gorm.DB, error) {
	log.Printf("env=%+v", env)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		env.PostgresHost,
		env.PostgresUsername,
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
