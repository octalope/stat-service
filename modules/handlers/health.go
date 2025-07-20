package handlers

import (
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	if r.Method == http.MethodGet {
		io.WriteString(w, "OK")
	} else {
		log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
