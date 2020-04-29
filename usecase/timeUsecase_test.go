package usecase

import (
	"testing"
	"time"
)

func TestAddAJob(t *testing.T) {
	t.Run("AddAJob 異常系 タイマがゼロ", func(t *testing.T) {
		timeJob := TimerJob{}

		result := AddAJob(timeJob)
		if result != false {
			t.Errorf("result should be false: %v", result)
		}
	})

	t.Run("AddAJob 正常系", func(t *testing.T) {
		resultCh := make(chan string)
		stopCh := make(chan struct{})

		counter := 0

		timeJob := TimerJob{
			Time:     (500 * time.Millisecond),
			Chan:     resultCh,
			StopChan: stopCh,
			Job: func(ch chan string) {
				counter++
			},
		}

		result := AddAJob(timeJob)
		if result != true {
			t.Errorf("result should be true: %v", result)
		}

		go func() {
			time.Sleep(time.Second * 1)

			// jobを停止
			close(stopCh)
		}()

		for {
			_, ok := <-resultCh
			if ok == false {
				break
			}
		}

		if counter < 2 {
			t.Errorf("counter invalid: %v", counter)
		}
	})
}
