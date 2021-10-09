package main

import (
	"fmt"
	"net/http"
)

const portNumber string = "8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	var sum float64

	for i := 0; i < 1000; i++ {
		sum = sum + float64(i)
	}

	n, err := fmt.Fprintf(w, fmt.Sprintf("The sum is equal to: %.2f", sum))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start web server:
	fmt.Println(fmt.Sprintf("Starting web server on port %s...", portNumber))
	_ = http.ListenAndServe(fmt.Sprintf(":%s", portNumber), nil)
}
