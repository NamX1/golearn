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
| `data-structures/structs` | 🚧 | Structs, literals, zero values, copy semantics, comparability |
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

### `data-structures/structs`

Not written yet. This is the next topic on the list, and the tasks below are the
brief. Structure it as a single `main.go`, following the conventions above: a
`section` helper, library data, one short doc comment per function.

Build `Book`, `Member` and `Shelf` types and use them to show:

- Exported (`Title`) vs unexported (`tags`) fields
- Named vs positional struct literals, and why positional is a trap
- The zero value: every field blank without any constructor
- Nested structs, and slices of structs
- **Copy semantics**: assignment copies the entire struct
- **Comparability**: `==` only works if every field is itself comparable

Methods and pointer receivers come after this, in part 2.

#### Tasks

Work through these in order. For the two marked **predict first**, write down what
you think the output will be *before* running it.

1. **Named literal.** Write a function `NewArrival(title, author string, copies int) Book`
   that returns a `Book` built with a named literal. Have it return
   `"The Go Programming Language"`, `"Donovan and Kernighan"`, 3 copies. Print
   the `Title`, `Author` and `Copies`.

2. **Zero value.** Declare `var draft Book` without assigning anything. Print it
   with `%+v` and confirm every field is blank, including the `[]string`.

3. **Copy semantics — predict first.**

   ```go
   a := Book{Title: "Dune", Copies: 1}
   b := a
   b.Copies = 5
   fmt.Println(a.Copies, b.Copies)
   ```

   Write down your prediction. It is `1 5`, not `5 5`. Then explain in one
   sentence why.

4. **The range trap — predict first.** Take one copy off every book in a
   `[]Book`. The obvious version silently does nothing:

   ```go
   for _, b := range books {
       b.Copies--
   }
   ```

   Predict whether `books` changes. Then fix it, and explain in one sentence why
   the broken version fails. **This is the one that matters** — the loop runs,
   nothing errors, and the data is just wrong.

5. **Comparability.** `Book` has an unexported `tags []string` field, so
   `bookA == bookB` is a compile error. Remove the field, then write one line
   comparing two `Book` values and print the result.

6. **Stretch — your first method.** Give `Book` a `Borrow()` method that
   refuses when `Copies` is 0, otherwise decrements `Copies` and sets
   `OnLoan`. Write it with a **pointer** receiver, and call it from a function
   in `main` rather than from `main` itself.

   Then rewrite it with a **value** receiver and watch `b.Copies--` stop having
   any effect on the original. That contrast is the whole point of part 2, so
   do not skip it.

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
