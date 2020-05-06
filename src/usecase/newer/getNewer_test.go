package newer

import (
	"testing"
	"time"

	"github.com/blacknikka/go-client-mock/usecase"
)

const (
	timeFormat string = "2006-01-02T15:04:05Z"
)

func TestGet(t *testing.T) {
	t.Run("Get正常系", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:20Z",
				98.31379745672704,
			},
			[]interface{} {
				"2020-05-05T05:49:30Z",
				98.72237582572546,
			},
		}

		// create a base time
		baseTime, _ := time.Parse(timeFormat, "2020-05-05T05:49:19Z")

		// create checker (set base time)
		checker := &usecase.CheckUpdater{}
		checker.CheckUpdate(baseTime)

		// create a test instance
		sut := NewGetNewer(checker)
		result, err := sut.Get(data)
		if err != nil {
			t.Errorf("error should be nil :%v", err)
		}

		if len(result) != 2 {
			t.Errorf("the length of returned variable is not valid :%v", len(result))
		}

		comparer := []struct {
			Time string
			Value float64
		} {
			{"2020-05-05T05:49:20Z", 98.31379745672704, },
			{"2020-05-05T05:49:30Z", 98.72237582572546, },
		}

		for idx, data := range comparer {
			if tm, _ := time.Parse(timeFormat, data.Time); tm.Equal(result[idx][0].(time.Time)) != true {
				t.Errorf("want %v, got %v", data.Time, result[idx][0])
			}

			if data.Value != result[idx][1].(float64) {
				t.Errorf("want %v, got %v", data.Value, result[idx][1])
			}
		}
	})

	t.Run("Get正常系_baseTimeと同じ時間が含まれる場合", func(t *testing.T) {
		// 基準時間と同じ時間は、newerとしては扱わない
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:20Z",
				98.31379745672704,
			},
			[]interface{} {
				"2020-05-05T05:49:30Z",
				98.72237582572546,
			},
		}

		// create a base time
		baseTime, _ := time.Parse(timeFormat, "2020-05-05T05:49:20Z")

		// create checker (set base time)
		checker := &usecase.CheckUpdater{}
		checker.CheckUpdate(baseTime)

		// create a test instance
		sut := NewGetNewer(checker)
		result, _ := sut.Get(data)

		comparer := []struct {
			Time string
			Value float64
		} {
			{"2020-05-05T05:49:30Z", 98.72237582572546, },
		}

		for idx, data := range comparer {
			if tm, _ := time.Parse(timeFormat, data.Time); tm.Equal(result[idx][0].(time.Time)) != true {
				t.Errorf("want %v, got %v", data.Time, result[idx][0])
			}

			if data.Value != result[idx][1].(float64) {
				t.Errorf("want %v, got %v", data.Value, result[idx][1])
			}
		}
	})

	t.Run("Get正常系_配列の長さが長い場合", func(t *testing.T) {
		// 基準時間と同じ時間は、newerとしては扱わない
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:20Z",
				1.0,
				2,
				-3,
				-4.0,
				0,
				4294967295,
			},
		}

		// create a base time
		baseTime, _ := time.Parse(timeFormat, "2020-05-05T05:49:19Z")

		// create checker (set base time)
		checker := &usecase.CheckUpdater{}
		checker.CheckUpdate(baseTime)

		// create a test instance
		sut := NewGetNewer(checker)
		result, _ := sut.Get(data)

		comparer := []struct {
			Time string
			Value1 interface{}
			Value2 interface{}
			Value3 interface{}
			Value4 interface{}
			Value5 interface{}
			Value6 interface{}
		} {
			{"2020-05-05T05:49:20Z", 1.0, 2, -3, -4.0, 0, 4294967295, },
		}

		for idx, data := range comparer {
			if tm, _ := time.Parse(timeFormat, data.Time); tm.Equal(result[idx][0].(time.Time)) != true {
				t.Errorf("want %v, got %v", data.Time, result[idx][0])
			}

			if data.Value1 != result[idx][1] {
				t.Errorf("want %v, got %v", data.Value1, result[idx][1])
			}

			if data.Value2 != result[idx][2] {
				t.Errorf("want %v, got %v", data.Value2, result[idx][2])
			}

			if data.Value3 != result[idx][3] {
				t.Errorf("want %v, got %v", data.Value3, result[idx][3])
			}

			if data.Value4 != result[idx][4] {
				t.Errorf("want %v, got %v", data.Value4, result[idx][4])
			}

			if data.Value5 != result[idx][5] {
				t.Errorf("want %v, got %v", data.Value5, result[idx][5])
			}

			if data.Value6 != result[idx][6] {
				t.Errorf("want %v, got %v", data.Value6, result[idx][6])
			}
		}
	})

	t.Run("Get正常系_過去時間のデータは対象にしない", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:10Z",
				1,
			},
			[]interface{} {
				"2020-05-05T05:49:20Z",
				2,
			},
			[]interface{} {
				"2020-05-05T05:49:30Z",
				3,
			},
		}

		// create a base time
		baseTime, _ := time.Parse(timeFormat, "2020-05-05T05:49:19Z")

		// create checker (set base time)
		checker := &usecase.CheckUpdater{}
		checker.CheckUpdate(baseTime)

		// create a test instance
		sut := NewGetNewer(checker)
		result, _ := sut.Get(data)

		comparer := []struct {
			Time string
			Value interface{}
		} {
			{"2020-05-05T05:49:20Z", 2, },
			{"2020-05-05T05:49:30Z", 3, },
		}

		for idx, data := range comparer {
			if tm, _ := time.Parse(timeFormat, data.Time); tm.Equal(result[idx][0].(time.Time)) != true {
				t.Errorf("want %v, got %v", data.Time, result[idx][0])
			}

			if data.Value != result[idx][1] {
				t.Errorf("want %v, got %v", data.Value, result[idx][1])
			}
		}
	})

}
