package crons

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lucasres/gilus/internal/domain/entities"
)

type PingCronUseCase struct{}

func (p *PingCronUseCase) Execute(ctx context.Context, e entities.PingCron) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		return fmt.Errorf("erro when create session for aws: %w", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	av := make(map[string]types.AttributeValue)
	av["pingAt"] = &types.AttributeValueMemberS{Value: e.PingAt}
	av["name"] = &types.AttributeValueMemberS{Value: e.Name}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("cron_executions"),
		Item:      av,
	}

	_, err = svc.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("erro when put item in dynamodb: %w", err)
	}

	return nil
}

func NewPingCronUseCase() *PingCronUseCase {
	return &PingCronUseCase{}
}
