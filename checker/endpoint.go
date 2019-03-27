package main

import (
	"log"
	"net/http"
)

type endpoint struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	Status       int    `json:"status,omitempty"`
	ResponseTime int    `json:"responsetime,omitempty"`
}

func (e *endpoint) update() {
	resp, err := http.Get(e.URL)
	if err != nil {
		log.Println("Error connecting to URL:", err)
	}
	defer resp.Body.Close()

	http.Post("http://api/v1/endpoint/", "application/json")
}
