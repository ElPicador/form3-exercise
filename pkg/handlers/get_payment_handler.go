package handlers

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

type getPaymentResponse struct {
	Payment payments.Payment `json:"payment"`
}

type GetPaymentHandler struct {
	repository *payments.Repository
}

func NewGetPaymentHandler(repository *payments.Repository) *GetPaymentHandler {
	return &GetPaymentHandler{
		repository: repository,
	}
}

func (h *GetPaymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := ValidateAndGetUUIDFromParams(w, r)
	if err != nil {
		return
	}

	exists, err := h.repository.Exists(id.String())
	err = ValidateExistence(exists, err, w)
	if err != nil {
		return
	}

	payment, err := h.repository.Get(id.String())
	if err != nil {
		log.Printf("[ERROR] cannot get payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusOK, &getPaymentResponse{Payment: *payment})
}
