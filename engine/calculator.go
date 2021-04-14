package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/log"
)

type Calculator struct {
	offset int
}

func (c *Calculator) Calculate(ctx context.Context, amount int) int {
	defer log.Trace(ctx, log.Entry(ctx, "engine.(*Calculator).Calculate", &map[string]interface{}{"amount": amount}))

	log.G(ctx).Info("Logging in Calculate")
	return c.offset + amount
}

func NewCalculator(offset int) *Calculator {
	return &Calculator{
		offset: offset,
	}
}
