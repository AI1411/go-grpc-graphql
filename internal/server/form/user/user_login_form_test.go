package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewLoginForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.LoginRequest
		expected *LoginForm
	}{
		{
			name: "Valid input",
			input: &grpc.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			expected: &LoginForm{
				Email:    "test@example.com",
				Password: "password123",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewLoginForm(tc.input)
			if result.Email != tc.expected.Email || result.Password != tc.expected.Password {
				t.Errorf("NewLoginForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
