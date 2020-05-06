package latest

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Run("Get正常系", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:10Z",
				nil,
			},
			[]interface{} {
				"2020-05-05T05:49:20Z",
				98.31379745672704,
			},
			[]interface{} {
				"2020-05-05T05:49:30Z",
				98.72237582572546,
			},
		}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != false {
			t.Errorf("result shouldn't be initialized time :%v", result)
		}

		compared, _ := time.Parse("2006-01-02T15:04:05Z", "2020-05-05T05:49:30Z")
		if result.Equal(compared) != true {
			t.Errorf("time is invalid :%v", result)
		}
	})

	t.Run("Get正常系_並び順を変更", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:30Z",
				98.72237582572546,
			},
			[]interface{} {
				"2020-05-05T05:49:10Z",
				nil,
			},
			[]interface{} {
				"2020-05-05T05:49:20Z",
				98.31379745672704,
			},
		}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != false {
			t.Errorf("result shouldn't be initialized time :%v", result)
		}

		compared, _ := time.Parse("2006-01-02T15:04:05Z", "2020-05-05T05:49:30Z")
		if result.Equal(compared) != true {
			t.Errorf("time is invalid :%v", result)
		}
	})

	t.Run("Get異常系_timeがすべてnil", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				"2020-05-05T05:49:10Z",
				nil,
			},
			[]interface{} {
				"2020-05-05T05:49:20Z",
				nil,
			},
		}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != true {
			t.Errorf("result should be initialized time :%v", result)
		}
	})

	t.Run("Get異常系_与えられた配列がおかしい", func(t *testing.T) {
		data := []interface{} {}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != true {
			t.Errorf("result should be initialized time :%v", result)
		}
	})

	t.Run("Get異常系_与えられた配列がおかしい2", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				nil,
			},
			[]interface{} {
				nil,
			},
		}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != true {
			t.Errorf("result should be initialized time :%v", result)
		}
	})

	t.Run("Get異常系_与えられた配列がおかしい3", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				nil,
				nil,
			},
			[]interface{} {
				nil,
				nil,
			},
		}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != true {
			t.Errorf("result should be initialized time :%v", result)
		}
	})

	t.Run("Get異常系_与えられた配列がおかしい4", func(t *testing.T) {
		data := []interface{} {
			[]interface{} {
				nil,
				1,
			},
			[]interface{} {
				nil,
				2,
			},
		}

		sut := GetLatest{}
		result := sut.Get(data)
		if result.IsZero() != true {
			t.Errorf("result should be initialized time :%v", result)
		}
	})
}

