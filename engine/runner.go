package engine

import (
	"context"

	"github.com/rakamoviz/logging-exp/util/log"
	"github.com/rakamoviz/logging-exp/util/runtimeflags"
)

type Runner struct {
	evaluator *Evaluator
}

func (r *Runner) Run(ctx context.Context) {
	if runtimeflags.Get(ctx).Trace() {
		defer log.FnExit(ctx, log.FnEntrance(ctx, "engine.(*Runner).Run", nil))
	}

	log.G(ctx).Info("Logging inside a code block in the Run method")

	r.evaluator.Evaluate(ctx, "rule1")
}

func NewRunner(evaluator *Evaluator) *Runner {
	return &Runner{
		evaluator: evaluator,
	}
}
