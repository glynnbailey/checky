package main

import (
	"html/template"
	"log"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	// pull data from API
	endpoints := apiSelectAllEndpoints()

	// pull HTML template
	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		log.Println(err)
	}

	// send HTTP response
	tmpl.Execute(w, endpoints)
}

func handlerFavicon(w http.ResponseWriter, r *http.Request) {
	return
}
