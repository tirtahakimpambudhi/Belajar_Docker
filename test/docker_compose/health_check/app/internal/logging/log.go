package logging

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() *logrus.Logger {
	log := logrus.New()
	lumberjackLogger := &lumberjack.Logger{
        // Log file abbsolute path, os agnostic
        Filename:   filepath.ToSlash("./temp/log/app.log"), 
        MaxSize:    5, // MB
        MaxBackups: 10,
        MaxAge:     30,   // days
        Compress:   true, // disabled by default
    }
	log.SetOutput(io.MultiWriter(os.Stdout,lumberjackLogger))
	if os.Getenv("APP_ENV") == "production" {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	log.AddHook(NewBugTrackerHooks())
	return log
}