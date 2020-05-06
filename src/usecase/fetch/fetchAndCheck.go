package fetch

import (
	"errors"
	"reflect"

	"github.com/blacknikka/go-client-mock/usecase"
	"github.com/blacknikka/go-client-mock/usecase/newer"
)

const (
	ErrorForRequest    string = "request failed"
	ErrorForDecode     string = "decode error"
	ErrorForLatestData string = "data is invalid (empty)"
)

type fetchAndCheck struct {
	contentUsecase *usecase.ContentUsecase
	newerChecker *newer.GetNewer
}

func NewFetchAndCheck(cu *usecase.ContentUsecase) *fetchAndCheck {
	return &fetchAndCheck{
		contentUsecase: cu,
		newerChecker: newer.NewGetNewer(&usecase.CheckUpdater{}),
	}
}

type InfluxStructure struct {
	Result []struct {
		Id     int `json:"statement_id"`
		Series []struct {
			Name    string        `json:"name"`
			Columns []string      `json:"columns"`
			Values  []interface{} `json:values`
		} `json:"series"`
	} `json:"results"`
}

func (fc fetchAndCheck) Exec() ([][]interface{}, error) {
	// fetch via http
	result, err := fc.contentUsecase.GetContent(`http://localhost:8086/query?db=telegraf&q=SELECT mean("usage_idle") FROM "cpu" WHERE time >= now() - 3m GROUP BY time(10s) fill(null)`)
	if err != nil {
		return nil, errors.New(ErrorForRequest)
	}

	// decode json string
	encoder := usecase.JsonEncoder{}
	content, err := encoder.Decode(result, reflect.TypeOf(InfluxStructure{}))
	if err != nil {
		return nil, errors.New(ErrorForDecode)
	}

	influx := content.(*InfluxStructure)

	// check the json structure is valid
	// check if having the specified data member.
	if ok := isValid(influx); ok == false {
		return nil, errors.New(ErrorForLatestData)
	}

	newerItems, err := fc.newerChecker.Get(influx.Result[0].Series[0].Values)
	if err == nil {
		return newerItems, nil
	}

	return nil, errors.New(ErrorForLatestData)
}

func isValid(influx *InfluxStructure) bool {
	if len(influx.Result) == 0 {
		return false
	}

	if len(influx.Result[0].Series) == 0 {
		return false
	}

	if len(influx.Result[0].Series[0].Values) == 0 {
		return false
	}

	return true
}
