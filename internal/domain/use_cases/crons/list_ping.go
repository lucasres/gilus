package crons

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/lucasres/gilus/internal/domain/entities"
)

type ListPingCronUseCase struct{}

func (l *ListPingCronUseCase) Execute(ctx context.Context) ([]entities.PingCron, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		return nil, fmt.Errorf("erro when create session for aws: %w", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	result, err := svc.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("cron_executions"),
	})
	if err != nil {
		return nil, fmt.Errorf("erro when scan dynamodb: %w", err)
	}

	rs := make([]entities.PingCron, 0)
	for _, v := range result.Items {
		var e entities.PingCron
		err := attributevalue.UnmarshalMap(v, &e)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal item: %w", err)
		}
		rs = append(rs, e)
	}
	// err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &rs)
	// if err != nil {
	// 	return nil, fmt.Errorf("erro when decode data from dynamodb: %w", err)
	// }

	return rs, nil
}

func NewListPingCronUseCase() *ListPingCronUseCase {
	return &ListPingCronUseCase{}
}
