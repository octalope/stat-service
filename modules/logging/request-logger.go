package logging

import (
	"github.com/rs/zerolog/hlog"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	log := Get()

	h := hlog.NewHandler(log)

	accessHandler := hlog.AccessHandler(
		func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status", status).
				Int("response_size", size).
				Dur("elapsed_ms", duration).
				Msg("incoming request")
		},
	)

	userAgentHandler := hlog.UserAgentHandler("http_user_agent")

	return h(accessHandler(userAgentHandler(next)))
}
