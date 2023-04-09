package category

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewGetCategoryForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.GetCategoryRequest
		expected *GetCategoryForm
	}{
		{
			name: "Valid input",
			input: &grpc.GetCategoryRequest{
				Id: "12345",
			},
			expected: &GetCategoryForm{
				ID: "12345",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewGetCategoryForm(tc.input)
			if result.ID != tc.expected.ID {
				t.Errorf("NewGetCategoryForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
