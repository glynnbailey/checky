package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type endpoint struct {
	ID           int           `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	URL          string        `json:"url,omitempty"`
	Status       int           `json:"status,omitempty"`
	ResponseTime time.Duration `json:"responsetime,omitempty"`
}

func (e *endpoint) update() {
	// GET request to check its availability
	startTime := time.Now()
	resp, err := http.Get(e.URL)
	elapsedTime := time.Since(startTime)
	if err != nil {
		log.Println("Error connecting to URL:", err)
	}
	defer resp.Body.Close()

	// update the endpoint data
	e.Status = resp.StatusCode
	e.ResponseTime = elapsedTime

	// marshal json to send to api
	endpointJSON, err := json.Marshal(e)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}

	// update the db via the API
	client := &http.Client{}
	reader := bytes.NewReader(endpointJSON)

	req, err := http.NewRequest("PUT", "http://api/v1/endpoint/"+strconv.Itoa(e.ID), reader)
	if err != nil {
		log.Println("Error creating HTTP PUT request:", err)
	}

	_, err = client.Do(req)
	if err != nil {
		log.Println("Error sending HTTP PUT request:", err)
	}
}
