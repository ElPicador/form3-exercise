package main

import (
	"github.com/ElPicador/form3-exercise/pkg/api"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

func main() {
	repository := payments.NewRepository("./data")
	generator := payments.NewPaymentIDGenerator()
	a := api.New(repository, generator)

	srv := &http.Server{
		Handler: a.Handler(),
		Addr:    "127.0.0.1:3000",
	}
	log.Println("using ./data as the directory for the storage")
	log.Println("starting server on http://localhost:3000")

	log.Fatal(srv.ListenAndServe())
}
