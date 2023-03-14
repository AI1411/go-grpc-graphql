package grpc

import (
	"google.golang.org/grpc"

	grpcClient "github.com/AI1411/go-grpc-graphql/grpc"
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

func ConnectUserServiceClient() (grpcClient.UserServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return grpcClient.NewUserServiceClient(conn), nil
}

func ConnectTweetServiceClient() (grpcClient.TweetServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return grpcClient.NewTweetServiceClient(conn), nil
}
