package engine

import (
	"context"
	"fmt"

	"github.com/rakamoviz/logging-exp/log"
	"github.com/rakamoviz/logging-exp/util"
)

type Runner struct {
	evaluator *Evaluator
}

func (r *Runner) Run(ctx context.Context) {
	fmt.Println(util.GetRuntimeFlags(ctx))
	if util.GetRuntimeFlags(ctx).Trace() {
		defer log.Trace(ctx, log.Entry(ctx, "engine.(*Runner).Run", nil))
	}

	log.G(ctx).Info("Logging in Run")

	r.evaluator.Evaluate(ctx, "rule1")
}

func NewRunner(evaluator *Evaluator) *Runner {
	return &Runner{
		evaluator: evaluator,
	}
}
