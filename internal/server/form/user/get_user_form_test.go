package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewGetUserForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.GetUserRequest
		expected *GetUserForm
	}{
		{
			name: "Valid input",
			input: &grpc.GetUserRequest{
				Id: "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
			},
			expected: &GetUserForm{
				ID: "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewGetUserForm(tc.input)
			if result.ID != tc.expected.ID {
				t.Errorf("NewGetUserForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
