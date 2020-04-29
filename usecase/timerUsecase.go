package usecase

import (
	"time"
)

type TimerJob struct {
	Time     time.Duration
	Chan     chan string
	StopChan chan struct{}
	Job      func(chan string)
}

// AddAJob adds a job and will run it.
func AddAJob(job TimerJob) bool {
	if job.Time == 0 {
		return false
	}

	go func() {
		defer func() {
			close(job.Chan)
		}()

		for range time.Tick(job.Time) {
			job.Job(job.Chan)

			select {
			case <-job.StopChan:
				println("stop request received.")
				return
			default:
			}
		}
	}()

	return true
}
