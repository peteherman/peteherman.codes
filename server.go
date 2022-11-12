package main

import (
	"fmt"
	"log"
	"net/http"
)

const SERVER_PORT = 8080

func main() {
	// Create new base Mux
	mux := http.NewServeMux()

	// Add handlers
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// Add logger middleware
	loggerMux := NewLoggerMiddleware(mux)

	// Configure webserver uri
	server_uri := fmt.Sprintf(":%v", SERVER_PORT)
	
	// Start the webserver	
	log.Println("Starting Web Server")	
	if err := http.ListenAndServe(server_uri, loggerMux); err != nil {
		log.Fatal(err)
	}
}
