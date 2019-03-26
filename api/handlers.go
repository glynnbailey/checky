package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// TODO not this
	http.Redirect(w, r, "/v1/", 301)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// returns JSON of all endpoints
func v1Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		endpointsJSON := dbSelectAllEndpoints()
		w.Header().Set("Content-Type", "application/json")
		w.Write(endpointsJSON)
	}
}

// GET - returns JSON of a individual endpoint, by ID
// POST - updates an existing endpoint entry, by ID
func v1EndpointHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		id := 0
		id, _ = strconv.Atoi(r.RequestURI[len("/v1/endpoint/"):]) // ignoring error, in the event of an error empty JSON is sent
		endpointJSON := dbSelectSingleEndpoint(id)
		w.Write(endpointJSON)

	case "POST":
		// decode POST data
		decoder := json.NewDecoder(r.Body)
		var e endpoint
		err := decoder.Decode(&e)
		if err != nil {
			log.Println("Error decoding JSON POST data:", err)
		}

		// update DB
		dbUpdateEndpoint(e)

		// send updated data
		id := 0
		id, _ = strconv.Atoi(r.RequestURI[len("/v1/endpoint/"):]) // ignoring error, in the event of an error empty JSON is sent
		endpointJSON := dbSelectSingleEndpoint(id)
		w.Write(endpointJSON)
	}
}
