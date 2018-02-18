package recorders

import (
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"github.com/hr3lxphr6j/bililive-go/src/lib/events"
	"os/exec"
	"time"
)

type Recorder struct {
	Live       api.Live
	OutPutFile string
	StartTime  time.Time

	cmd    *exec.Cmd
	ed     events.IEventDispatcher
	stop   chan struct{}
	isStop bool
}

func (r *Recorder) run() {
Loop:
	for !r.isStop {
		select {
		case <-r.stop:
			break Loop
		default:
			urls, _ := r.Live.GetUrls()
			r.cmd = exec.Command(
				"ffmpeg",
				"-y", "-re",
				"-i", urls[0].String(),
				"-c", "copy",
				"-bsf:a", "aac_adtstoasc",
				r.OutPutFile,
			)
			r.cmd.Run()
		}
	}
}

func (r *Recorder) Start() error {
	r.StartTime = time.Now()
	r.stop = make(chan struct{})
	r.isStop = false
	go r.run()
	r.ed.DispatchEvent(events.NewEvent(RecordeStart, r.Live))
	return nil
}

func (r *Recorder) Close() {
	close(r.stop)
	stdIn, err := r.cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdIn.Write([]byte("q"))
	r.isStop = true
	r.ed.DispatchEvent(events.NewEvent(RecordeStop, r.Live))
}
