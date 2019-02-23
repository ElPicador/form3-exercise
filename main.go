package main

import (
	"github.com/ElPicador/form3-exercise/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Path("/isalive").HandlerFunc(handlers.PingHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
	}
	log.Println("starting server on http://localhost:3000")

	log.Fatal(srv.ListenAndServe())
}
