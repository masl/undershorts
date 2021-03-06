package handler

import (
	"net/http"

	"github.com/go-yaml/yaml"
)

type Y []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// Parse YAML files to http handler which maps path keys to url values
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYAML)
	if err != nil {
		return nil, err
	}

	return MapHandler(pathMap, fallback), nil
}

// Parse YAML into struct
func parseYAML(yml []byte) (y Y, err error) {
	err = yaml.Unmarshal(yml, &y)
	if err != nil {
		return
	}

	return
}

// Create map from parsed YAML
func buildMap(yml Y) (pathMap map[string]string) {
	pathMap = make(map[string]string)
	for _, v := range yml {
		pathMap[v.Path] = v.URL
	}

	return
}
