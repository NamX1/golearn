# golearn

A personal repository for learning Go, one topic at a time. Each folder is a self-contained, runnable example built around a single concept, and they all share one running example: a small lending library.

## Getting Started

### Prerequisites

- Go (see `go.mod` for the required version)

### Running an example

Every topic lives in its own folder with a `main.go`. Run any example directly with `go run`:

```sh
go run ./basics/hello-world
go run ./basics/variables-and-types
go run ./stdlib/net/http/server
```

Some examples are designed to work together. For instance, the catalogue server and its client:

```sh
# terminal 1 — start the server
go run ./stdlib/net/http/server

# terminal 2 — call it with the client
go run ./stdlib/net/http/client
```

## Layout

| Path | Purpose |
|---|---|
| `basics/` | The language itself: syntax, types, control flow |
| `data-structures/` | The built-in composite types |
| `interfaces-and-errors/` | Behaviour and failure |
| `concurrency/` | Goroutines, channels, and synchronisation |
| `stdlib/` | One folder per standard-library area, mirroring the real import path |
| `projects/` | Larger programs that combine several topics |

`stdlib/` paths deliberately match Go's own package paths, so `stdlib/net/http/server` is the server built from `net/http`.

## Progress

| Topic | Status | Description |
|---|---|---|
| `basics/hello-world` | ✅ | The classic first Go program |
| `basics/variables-and-types` | ✅ | Every basic type, `var` / `:=` / `const`, zero values |
| `basics/functions` | ✅ | Parameters and multiple return values |
| `basics/control-flow` | ✅ | if/else, switch, for, range, nested loops |
| `data-structures/slices` | 🚧 | Slices |
| `data-structures/maps` | 🚧 | Maps |
| `data-structures/structs` | 🚧 | Structs and methods |
| `interfaces-and-errors/interfaces` | 🚧 | Interfaces |
| `interfaces-and-errors/error-handling` | 🚧 | Wrapping and inspecting errors |
| `interfaces-and-errors/custom-errors` | 🚧 | Custom error types |
| `concurrency/goroutines` | 🚧 | Goroutines |
| `concurrency/channels` | 🚧 | Channels |
| `concurrency/waitgroups` | 🚧 | `sync.WaitGroup` |
| `stdlib/net/http/server` | ✅ | Routing, JSON responses, mandatory timeouts |
| `stdlib/net/http/client` | ✅ | Requests, JSON decoding, error wrapping |
| `projects` | 🚧 | Larger, combined projects |

✅ implemented &nbsp;&nbsp; 🚧 planned / in progress

## Conventions used throughout

- Functions that only produce output are named `printSomething`; functions that
  return a value are named without a `get` prefix.
- Example data is always library data: members, books, shelves, fines.
- Each example carries a small `section(title)` helper that prints an
  `== Title ==` heading before every demo, so grouped output stays readable
  instead of running together into one wall of lines.
- Mistakes made while writing an example are recorded in the Highlights section
  below rather than left as comments in the code, so the code reads as a
  finished example while the reasoning stays discoverable.

## Highlights

### `basics/control-flow`

Covering if/else, an error-returning `if`, tagged and tagless `switch`, the
three-clause `for`, `range` over a slice, and a nested loop over a slice of
slices.

It also keeps a written record of the bugs it went through, marked in place as
`BUG, now fixed`. Worth reading, because three of them failed *quietly* rather
than loudly:

- A `for` loop using `i == shelves` instead of `<=`. Go tests the condition
  before the first iteration, and `i` starts at `0`, so the body never ran even
  once. A dead loop, not an infinite one — nothing hung, the output was just
  missing.
- A `switch` whose cases overlapped (`days <= 30` and `days >= 30`), so day 30
  matched twice and `default` was unreachable. Moving the last case to `>= 90`
  fixed the overlap but opened a hole: days 31-89 matched nothing and reported
  "Invalid date".
- A nested loop printing the shelf number from a function parameter instead of
  the outer loop's index, so every book came out labelled with the same shelf.

The rule behind all three: make the branches cover every input, and check that
the value you're printing is the one that actually varies.

```
$ go run ./basics/control-flow

== Card type ==
Teen Card

== Membership status ==
Welcome back!

== Overdue fine ==
Small fine

== Shelf numbers ==
1
2
3
4
5

== Book titles ==
0 Ghost in the wires
1 Psychology of money
2 White nights

== Books by shelf ==
1 White Night
2 Ghost In The Wire
3 Atomic Habits
```

### `stdlib/net/http/server`

A minimal catalogue server built on the standard `net/http` package:

- Routing with `http.ServeMux`
- Plain-text and JSON responses, via struct tags
- `http.Server` with read and write timeouts
- The `errors.Is(err, http.ErrServerClosed)` filter

Graceful shutdown is **not** implemented yet — nothing calls `server.Shutdown`,
so the `ErrServerClosed` branch is there to show the correct shape before the
`concurrency/` topics land.

### `stdlib/net/http/client`

A minimal client for that server:

- `http.Client` with a client-side timeout
- Checking `resp.StatusCode` by hand, since `net/http` will not do it for you
- Decoding JSON into a struct
- Wrapping failures with `fmt.Errorf` and `%w`
- Separating request logic from presentation: `fetchStatus` does the work, `main` does the printing

## Module

```
module github.com/NamX1/golearn
```

## License

MIT. See [LICENSE](LICENSE).
