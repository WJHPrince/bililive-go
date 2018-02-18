package listeners

import (
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"github.com/hr3lxphr6j/bililive-go/src/lib/events"
	"time"
)

type Listener struct {
	Live api.Live

	ticker *time.Ticker
	ed     events.IEventDispatcher
	stop   chan struct{}
	isStop bool
	status bool
}

func (l *Listener) Start() error {
	go l.run()
	l.isStop = false
	return nil
}

func (l *Listener) Close() {
	l.isStop = true
	close(l.stop)
}

func (l *Listener) run() {
	defer func() {
		l.ticker.Stop()
	}()
Loop:
	for !l.isStop {
		select {
		case <-l.stop:
			break Loop
		case <-l.ticker.C:
			info, err := l.Live.GetRoom()
			if err != nil {
				continue Loop
			}
			if info.Status == l.status {
				continue Loop
			}
			l.status = info.Status
			if l.status {
				l.ed.DispatchEvent(events.NewEvent(LiveStart, l.Live))
			} else {
				l.ed.DispatchEvent(events.NewEvent(LiveEnd, l.Live))
			}
		}
	}
}
