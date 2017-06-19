package http_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eggsbenjamin/piemapping/commons"
	"github.com/eggsbenjamin/piemapping/repository"
	"github.com/gorilla/mux"
)

type DriverAvailabilityHandler struct {
	jRepo repository.JourneyRepositor
	log   commons.LevelledLogWriter
}

//	constructor
func NewDriverAvailabilityHandler(jRepo repository.JourneyRepositor, llw commons.LevelledLogWriter) *DriverAvailabilityHandler {
	return &DriverAvailabilityHandler{
		jRepo: jRepo,
		log:   llw,
	}
}

func (h *DriverAvailabilityHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["driverId"]
	jrnys, err := h.jRepo.GetByDriverAvailability(id)
	if err != nil {
		h.log.Errorf("Repo error: %s", err.Error())
		serverError(w, r)
		return
	}
	body, err := json.Marshal(jrnys)
	if err != nil {
		h.log.Errorf("Marshaling error: %s", err.Error())
		serverError(w, r)
		return
	}
	ok(w, r, body)
}

type DepLocWithinRangeHandler struct {
	jRepo repository.JourneyRepositor
	log   commons.LevelledLogWriter
}

//	constructor
func NewDepLocWithinRangeHandler(jRepo repository.JourneyRepositor, llw commons.LevelledLogWriter) *DepLocWithinRangeHandler {
	return &DepLocWithinRangeHandler{
		jRepo: jRepo,
		log:   llw,
	}
}

func (h *DepLocWithinRangeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	depLoc := vars["depLoc"]
	st := vars["start"]
	end := vars["end"]
	jrnys, err := h.jRepo.GetByDepLocWithinRange(depLoc, st, end)
	if err != nil {
		h.log.Errorf("Repo error: %s", err.Error())
		serverError(w, r)
		return
	}
	body, err := json.Marshal(jrnys)
	if err != nil {
		h.log.Errorf("Marshaling error: %s", err.Error())
		serverError(w, r)
		return
	}
	ok(w, r, body)
}
