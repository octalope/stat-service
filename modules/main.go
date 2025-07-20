package main

import (
	"fmt"
	"github.com/octalope/stat-service/modules/handlers"
	"github.com/octalope/stat-service/modules/logging"
	"net/http"
)

const port uint16 = 8080
const applicationName = "stat-service"

func main() {
	log := logging.Get()
	log.Info().Msgf("Starting %v", applicationName)

	mux := http.NewServeMux()
	handlers.Register(mux)

	log.Info().
		Msgf("listening on port %v", port)

	log.Fatal().
		Err(http.ListenAndServe(fmt.Sprintf(":%v", port), logging.RequestLogger(mux))).
		Msg(fmt.Sprintf("%v terminated", applicationName))
}
