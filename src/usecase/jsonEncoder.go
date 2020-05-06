package usecase

import (
	"encoding/json"
	"errors"
	"reflect"
)

type JsonEncoder struct{}

func (c *JsonEncoder) Decode(str string, t reflect.Type) (interface{}, error) {
	structure := reflect.New(t).Interface()
	err := json.Unmarshal([]byte(str), structure)
	if err != nil {
		return nil, errors.New("json unmarshal failed")
	}
	return structure, nil
}
