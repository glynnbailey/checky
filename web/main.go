package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Web Service Starting")

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/favicon.ico", handlerFavicon)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":80", nil)
}
