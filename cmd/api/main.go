package api

import (
	"github.com/ElPicador/form3-exercise/pkg/handlers"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	repository := payments.NewRepository(".")
	generator := payments.NewPaymentIDGenerator()

	r := mux.NewRouter()
	r.
		Path("/isalive").
		HandlerFunc(handlers.PingHandler)
	r.
		Path("/1/payments").
		Methods(http.MethodPost).
		Handler(handlers.NewCreatePaymentHandler(repository, generator))
	r.
		Path("/1/payments/{uuid}").
		Methods(http.MethodGet).
		Handler(handlers.NewGetPaymentHandler(repository))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
	}
	log.Println("starting server on http://localhost:3000")

	log.Fatal(srv.ListenAndServe())
}
