package main

import (
	"flag"
	"log/slog"
	"net/http"

	"github.com/srgssr/pillarbox-event-dispatcher/api/handler"
)

func main() {
	// Command-line available flags
	addr := flag.String("port", ":3569", "HTTP server port")
	debug := flag.Bool("debug", false, "active debug logging, default value false")

	flag.Parse()

	// Set the debug mode if true
	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Debug mode activated")
	}

	slog.Info("Server start", "PORT", *addr)

	// HTTP request multiplexer
	serveMux := http.NewServeMux()
	// HTTP server parameters
	metricsServer := &http.Server{
		Addr:    *addr,
		Handler: serveMux,
	}

	// Endpoint used by user to send data
	serveMux.HandleFunc("POST /metrics", handler.PostMetrics)
	// Endpoint used by clients to connect to the SSE and consume data
	serveMux.HandleFunc("GET /metrics", handler.GetMetrics)
	// Endpoint to check the service health
	serveMux.HandleFunc("GET /health", handler.Health)

	// Run the HTTP server
	if err := metricsServer.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
