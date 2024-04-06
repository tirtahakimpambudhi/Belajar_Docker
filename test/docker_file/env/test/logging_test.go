package test

import (
	"go_env/internal/log"
	"log/slog"
	"testing"
	"os"
)

func TestLoggingLevelInfo(t *testing.T) {
	logging := slog.New(log.NewHandler(nil))
	logging.Info("Testing Logging Level Info")
	t.Log("finish")
}

type TeeWriter struct {
    stdout *os.File
    file   *os.File
}

func (t *TeeWriter) Write(p []byte) (n int, err error) {
    n, err = t.stdout.Write(p)
    if err != nil {
        return n, err
    }
    n, err = t.file.Write(p)
    return n, err
}

func TestLogFile(t *testing.T) {
    file, _ := os.Create("output.txt")
    writer := &TeeWriter{
        stdout: os.Stdout,
        file:   file,
    }
    h := slog.NewJSONHandler(writer, nil)
    logger := slog.New(h)
    logger.Info("Hello, World!")
}