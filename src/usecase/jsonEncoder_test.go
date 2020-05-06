package usecase

import (
	"testing"
	"reflect"
)

type influxStructure struct {
	Result []struct {
		Id     int `json:"statement_id"`
		Series []struct {
			Name    string        `json:"name"`
			Columns []string      `json:"columns"`
			Values  []interface{} `json:values`
		} `json:"series"`
	} `json:"results"`
}

const checkedStr string = `{
	"results": [
		{
			"statement_id": 0,
			"series": [
				{
					"name": "cpu",
					"columns": [
						"time",
						"mean"
					],
					"values": [
						[
							"2020-05-05T05:49:10Z",
							null
						],
						[
							"2020-05-05T05:49:20Z",
							98.31379745672704
						],
						[
							"2020-05-05T05:49:30Z",
							98.72237582572546
						]
					]
				}
			]
		}
	]
}`

func TestDecode(t *testing.T) {
	t.Run("Decode正常系", func(t *testing.T) {
		sut := &JsonEncoder{}
		result, err := sut.Decode(checkedStr, reflect.TypeOf(influxStructure{}))
		if err != nil {
			t.Errorf("error should be nil :%v", err)
		}

		if result == nil {
			t.Errorf("result shouldn't be nil: %v", result)
		}

		if reflect.TypeOf(result) != reflect.TypeOf(&influxStructure{}) {
			t.Errorf("returned type is invalid. want :%T, got: %T", &influxStructure{}, result)
		}
	})

	t.Run("Decode異常系", func(t *testing.T) {
		sut := &JsonEncoder{}

		// path the invalid json structure string.
		result, err := sut.Decode("{", reflect.TypeOf(influxStructure{}))
		if err == nil {
			t.Errorf("error shouldn't be nil: %v", err)
		}

		if result != nil {
			t.Errorf("result should be nil: %v", result)
		}
	})
}
