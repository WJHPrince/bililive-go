package log

import (
	"context"
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"github.com/hr3lxphr6j/bililive-go/src/instance"
	"github.com/hr3lxphr6j/bililive-go/src/interfaces"
	"github.com/hr3lxphr6j/bililive-go/src/lib/events"
	"github.com/hr3lxphr6j/bililive-go/src/listeners"
	"github.com/hr3lxphr6j/bililive-go/src/recorders"
	"github.com/sirupsen/logrus"
	"os"
)

func info2Fields(info *api.Info) logrus.Fields {
	return logrus.Fields{
		"Url":      info.Url,
		"HostName": info.HostName,
		"RoomName": info.RoomName,
	}
}

func registerEventLog(ed events.IEventDispatcher, logger *interfaces.Logger) {

	ed.AddEventListener(listeners.ListenStart, events.NewEventListener(func(event *events.Event) {
		logger.WithFields(info2Fields(event.Object.(*api.Info))).Info(event.Type)
	}))

	ed.AddEventListener(listeners.LiveStart, events.NewEventListener(func(event *events.Event) {
		logger.WithFields(info2Fields(event.Object.(*api.Info))).Info(event.Type)
	}))

	ed.AddEventListener(listeners.LiveEnd, events.NewEventListener(func(event *events.Event) {
		logger.WithFields(info2Fields(event.Object.(*api.Info))).Info(event.Type)
	}))

	ed.AddEventListener(recorders.RecordeStart, events.NewEventListener(func(event *events.Event) {
		logger.WithFields(info2Fields(event.Object.(*api.Info))).Info(event.Type)
	}))

	ed.AddEventListener(recorders.RecordeStop, events.NewEventListener(func(event *events.Event) {
		logger.WithFields(info2Fields(event.Object.(*api.Info))).Info(event.Type)
	}))

}

func NewLogger(ctx context.Context) *interfaces.Logger {
	inst := instance.GetInstance(ctx)

	logLevel := logrus.InfoLevel
	switch inst.Config.LogLevel {
	case "panic":
		logLevel = logrus.PanicLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "debug":
		logLevel = logrus.DebugLevel
	default:
		logLevel = logrus.InfoLevel
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

	registerEventLog(inst.EventDispatcher.(events.IEventDispatcher), logger)

	return logger
}
