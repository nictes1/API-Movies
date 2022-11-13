package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func ClientRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	// request + headers
	var req = httptest.NewRequest(method, url,  bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	// response
	var res = httptest.NewRecorder()

	return req, res
}