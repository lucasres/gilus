package crons

import (
	"context"
	"time"

	"github.com/lucasres/gilus/internal/domain/entities"
)

type ListPingCronUseCase struct{}

func (l *ListPingCronUseCase) Execute(ctx context.Context) ([]*entities.PingCron, error) {
	rs := make([]*entities.PingCron, 0)
	rs = append(rs, entities.NewPingCron("teste-1", time.Now()))
	rs = append(rs, entities.NewPingCron("teste-1", time.Now()))
	rs = append(rs, entities.NewPingCron("teste-1", time.Now()))
	rs = append(rs, entities.NewPingCron("teste-1", time.Now()))
	return rs, nil
}

func NewListPingCronUseCase() *ListPingCronUseCase {
	return &ListPingCronUseCase{}
}
