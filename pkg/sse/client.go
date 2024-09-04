package sse

import (
	"log"
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
		log.Println("Try to close a client that doesn't exist", "clientId", clientId)
		return
	}

	close(client)
	client = nil
	delete(clients, clientId)
}

// Broadcast metrics data to connected clients
func Broadcast(data string) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	for _, client := range clients {
		client <- data
	}
}
