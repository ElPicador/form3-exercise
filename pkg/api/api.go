package api

import (
	"github.com/ElPicador/form3-exercise/pkg/handlers"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	repository *payments.Repository
	generator  payments.PaymentIDGenerator
}

func New(repository *payments.Repository, generator payments.PaymentIDGenerator) *API {
	return &API{
		repository: repository,
		generator:  generator,
	}
}

func (a *API) Handler() *mux.Router {
	r := mux.NewRouter()
	r.
		Path("/isalive").
		HandlerFunc(handlers.PingHandler)
	r.
		Path("/1/payments").
		Methods(http.MethodGet).
		Handler(handlers.NewListPaymentsHandler(a.repository))
	r.
		Path("/1/payments").
		Methods(http.MethodPost).
		Handler(handlers.NewCreatePaymentHandler(a.repository, a.generator))
	r.
		Path("/1/payments/{uuid}").
		Methods(http.MethodGet).
		Handler(handlers.NewGetPaymentHandler(a.repository))
	r.
		Path("/1/payments/{uuid}").
		Methods(http.MethodDelete).
		Handler(handlers.NewDeletePaymentHandler(a.repository))

	return r
}
