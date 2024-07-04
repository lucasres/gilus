package entities

import "time"

type PingCron struct {
	Name   string
	PingAt time.Time
}

func NewPingCron(n string, p time.Time) *PingCron {
	return &PingCron{
		Name:   n,
		PingAt: p,
	}
}
