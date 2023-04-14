package main

import (
	"context"
	"log"

	"github.com/AI1411/go-grpc-graphql/internal/env"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/infra/logger"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/user"
)

func main() {
	err := distributePointAllUsers()
	if err != nil {
		return
	}
}

func distributePointAllUsers() error {
	e, err := env.NewValue()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	zapLogger, _ := logger.NewLogger(e.Debug)
	zapLogger.Info("DistributePointAllUsers Start")

	dbClient := db.NewClient(&e.DB, zapLogger)

	userRepo := repository.NewUserRepository(dbClient, nil)
	userPointRepo := repository.NewUserPointRepository(dbClient)

	usecase := user.NewDistributePointAllUsersImpl(userRepo, userPointRepo)
	if err = usecase.Exec(context.Background(), ""); err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	zapLogger.Info("DistributePointAllUsers End")

	return nil
}
