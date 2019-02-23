package handlers

import (
	"github.com/ElPicador/form3-exercise/payments"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	id := vars["uuid"]

	_, err := uuid.Parse(id)
	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	exists, err := h.repository.Exists(id)
	if err != nil {
		log.Printf("[ERROR] cannot get payment: %s\n", err)
		Write500(w)
		return
	}

	if !exists {
		WriteJSONMessage(w, http.StatusNotFound, "payment doesnt exist")
		return
	}

	payment, err := h.repository.Get(id)
	if err != nil {
		log.Printf("[ERROR] cannot get payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusOK, &getPaymentResponse{Payment: *payment})
}
