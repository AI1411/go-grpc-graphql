package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/AI1411/go-grpc-graphql/grpc/grpcconnect"
	"github.com/AI1411/go-grpc-graphql/internal/env"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/infra/logger"
	chatRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/chat"
	roomRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/room"
	tweetRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/tweet"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/server"
)

const defaultPort = "8081"

func main() {
	e, err := env.NewValue()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	zapLogger, _ := logger.NewLogger(e.Debug)

	dbClient, err := db.NewClient(&e.DB, zapLogger)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	userRepo := repository.NewUserRepository(dbClient)
	tweetRepo := tweetRepo.NewTweetRepository(dbClient)
	chatRepo := chatRepo.NewChatRepository(dbClient)
	roomRepo := roomRepo.NewRoomRepository(dbClient)
	userPointRepo := repository.NewUserPointRepository(dbClient)

	userServer := server.NewUserServer(dbClient, zapLogger, userRepo)
	tweetServer := server.NewTweetServer(dbClient, zapLogger, userRepo, tweetRepo)
	chatServer := server.NewChatServer(dbClient, zapLogger, userRepo, chatRepo)
	roomServer := server.NewRoomServer(dbClient, zapLogger, userRepo, roomRepo)
	userPointServer := server.NewUserPointServer(dbClient, zapLogger, userRepo, userPointRepo)

	mux := http.NewServeMux()

	mux.Handle(grpcconnect.NewChatServiceHandler(chatServer))
	mux.Handle(grpcconnect.NewUserServiceHandler(userServer))
	mux.Handle(grpcconnect.NewTweetServiceHandler(tweetServer))
	mux.Handle(grpcconnect.NewRoomServiceHandler(roomServer))
	mux.Handle(grpcconnect.NewUserPointServiceHandler(userPointServer))

	zapLogger.Info("start grpc Server port: " + e.ServerPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", e.ServerPort), h2c.NewHandler(mux, &http2.Server{})); err != nil {
		zapLogger.Fatal("failed to serve")
	}
}
