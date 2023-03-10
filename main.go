package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcServer "github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/env"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	"github.com/AI1411/go-grpc-praphql/internal/infra/logger"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	interceptor "github.com/AI1411/go-grpc-praphql/internal/intorceptor"
	"github.com/AI1411/go-grpc-praphql/internal/server"
)

func main() {
	e, err := env.NewValue()
	if err != nil {
		panic("Error loading .env file")
	}

	zapLogger, _ := logger.NewLogger(e.Debug)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", e.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dbClient, err := db.NewClient(&e.DB, zapLogger)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	userRepo := repository.NewUserRepository(dbClient)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptor.ZapLoggerInterceptor(zapLogger),
		),
	)

	userServer := server.NewUserServer(
		dbClient,
		zapLogger,
		userRepo,
	)
	grpcServer.RegisterUserServiceServer(s, userServer)

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("lister to sever port %s", e.Port)
}
