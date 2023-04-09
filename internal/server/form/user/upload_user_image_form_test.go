package user

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewUploadUserImageForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.UploadUserImageRequest
		expected *UploadUserImageForm
	}{
		{
			name: "Valid input",
			input: &grpc.UploadUserImageRequest{
				UserId: "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Image:  "base64-encoded-image",
			},
			expected: &UploadUserImageForm{
				UserId: "e2a0b779-91c0-4f84-9d58-ef5ea5e5b5a9",
				Image:  "base64-encoded-image",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewUploadUserImageForm(tc.input)
			if result.UserId != tc.expected.UserId || result.Image != tc.expected.Image {
				t.Errorf("NewUploadUserImageForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
