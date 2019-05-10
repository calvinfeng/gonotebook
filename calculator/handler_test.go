package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func addQueryParam(r *http.Request, lop, rop string) {
	q := r.URL.Query()
	q.Add("lop", lop)
	q.Add("rop", rop)
	r.URL.RawQuery = q.Encode()
}

func TestAddHandler(t *testing.T) {
	handler := CreateOperationHandler(Addition)

	r := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	addQueryParam(r, "1", "1")
	w := httptest.NewRecorder()

	handler(w, r)

	// Extract response from the handler.
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Error("bad status code from handler")
	}

	if string(body) != "1 + 1 = 2" {
		t.Error("wrong result from handler")
	}
}
