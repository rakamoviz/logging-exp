package engine

import (
	"context"

	log "github.com/rakamoviz/loghelpr"
)

type Runner struct {
	evaluator *Evaluator
}

func (r *Runner) Run(ctx context.Context) {
	defer log.Fn(
		ctx,
		"engine.(*Runner).Run",
		nil,
	)()

	//log.Info(ctx, "Logging inside a code block in the Run method")()

	r.evaluator.Evaluate(ctx, "rule1")
}

func NewRunner(evaluator *Evaluator) *Runner {
	return &Runner{
		evaluator: evaluator,
	}
}
