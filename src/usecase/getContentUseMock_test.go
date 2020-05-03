package usecase

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetContentUseMock(t *testing.T) {
	t.Run("GetContentUseMock正常系", func(t *testing.T) {
		// httpmockの有効化
		httpmock.Activate()
		// 処理後にリセットを行う
		defer httpmock.DeactivateAndReset()

		// mock化したいRequestを登録
		// (requestするhttpメソッド名, requestするURL, 期待する戻り値)を引数に設定する
		httpmock.RegisterResponder("GET", "http://json/comments",
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(200, "ok")
				return res, nil
			},
		)

		content, err := GetContentUseMock()
		if err != nil {
			t.Errorf("error should be nil: %v", err)
		}

		if content != "ok" {
			t.Errorf("content invalid: %v", content)
		}
	})

	t.Run("GetContentUseMock異常系", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "http://json/comments",
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(500, "")
				return res, nil
			},
		)

		content, err := GetContentUseMock()
		if err == nil {
			t.Errorf("error shouldn't be nil: %v", err)
		}

		if err.Error() != IntervanServerError {
			t.Errorf("error message invalid want: %v, got: %v", IntervanServerError, err.Error())
		}

		if content != "" {
			t.Errorf("content invalid: %v", content)
		}
	})
}
