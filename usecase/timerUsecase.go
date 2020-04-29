package usecase

import (
	"fmt"
	"time"
)

type TimerJob struct {
	Time time.Duration
	Chan chan string
	Job  func(chan string)
}

// AddAJob adds a job and will run it.
func AddAJob(job TimerJob) bool {
	if job.Time == 0 {
		return false
	}

	go func() {
		for range time.Tick(job.Time) {
			fmt.Println("Tick!!")
			job.Job(job.Chan)
		}
	}()

	return true
}
