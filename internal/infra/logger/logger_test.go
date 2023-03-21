package logger

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	testCases := []struct {
		name          string
		debug         bool
		expectedLevel zapcore.Level
	}{
		{
			name:          "Debug mode enabled",
			debug:         true,
			expectedLevel: zapcore.DebugLevel,
		},
		{
			name:          "Debug mode disabled",
			debug:         false,
			expectedLevel: zapcore.InfoLevel,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger, err := NewLogger(tc.debug)
			if err != nil {
				t.Fatalf("Unexpected error while creating logger: %v", err)
			}

			if logger == nil {
				t.Fatal("Expected logger to be non-nil")
			}

			core := logger.Core()
			if core == nil {
				t.Fatal("Expected core to be non-nil")
			}

			if core.Enabled(tc.expectedLevel) != true {
				t.Errorf("Expected logger to be enabled for level %v", tc.expectedLevel)
			}

			// Check if the level below the expected level is disabled
			levelBelow := tc.expectedLevel - 1
			if core.Enabled(levelBelow) != false {
				t.Errorf("Expected logger to be disabled for level %v", levelBelow)
			}
		})
	}
}
