package handlers

import (
	"net/http"
)

var pingResponse = []byte(`{"yes": "i_am"}`)

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	WriteJSON(w, http.StatusOK, pingResponse)
}
