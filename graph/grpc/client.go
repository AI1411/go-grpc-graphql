package grpc

import (
	"google.golang.org/grpc"

	grpcUser "github.com/AI1411/go-grpc-praphql/grpc"
)

func connect() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	opts = []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConnectUserServiceClient() (grpcUser.UserServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return grpcUser.NewUserServiceClient(conn), nil
}
