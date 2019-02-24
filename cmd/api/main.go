package api

import (
	"github.com/ElPicador/form3-exercise/pkg/api"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"log"
	"net/http"
)

func main() {
	repository := payments.NewRepository(".")
	generator := payments.NewPaymentIDGenerator()
	a := api.New(repository, generator)

	srv := &http.Server{
		Handler: a.Handler(),
		Addr:    "127.0.0.1:3000",
	}
	log.Println("starting server on http://localhost:3000")

	log.Fatal(srv.ListenAndServe())
}
