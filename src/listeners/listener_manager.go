package listeners

import (
	"context"
	"github.com/hr3lxphr6j/bililive-go/src"
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"sync"
	"time"
)

type IListenerManager interface {
	AddListener(live api.Live) error
	RemoveListener(live api.Live) error
	HasListener(live api.Live) bool
}

func NewIListenerManager(ctx context.Context) IListenerManager {
	lm := new(ListenerManager)
	lm.ctx = ctx
	return lm
}

type ListenerManager struct {
	ctx    context.Context
	savers map[api.Live]*Listener
	lock   sync.RWMutex
}

func (l *ListenerManager) verifyLive(live api.Live) bool {
	for i := 0; i < 3; i++ {
		_, err := live.GetRoom()
		if err == nil {
			return true
		}
		if api.IsRoomNotExistsError(err) {
			return false
		}
	}
	return false
}

func (l *ListenerManager) AddListener(live api.Live) error {

	if !l.verifyLive(live) {
		return roomNotExistError
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	if _, ok := l.savers[live]; ok {
		return listenerExistError
	}

	listener := &Listener{
		Live:   live,
		ticker: time.NewTicker(time.Duration(core.GetInstance(l.ctx).Config.Interval) * time.Second),
		ed:     core.GetInstance(l.ctx).EventDispatcher,
		stop:   make(chan struct{}),
	}

	l.savers[live] = listener
	listener.Start()
	return nil
}

func (l *ListenerManager) RemoveListener(live api.Live) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if listener, ok := l.savers[live]; !ok {
		return listenerNotExistError
	} else {
		listener.Close()
		delete(l.savers, live)
		return nil
	}
}

func (l *ListenerManager) HasListener(live api.Live) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	_, ok := l.savers[live]
	return ok
}
