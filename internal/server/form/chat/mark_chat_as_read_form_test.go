package chat

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewMarkChatAsReadForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.MarkChatAsReadRequest
		expected *MarkChatAsReadForm
	}{
		{
			name: "Valid input",
			input: &grpc.MarkChatAsReadRequest{
				Id: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
			expected: &MarkChatAsReadForm{
				ID: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewMarkChatAsReadForm(tc.input)
			if result.ID != tc.expected.ID {
				t.Errorf("NewMarkChatAsReadForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
