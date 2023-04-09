package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewDeleteUserHobbyForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.DeleteUserHobbyRequest
		expected *DeleteUserHobbyForm
	}{
		{
			name: "Valid input",
			input: &grpc.DeleteUserHobbyRequest{
				UserId:  "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				HobbyId: "d6f1b3f8-6ebd-4d1b-80e9-2b4e846aa0a0",
			},
			expected: &DeleteUserHobbyForm{
				UserID:  "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				HobbyID: "d6f1b3f8-6ebd-4d1b-80e9-2b4e846aa0a0",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewDeleteUserHobbyForm(tc.input)
			if result.UserID != tc.expected.UserID || result.HobbyID != tc.expected.HobbyID {
				t.Errorf("NewDeleteUserHobbyForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
