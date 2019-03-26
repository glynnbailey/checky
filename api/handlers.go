package main

import (
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
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

// GET - returns JSON of a individual endpoint by ID
// POST - updates a individual endpoint by ID
func v1EndpointHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		endpointJSON := dbSelectSingleEndpoint(r.RequestURI[len("/v1/endpoint/"):])
		w.Header().Set("Content-Type", "application/json")
		w.Write(endpointJSON)

	case "POST":
		return
	}
}
