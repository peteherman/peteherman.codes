package main

import (
	"log"
	"net/http"
)

func main() {
	// Create new base Mux
	mux := http.NewServeMux()

	// Add handlers
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// Add logger middleware
	loggerMux := NewLoggerMiddleware(mux)

	// Start the webserver
	log.Println("Starting Web Server")
	if err := http.ListenAndServe(":8080", loggerMux); err != nil {
		log.Fatal(err)
	}
}
