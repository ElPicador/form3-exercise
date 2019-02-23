package handlers

import (
	"net/http"
	"net/http/httptest"
)

func ServeAndRecord(handler http.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	return rr
}
