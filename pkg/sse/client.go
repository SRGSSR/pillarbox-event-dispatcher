package sse

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

type eventChannel chan string

var clients = make(map[string]eventChannel)
var clientMutex = &sync.Mutex{}

// Create a new client, assign a unique client ID and store the client in the clients map
func CreateClient() (string, eventChannel) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	clientId := uuid.NewString()
	slog.Debug("Generate new client ID", "clientId", clientId)

	if _, ok := clients[clientId]; !ok {
		slog.Debug("Create new client", "clientId", clientId)
		clients[clientId] = make(eventChannel)
	}

	return clientId, clients[clientId]
}

// Close the client connection and the client from the clients map
func CloseClient(clientId string) {
	client, ok := clients[clientId]

	if !ok {
		slog.Warn("Try to close a client that doesn't exist", "clientId", clientId)
		return
	}

	slog.Debug("Close client with ID", "clientId", clientId)

	close(client)
	client = nil
	delete(clients, clientId)
}

// Broadcast metrics data to connected clients
func Broadcast(data string) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	for id, client := range clients {
		slog.Debug("Write data to client", "clientId", id)
		client <- data
	}
}
