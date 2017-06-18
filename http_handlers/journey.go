package http_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eggsbenjamin/piemapping/commons"
	"github.com/eggsbenjamin/piemapping/repository"
	"github.com/gorilla/mux"
)

type DriverAvailablityHandler struct {
	jRepo repository.JourneyRepositor
	log   commons.LevelledLogWriter
}

//	constructor
func NewDriverAvailabilityHandler(jRepo repository.JourneyRepositor, llw commons.LevelledLogWriter) *DriverAvailablityHandler {
	return &DriverAvailablityHandler{
		jRepo: jRepo,
		log:   llw,
	}
}

func (h *DriverAvailablityHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["driverId"]
	jrnys, err := h.jRepo.GetByDriverAvailability(id)
	if err != nil {
		h.log.Errorf("Repo error: %s", err.Error())
		serverError(w)
		return
	}
	body, err := json.Marshal(jrnys)
	if err != nil {
		h.log.Errorf("Marshaling error: %s", err.Error())
		serverError(w)
		return
	}
	ok(w, body)
}
