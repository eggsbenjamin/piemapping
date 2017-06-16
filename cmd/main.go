package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:3030",
	}

	log.Fatal(srv.ListenAndServe())
}
