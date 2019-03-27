package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func apiSelectAllEndpoints() []endpoint {
	resp, err := http.Get("http://api/v1/")
	if err != nil {
		log.Println("Error accessing API:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading API HTTP response body:", err)
	}

	var endpoints []endpoint
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}

	return endpoints
}
