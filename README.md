# Key-Value Store

This repository contains a simple key-value store (KVS) implementation in Go. The KVS includes TCP, HTTP, and gRPC server implementations for interacting with the store. Note that the current implementation is not distributed and operates as a single-instance, in-memory store.

## Overview

The key-value store supports the following operations:

- **GET**: Retrieve the value associated with a given key.
- **SET**: Store a value with a given key.
- **DELETE**: Remove a key-value pair from the store.

### Features

- **In-Memory Storage**: The KVS stores data in memory and does not persist data across server restarts.
- **Single-Instance**: The current implementation operates as a single instance with no built-in support for distributed or replicated storage.

## Running the Project

To run each server or client, navigate to the respective command directory and use `go run`:

```bash
cd cmd/tcpserver
go run main.go
```

```bash
go run cmd/tcpclient/main.go
```

## NOTE

The project is designed for personal learning and experimentation. It demonstrates basic networking and API design in Go.
