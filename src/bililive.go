package core

import (
	"github.com/hr3lxphr6j/bililive-go/src/configs"
	"github.com/hr3lxphr6j/bililive-go/src/lib/events"
	"github.com/hr3lxphr6j/bililive-go/src/listeners"
)

type Instance struct {
	Config          *configs.Config
	EventDispatcher events.IEventDispatcher
	ListenerManager listeners.IListenerManager
}
