package resolver

import "github.com/AI1411/go-grpc-graphql/grpc"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserClient  grpc.UserServiceClient
	TweetClient grpc.TweetServiceClient
	ChatClient  grpc.ChatServiceClient
	RoomClient  grpc.RoomServiceClient
}
