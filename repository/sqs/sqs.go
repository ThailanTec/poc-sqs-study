package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSRepository struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSRepository(queueURL string) (*SQSRepository, error) {
	// Load configuration with the correct endpoint and credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"), // Replace with your desired region
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("your_access_key", "your_secret_key", "your_session_token")),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == sqs.ServiceID && region == "us-east-1" {
					return aws.Endpoint{
						URL: "http://localhost:4566", // Endpoint of LocalStack
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			})),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := sqs.NewFromConfig(cfg)

	return &SQSRepository{
		client:   client,
		queueURL: queueURL,
	}, nil
}
