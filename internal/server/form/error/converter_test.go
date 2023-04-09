package error

import (
	"testing"
)

func TestConvertFieldErrorDescription(t *testing.T) {
	testCases := []struct {
		name           string
		defaultMessage string
		tag            string
		field          string
		param          string
		expected       string
	}{
		{
			name:           "Required error",
			defaultMessage: "Default message",
			tag:            "required",
			field:          "フィールド",
			param:          "",
			expected:       "フィールドは必須です",
		},
		{
			name:           "UUID4 error",
			defaultMessage: "Default message",
			tag:            "uuid4",
			field:          "フィールド",
			param:          "",
			expected:       "フィールドはUUID(v4)でなければなりません",
		},
		{
			name:           "Unknown error",
			defaultMessage: "Default message",
			tag:            "unknown",
			field:          "フィールド",
			param:          "",
			expected:       "Default message",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ConvertFieldErrorDescription(tc.defaultMessage, tc.tag, tc.field, tc.param)
			if result != tc.expected {
				t.Errorf("ConvertFieldErrorDescription(%v, %v, %v, %v) = %v; want %v", tc.defaultMessage, tc.tag, tc.field, tc.param, result, tc.expected)
			}
		})
	}
}
