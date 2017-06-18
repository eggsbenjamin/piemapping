package main

import (
	"log"
	"net/http"

	"github.com/eggsbenjamin/piemapping/http_handlers"
	"github.com/eggsbenjamin/piemapping/repository"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

//	'run' command to spin up server
func run() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run",
		Run: func(cmd *cobra.Command, args []string) {
			r := mux.NewRouter()
			wrapr := http_handlers.NewHandlerWrapper(logr)
			conn := repository.NewConnection(logr, nil)
			db := repository.NewDBWrapper(conn)
			jRepo := repository.NewJourneyRepository(db, logr)
			http_handlers.Register(r, logr, jRepo)
			http.Handle("/", wrapr.Init(r))
			log.Fatal(http.ListenAndServe(":3030", nil))
		},
	}
}
