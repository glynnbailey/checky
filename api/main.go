package main

import (
	"fmt"
	"net/http"
)

// TODO auth
// TODO move all DB access to the API service

func main() {
	fmt.Println("API Service Starting")

	// setup DB tables if they dont exist
	dbInit()

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/favicon.ico", handlerFavicon)
	http.HandleFunc("/v1/", handlerV1)
	http.HandleFunc("/v1/endpoint/", handlerV1Endpoint)

	http.ListenAndServe(":80", nil)
}
