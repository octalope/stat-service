package main

import (
	"fmt"
	"net/http"

	"github.com/octalope/stat-service/modules/handlers"
	"github.com/octalope/stat-service/modules/util"
)

const port uint16 = 8080

func main() {
	log := util.Log()
	log.Info().Msg("Starting stat-service")

	mux := http.NewServeMux()
	handlers.Register(mux)

	log.Info().
		Str("port", fmt.Sprintf("%d", port)).
		Msg("listening on port")

	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
