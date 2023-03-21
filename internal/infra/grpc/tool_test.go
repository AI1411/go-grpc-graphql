package grpc

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestToGRPCEnum(t *testing.T) {
	testCases := []struct {
		id               int
		name             string
		grpcConvertible  EnumConvertible
		expectedGRPCEnum int32
	}{
		{
			id:               1,
			name:             "通常会員",
			grpcConvertible:  entity.UserStatusValue["通常会員"],
			expectedGRPCEnum: int32(grpc.Status_ACTIVE),
		},
		{
			id:               2,
			name:             "退会済",
			grpcConvertible:  entity.UserStatusValue["退会済"],
			expectedGRPCEnum: int32(grpc.Status_RESIGNED),
		},
		{
			id:               3,
			name:             "アカウント停止",
			grpcConvertible:  entity.UserStatusValue["アカウント停止"],
			expectedGRPCEnum: int32(grpc.Status_BANDED),
		},
		{
			id:               4,
			name:             "プレミアム",
			grpcConvertible:  entity.UserStatusValue["プレミアム"],
			expectedGRPCEnum: int32(grpc.Status_PREMIUM),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToGRPCEnum(grpc.Status_value, tc.grpcConvertible)
			if result != tc.expectedGRPCEnum {
				t.Errorf("Expected %d, got %d", tc.expectedGRPCEnum, result)
			}
		})
	}
}
