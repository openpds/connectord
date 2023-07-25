package connector

import "context"

type CheckHealthOutput struct{}

type HealthChecker interface {
	CheckHealth(context.Context) (*CheckHealthOutput, error)
}

type HealthCheckerFunc func(context.Context) (*CheckHealthOutput, error)

func (h HealthCheckerFunc) CheckHelth(ctx context.Context) (*CheckHealthOutput, error) {
	return h(ctx)
}
