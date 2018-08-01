package urlshort

import (
	"net/http"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to ther corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls MappedUrl, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements90 http.Handler)
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

type MappedUrl map[string]string

func YAMLHandler(yml []byte, fallback http.Handler) (http.Handler, error) {
	var l []MappedUrl
	err := yaml.Unmarshal(yml, &l)
	if err != nil {
		return nil, err
	}
	m := make(MappedUrl)
	for _, v := range l {
		m[v["path"]] = v["url"]
	} 
	return MapHandler(m, fallback), nil
}

func JSONHandler(urls []byte, fallback http.Handler) (http.Handler, error) {
	var m []MappedUrl
	err := json.Unmarshal(urls, &m)
	if err != nil {
		return nil, err
	}
	u := make(MappedUrl)
	for _, v := range m {
		u[v["path"]] = v["url"]
	}
	
	return MapHandler(u, fallback), nil
}