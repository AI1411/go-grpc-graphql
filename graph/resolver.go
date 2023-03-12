package graph

import "github.com/AI1411/go-grpc-praphql/grpc"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserClient grpc.UserServiceClient
}
