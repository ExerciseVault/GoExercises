package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler for the root URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //used to construct the http response and contains info about incoming request
		// Set response content type to HTML
		w.Header().Set("Content-Type", "text/html")

		// Write HTML content
		fmt.Fprintln(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Welcome</title>
			</head>
			<body>
				<h1>Hello from Go!</h1>
			</body>
			</html>
		`)
	})

	// Start the server on port 8080
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil) //Indicates that the default ServeMux should be used to handle incoming requests.
}
