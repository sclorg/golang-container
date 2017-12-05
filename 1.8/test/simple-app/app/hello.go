package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Welcome to the Go toolset.")

}

func main() {

	fmt.Println("Hello.")
	fmt.Println("Starting http server.")

	// Register handler function

	http.HandleFunc("/welcome", Welcome)

	fmt.Println("Go to localhost:8080/welcome To terminate press CTRL+C")

	// Start server

	http.ListenAndServe(":8080", nil)
}
