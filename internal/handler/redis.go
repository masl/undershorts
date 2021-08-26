package handler

import (
	"net/http"

	"github.com/masl/undershorts/internal/db"
)

// Parse Redis paths to http handler
func RedisHandler(urls []string, fallback http.Handler) (http.HandlerFunc, error) {
	pathMap, err := getPaths(urls)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathMap, fallback), nil
}

// Maps path keys to url values
func getPaths(urls []string) (pathMap map[string]string, err error) {
	pathMap = make(map[string]string)
	for _, v := range urls {
		path, err := db.GetURL(db.RedisClient, v)
		if err != nil {
			continue
		}
		pathMap[v] = path
	}
	return
}
