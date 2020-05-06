package fetch

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/blacknikka/go-client-mock/client"
	"github.com/blacknikka/go-client-mock/usecase"
)

const responseString string = `{
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

type clientMock struct {
	client.HttpClient
	MockedDo func(req *http.Request) (*http.Response, error)
}

func (c *clientMock) Do(req *http.Request) (*http.Response, error) {
	return c.MockedDo(req)
}

func TestExec(t *testing.T) {
	t.Run("Exec正常系", func(t *testing.T) {
		spyMocked := &clientMock{
			MockedDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(responseString)),
				}, nil
			},
		}
		contentUsecase := usecase.NewContentUsecase(spyMocked)

		fetchAndCheck := NewFetchAndCheck(contentUsecase)
		result, err := fetchAndCheck.Exec()

		if err != nil {
			t.Errorf("err should be nil: %v", nil)
		}

		if result != true {
			t.Errorf("result should be true: %v", result)
		}
	})

	t.Run("Exec異常系_request error", func(t *testing.T) {
		spyMocked := &clientMock{
			MockedDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("something error")
			},
		}
		contentUsecase := usecase.NewContentUsecase(spyMocked)

		fetchAndCheck := NewFetchAndCheck(contentUsecase)
		result, err := fetchAndCheck.Exec()
		if result != false {
			t.Errorf("result should be false: %v", result)
		}

		if err != nil {
			if err.Error() != ErrorForRequest {
				t.Errorf("err message is invalid: %v", err.Error())
			}
		} else {
			t.Errorf("err shouldn't be nil: %v", err)
		}
	})

	t.Run("Exec異常系_json decode error", func(t *testing.T) {
		spyMocked := &clientMock{
			MockedDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{")),
				}, nil
			},
		}
		contentUsecase := usecase.NewContentUsecase(spyMocked)

		fetchAndCheck := NewFetchAndCheck(contentUsecase)
		result, err := fetchAndCheck.Exec()
		if result != false {
			t.Errorf("result should be false: %v", result)
		}

		if err != nil {
			if err.Error() != ErrorForDecode {
				t.Errorf("err message is invalid: %v", err.Error())
			}
		} else {
			t.Errorf("err shouldn't be nil: %v", err)
		}
	})
}
