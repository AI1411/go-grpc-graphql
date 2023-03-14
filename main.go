package main

import (
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	grpcServer "github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/env"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/infra/logger"
	tweetRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	interceptor "github.com/AI1411/go-grpc-graphql/internal/intorceptor"
	"github.com/AI1411/go-grpc-graphql/internal/server"
)

const defaultPort = "8081"

func main() {
	e, err := env.NewValue()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", e.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	zapLogger, _ := logger.NewLogger(e.Debug)

	dbClient, err := db.NewClient(&e.DB, zapLogger)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	userRepo := repository.NewUserRepository(dbClient)
	tweetRepo := tweetRepo.NewTweetRepository(dbClient)

	s := grpc.NewServer(
		grpcMiddleware.WithUnaryServerChain(
			interceptor.ZapLoggerInterceptor(zapLogger),
		),
	)

	userServer := server.NewUserServer(
		dbClient,
		zapLogger,
		userRepo,
	)

	tweetServer := server.NewTweetServer(dbClient, zapLogger, userRepo, tweetRepo)

	grpcServer.RegisterUserServiceServer(s, userServer)
	grpcServer.RegisterTweetServiceServer(s, tweetServer)

	zapLogger.Info("start grpc Server port: " + e.ServerPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
