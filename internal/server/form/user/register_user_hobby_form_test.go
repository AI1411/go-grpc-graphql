package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewRegisterUserHobbyForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.RegisterUserHobbyRequest
		expected *RegisterUserHobbyForm
	}{
		{
			name: "Valid input",
			input: &grpc.RegisterUserHobbyRequest{
				UserId:  "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				HobbyId: "a5e5b5a9-8d3a-44e2-bf5d-8d3a4b7a0b1c",
			},
			expected: &RegisterUserHobbyForm{
				UserID:  "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				HobbyID: "a5e5b5a9-8d3a-44e2-bf5d-8d3a4b7a0b1c",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewRegisterUserHobbyForm(tc.input)
			if result.UserID != tc.expected.UserID || result.HobbyID != tc.expected.HobbyID {
				t.Errorf("NewRegisterUserHobbyForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
