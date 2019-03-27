package main

import "net/http"

// returns JSON of all endpoints
func handlerV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlerV1GET(w, r)
	default:
		// TODO send "resource not available" or relevant http code or whatever
	}
}

func handlerV1GET(w http.ResponseWriter, r *http.Request) {
	endpointsJSON := dbSelectAllEndpoints()
	w.Header().Set("Content-Type", "application/json")
	w.Write(endpointsJSON)
}
