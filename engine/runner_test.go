package engine

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/log"
	"github.com/rakamoviz/logging-exp/util"
	"github.com/sirupsen/logrus"

	"github.com/sirupsen/logrus/hooks/test"
	//"github.com/stretchr/testify/assert"
)

func buildExecutionContext(runtimeConfig map[string]string, logger *logrus.Logger) context.Context {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	executionId := uuid.New()
	runtimeFlags := util.NewRuntimeFlags(runtimeConfig)

	baseContext := context.Background()
	baseContext = context.WithValue(
		baseContext, util.ExecutionIdKey{}, executionId,
	)
	baseContext = util.WithRuntimeFlags(baseContext, runtimeFlags)

	return log.WithLogger(
		baseContext,
		logger.WithFields(logrus.Fields{"executionId": executionId}),
	)
}

func BenchmarkRunWithTrace(b *testing.B) {
	executionContext := buildExecutionContext(map[string]string{
		"trace": "true",
	}, logrus.New())

	calculator := NewCalculator(4)
	evaluator := NewEvaluator(calculator)
	runner := NewRunner(evaluator)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		runner.Run(executionContext)
	}
}

func BenchmarkRunWithoutTrace(b *testing.B) {
	executionContext := buildExecutionContext(map[string]string{
		"trace": "false",
	}, logrus.New())

	calculator := NewCalculator(4)
	evaluator := NewEvaluator(calculator)
	runner := NewRunner(evaluator)

	// run the Run function b.N times
	for n := 0; n < b.N; n++ {
		runner.Run(executionContext)
	}
}

func TestRun(t *testing.T) {
	logger, hook := test.NewNullLogger()

	executionContext := buildExecutionContext(map[string]string{
		"trace": "true",
	}, logger)

	calculator := NewCalculator(4)
	evaluator := NewEvaluator(calculator)
	runner := NewRunner(evaluator)

	runner.Run(executionContext)

	fmt.Println("==============> MYLOG", len(hook.Entries), hook.Entries[0].Data)
}
