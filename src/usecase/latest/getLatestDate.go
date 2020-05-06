package latest

import (
	"time"
)

type GetLatest struct{}

func (GetLatest) Get(data []interface{}) time.Time {
	var latest time.Time

	for _, d := range data {
		if timeData, ok := d.([]interface{}); ok {
			if len(timeData) == 2 {
				if timeData[1] != nil {
					// if the data is not nil, timeData[0] should be time.Time
					if timeString, ok := timeData[0].(string); ok {
						if t, ok := time.Parse("2006-01-02T15:04:05Z", timeString); ok == nil {
							if t.After(latest) {
								latest = t
							}
						}
					}
				}
			}
		}
	}
	return latest
}
