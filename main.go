package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/engine"
	"github.com/rakamoviz/logging-exp/log"
	"github.com/rakamoviz/logging-exp/util"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	executionId := uuid.New()
	runtimeFlags := util.NewRuntimeFlags(map[string]string{
		"trace": "true",
	})

	baseContext := context.Background()
	baseContext = context.WithValue(
		baseContext, util.ExecutionIdKey{}, executionId,
	)
	baseContext = util.WithRuntimeFlags(baseContext, runtimeFlags)

	executionContext := log.WithLogger(
		baseContext,
		logger.WithFields(logrus.Fields{"executionId": executionId}),
	)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	runner.Run(executionContext)
}
