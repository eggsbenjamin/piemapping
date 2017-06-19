package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eggsbenjamin/piemapping/http_handlers"
	"github.com/eggsbenjamin/piemapping/repository"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//	'run' command to spin up server
func run() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run",
		Run: func(cmd *cobra.Command, args []string) {
			r := mux.NewRouter()
			conn := repository.NewConnection(logr, nil)
			defer conn.Close()
			db := repository.NewDBWrapper(conn)
			jRepo := repository.NewJourneyRepository(db, logr)
			http_handlers.Register(r, logr, jRepo)
			initServer(r)
		},
	}
}

//	initialise the server
func initServer(r *mux.Router) {
	wrapr := http_handlers.NewHandlerWrapper(logr)
	port := viper.GetString("port")
	addr := fmt.Sprintf(":%s", port)
	hdlr := wrapr.Init(r)
	srv := &http.Server{
		Addr:    addr,
		Handler: hdlr,
	}
	logr.Infof("Server listening on port: '%s'", port)
	log.Fatal(srv.ListenAndServe())
}
