package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewUpdateUserProfileForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.UpdateUserProfileRequest
		expected *UpdateUserProfileForm
	}{
		{
			name: "Valid input",
			input: &grpc.UpdateUserProfileRequest{
				Id:           "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Username:     "testuser",
				Prefecture:   grpc.Prefecture_TOKYO,
				Introduction: "Hello, I'm a test user.",
				BloodType:    grpc.BloodType_A,
			},
			expected: &UpdateUserProfileForm{
				ID:           "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Username:     "testuser",
				Prefecture:   grpc.Prefecture_TOKYO,
				Introduction: "Hello, I'm a test user.",
				BloodType:    grpc.BloodType_A,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewUpdateUserProfileForm(tc.input)
			if result.ID != tc.expected.ID ||
				result.Username != tc.expected.Username ||
				result.Prefecture != tc.expected.Prefecture ||
				result.Introduction != tc.expected.Introduction ||
				result.BloodType != tc.expected.BloodType {
				t.Errorf("NewUpdateUserProfileForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
