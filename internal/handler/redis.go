package handler

import (
	"net/http"

	"github.com/masl/undershorts/internal/db"
)

// Parse Redis paths to http handler
func RedisHandler(urls []string, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrlMap, err := getPathUrlMap(urls)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathUrlMap, fallback), nil
}

// Maps path keys to url values
func getPathUrlMap(paths []string) (pathMap map[string]string, err error) {
	pathMap = make(map[string]string)
	for _, path := range paths {
		url, err := db.GetURL(path)
		if err != nil {
			continue
		}
		pathMap[path] = url
	}
	return
}
