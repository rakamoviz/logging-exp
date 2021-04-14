package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/engine"
	"github.com/rakamoviz/logging-exp/log"

	//"github.com/rakamoviz/logging-exp/util"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	executionId := uuid.New()

	executionContext := log.WithLogger(
		context.WithValue(context.Background(), "executionId", executionId),
		logrus.WithFields(logrus.Fields{"executionId": executionId}),
	)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	runner.Run(executionContext)
}
