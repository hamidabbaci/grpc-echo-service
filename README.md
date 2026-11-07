# gRPC Echo Service

A simple gRPC service written in Go to learn the basics of gRPC, including a server and a client.

## About

This project is a simple Echo service: the client sends a message to the server, and the server returns the same message with some extra text appended. The goal of this project is to learn how to define a service with Protocol Buffers and implement a Unary RPC call between a client and a server.

## Tech Stack

- Go
- gRPC
- Protocol Buffers (protobuf)

## Project Structure



## How to Run

### 1. Run the server

```bash
go run server/main.go
```

### 2. Run the client (in a new terminal)

```bash
go run client/main.go
```

## Sample Output
