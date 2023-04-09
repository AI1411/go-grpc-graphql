package hobby

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewCreateHobbyForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.CreateHobbyRequest
		expected *CreateHobbyForm
	}{
		{
			name: "Valid input",
			input: &grpc.CreateHobbyRequest{
				Name:        "テスト趣味",
				Description: "これはテスト趣味です",
				CategoryId:  "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
			expected: &CreateHobbyForm{
				Name:        "テスト趣味",
				Description: "これはテスト趣味です",
				CategoryID:  "6e70a9c9-3b1d-49c5-91e1-684f885e5b16",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewCreateHobbyForm(tc.input)
			if result.Name != tc.expected.Name || result.Description != tc.expected.Description || result.CategoryID != tc.expected.CategoryID {
				t.Errorf("NewCreateHobbyForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
