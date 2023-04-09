package util

import (
	"errors"
	"testing"
	"time"
)

func TestIsNilOrEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"Nil", nil, true},
		{"Empty string", "", true},
		{"Non-empty string", "abc", false},
		{"Zero rune", rune(0), false},
		{"Non-zero rune", 'a', false},
		{"Zero int", 0, false},
		{"Non-zero int", 42, false},
		{"Zero float", 0.0, false},
		{"Non-zero float", 3.14, false},
		{"Nil error", error(nil), true},
		{"Non-nil error", errors.New("error"), false},
		{"Zero time", time.Time{}, true},
		{"Non-zero time", time.Now(), false},
		{"Nil slice", ([]int)(nil), true},
		{"Empty slice", []int{}, true},
		{"Non-empty slice", []int{1, 2, 3}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsNilOrEmpty(tc.input)
			if result != tc.expected {
				t.Errorf("IsNilOrEmpty(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
