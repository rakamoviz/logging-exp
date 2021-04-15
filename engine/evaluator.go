package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/util/log"
	"github.com/rakamoviz/logging-exp/util/runtimeflags"
)

type Evaluator struct {
	calculator *Calculator
}

func (e *Evaluator) Evaluate(ctx context.Context, rule string) bool {
	if runtimeflags.Get(ctx).Trace() {
		defer log.FnExit(ctx, log.FnEntrance(ctx, "engine.(*Evaluator).Evaluate", &map[string]interface{}{"rule": rule}))
	}

	log.G(ctx).Info("Logging inside a code block in the Evaluate method")

	e.calculator.Calculate(ctx, 6)
	return true
}

func NewEvaluator(calculator *Calculator) *Evaluator {
	return &Evaluator{
		calculator: calculator,
	}
}
