package form

import (
	"errors"
	"testing"

	"github.com/AI1411/go-grpc-graphql/grpc"
)

type sampleForm struct {
	Prefecture grpc.Prefecture `validate:"required,prefecture"`
	BloodType  grpc.BloodType  `validate:"required,bloodType"`
}

func TestFormValidator_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		input    *sampleForm
		expected error
	}{
		{
			name: "Valid input",
			input: &sampleForm{
				Prefecture: grpc.Prefecture_AICHI,
				BloodType:  grpc.BloodType_A,
			},
			expected: nil,
		},
		{
			name: "Invalid input",
			input: &sampleForm{
				Prefecture: grpc.Prefecture(-1),
				BloodType:  grpc.BloodType(-1),
			},
			expected: errors.New("invalid input"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewFormValidator(tc.input)
			err := validator.Validate()

			if tc.expected == nil && err != nil {
				t.Errorf("Expected no error, got error: %v", err)
			} else if tc.expected != nil && err == nil {
				t.Errorf("Expected an error, got no error")
			}
		})
	}
}
