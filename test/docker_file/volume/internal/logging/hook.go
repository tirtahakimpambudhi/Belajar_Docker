package logging

import (
	"os"
	"path/filepath"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/sirupsen/logrus"
	"io"
)

type BugTrackerHooks struct {
	logpath string
}

func NewBugTrackerHooks() *BugTrackerHooks {
	return &BugTrackerHooks{
		logpath: "./temp/log/exception",
	}
}

func (hook *BugTrackerHooks) Levels() []logrus.Level {
	return []logrus.Level{
	 logrus.WarnLevel,
	 logrus.ErrorLevel,
	 logrus.FatalLevel,
	 logrus.PanicLevel,
	}
   }
   
   func (hook *BugTrackerHooks) Fire(entry *logrus.Entry) error {
	defaultOutput := entry.Logger.Out
	logFileName := ""
	switch entry.Level {
	case logrus.ErrorLevel:
		logFileName = "error.log"
	case logrus.FatalLevel:
		logFileName = "fatal.log"
	case logrus.PanicLevel:
		logFileName = "panic.log"
	default:
		logFileName = "warn.log"
	}
	// writer, err := os.OpenFile(filepath.Join(hook.logpath, logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil && errors.Is(err,os.ErrNotExist){
	// 	return fmt.Errorf("file %s not exist",filepath.Join(hook.logpath, logFileName))
	// }
	lumberjackLogger := &lumberjack.Logger{
        // Log file abbsolute path, os agnostic
        Filename:   filepath.ToSlash(filepath.Join(hook.logpath, logFileName)), 
        MaxSize:    5, // MB
        MaxBackups: 10,
        MaxAge:     30,   // days
        Compress:   true, // disabled by default
    }

    // Fork writing into two outputs
    multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)

	entry.Logger.SetOutput(multiWriter)
	defer entry.Logger.SetOutput(defaultOutput)
	entry.Logger.Printf("time='%v' level='%v' message='%s'\n", entry.Time, entry.Level, entry.Message)
	return nil
}