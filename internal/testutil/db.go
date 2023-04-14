package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/env"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/infra/logger"
)

func TestConnection(t *testing.T) (db.Client, error) {
	a := assert.New(t)

	e, err := env.NewValue()
	a.NoError(err)

	zapLogger, err := logger.NewLogger(true)
	a.NoError(err)

	dbClient := db.NewClient(&e.DB, zapLogger)
	return dbClient, err
}
