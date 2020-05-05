package fetch

import (
	"errors"

	"github.com/blacknikka/go-client-mock/usecase"
)

const (
	ErrorForRequest string = "request failed"
	ErrorForDecode  string = "decode error"
)

type fetchAndCheck struct {
	contentUsecase *usecase.ContentUsecase
}

func NewFetchAndCheck(cu *usecase.ContentUsecase) *fetchAndCheck {
	return &fetchAndCheck{
		contentUsecase: cu,
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

func (fc fetchAndCheck) Exec() (bool, error) {
	result, err := fc.contentUsecase.GetContent(`http://localhost:8086/query?db=telegraf&q=SELECT mean("usage_idle") FROM "cpu" WHERE time >= now() - 3m GROUP BY time(10s) fill(null)`)
	if err != nil {
		return false, errors.New(ErrorForRequest)
	}

	encoder := usecase.JsonEncoder{}
	var structure InfluxStructure
	_, err = encoder.Decode(result, &structure)
	if err != nil {
		return false, errors.New(ErrorForDecode)
	}

	return true, nil
}
