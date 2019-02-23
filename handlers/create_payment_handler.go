package handlers

import (
	"encoding/json"
	"github.com/ElPicador/form3-exercise/payments"
	"net/http"
)

type createPaymentRequest struct {
	Payment payments.Payment `json:"payment"`
}

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		WriteJSONMessage(w, http.StatusBadRequest, "body of request must be a json object")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody createPaymentRequest
	err := decoder.Decode(&requestBody)

	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid json")
		return
	}

	WriteJSON(w, http.StatusCreated, []byte(`{"payment_id":"uuidv4"}`))
}
