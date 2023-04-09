package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewUpdateUserStatusForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.UpdateUserStatusRequest
		expected *UpdateUserStatusForm
	}{
		{
			name: "Valid input",
			input: &grpc.UpdateUserStatusRequest{
				Id:     "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Status: grpc.Status_ACTIVE,
			},
			expected: &UpdateUserStatusForm{
				ID:     "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Status: grpc.Status_ACTIVE,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewUpdateUserStatusForm(tc.input)
			if result.ID != tc.expected.ID || result.Status != tc.expected.Status {
				t.Errorf("NewUpdateUserStatusForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
