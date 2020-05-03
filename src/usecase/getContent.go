package usecase

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/blacknikka/go-client-mock/client"
)

const (
	InternalServerError string = "internal error"
)

func NewContentUsecase(c client.HttpClient) *ContentUsecase {
	return &ContentUsecase{
		httpClient: c,
	}
}

type ContentUsecase struct {
	httpClient client.HttpClient
}

func (content ContentUsecase) GetContent() (string, error) {
	// get client

	request, err := http.NewRequest("GET", "http://json/posts", nil)
	if err != nil {
		return "", errors.New("request creation error")
	}

	resp, err := content.httpClient.Do(request)
	if err != nil {
		return "", errors.New("request error")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return "", errors.New(InternalServerError)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("an error of reading request body")
	}

	return string(byteArray), nil
}
