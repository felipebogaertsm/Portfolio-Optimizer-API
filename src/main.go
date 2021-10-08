package main

import (
	"fmt"
	"net/http"
)

func main() {
	var sum float64

	for i := 0; i < 1000; i++ {
		sum = sum + float64(i)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, fmt.Sprintf("%.2f", sum))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// Start web server:
	fmt.Println("Starting web server...")
	http.ListenAndServe(":8080", nil)
}
