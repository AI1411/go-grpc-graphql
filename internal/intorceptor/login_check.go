package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func AuthUnaryInterceptor(ctx context.Context) (context.Context, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing context metadata")
	}

	token := meta["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization token")
	}

	token[0] = removeBearerFromToken(token[0])

	_, err := util.ValidateJWT(token[0])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	userID, err := util.GetUserIDFromJWT(token[0])

	context.WithValue(ctx, "userID", userID)

	return ctx, nil
}

func removeBearerFromToken(header string) string {
	token := strings.TrimSpace(header)
	return strings.TrimPrefix(token, "Bearer ")
}
