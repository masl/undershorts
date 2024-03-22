package shorten

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/masl/undershorts/internal/db"
)

func Handle(postgres *db.PostgresClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: auth middleware

		requestBytes, err := io.ReadAll(io.LimitReader(r.Body, 4096))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = postgres.AddURL(requestBody.ShortURL, requestBody.LongURL)
		if err != nil {
			// TODO: more meaningful messages
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
