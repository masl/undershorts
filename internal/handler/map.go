package handler

import "net/http"

// Map path keys from map to URL values
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		longPath, exists := pathsToUrls[r.RequestURI]
		if exists {
			http.Redirect(w, r, longPath, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
