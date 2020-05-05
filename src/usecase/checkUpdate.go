package usecase

import (
	"time"
	"errors"
)

type CheckUpdater struct {
	latestTimestamp time.Time
}

func (c *CheckUpdater) CheckUpdate(checkedTime time.Time) (bool, error) {
	if checkedTime.IsZero() == true {
		return false, errors.New("Checked time is zero")
	}

	if checkedTime.After(c.latestTimestamp) {
		// given time is a newer one than latestTimestamp
		c.latestTimestamp = checkedTime
		return true, nil
	}

	// given time is not a newer one than latestTimestamp
	return false, nil
}
