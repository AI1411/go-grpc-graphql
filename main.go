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
	"github.com/AI1411/go-grpc-graphql/internal/infra/redis"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	redisRepository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/redis"
	roomRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
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
	redisClient := redis.NewRedisClient(&e.Redis)

	userRepo := repository.NewUserRepository(dbClient)
	tweetRepo := tweetRepo.NewTweetRepository(dbClient)
	chatRepo := chatRepo.NewChatRepository(dbClient)
	roomRepo := roomRepo.NewRoomRepository(dbClient)
	userPointRepo := repository.NewUserPointRepository(dbClient)
	redisRepo := redisRepository.NewRedisRepository(redisClient)

	s := grpc.NewServer(
		grpcMiddleware.WithUnaryServerChain(
			interceptor.ZapLoggerInterceptor(zapLogger),
		),
	)

	userServer := server.NewUserServer(dbClient, zapLogger, userRepo, redisRepo)
	tweetServer := server.NewTweetServer(dbClient, zapLogger, userRepo, tweetRepo)
	chatServer := server.NewChatServer(dbClient, zapLogger, userRepo, chatRepo)
	roomServer := server.NewRoomServer(dbClient, zapLogger, userRepo, roomRepo)
	userPointServer := server.NewUserPointServer(dbClient, zapLogger, userRepo, userPointRepo)

	grpcServer.RegisterUserServiceServer(s, userServer)
	grpcServer.RegisterTweetServiceServer(s, tweetServer)
	grpcServer.RegisterChatServiceServer(s, chatServer)
	grpcServer.RegisterRoomServiceServer(s, roomServer)
	grpcServer.RegisterUserPointServiceServer(s, userPointServer)

	zapLogger.Info("start grpc Server port: " + e.ServerPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
