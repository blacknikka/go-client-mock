package fetch

import (
	"net/http"
	"errors"

	"github.com/blacknikka/go-client-mock/usecase"
)

type FetchAndCheck struct{}

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

func (FetchAndCheck) Exec() (bool, error) {
	client := &http.Client{}
	contentUsecase := usecase.NewContentUsecase(client)
	result, err := contentUsecase.GetContent(`http://localhost:8086/query?db=telegraf&q=SELECT mean("usage_idle") FROM "cpu" WHERE time >= now() - 3m GROUP BY time(10s) fill(null)`)
	if err != nil {
		return false, errors.New("request failed")
	}

	encoder := usecase.JsonEncoder{}
	var structure InfluxStructure
	_, err = encoder.Decode(result, &structure)
	if err != nil {
		return false, errors.New("error for request")
	}

	return true, nil
}
