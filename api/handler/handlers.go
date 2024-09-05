package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"

	"github.com/srgssr/pillarbox-event-dispatcher/pkg/sse"
)

// Metrics accepts any JSON payload entry and forwards it to all connected clients.
//
// This service acts as a passthrough, with the only check being the that the payload must be valid JSON.
// If the payload is not a valid JSON, an error response is returned.
func Metrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !json.Valid(body) {
		err = fmt.Errorf("Body (%v) is not valid JSON", body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Forward the JSON payload to connected clients
	sse.Broadcast(string(body))
}

// EventDispatcher enables clients to listen for events via Server-Sent Events.
//
// This endpoint creates a client connection and streams data to connected clients.
//
// When a client disconnects, the associated client ID is closed and the associated channel is distroyed.
func EventDispatcher(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientId, clientChannel := sse.CreateClient()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		select {
		case data := <-clientChannel:
			fmt.Fprintf(w, "data: %s\n\n", data)

			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			// Client disconnected
			sse.CloseClient(clientId)

			return
		}
	}
}

// The Health endpoint returns a JSON payload containing memory and goroutine statistics. This will allows to monitore the healthiness of the service.
func Health(w http.ResponseWriter, r *http.Request) {
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)

	data := map[string]any{
		"message":   "PONG",
		"timestamp": time.Now().UnixMilli(),
		// Alloc is bytes of allocated heap objects.
		"alloc": fmt.Sprintf("%v MiB", memStats.Alloc/1024/1024),
		// The number of live objects is Mallocs - Frees.
		"live_objects": fmt.Sprintf("%v", memStats.Mallocs-memStats.Frees),
		// NumGC is the number of completed GC cycles.
		"num_gc": fmt.Sprintf("%v", memStats.NumGC),
		// NumGoroutine returns the number of goroutines that currently exist.
		"num_goroutines": fmt.Sprintf("%v", runtime.NumGoroutine()),
		// stop-the-world pause since the program started.
		"pause_total_ns": fmt.Sprintf("%v", memStats.PauseTotalNs),
		// Sys is the total bytes of memory obtained from the OS.
		"sys": fmt.Sprintf("%v MiB", memStats.Sys/1024/1024),
		// TotalAlloc is cumulative bytes allocated for heap objects.
		// Will never decrease
		"total_alloc": fmt.Sprintf("%v MiB", memStats.TotalAlloc/1024/1024),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
