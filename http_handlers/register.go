package http_handlers

import (
	"github.com/eggsbenjamin/piemapping/commons"
	"github.com/eggsbenjamin/piemapping/repository"
	"github.com/gorilla/mux"
)

//	http handler registration
func Register(r *mux.Router, llw commons.LevelledLogWriter, jRepo repository.JourneyRepositor) {
	driAv := NewDriverAvailabilityHandler(jRepo, llw)
	dLWR := NewDepLocWithinRangeHandler(jRepo, llw)

	r.HandleFunc("/journeys/driver/{driverId}/availability", driAv.Handle)
	r.HandleFunc("/journeys/start/{depLoc}/range/{start}/to/{end}", dLWR.Handle)
}
