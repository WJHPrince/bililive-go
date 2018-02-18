package core

import (
	"context"
	"github.com/hr3lxphr6j/bililive-go/src/configs"
	"github.com/hr3lxphr6j/bililive-go/src/lib/events"
	"github.com/hr3lxphr6j/bililive-go/src/listeners"
	"github.com/hr3lxphr6j/bililive-go/src/recorders"
)

type Instance struct {
	Config          *configs.Config
	EventDispatcher events.IEventDispatcher
	ListenerManager listeners.IListenerManager
	RecorderManager recorders.IRecorderManager
}

func NewInstance(config *configs.Config) (*Instance, error) {
	instance := new(Instance)
	ctx := context.WithValue(context.Background(), InstanceKey, instance)
	instance.Config = config
	instance.EventDispatcher = events.NewIEventDispatcher(ctx)
	instance.ListenerManager = listeners.NewIListenerManager(ctx)
	instance.RecorderManager = recorders.NewIRecorderManager(ctx)
	return instance, nil
}
