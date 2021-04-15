package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/util/log"
)

type Calculator struct {
	offset int
}

func (c *Calculator) Calculate(ctx context.Context, amount int) int {
	defer log.LogFn(log.Fn(
		ctx,
		"engine.(*Calculator).Calculate",
		&map[string]interface{}{"amount": amount},
	)())()

	log.Info(ctx, "Logging inside a code block in the Calculate method")()
	return c.offset + amount
}

func NewCalculator(offset int) *Calculator {
	return &Calculator{
		offset: offset,
	}
}
