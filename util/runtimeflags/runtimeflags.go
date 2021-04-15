package runtimeflags

import (
	"context"

	"github.com/rakamoviz/logging-exp/util/contextkeys"
)

type RuntimeFlags struct {
	trace bool
}

func (rf RuntimeFlags) Trace() bool {
	return rf.trace
}

func New(config map[string]string) *RuntimeFlags {
	traceStr, ok := config["trace"]
	trace := false
	if ok {
		if traceStr == "true" {
			trace = true
		}
	}

	return &RuntimeFlags{
		trace: trace,
	}
}

func Get(ctx context.Context) *RuntimeFlags {
	runtimeFlagsCopy := RuntimeFlags{}

	runtimeFlags := ctx.Value(contextkeys.RuntimeFlags{})
	if runtimeFlags != nil {
		runtimeFlagsCopy.trace = runtimeFlags.(RuntimeFlags).trace
	}

	return &runtimeFlagsCopy
}

func BuildContext(ctx context.Context, runtimeFlags *RuntimeFlags) context.Context {
	runtimeFlagsCopy := RuntimeFlags{}

	runtimeFlagsCopy.trace = runtimeFlags.trace

	return context.WithValue(ctx, contextkeys.RuntimeFlags{}, runtimeFlagsCopy)
}
