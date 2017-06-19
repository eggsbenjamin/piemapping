package http_handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//	send a 500 with a generic message
func serverError(w http.ResponseWriter, r *http.Request) {
	body := `{ "message": "internal server error" }`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, body)
}

//	send a 200 with a response body
func ok(w http.ResponseWriter, r *http.Request, raw []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

//	sets the start time for a request and returns a context containing it
func setStartTime(ctx context.Context, r *http.Request) context.Context {
	n := time.Now()
	return context.WithValue(ctx, "start", n)
}
