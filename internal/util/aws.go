package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/AI1411/go-grpc-graphql/internal/env"
)

func NewAWSSession(e *env.Values) *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(e.AwsRegion),
		Endpoint:         aws.String("http://localhost:9000"),
		Credentials:      credentials.NewStaticCredentials("root", "password", ""),
		S3ForcePathStyle: aws.Bool(true),
	}))
}
