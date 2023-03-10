package interceptor

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	grpcLogging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
		if contains(ignoreServices, serviceName) {
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

// Contains from: [Goで型に縛られない良い感じのContains関数の書き方](https://zenn.dev/glassonion1/articles/7c7830a269909c)
func contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()
			// 型変換可能か確認する
			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}
			// 型変換する
			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
			// 等価判定をする
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}
