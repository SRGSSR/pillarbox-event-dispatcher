package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/srgssr/pillarbox-event-dispatcher/api/handler"
)

func main() {
	// Command-line available flags
	addr := flag.String("port", ":3569", "HTTP server port")

	flag.Parse()

	// HTTP request multiplexer
	serveMux := http.NewServeMux()
	// HTTP server parameters
	metricsServer := &http.Server{
		Addr:    *addr,
		Handler: serveMux,
	}

	// Endpoint used by user to send data and consumers to receive data
	serveMux.HandleFunc("/metrics", handler.Metrics)
	// Endpoint to check the service health
	serveMux.HandleFunc("/health", handler.Health)

	// Run the HTTP server
	if err := metricsServer.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
