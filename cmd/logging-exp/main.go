package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/internal/contextkeys"
	"github.com/rakamoviz/logging-exp/internal/engine"
	log "github.com/rakamoviz/loghelpr"
	"github.com/sirupsen/logrus"
)

func sampleLog() bool {
	return true
}

func main() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	executionId := uuid.New()

	baseContext := context.Background()
	baseContext = context.WithValue(
		baseContext, contextkeys.ExecutionId{}, executionId,
	)

	executionContext := log.BuildContext(
		baseContext,
		logger.WithFields(logrus.Fields{"executionId": executionId}),
		sampleLog(),
	)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	runner.Run(executionContext)
}
