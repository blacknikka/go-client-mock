package usecase

import (
	"context"
	"time"
)

type TimerJob struct {
	Time time.Duration
	Chan chan string
	Job  func(chan string)
}

// AddAJob adds a job and will run it.
func AddAJob(ctx context.Context, job TimerJob) bool {
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
			case <-ctx.Done():
				println(ctx.Err())
				return
			default:
			}
		}
	}()

	return true
}
