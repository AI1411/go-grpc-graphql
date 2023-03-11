package grpc

import (
	"reflect"
)

type EnumConvertible interface {
	String() string
}

func ToGRPCEnum(grpcEnumValue map[string]int32, grpcConvertible EnumConvertible) int32 {
	if reflect.TypeOf(grpcConvertible).Kind() == reflect.String {
		return grpcEnumValue[grpcConvertible.String()]
	}

	if grpcConvertible == nil || reflect.ValueOf(grpcConvertible).IsNil() {
		return 0
	}

	return grpcEnumValue[grpcConvertible.String()]
}
