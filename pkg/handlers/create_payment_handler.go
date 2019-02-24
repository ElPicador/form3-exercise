package handlers

import (
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
	requestBody, err := ValidateAndGetPaymentFromBody(w, r)
	if err != nil {
		return
	}

	id, err := h.generator.GenerateUniqueID()
	if err != nil {
		log.Printf("[ERROR] cannot generate UUIDv4: %s\n", err)
		Write500(w)
		return
	}

	requestBody.ID = id.String()
	err = h.repository.Save(id.String(), requestBody)
	if err != nil {
		log.Printf("[ERROR] cannot save payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusCreated, &createPaymentResponse{PaymentID: id.String()})
}
