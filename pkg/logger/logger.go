package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

type Config struct {
	Level logrus.Level
	Hook  logrus.Hook
}

func NewLogger(c *Config) {
	Logger = &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		ReportCaller: true,
		Level:        c.Level,
		ExitFunc:     os.Exit,
	}
	if c.Hook != nil {
		Logger.Hooks.Add(c.Hook)
	}
}
