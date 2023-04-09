package category

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewDeleteCategoryForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.DeleteCategoryRequest
		expected *DeleteCategoryForm
	}{
		{
			name: "Valid input",
			input: &grpc.DeleteCategoryRequest{
				Id: "12345",
			},
			expected: &DeleteCategoryForm{
				ID: "12345",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewDeleteCategoryForm(tc.input)
			if result.ID != tc.expected.ID {
				t.Errorf("NewDeleteCategoryForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
