package recorders

import (
	"github.com/hr3lxphr6j/bililive-go/src/api"
	"os/exec"
	"time"
)

type Recorder struct {
	Live       api.Live
	OutPutFile string
	StartTime  time.Time
	cmd        *exec.Cmd
	stop       chan struct{}
}

func (r *Recorder) run() {
Loop:
	for {
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
	go r.run()
	return nil
}

func (r *Recorder) Close() {
	close(r.stop)
	stdIn, err := r.cmd.StdinPipe()
	if err != nil {
		return
	}
	stdIn.Write([]byte("q"))
}
