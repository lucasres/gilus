package crons

import (
	"context"

	"github.com/lucasres/gilus/internal/domain/entities"
)

type PingCronUseCase struct{}

func (p *PingCronUseCase) Execute(ctx context.Context, e *entities.PingCron) error {
	return nil
}

func NewPingCronUseCase() *PingCronUseCase {
	return &PingCronUseCase{}
}
