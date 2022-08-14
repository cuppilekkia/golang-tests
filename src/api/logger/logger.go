package logger

import (
	"fmt"
	"microservices-course/src/api/config"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: logLevel,
		Out:   os.Stdout,
	}

	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}
}

func Info(message string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}

	Log.WithFields(parseFields(tags)).Info(message)
}

func Error(message string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg := fmt.Sprintf("%s - ERROR - %v", message, err)
	Log.WithFields(parseFields(tags)).Error(msg)
}

func parseFields(tags []string) logrus.Fields {
	fields := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		kv := strings.Split(tag, ":")
		fields[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}
	return fields
}
