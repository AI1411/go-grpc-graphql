package interceptor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// mockUnaryHandler is a dummy grpc.UnaryHandler function for testing.
func mockUnaryHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return "response", nil
}

func TestZapLoggerInterceptor(t *testing.T) {
	// Create a buffer to capture logs
	loggerBuffer, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("failed to create buffer logger: %v", err)
	}

	// Create the ZapLoggerInterceptor
	interceptor := ZapLoggerInterceptor(loggerBuffer)

	// Create a context and request
	ctx := context.Background()
	req := "request"

	// Create mock server info
	info := &grpc.UnaryServerInfo{
		FullMethod: "/api.ExampleService/TestMethod",
	}

	// Call the interceptor
	res, err := interceptor(ctx, req, info, mockUnaryHandler)

	// Check if the response is as expected
	expectedResponse := "response"
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, res)

	// Check if logs were generated
	logs := loggerBuffer.Check(zap.InfoLevel, "request")
	assert.NotNil(t, logs)

	logs = loggerBuffer.Check(zap.InfoLevel, "response")
	assert.NotNil(t, logs)
}
