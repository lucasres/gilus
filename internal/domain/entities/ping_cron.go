package entities

type PingCron struct {
	PingAt string `json:"pingAt"`
	Name   string `json:"name"`
}

func NewPingCron(n, p string) *PingCron {
	return &PingCron{
		Name:   n,
		PingAt: p,
	}
}
