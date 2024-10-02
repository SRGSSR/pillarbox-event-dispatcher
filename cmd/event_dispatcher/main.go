package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/srgssr/pillarbox-event-dispatcher/api/handler"
)

func main() {
	// Command-line available flags
	addr := flag.String("port", ":8080", "HTTP server port")

	flag.Parse()

	// HTTP request multiplexer
	serveMux := http.NewServeMux()
	// HTTP server parameters
	dispatcherServer := &http.Server{
		Addr:    *addr,
		Handler: serveMux,
	}

	// Endpoint used by user to send data
	serveMux.HandleFunc("/api/events", handler.EventReceiver)
	// Endpoint used by clients to connect to the SSE and consume data
	serveMux.HandleFunc("/events", handler.EventDispatcher)
	// Endpoint to check the service health
	serveMux.HandleFunc("/health", handler.Health)

	// Run the HTTP server
	if err := dispatcherServer.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
