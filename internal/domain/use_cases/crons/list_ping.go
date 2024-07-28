package crons

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/lucasres/gilus/internal/domain/entities"
)

type ListPingCronUseCase struct{}

func (l *ListPingCronUseCase) Execute(ctx context.Context, name string) ([]entities.PingCron, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		return nil, fmt.Errorf("erro when create session for aws: %w", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	result, err := svc.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String("cron_executions"),
		IndexName:              aws.String("name-index"),
		Limit:                  aws.Int32(10),
		KeyConditionExpression: aws.String("#name = :name"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: name},
		},
		ExpressionAttributeNames: map[string]string{
			"#name": "name",
		},
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

	return rs, nil
}

func NewListPingCronUseCase() *ListPingCronUseCase {
	return &ListPingCronUseCase{}
}
