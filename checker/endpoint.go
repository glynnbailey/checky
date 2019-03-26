package main

import (
	"log"
	"net/http"
)

type endpoint struct {
	ID     int    `json: "id"`
	Name   string `json: "name"`
	URL    string `json: "url"`
	Status int    `json: "status"`
}

func (e *endpoint) checkStatus() {
	resp, err := http.Get(e.URL)
	if err != nil {
		log.Println("Error connecting to URL:", err)
	}
	defer resp.Body.Close()

	http.Post("http://api/v1/endpoint/", "application/json")
}
