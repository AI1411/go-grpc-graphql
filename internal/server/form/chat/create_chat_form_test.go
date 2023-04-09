package chat

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewCreateChatForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.CreateChatRequest
		expected *CreateChatForm
	}{
		{
			name: "Valid input",
			input: &grpc.CreateChatRequest{
				RoomId:     "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
				FromUserId: "aa4f4a6d-4c77-4e2f-816c-7a8860f8a80d",
				ToUserId:   "0a74eb8f-7e21-4c23-8913-3bae8d7eb658",
				Body:       "Hello, how are you?",
			},
			expected: &CreateChatForm{
				roomID:     "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
				fromUserID: "aa4f4a6d-4c77-4e2f-816c-7a8860f8a80d",
				toUserID:   "0a74eb8f-7e21-4c23-8913-3bae8d7eb658",
				body:       "Hello, how are you?",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewCreateChatForm(tc.input)
			if result.roomID != tc.expected.roomID || result.fromUserID != tc.expected.fromUserID || result.toUserID != tc.expected.toUserID || result.body != tc.expected.body {
				t.Errorf("NewCreateChatForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
