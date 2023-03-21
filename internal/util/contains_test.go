package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		list     interface{}
		elem     interface{}
		expected bool
	}{
		{
			name:     "Valid slice contains element",
			list:     []int{1, 2, 3, 4, 5},
			elem:     3,
			expected: true,
		},
		{
			name:     "Valid slice does not contain element",
			list:     []int{1, 2, 3, 4, 5},
			elem:     7,
			expected: false,
		},
		{
			name:     "Empty slice does not contain element",
			list:     []int{},
			elem:     3,
			expected: false,
		},
		{
			name:     "Non-slice input returns false",
			list:     42,
			elem:     3,
			expected: false,
		},
		{
			name:     "Mismatched element type returns false",
			list:     []int{1, 2, 3, 4, 5},
			elem:     "3",
			expected: false,
		},
		{
			name:     "Valid slice of strings contains element",
			list:     []string{"apple", "banana", "cherry"},
			elem:     "banana",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Contains(tc.list, tc.elem)
			assert.Equal(t, tc.expected, result, "Unexpected result for test case: %s", tc.name)
		})
	}
}
