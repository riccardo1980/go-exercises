package urlshort

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func HelpMe() string {
	return "me"
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	myHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		url := pathsToUrls[r.URL.Path]

		if url != "" {
			log.Printf("Redirecting %s to %s", r.URL.Path, url)
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
	return myHandlerFunc
}

type pathToURL struct {
	Path string
	URL  string
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

	log.Println("Reading YAML")

	var paths []pathToURL
	err := yaml.Unmarshal(yml, &paths)

	pathsToUrls := make(map[string]string)

	for _, path := range paths {
		pathsToUrls[path.Path] = path.URL
	}
	return MapHandler(pathsToUrls, fallback), err
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//	[
// 		{
// 			"path" : "/ansa",
// 			"url": "https://ansa.it"
// 		}
// 	]
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {

	log.Println("Reading JSON")
	var paths []pathToURL
	err := json.Unmarshal(jsn, &paths)

	pathsToUrls := make(map[string]string)

	for _, path := range paths {
		pathsToUrls[path.Path] = path.URL
	}

	return MapHandler(pathsToUrls, fallback), err

}
