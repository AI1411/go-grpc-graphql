package category

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewListCategoryForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.ListCategoryRequest
		expected *ListCategoryForm
	}{
		{
			name: "Valid input",
			input: &grpc.ListCategoryRequest{
				Name:   "TestCategory",
				Order:  "asc",
				Limit:  10,
				Offset: 0,
			},
			expected: &ListCategoryForm{
				Name:   "TestCategory",
				Order:  "asc",
				Limit:  10,
				Offset: 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewListCategoryForm(tc.input)
			if result.Name != tc.expected.Name || result.Order != tc.expected.Order || result.Limit != tc.expected.Limit || result.Offset != tc.expected.Offset {
				t.Errorf("NewListCategoryForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
