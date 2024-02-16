package test

import (
	"go_expose/internal/log"
	"log/slog"
	"testing"
)

func TestLoggingLevelInfo(t *testing.T) {
	logging := slog.New(log.NewHandler(nil))
	logging.Info("Testing Logging Level Info")
	t.Log("finish")
}