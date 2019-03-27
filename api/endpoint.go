package main

import "time"

type endpoint struct {
	ID           int           `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	URL          string        `json:"url,omitempty"`
	Status       int           `json:"status,omitempty"`
	ResponseTime time.Duration `json:"responsetime,omitempty"`
}
