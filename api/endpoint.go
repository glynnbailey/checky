package main

type endpoint struct {
	ID           int    `json:"name"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Status       int    `json:"status"`
	ResponseTime int    `json: "responsetime"`
}
