package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcClient "github.com/AI1411/go-grpc-graphql/grpc"
)

func connect() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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

func ConnectChatServiceClient() (grpcClient.ChatServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return grpcClient.NewChatServiceClient(conn), nil
}

func ConnectRoomServiceClient() (grpcClient.RoomServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return grpcClient.NewRoomServiceClient(conn), nil
}
