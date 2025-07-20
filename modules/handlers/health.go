package handlers

import (
	"io"
	"net/http"

	"github.com/octalope/stat-service/modules/util"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log := util.Log()
	log.Info().
		Str("verb", r.Method).
		Str("path", "/health").
		Msg("request")

	if r.Method == http.MethodGet {
		io.WriteString(w, "OK")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
