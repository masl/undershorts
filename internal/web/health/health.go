package health

import "net/http"

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add DB health checks
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}
