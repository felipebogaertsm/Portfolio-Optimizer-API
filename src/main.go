package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber string = "8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}
}

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start web server:
	fmt.Println(fmt.Sprintf("Starting web server on port %s...", portNumber))
	_ = http.ListenAndServe(fmt.Sprintf(":%s", portNumber), nil)
}
