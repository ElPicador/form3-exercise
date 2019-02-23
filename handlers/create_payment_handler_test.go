package handlers

import "net/http"

func CreatePaymentHandler(w http.ResponseWriter, _ *http.Request) {
	WriteJSON(w, http.StatusOK, pingResponse)
}
