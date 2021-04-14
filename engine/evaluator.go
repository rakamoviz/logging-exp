package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/log"
	"github.com/rakamoviz/logging-exp/util"
)

type Evaluator struct {
	calculator *Calculator
}

func (e *Evaluator) Evaluate(ctx context.Context, rule string) bool {
	if util.GetRuntimeFlags(ctx).Trace() {
		defer log.Trace(ctx, log.Entry(ctx, "engine.(*Evaluator).Evaluate", &map[string]interface{}{"rule": rule}))
	}

	log.G(ctx).Info("Logging in Evaluate")

	e.calculator.Calculate(ctx, 6)
	return true
}

func NewEvaluator(calculator *Calculator) *Evaluator {
	return &Evaluator{
		calculator: calculator,
	}
}
