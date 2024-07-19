package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./templates")))

	// Start the HTTP server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
