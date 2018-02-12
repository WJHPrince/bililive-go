package recorders

import (
	"context"
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"sync"
	"time"
)

type IRecorderManager interface {
	AddRecorder(ctx context.Context, live api.Live) error
	GetRecorder(ctx context.Context, live api.Live) *Recorder
	RemoveRecorder(ctx context.Context, live api.Live) (time.Duration, error)
}

type RecorderManager struct {
	saver map[api.Live]*Recorder
	lock  sync.RWMutex
}

func (r *RecorderManager) AddRecorder(ctx context.Context, live api.Live) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return nil

}

func (r *RecorderManager) GetRecorder(ctx context.Context, live api.Live) *Recorder {
	return nil
}

func (r *RecorderManager) RemoveRecorder(ctx context.Context, live api.Live) (time.Duration, error) {
	return 0, nil
}
