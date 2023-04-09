package hobby

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewDeleteHobbyForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.DeleteHobbyRequest
		expected *DeleteHobbyForm
	}{
		{
			name: "Valid input",
			input: &grpc.DeleteHobbyRequest{
				Id: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
			expected: &DeleteHobbyForm{
				ID: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewDeleteHobbyForm(tc.input)
			if result.ID != tc.expected.ID {
				t.Errorf("NewDeleteHobbyForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
