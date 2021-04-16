package engine

import (
	"context"

	log "github.com/rakamoviz/loghelpr"
)

type Evaluator struct {
	calculator *Calculator
}

func (e *Evaluator) Evaluate(ctx context.Context, rule string) bool {
	defer log.LogFn(log.Fn(
		ctx,
		"engine.(*Evaluator).Evaluate",
		&map[string]interface{}{"rule": rule},
	)())()

	log.Info(ctx, "Logging inside a code block in the Evaluate method")()

	e.calculator.Calculate(ctx, 6)
	return true
}

func NewEvaluator(calculator *Calculator) *Evaluator {
	return &Evaluator{
		calculator: calculator,
	}
}
