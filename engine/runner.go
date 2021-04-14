package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/log"
)

type Runner struct {
	evaluator *Evaluator
}

func (r *Runner) Run(ctx context.Context) {
	defer log.Trace(ctx, log.Entry(ctx, "engine.(*Runner).Run", nil))

	log.G(ctx).Info("Logging in Run")

	r.evaluator.Evaluate(ctx, "rule1")
}

func NewRunner(evaluator *Evaluator) *Runner {
	return &Runner{
		evaluator: evaluator,
	}
}
