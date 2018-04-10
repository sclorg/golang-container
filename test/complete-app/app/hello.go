package main

import (
	"complete-app/welcome"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Hello.")
	fmt.Println("Starting http server.")

	// Register handler function

	http.HandleFunc("/welcome", welcome.Welcome)

	fmt.Println("Go to localhost:8080/welcome To terminate press CTRL+C")

	// Start server

	http.ListenAndServe(":8080", nil)
}
