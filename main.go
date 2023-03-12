package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	grpcServer "github.com/AI1411/go-grpc-praphql/grpc"
	"github.com/AI1411/go-grpc-praphql/internal/env"
	"github.com/AI1411/go-grpc-praphql/internal/infra/db"
	"github.com/AI1411/go-grpc-praphql/internal/infra/logger"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	interceptor "github.com/AI1411/go-grpc-praphql/internal/intorceptor"
	"github.com/AI1411/go-grpc-praphql/internal/server"
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

	zapLogger.Info("start grpc Server port: " + e.ServerPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
