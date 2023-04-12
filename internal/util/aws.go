package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/AI1411/go-grpc-graphql/internal/env"
)

func NewAWSSession(e *env.Values) (*session.Session, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region:      &e.AwsRegion,
		Endpoint:    &e.AwsS3Endpoint,
		Credentials: credentials.NewStaticCredentials(e.AwsAccessKeyID, e.AwsSecretAccessKey, ""),
	})

	if err != nil {
		return nil, err
	}

	return awsSession, nil
}
