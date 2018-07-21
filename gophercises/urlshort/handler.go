package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrlsStruct := make([]PathUrl, 1)
	err := yaml.Unmarshal(yml, &pathToUrlsStruct)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range pathToUrlsStruct {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), err
}
