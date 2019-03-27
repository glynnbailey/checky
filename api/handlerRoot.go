package main

import (
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	// TODO not this
	http.Redirect(w, r, "/v1/", 301)
}

func handlerFavicon(w http.ResponseWriter, r *http.Request) {
	return
}
