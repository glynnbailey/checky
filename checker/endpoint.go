package main

import (
	"log"
	"net/http"
)

type endpoint struct {
	Name   string
	URL    string
	Status int
}

func (e *endpoint) checkStatus() {
	resp, err := http.Get(e.URL)
	if err != nil {
		log.Println("Error connecting to URL:", err)
	}
	defer resp.Body.Close()
	e.Status = resp.StatusCode
}
