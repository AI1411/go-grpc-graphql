package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewLogoutForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.LogoutRequest
		expected *LogoutForm
	}{
		{
			name: "Valid input",
			input: &grpc.LogoutRequest{
				Token:  "example_token",
				UserId: "6d3e3e69-0c74-4fde-af49-9c44a1a4c4b4",
			},
			expected: &LogoutForm{
				Token:  "example_token",
				UserID: "6d3e3e69-0c74-4fde-af49-9c44a1a4c4b4",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewLogoutForm(tc.input)
			if result.Token != tc.expected.Token || result.UserID != tc.expected.UserID {
				t.Errorf("NewLogoutForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
