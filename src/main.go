package main

import (
	"fmt"
	"net/http"
	"webapi/views"
)

const portNumber string = "8080"

// main is the main function
func main() {
	http.HandleFunc("/", views.Home)
	http.HandleFunc("/about", views.About)

	// Start web server:
	fmt.Println(fmt.Sprintf("Starting web server on port %s...", portNumber))
	_ = http.ListenAndServe(fmt.Sprintf(":%s", portNumber), nil)
}
