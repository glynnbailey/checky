package main

type endpoint struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status int    `json:"status"`
}
