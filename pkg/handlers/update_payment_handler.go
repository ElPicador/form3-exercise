package handlers

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

type UpdatePaymentHandler struct {
	repository *payments.Repository
}

func NewUpdatePaymentHandler(repository *payments.Repository) *UpdatePaymentHandler {
	return &UpdatePaymentHandler{
		repository: repository,
	}
}

func (h *UpdatePaymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ValidateAndGetPaymentFromBody(w, r)
	if err != nil {
		return
	}

	id, err := ValidateAndGetUUIDFromParams(w, r)
	if err != nil {
		return
	}

	exists, err := h.repository.Exists(id.String())
	err = ValidateExistence(exists, err, w)
	if err != nil {
		return
	}

	requestBody.ID = id.String()
	err = h.repository.Save(id.String(), requestBody)
	if err != nil {
		log.Printf("[ERROR] cannot update payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
