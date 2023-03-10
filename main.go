package main

import (
	"fmt"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcServer "github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/env"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/server"
)

func main() {
	e, err := env.NewValue()
	if err != nil {
		fmt.Println(err.Error())
		panic("Error loading .env file")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", e.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dbClient, err := db.NewClient(&e.DB, &zap.Logger{})
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	userRepo := repository.NewUserRepository(dbClient)

	s := grpc.NewServer()
	userServer := server.NewUserServer(
		dbClient,
		&zap.Logger{},
		userRepo,
	)
	grpcServer.RegisterUserServiceServer(s, userServer)

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("lister to sever port %s", e.Port)
}
