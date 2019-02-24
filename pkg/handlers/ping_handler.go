package handlers

import (
	"net/http"
)

var pingResponse = []byte(`{"yes": "i_am"}`)

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	WriteContentTypeJSON(w, http.StatusOK, pingResponse)
}
