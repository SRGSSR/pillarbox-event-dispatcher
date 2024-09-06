# Pillarbox Event Dispatcher

A Go stateless micro-service that receives JSON data via a HTTP POST request and
leverage the power of SSE to broadcast JSON data to multiple consumers.

The goal of this service is to receive events from Pillarbox players to feed into tools that can address a wide range of use cases, such as providing a general overview of the health of our offering, or helping diagnose potential problems.

> [!IMPORTANT]
>  What this service doesn't do
>
> - It does not store events, even temporarily.
> - None of the data is critical, so there's no mechanism for resend events that haven't been received *(if the data is being used to monitor products you consider critical, it's probably time to monitor those products instead)*.

## Available URLs

To send events from the player, use the URL:

- `https://zdkimhgwhh.eu-central-1.awsapprunner.com/metrics`

  Expects a `POST` request with a valid JSON payload

To listen to events to feed any tool, use the URL:

- `https://zdkimhgwhh.eu-central-1.awsapprunner.com/event-dispatcher`

To check health status, use URL:

- `https://zdkimhgwhh.eu-central-1.awsapprunner.com/health`

## How to run locally

To run this project on your machine you need to install the Go programming language. You'll find the installation instruction on the following [link](https://go.dev/doc/install).

Once the installation completed, on your terminal, run the command below to start the HTTP server:

- `go run cmd/event_dispatcher/main.go`

### Receive and send data from and to the server

To receive data, you need to connect to the SSE server. To do this, in your terminal run:

- `curl -n http://localhost:3569/event-dispatcher`

*You can create as many clients as you need by simply opening as many terminal tabs as you need and running the command shown above.*

To send data to the server. In your terminal run:

- `curl -X POST http://localhost:3569/metrics -H 'Content-Type: application/json' -d "{\"msg\": \"data\", \"timestamp\": \"$EPOCHSECONDS\"}"`

## Application configuration

This application allows to customize the port on which it runs using the flag below:

- `port` which allows to redefine the port used by the HTTP server, default value is `:3569`

### Example

Change the default port on which the application runs:

- `go run cmd/event_dispatcher/main.go -port ":35420"`

## Compilation

The [Go](https://go.dev/) language makes it easy to [produce binaries](https://go.dev/doc/tutorial/compile-install) compatible with a multitude of operating systems and processor architectures. The advantage is that everything is self-contained, so there's no need to install runtimes, unlike java, javascript, php etc...

To display the list of available operating systems and processor architectures:

- `go tool dist list`

To display the current type of operating system and processor architecture:

- `go env GOOS GOARCH`

To produce a MacOS ARM64-compatible binary from Linux:

- `CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o pillarboxEventDispatcher cmd/event_dispatcher/main.go`

*`CGO_ENABLED=0` means that the binary produced is statically-linked and needs no external dependencies to run.*
