package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewUpdateUserPasswordForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.UpdateUserPasswordRequest
		expected *UpdateUserPasswordForm
	}{
		{
			name: "Valid input",
			input: &grpc.UpdateUserPasswordRequest{
				Id:                   "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			expected: &UpdateUserPasswordForm{
				ID:                   "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewUpdateUserPasswordForm(tc.input)
			if result.ID != tc.expected.ID ||
				result.Password != tc.expected.Password ||
				result.PasswordConfirmation != tc.expected.PasswordConfirmation {
				t.Errorf("NewUpdateUserPasswordForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
