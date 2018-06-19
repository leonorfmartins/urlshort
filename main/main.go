package main

import (
	"fmt"
	"net/http"
	"flag"
	"io/ioutil"
	"github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()
	yaml := flag.String("yaml", "../url.yaml", "Yaml file with ulrs")
	json := flag.String("json", "../url.json", "Url short in json format")
	flag.Parse()

	yamlContent, err := ioutil.ReadFile(*yaml)
	if err != nil {
		panic(err)
	}
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlContent, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
