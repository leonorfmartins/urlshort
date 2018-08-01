package main

import (
	"fmt"
	"net/http"
	"flag"
	"io/ioutil"
	"github.com/gophercises/urlshort"
)

func main() {
	yamlPath := flag.String("yaml", "../url.yaml", "Yaml file with ulrs")
	jsonPath := flag.String("json", "", "Path to json file")
	flag.Parse()

	fmt.Println("Starting the server on :8080")
	mux := handleDataSource(*jsonPath, *yamlPath)
	http.ListenAndServe(":8080", mux)
}

func handleDataSource(jsonPath string, yamlPath string) http.Handler {
	mux := defaultMux()
	pathsToUrls := urlshort.MappedUrl{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	if jsonPath != "" {
		return handleJson(jsonPath, mapHandler)
	}
	return handleYaml(yamlPath, mapHandler)
}

func handleYaml(yamlPath string, mapHandler http.HandlerFunc) http.Handler {
	yamlContent, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(yamlContent, mapHandler)
	if err != nil {
		panic(err)
	}
	return yamlHandler
}

func handleJson(jsonPath string, mapHandler http.HandlerFunc) http.Handler {
	jsonFile, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.YAMLHandler(jsonFile, mapHandler)
	if err != nil {
		panic(err)
	}
	return jsonHandler
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
