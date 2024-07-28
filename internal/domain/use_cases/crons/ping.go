package crons

import (
	"context"
	"fmt"
	"time"

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

	e.PingAt = time.Now().Format("2006-01-02 15:04") + ":00"

	av := make(map[string]types.AttributeValue)
	av["pingAt"] = &types.AttributeValueMemberS{Value: e.PingAt}
	av["name"] = &types.AttributeValueMemberS{Value: e.Name}

	_, err = svc.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("cron_executions"),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("erro when put item in ping dynamodb: %w", err)
	}

	has, err := svc.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String("crons"),
		Limit:                  aws.Int32(1),
		KeyConditionExpression: aws.String("#name = :name"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: e.Name},
		},
		ExpressionAttributeNames: map[string]string{
			"#name": "name",
		},
	})
	if err != nil {
		return fmt.Errorf("erro when check if has cron name saved: %w", err)
	}

	fmt.Println("total")
	fmt.Println(len(has.Items))

	if len(has.Items) == 0 {
		av := make(map[string]types.AttributeValue)
		av["name"] = &types.AttributeValueMemberS{Value: e.Name}
		av["createdAt"] = &types.AttributeValueMemberS{Value: time.Now().Format("2006-01-02 15:04:05")}

		_, err = svc.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String("crons"),
			Item:      av,
		})
		if err != nil {
			return fmt.Errorf("erro when put item in cron names dynamodb: %w", err)
		}
	}

	return nil
}

func NewPingCronUseCase() *PingCronUseCase {
	return &PingCronUseCase{}
}
