package handlers

import "net/http"

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.HandleFunc("/lsf", LsfHandler)
}
