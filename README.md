# golearn

A personal repository for learning Go (Golang), one topic at a time. Each folder is a self-contained, runnable example focused on a single concept.

## Getting Started

### Prerequisites

- Go (see `go.mod` for the required version)

### Running an example

Every topic lives in its own folder with a `main.go`. Run any example directly with `go run`:

```sh
go run ./basics/hello-world
go run ./networking/http-server
```

Some examples are designed to work together. For instance, the HTTP server and client:

```sh
# terminal 1 — start the server
go run ./networking/http-server

# terminal 2 — call it with the client
go run ./networking/http-client
```

## Structure

| Topic | Status | Description |
|---|---|---|
| `basics/hello-world` | ✅ | The classic first Go program |
| `basics/variables` | ✅ | Declaring variables and basic types |
| `basics/functions` | ✅ | Defining and calling functions |
| `basics/control-flow` | 🚧 | if/else, switch, loops |
| `data-structures/slices` | 🚧 | Slices |
| `data-structures/maps` | 🚧 | Maps |
| `data-structures/structs` | 🚧 | Structs |
| `interfaces-and-errors/interfaces` | 🚧 | Interfaces |
| `interfaces-and-errors/error-handling` | 🚧 | Error handling patterns |
| `interfaces-and-errors/custom-errors` | 🚧 | Custom error types |
| `concurrency/goroutines` | 🚧 | Goroutines |
| `concurrency/channels` | 🚧 | Channels |
| `concurrency/waitgroups` | 🚧 | `sync.WaitGroup` |
| `networking/http-server` | ✅ | Basic HTTP server: routing, JSON responses, graceful timeouts |
| `networking/http-client` | ✅ | Basic HTTP client: requests, JSON decoding, error handling |
| `projects` | 🚧 | Larger, combined projects |

✅ implemented &nbsp;&nbsp; 🚧 planned / in progress

## Highlights

### `networking/http-server`

A minimal HTTP server built with the standard library `net/http` package, demonstrating:

- Routing with `http.ServeMux`
- Plain text and JSON responses
- Configuring `http.Server` with read/write timeouts
- Graceful handling of `http.ErrServerClosed`

### `networking/http-client`

A minimal HTTP client demonstrating:

- Using `http.Client` with a timeout
- Decoding JSON responses into structs
- Wrapping errors with `fmt.Errorf` and `%w`
- Separating request logic from presentation (a dedicated function does the work; `main` handles output)

## Module

```
module github.com/NamX1/golearn
```
