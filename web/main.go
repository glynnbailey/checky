package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("API Service Starting")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":80", nil)
}
