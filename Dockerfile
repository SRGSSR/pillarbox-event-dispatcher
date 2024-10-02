# First Stage: Build the Go application
FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build -o main cmd/event_dispatcher/main.go

# Second Stage: Create the runtime image
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/main .

# Expose the port
EXPOSE 8080

# Set default values if environment variables are not provided
CMD ["sh", "-c", "./main"]
