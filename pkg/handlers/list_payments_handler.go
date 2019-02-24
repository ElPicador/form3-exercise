package handlers

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

type listPaymentsResponse struct {
	Payments []*payments.Payment `json:"payments"`
}

type ListPaymentsHandler struct {
	repository *payments.Repository
}

func NewListPaymentsHandler(repository *payments.Repository) *ListPaymentsHandler {
	return &ListPaymentsHandler{
		repository: repository,
	}
}

func (h *ListPaymentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := h.repository.GetAll()
	if err != nil {
		log.Printf("[ERROR] cannot get all payments: %s\n", err)
		Write500(w)
		return
	}

	WriteJSON(w, http.StatusOK, &listPaymentsResponse{Payments: p})
}
