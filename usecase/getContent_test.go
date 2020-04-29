package usecase

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/blacknikka/go-client-mock/client"
)

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()

	os.Exit(ret)
}

type clientMock struct {
	client.HttpClient
	MockedDo func(req *http.Request) (*http.Response, error)
}

func (c *clientMock) Do(req *http.Request) (*http.Response, error) {
	return c.MockedDo(req)
}

func TestNewContentUsecase(t *testing.T) {
	t.Run("Client Do正常系", func(t *testing.T) {
		spyMocked := &clientMock{
			MockedDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
				}, nil
			},
		}
		contentUsecase := NewContentUsecase(spyMocked)
		content, err := contentUsecase.GetContent()

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		if content != "OK" {
			t.Errorf("Returned Content invalid: %v", content)
		}
	})
}
