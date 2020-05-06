package newer

import (
	"time"

	"github.com/blacknikka/go-client-mock/usecase"
)

type GetNewer struct {
	updateChecker *usecase.CheckUpdater
}

func NewGetNewer(checker *usecase.CheckUpdater) *GetNewer {
	return &GetNewer{
		updateChecker: checker,
	}
}

func (newer *GetNewer) Get(data []interface{}) ([][]interface{}, error) {
	result := [][]interface{}{}
	for _, d := range data {
		if timeData, ok := d.([]interface{}); ok {
			if timeString, ok := timeData[0].(string); ok {
				if t, ok := time.Parse("2006-01-02T15:04:05Z", timeString); ok == nil {
					if isNewer, _ := newer.updateChecker.CheckUpdate(t); isNewer {
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
