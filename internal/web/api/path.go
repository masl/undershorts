package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masl/undershorts/internal/db"
)

// GET shorts data
func PathEndpoint(router *mux.Router) {
	router.HandleFunc("/{path}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortPath := vars["path"]

		// Get all paths
		all, err := db.GetAllURLS()
		if err != nil {
			fmt.Println("Error while getting keys")
		}

		// Check for paths
		for k, v := range all {
			if v == shortPath {
				longUrl, err := db.GetURL(shortPath)
				if err != nil {
					continue
				}
				rw.WriteHeader(http.StatusOK)
				json.NewEncoder(rw).Encode(map[string]string{"path": shortPath, "url": longUrl})
				break
			}
			if k >= len(all)-1 {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write([]byte("this path does not exist"))
			}
		}
	}).Methods("GET")
}
