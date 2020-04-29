package usecase

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	IntervanServerError string = "internal error"
)

func GetContentUseMock() (string, error) {
	// get client

	request, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		return "", errors.New("request creation error")
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", errors.New("request error")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return "", errors.New(IntervanServerError)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("an error of reading request body")
	}

	return string(byteArray), nil
}
