package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-praphql/internal/env"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	"github.com/AI1411/go-grpc-praphql/internal/infra/logger"
)

func TestConnection(t *testing.T) (*db.Client, error) {
	a := assert.New(t)

	e, err := env.NewValue()
	a.NoError(err)

	zapLogger, err := logger.NewLogger(true)
	a.NoError(err)

	dbClient, err := db.NewClient(&e.DB, zapLogger)
	a.NoError(err)

	return dbClient, err
}
