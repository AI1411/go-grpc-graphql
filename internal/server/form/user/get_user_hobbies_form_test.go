package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewGetUserHobbiesForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.GetUserHobbiesRequest
		expected *GetUserHobbiesForm
	}{
		{
			name: "Valid input",
			input: &grpc.GetUserHobbiesRequest{
				UserId: "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
			},
			expected: &GetUserHobbiesForm{
				UserID: "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewGetUserHobbiesForm(tc.input)
			if result.UserID != tc.expected.UserID {
				t.Errorf("NewGetUserHobbiesForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
