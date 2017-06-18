package http_handlers

import (
	"fmt"
	"net/http"
)

//	send a 500 with a generic message
func serverError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `{ "message": "internal server error" }`)
}

//	send a 200 with a response body
func ok(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
