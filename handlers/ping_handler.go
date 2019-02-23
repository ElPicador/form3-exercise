package handlers

import (
	"net/http"
)

var pingResponse = []byte(`{"yes": "i_am"}`)

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(pingResponse)
}
