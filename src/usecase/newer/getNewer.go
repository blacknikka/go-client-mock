package newer

import (
	"time"
)

type GetNewer struct{}

func (GetNewer) Get(data []interface{}, baseTime time.Time) ([][]interface{}, error) {
	result := [][]interface{}{}
	for _, d := range data {
		if timeData, ok := d.([]interface{}); ok {
			if timeString, ok := timeData[0].(string); ok {
				if t, ok := time.Parse("2006-01-02T15:04:05Z", timeString); ok == nil {
					if t.After(baseTime) {
						array := []interface{}{}
						array = append(array, t)
						// if this dataset is newer than the base time,
						// it should be transformed to be value(float or string).
						for i := 1; i < len(timeData); i++ {
							if casted, ok := timeData[i].(int); ok {
								array = append(array, casted)
							} else if casted, ok := timeData[i].(float64); ok {
								array = append(array, casted)
							} else if casted, ok := timeData[i].(string); ok {
								array = append(array, casted)
							}
						}

						result = append(result, array)
					}
				}
			}
		}
	}

	return result, nil
}
