package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewCreateUserForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.CreateUserRequest
		expected *CreateUserForm
	}{
		{
			name: "Valid input",
			input: &grpc.CreateUserRequest{
				Username:     "TestUser",
				Email:        "test@example.com",
				Password:     "password123",
				Prefecture:   grpc.Prefecture_TOKYO,
				Introduction: "こんにちは、私はテストユーザーです。",
				BloodType:    grpc.BloodType_A,
			},
			expected: &CreateUserForm{
				Username:     "TestUser",
				Email:        "test@example.com",
				Password:     "password123",
				Prefecture:   grpc.Prefecture_TOKYO,
				Introduction: "こんにちは、私はテストユーザーです。",
				BloodType:    grpc.BloodType_A,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewCreateUserForm(tc.input)
			if result.Username != tc.expected.Username ||
				result.Email != tc.expected.Email ||
				result.Password != tc.expected.Password ||
				result.Prefecture != tc.expected.Prefecture ||
				result.Introduction != tc.expected.Introduction ||
				result.BloodType != tc.expected.BloodType {
				t.Errorf("NewCreateUserForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
