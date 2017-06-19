package http_handlers

import (
	"net/http"
	"time"

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
		ctx := setStartTime(r.Context(), r)
		defer hw.post(w, r.WithContext(ctx))
		hw.pre(w, r.WithContext(ctx))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (hw *HandlerWrapper) pre(w http.ResponseWriter, r *http.Request) {
	hw.log.Infof(
		"Incoming request: %s %s %s",
		r.Method,
		r.URL,
		r.RemoteAddr,
	)
}

func (hw *HandlerWrapper) post(w http.ResponseWriter, r *http.Request) {
	if p := recover(); p != nil {
		hw.log.Errorf("Critical error: %s", p)
		hw.log.Info("Recovered")
		serverError(w, r)
	}
	st := r.Context().Value("start").(time.Time)
	dur := time.Since(st).Seconds()
	hw.log.Infof(
		"Response sent: %d %fs | %s %s %s",
		1,
		dur,
		r.Method,
		r.URL,
		r.RemoteAddr,
	)
}
