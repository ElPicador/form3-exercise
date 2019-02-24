package handlers

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	id := vars["uuid"]

	_, err := uuid.Parse(id)
	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	exists, err := h.repository.Exists(id)
	if err != nil {
		log.Printf("[ERROR] cannot delete payment: %s\n", err)
		Write500(w)
		return
	}

	if !exists {
		WriteJSONMessage(w, http.StatusNotFound, "payment doesnt exist")
		return
	}

	err = h.repository.Delete(id)
	if err != nil {
		log.Printf("[ERROR] cannot delete payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
