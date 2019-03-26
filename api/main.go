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

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/v1/", v1Handler)
	http.HandleFunc("/v1/endpoint/", v1EndpointHandler)

	http.ListenAndServe(":80", nil)
}
