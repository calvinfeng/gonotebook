package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type QueryParam struct {
	Name  string
	Value string
}

func addQueryParam(r *http.Request, params ...QueryParam) {
	q := r.URL.Query()
	for _, param := range params {
		q.Add(param.Name, param.Value)
	}
	r.URL.RawQuery = q.Encode()
}

type HandlerTestCase struct {
	ExpectedStatus int
	LeftOperand    string
	RightOperand   string
	ExpectedBody   string
}

// TestAddHandler is an example of how to use httptest package to test a handler.
func TestAddHandler(t *testing.T) {
	testCases := []HandlerTestCase{
		HandlerTestCase{
			ExpectedStatus: http.StatusOK,
			LeftOperand:    "1",
			RightOperand:   "1",
			ExpectedBody:   "1 + 1 = 2",
		},
		HandlerTestCase{
			ExpectedStatus: http.StatusBadRequest,
			LeftOperand:    "foo",
			RightOperand:   "2",
			ExpectedBody:   "invalid lop query parameter",
		},
		HandlerTestCase{
			ExpectedStatus: http.StatusBadRequest,
			LeftOperand:    "1",
			RightOperand:   "foo",
			ExpectedBody:   "invalid rop query parameter",
		},
	}

	handler := CreateOperationHandler(Addition)

	for _, tc := range testCases {
		r := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
		addQueryParam(r, QueryParam{"lop", tc.LeftOperand}, QueryParam{"rop", tc.RightOperand})
		w := httptest.NewRecorder()

		handler(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != tc.ExpectedStatus {
			t.Errorf("actual status %d != expected status %d", resp.StatusCode, tc.ExpectedStatus)
		}

		if string(body) != tc.ExpectedBody {
			t.Errorf("actual body %s != expected body %s", body, tc.ExpectedBody)
		}
	}
}
