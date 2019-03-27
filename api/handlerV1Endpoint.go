package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// GET - returns JSON of a individual endpoint, by ID
// POST - updates an existing endpoint entry, by ID
func handlerV1Endpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		handlerV1EndpointGET(w, r)
	case "POST":
		return
	case "PUT":
		handlerV1EndpointPUT(w, r)
	case "DELETE":
		return
	default:
		return
	}
}

func handlerV1EndpointGET(w http.ResponseWriter, r *http.Request) {
	id := 0
	fmt.Println(r.RequestURI)
	id, _ = strconv.Atoi(r.RequestURI[len("/v1/endpoint/"):]) // ignoring error, in the event of an error empty JSON is sent
	endpointJSON := dbSelectSingleEndpoint(id)
	w.Write(endpointJSON)
}

func handlerV1EndpointPUT(w http.ResponseWriter, r *http.Request) {
	// decode POST data
	decoder := json.NewDecoder(r.Body)
	var e endpoint
	err := decoder.Decode(&e)
	if err != nil {
		log.Println("Error decoding JSON POST data:", err)
	}

	// update DB
	id := r.RequestURI[len("/v1/endpoint/"):]
	dbUpdateEndpoint(id, e)

	// send updated data
	handlerV1EndpointGET(w, r)
}
