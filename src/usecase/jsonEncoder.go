package usecase

import (
	"encoding/json"
	"errors"
)

type JsonEncoder struct{}

func (c *JsonEncoder) Decode(str string, structure interface{}) (interface{}, error) {
	err := json.Unmarshal([]byte(str), structure)
	if err != nil {
		return nil, errors.New("json unmarshal failed")
	}
	return &structure, nil
}
