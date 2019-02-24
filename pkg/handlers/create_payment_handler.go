package handlers

import (
	"encoding/json"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

type createPaymentResponse struct {
	PaymentID string `json:"payment_id"`
}

type CreatePaymentHandler struct {
	repository *payments.Repository
	generator  payments.PaymentIDGenerator
}

func NewCreatePaymentHandler(repository *payments.Repository, generator payments.PaymentIDGenerator) *CreatePaymentHandler {
	return &CreatePaymentHandler{
		repository: repository,
		generator:  generator,
	}
}

func (h *CreatePaymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		WriteJSONMessage(w, http.StatusBadRequest, "body of request must be a json object")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody payments.Payment
	err := decoder.Decode(&requestBody)

	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid json")
		return
	}

	id, err := h.generator.GenerateUniqueID()
	if err != nil {
		log.Printf("[ERROR] cannot generate UUIDv4: %s\n", err)
		Write500(w)
		return
	}

	requestBody.ID = id.String()
	err = h.repository.Save(id.String(), &requestBody)
	if err != nil {
		log.Printf("[ERROR] cannot save payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusCreated, &createPaymentResponse{PaymentID: id.String()})
}
