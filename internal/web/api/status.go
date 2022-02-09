package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET status
func StatusEndpoint(router *mux.Router) {
	router.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(rw).Encode(map[string]bool{"ok": true})
	}).Methods("GET")
}
