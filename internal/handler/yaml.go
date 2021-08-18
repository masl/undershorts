package handler

import (
	"net/http"

	"github.com/go-yaml/yaml"
)

type Y []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)
	if err != nil {
		return nil, err
	}

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) (y Y, err error) {
	err = yaml.Unmarshal(yml, &y)
	if err != nil {
		return
	}

	return
}

func buildMap(yml Y) (pathMap map[string]string) {
	pathMap = make(map[string]string)
	for _, v := range yml {
		pathMap[v.Path] = v.URL
	}

	return
}
