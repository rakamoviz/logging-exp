package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/util/log"
	"github.com/rakamoviz/logging-exp/util/runtimeflags"
)

type Calculator struct {
	offset int
}

func (c *Calculator) Calculate(ctx context.Context, amount int) int {
	if runtimeflags.Get(ctx).Trace() {
		defer log.FnExit(ctx, log.FnEntrance(ctx, "engine.(*Calculator).Calculate", &map[string]interface{}{"amount": amount}))
	}

	log.G(ctx).Info("Logging inside a code block in the Calculate method")
	return c.offset + amount
}

func NewCalculator(offset int) *Calculator {
	return &Calculator{
		offset: offset,
	}
}
