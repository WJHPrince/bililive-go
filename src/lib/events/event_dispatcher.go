package events

import (
	"context"
	"github.com/hr3lxphr6j/bililive-go/src/instance"
	"sync"
)

func NewIEventDispatcher(ctx context.Context) IEventDispatcher {
	ed := &EventDispatcher{
		saver: make(map[EventType]eventListenerSet),
		lock:  new(sync.RWMutex),
	}
	instance.GetInstance(ctx).EventDispatcher = ed
	return ed
}

// 事件分发器接口
type IEventDispatcher interface {
	AddEventListener(eventType EventType, listener *EventListener)
	RemoveEventListener(eventType EventType, listener *EventListener)
	RemoveAllEventListener(eventType EventType)
	DispatchEvent(event *Event)
}

// 事件分发器
type EventDispatcher struct {
	saver map[EventType]eventListenerSet
	lock  *sync.RWMutex
}

func (e *EventDispatcher) Start(ctx context.Context) error {
	return nil
}

func (e *EventDispatcher) Close(ctx context.Context) {

}

func (e *EventDispatcher) AddEventListener(eventType EventType, listener *EventListener) {
	//e.lock.Lock()
	//defer e.lock.Unlock()
	_, isExist := e.saver[eventType]
	if !isExist {
		e.saver[eventType] = make(map[*EventListener]bool)
	}
	e.saver[eventType][listener] = true
}

func (e *EventDispatcher) RemoveEventListener(eventType EventType, listener *EventListener) {
	//e.lock.Lock()
	//defer e.lock.Unlock()
	delete(e.saver[eventType], listener)
}

func (e *EventDispatcher) RemoveAllEventListener(eventType EventType) {
	//e.lock.Lock()
	//defer e.lock.Unlock()
	delete(e.saver, eventType)
}

func (e *EventDispatcher) DispatchEvent(event *Event) {
	//e.lock.RLock()
	//defer e.lock.RLock()
	for l := range e.saver[event.Type] {
		l.Handler(event)
	}
}
