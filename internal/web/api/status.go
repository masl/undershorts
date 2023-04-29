package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET health
func HealthCheckEndpoint(router *mux.Router) {
	router.HandleFunc("/health", HealthCheckHandler)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: Add DB health checks
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
