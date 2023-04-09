package tweet

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewCreateTweetForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.CreateTweetRequest
		expected *CreateTweetForm
	}{
		{
			name: "Valid input",
			input: &grpc.CreateTweetRequest{
				Body:   "これはテストツイートです",
				UserId: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
			expected: &CreateTweetForm{
				Body:   "これはテストツイートです",
				UserID: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewCreateTweetForm(tc.input)
			if result.Body != tc.expected.Body || result.UserID != tc.expected.UserID {
				t.Errorf("NewCreateTweetForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
