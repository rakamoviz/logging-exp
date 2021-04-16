package engine

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rakamoviz/logging-exp/internal/contextkeys"
	"github.com/rakamoviz/logging-exp/internal/engine"
	log "github.com/rakamoviz/loghelpr"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
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

func TestRun(t *testing.T) {
	logger, hook := test.NewNullLogger()

	executionContext := buildExecutionContext(logger, true)

	calculator := engine.NewCalculator(4)
	evaluator := engine.NewEvaluator(calculator)
	runner := engine.NewRunner(evaluator)

	runner.Run(executionContext)

	assert.Equal(t, 9, len(hook.Entries))
}
