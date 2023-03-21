package interceptor

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	grpcLogging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/AI1411/go-grpc-graphql/internal/util"
)

// ZapLoggerInterceptor returns a new unary server interceptor for logging the execution of the unary handler
func ZapLoggerInterceptor(zapLogger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		startTime := time.Now()

		traceID, _ := uuid.NewRandom()

		// trace_idでrequestとresponseを紐付ける
		loggerWithField := zapLogger.With(
			zap.String("trace_id", traceID.String()),
		)

		// ログに出力したくないサービスを指定
		ignoreServices := []string{""}

		serviceName := strings.Split(info.FullMethod, "/")[2]
		if util.Contains(ignoreServices, serviceName) {
			return nil, nil
		}

		buffReq, _ := json.Marshal(req)
		jsonReq := map[string]interface{}{}
		_ = json.Unmarshal(buffReq, &jsonReq)

		loggerWithField.Info(
			"request",
			zap.String("grpc.start_time", startTime.Format(time.RFC3339Nano)),
			zap.String("service_name", serviceName),
			zap.Any("req", jsonReq),
		)

		res, respErr := handler(ctx, req)
		buffRes, _ := json.Marshal(res)

		jsonRes := map[string]interface{}{}
		_ = json.Unmarshal(buffRes, &jsonRes)

		// エラーが発生した時もresponseとerrを表示したいので、この行でdefer実行
		defer loggerWithField.Info(
			"response",
			zap.String("grpc.end_time", time.Now().Format(time.RFC3339Nano)),
			zap.Int64("grpc.duration_nano", time.Since(startTime).Nanoseconds()),
			zap.String("grpc.code", grpcLogging.DefaultErrorToCode(respErr).String()),
			zap.Any("response", jsonRes),
			zap.String("service_name", serviceName),
			zap.Error(respErr),
		)

		if respErr != nil {
			return nil, respErr
		}
		return res, nil
	}
}
