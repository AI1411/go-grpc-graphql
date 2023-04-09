package category

import (
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

func TestNewCreateCategoryForm(t *testing.T) {
	testCases := []struct {
		name     string
		input    *grpc.CreateCategoryRequest
		expected *CreateCategoryForm
	}{
		{
			name: "Valid input",
			input: &grpc.CreateCategoryRequest{
				Name:        "TestCategory",
				Description: "TestCategoryDescription",
			},
			expected: &CreateCategoryForm{
				Name:        "TestCategory",
				Description: "TestCategoryDescription",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewCreateCategoryForm(tc.input)
			if result.Name != tc.expected.Name || result.Description != tc.expected.Description {
				t.Errorf("NewCreateCategoryForm(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
