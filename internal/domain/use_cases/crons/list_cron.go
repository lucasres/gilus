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

type ListCronUseCase struct{}

func (l *ListCronUseCase) Execute(ctx context.Context) ([]entities.Cron, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		return nil, fmt.Errorf("erro when create session for aws: %w", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	result, err := svc.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("crons"),
	})
	if err != nil {
		return nil, fmt.Errorf("erro when scan dynamodb: %w", err)
	}

	rs := make([]entities.Cron, 0)
	for _, v := range result.Items {
		var e entities.Cron
		err := attributevalue.UnmarshalMap(v, &e)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal item: %w", err)
		}
		rs = append(rs, e)
	}

	return rs, nil
}

func NewListCronUseCase() *ListCronUseCase {
	return &ListCronUseCase{}
}
