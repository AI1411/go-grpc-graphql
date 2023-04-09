package chat

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewListChatForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.ListChatRequest
		expected *ListChatForm
	}{
		{
			name: "Valid input",
			input: &grpc.ListChatRequest{
				RoomId: "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
				UserId: "aa4f4a6d-4c77-4e2f-816c-7a8860f8a80d",
			},
			expected: &ListChatForm{
				RoomID:     "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
				FromUserID: "aa4f4a6d-4c77-4e2f-816c-7a8860f8a80d",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewListChatForm(tc.input)
			if result.RoomID != tc.expected.RoomID || result.FromUserID != tc.expected.FromUserID {
				t.Errorf("NewListChatForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
