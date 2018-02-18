package recorders

import (
	"context"
	"fmt"
	"github.com/hr3lxphr6j/bililive-go/src"
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"path/filepath"
	"sync"
	"time"
)

type IRecorderManager interface {
	AddRecorder(live api.Live) error
	GetRecorder(live api.Live) (*Recorder, error)
	RemoveRecorder(live api.Live) (time.Duration, error)
}

func NewIRecorderManager(ctx context.Context) IRecorderManager {
	rm := new(RecorderManager)
	rm.ctx = ctx
	return rm
}

type RecorderManager struct {
	ctx   context.Context
	saver map[api.Live]*Recorder
	lock  sync.RWMutex
}

func (r *RecorderManager) AddRecorder(live api.Live) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.saver[live]; ok {
		return recorderExistError
	}
	info, err := live.GetRoom()
	if err != nil {
		return err
	}
	t := time.Now()
	recorder := &Recorder{
		Live: live,
		OutPutFile: filepath.Join(
			core.GetInstance(r.ctx).Config.OutPutPath,
			fmt.Sprintf("[%d-%d-%d %d-%d-%d][%s]%s.flv",
				t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), info.HostName, info.RoomName)),
		ed: core.GetInstance(r.ctx).EventDispatcher,
	}
	r.saver[live] = recorder
	return nil

}

func (r *RecorderManager) GetRecorder(live api.Live) (*Recorder, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if r, ok := r.saver[live]; !ok {
		return nil, recorderNotExistError
	} else {
		return r, nil
	}
}

func (r *RecorderManager) RemoveRecorder(live api.Live) (time.Duration, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if recorder, ok := r.saver[live]; ok {
		return 0, recorderExistError
	} else {
		recorder.Close()
		delete(r.saver, live)
		return time.Now().Sub(recorder.StartTime), nil
	}
}
