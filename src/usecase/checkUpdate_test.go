package usecase

import (
	"testing"
	"time"
)

func TestCheckUpdate(t *testing.T) {
	t.Run("CheckUpdate異常系", func(t *testing.T) {
		checker := CheckUpdater{}
		result, err := checker.CheckUpdate(time.Time{})

		if err == nil {
			t.Errorf("error shouldn't be nil: %v", err)
		}

		if result != false {
			t.Errorf("result should be false: %v", result)
		}
	})

	t.Run("CheckUpdate正常系", func(t *testing.T) {
		checker := CheckUpdater{}
		result, err := checker.CheckUpdate(time.Now())

		if err != nil {
			t.Errorf("error should be nil: %v", err)
		}

		if result != true {
			t.Errorf("result should be true: %v", result)
		}
	})

	t.Run("CheckUpdate_未来時間を与えた場合", func(t *testing.T) {
		checker := CheckUpdater{}

		now := time.Now()
		after_than_now := now.Add(time.Second)

		result, _ := checker.CheckUpdate(now)
		if result != true {
			t.Errorf("result should be true: %v", result)
		}

		result, _ = checker.CheckUpdate(after_than_now)
		if result != true {
			t.Errorf("result should be true: %v", result)
		}
	})

	t.Run("CheckUpdate_同じ時間を与えた場合", func(t *testing.T) {
		checker := CheckUpdater{}

		now := time.Now()

		result, _ := checker.CheckUpdate(now)
		if result != true {
			t.Errorf("result should be true: %v", result)
		}

		// if the time is same as the time stored, it should return false (not new)
		result, _ = checker.CheckUpdate(now)

		if result != false {
			t.Errorf("result should be true: %v", result)
		}
	})

	t.Run("CheckUpdate_過去時間を与えた場合", func(t *testing.T) {
		checker := CheckUpdater{}

		now := time.Now()
		before_than_now := now.Add(-time.Second)

		result, _ := checker.CheckUpdate(now)
		if result != true {
			t.Errorf("result should be true: %v", result)
		}

		// if the time is before than the time stored, it should return false
		result, _ = checker.CheckUpdate(before_than_now)

		if result != false {
			t.Errorf("result should be true: %v", result)
		}
	})
}
