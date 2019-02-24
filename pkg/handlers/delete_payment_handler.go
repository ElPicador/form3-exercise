package handlers

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

type DeletePaymentHandler struct {
	repository *payments.Repository
}

func NewDeletePaymentHandler(repository *payments.Repository) *DeletePaymentHandler {
	return &DeletePaymentHandler{
		repository: repository,
	}
}

func (h *DeletePaymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := ValidateAndGetUUIDFromParams(w, r)
	if err != nil {
		return
	}

	exists, err := h.repository.Exists(id.String())
	err = ValidateExistence(exists, err, w)
	if err != nil {
		return
	}

	err = h.repository.Delete(id.String())
	if err != nil {
		log.Printf("[ERROR] cannot delete payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
