package short

import (
	"encoding/json"
	"net/http"

	"github.com/masl/undershorts/internal/db"
)

func Handle(postgres *db.PostgresClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.PathValue("shortURL")

		short, err := postgres.GetShortByShortURL(shortURL)
		if err != nil {
			// TODO: handle different error types
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// TODO: filter what should be in the response

		responseBytes, err := json.Marshal(short)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	}
}
