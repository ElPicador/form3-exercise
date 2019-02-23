package main

import (
	"github.com/ElPicador/form3-exercise/handlers"
	"github.com/ElPicador/form3-exercise/payments"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	repository := payments.NewRepository(".")

	r := mux.NewRouter()
	r.
		Path("/isalive").
		HandlerFunc(handlers.PingHandler)
	r.
		Path("/1/payments").
		Methods(http.MethodPost).
		Handler(handlers.NewCreatePaymentHandler(repository))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
	}
	log.Println("starting server on http://localhost:3000")

	log.Fatal(srv.ListenAndServe())
}
