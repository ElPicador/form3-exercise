package handlers

import (
	"encoding/json"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	id := vars["uuid"]

	_, err := uuid.Parse(id)
	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	if r.Body == nil {
		WriteJSONMessage(w, http.StatusBadRequest, "body of request must be a json object")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody payments.Payment
	err = decoder.Decode(&requestBody)

	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid json")
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

	requestBody.ID = id
	err = h.repository.Save(id, &requestBody)
	if err != nil {
		log.Printf("[ERROR] cannot update payment: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
