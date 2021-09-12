package logging

import (
	"testing"

	"go.uber.org/zap"
)

func TestGetLogger(t *testing.T) {
	logger, err := zap.NewProduction()
	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	logger.Info("successfully initialized a *zap.Logger!")
}
