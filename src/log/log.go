package log

import (
	"context"
	"github.com/hr3lxphr6j/bililive-go/src/instance"
	"github.com/hr3lxphr6j/bililive-go/src/interfaces"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(ctx context.Context) *interfaces.Logger {
	inst := instance.GetInstance(ctx)
	logLevel := logrus.InfoLevel
	if inst.Config.Debug {
		logLevel = logrus.DebugLevel
	}
	logger := &interfaces.Logger{Logger: &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Hooks: make(logrus.LevelHooks),
		Level: logLevel,
	}}

	inst.Logger = logger

	return logger
}
