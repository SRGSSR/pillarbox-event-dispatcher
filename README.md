# Pillarbox Event Dispatcher

A Go stateless micro-service that receives JSON data via a HTTP POST request and
leverage the power of SSE to broadcast JSON data to multiple consumers.

## How to run

To run this project on your machine you need to install the Go programming language. You'll find the installation instruction on the following [link](https://go.dev/doc/install).

Once the installation completed, on your terminal, run the command below to start the HTTP server:

Once the installation completed, on your terminal run the command below to start the HTTP server:

- `go run cmd/event_dispatcher/main.go`

### Server flags

You have access to two flags:

- `port` which allows to redefine the port used by the HTTP server, default value is `:3569`
- `debug` which allows to activate the debug mode, default value is `false`

#### Example

Run the server on port `:35420` in debug mode.

- `go run cmd/event_dispatcher/main.go -port ":35420" -debug true`

### Connect a client to the server

To listen for server events you'll to connect to the SSE server. In your terminal run:

- `curl -n http://localhost:3569/metrics`

*You can create as many clients as you need by simply opening as many terminal tabs as you need and running the command shown above.*

### Send data to the server

To send JSON data to the server. In your terminal run:

- `curl -X POST http://localhost:3569/metrics -H 'Content-Type: application/json' -d "{\"msg\": \"data\", \"timestamp\": \"$EPOCHSECONDS\"}"`

## Build

- `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pillarboxEventDispatcher cmd/event_dispatcher/main.go`
