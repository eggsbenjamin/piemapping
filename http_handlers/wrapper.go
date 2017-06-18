package http_handlers

import (
	"net/http"

	"github.com/eggsbenjamin/piemapping/commons"
)

type HandlerWrapper struct {
	log commons.LevelledLogWriter
}

//	constructor
func NewHandlerWrapper(llw commons.LevelledLogWriter) *HandlerWrapper {
	return &HandlerWrapper{
		log: llw,
	}
}

func (hw *HandlerWrapper) Init(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer hw.post(w, r)
		hw.pre(w, r)
		h.ServeHTTP(w, r)
	})
}

func (hw *HandlerWrapper) pre(w http.ResponseWriter, r *http.Request) {
	hw.log.Infof("Incoming request: %s '%s'", r.Method, r.URL)
}

func (hw *HandlerWrapper) post(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		hw.log.Errorf("Critical error: %s", r)
		hw.log.Info("Recovered")
		serverError(w)
	}
	hw.log.Infof("Request Complete")
}
