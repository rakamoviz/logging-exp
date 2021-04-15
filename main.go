package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/engine"
	"github.com/rakamoviz/logging-exp/util/contextkeys"
	"github.com/rakamoviz/logging-exp/util/log"
	"github.com/rakamoviz/logging-exp/util/runtimeflags"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	executionId := uuid.New()
	runtimeFlags := runtimeflags.New(map[string]string{
		"trace": "true",
	})

	baseContext := context.Background()
	baseContext = context.WithValue(
		baseContext, contextkeys.ExecutionId{}, executionId,
	)
	baseContext = runtimeflags.BuildContext(baseContext, runtimeFlags)

	executionContext := log.BuildContext(
		baseContext,
		logger.WithFields(logrus.Fields{"executionId": executionId}),
	)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	runner.Run(executionContext)
}
