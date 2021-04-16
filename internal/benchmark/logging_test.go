package benchmark

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/internal/contextkeys"
	"github.com/rakamoviz/logging-exp/internal/engine"
	log "github.com/rakamoviz/loghelpr"
	"github.com/sirupsen/logrus"
)

func buildExecutionContext(logger *logrus.Logger, sampleLog bool) context.Context {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	executionId := uuid.New()

	baseContext := context.Background()
	baseContext = context.WithValue(
		baseContext, contextkeys.ExecutionId{}, executionId,
	)

	return log.BuildContext(
		baseContext,
		logger.WithFields(logrus.Fields{"executionId": executionId}),
		sampleLog,
	)
}

func BenchmarkRunWithLogging(b *testing.B) {
	executionContext := buildExecutionContext(logrus.New(), true)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		runner.Run(executionContext)
	}
}

func BenchmarkRunWithoutLogging(b *testing.B) {
	executionContext := buildExecutionContext(logrus.New(), false)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		runner.Run(executionContext)
	}
}

func BenchmarkCalculateWithLogging(b *testing.B) {
	executionContext := buildExecutionContext(logrus.New(), true)

	calculator := engine.NewCalculator(4)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		calculator.Calculate(executionContext, 2)
	}
}

func BenchmarkCalculateWithoutLogging(b *testing.B) {
	executionContext := buildExecutionContext(logrus.New(), false)

	calculator := engine.NewCalculator(4)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		calculator.Calculate(executionContext, 2)
	}
}
