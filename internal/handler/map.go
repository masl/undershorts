package handler

import (
	"net/http"
	"strings"
)

// Map path keys from map to URL values
func MapHandler(pathsToURLS map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, exists := pathsToURLS[strings.ReplaceAll(r.RequestURI, "/", "")]
		if exists {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
