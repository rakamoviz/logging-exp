package engine

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/log"
	"github.com/rakamoviz/logging-exp/util"
	"github.com/sirupsen/logrus"
)

func buildExecutionContext(trace string) context.Context {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	executionId := uuid.New()
	runtimeFlags := util.NewRuntimeFlags(map[string]string{
		"trace": trace,
	})

	baseContext := context.Background()
	baseContext = context.WithValue(
		baseContext, util.ExecutionIdKey{}, executionId,
	)
	baseContext = util.WithRuntimeFlags(baseContext, runtimeFlags)

	return log.WithLogger(
		baseContext,
		logrus.WithFields(logrus.Fields{"executionId": executionId}),
	)
}

func BenchmarkRunWithTrace(b *testing.B) {
	executionContext := buildExecutionContext("true")

	calculator := NewCalculator(4)
	evaluator := NewEvaluator(calculator)
	runner := NewRunner(evaluator)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		runner.Run(executionContext)
	}
}

func BenchmarkRunWithoutTrace(b *testing.B) {
	executionContext := buildExecutionContext("false")

	calculator := NewCalculator(4)
	evaluator := NewEvaluator(calculator)
	runner := NewRunner(evaluator)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		runner.Run(executionContext)
	}
}

func TestRun(t *testing.T) {

}
