package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TODO auth
// TODO move all DB access to the API service

func main() {
	fmt.Println("API Service Starting")

	// setup DB tables if they dont exist
	dbInit()
	fmt.Println("DB initialized")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/v1/", v1Handler)
	http.HandleFunc("/v1/endpoint/", v1EndpointHandler)

	http.ListenAndServe(":80", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v1/", 301)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func v1Handler(w http.ResponseWriter, r *http.Request) {
	endpoints := dbSelectAllEndpoints()
	endpointsJSON, err := json.Marshal(endpoints)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(endpointsJSON)
}

func v1EndpointHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := dbSelectSingleEndpoints(r.RequestURI[len("/v1/endpoint/"):])
	endpointJSON, err := json.Marshal(endpoint)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(endpointJSON)
}
