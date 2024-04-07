package test

import (
	"io"
	"log/slog"
	"os"
	"testing"
    "github.com/sirupsen/logrus"
    "time"
    "gopkg.in/natefinch/lumberjack.v2"
    "path/filepath"
)

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
    file, _ := os.OpenFile("temp/log/output_test.log",os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_RDONLY,0644)
    h := slog.NewJSONHandler(io.MultiWriter(os.Stdout,file), nil)
    logger := slog.New(h)
    logger.Info("Hello, World!")
}

func TestLogrus(t *testing.T)  {
    log := logrus.New()
    file, _ := os.OpenFile("./temp/log/output_test.log",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
    writer := io.MultiWriter(os.Stdout,file)
    log.SetOutput(writer)
    log.SetFormatter(&logrus.JSONFormatter{})
    log.WithFields(logrus.Fields{
        "message" : "Hello World",
        "data" : []string{"hello","world"},
    }).Info("method GET")

}

func TestRotateLog(t *testing.T) {
    lumberjackLogger := &lumberjack.Logger{
        // Log file abbsolute path, os agnostic
        Filename:   filepath.ToSlash("./temp/log/rotate.log"), 
        MaxSize:    5, // MB
        MaxBackups: 10,
        MaxAge:     30,   // days
        Compress:   true, // disabled by default
    }

    // Fork writing into two outputs
    multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

    logFormatter := new(logrus.TextFormatter)
    logFormatter.TimestampFormat = time.RFC1123Z // or RFC3339
    logFormatter.FullTimestamp = true
   log := logrus.New()
   log.SetFormatter(logFormatter)
   log.SetLevel(logrus.InfoLevel)
   log.SetOutput(multiWriter)

   log.Info("Info ne masseh")
}