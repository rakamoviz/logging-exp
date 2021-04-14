package util

import (
	"context"
)

type RuntimeFlags struct {
	trace bool
}

func (rf RuntimeFlags) Trace() bool {
	return rf.trace
}

func NewRuntimeFlags(config map[string]string) *RuntimeFlags {
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

func GetRuntimeFlags(ctx context.Context) *RuntimeFlags {
	runtimeFlagsCopy := RuntimeFlags{}

	runtimeFlags := ctx.Value(RuntimeFlagsKey{})
	if runtimeFlags != nil {
		runtimeFlagsCopy.trace = runtimeFlags.(RuntimeFlags).trace
	}

	return &runtimeFlagsCopy
}

func WithRuntimeFlags(ctx context.Context, runtimeFlags *RuntimeFlags) context.Context {
	runtimeFlagsCopy := RuntimeFlags{}

	runtimeFlagsCopy.trace = runtimeFlags.trace

	return context.WithValue(ctx, RuntimeFlagsKey{}, runtimeFlagsCopy)
}
