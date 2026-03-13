# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

Teaching repository for a university Web Technologies course (3rd-year CS students). Contains progressive Go networking examples that build from raw TCP sockets up to REST APIs with databases. Examples are intentionally simple and self-contained for pedagogical clarity.

## Repository Architecture

Each subdirectory is an **independent Go module** with its own `go.mod`. There is no shared code between examples.

**Progression path (intended learning order):**

1. **telnet-simple** — Single-client TCP echo server (blocking)
2. **telnet-multi** — Multi-client TCP chat with goroutines and broadcast
3. **telnet-http** — HTTP server built from raw TCP (manual request parsing, no net/http)
4. **http-multi** — Multi-route HTTP server with forms/JS, still using raw TCP
5. **http-rest** — JSON REST API over raw TCP with CORS headers
6. **http-stdlib** — Same as http-rest but using `net/http` (shows the stdlib advantage)
7. **http-chat** — Simple in-memory chat API (GET/POST /messages) using `net/http` and `encoding/json`
8. **http-chat-sqlite** — Same chat API but persisted to SQLite with `modernc.org/sqlite`
9. **http-chat-client** — Frontend + proxy on :8000, forwards /api/messages to the chat API on :3000; serves static HTML/JS
10. **dns-server** — Custom DNS server using `miekg/dns`
8. **tcp-capture** — Packet sniffer using `google/gopacket`
9. **sqlite-simple** — Basic SQLite CRUD with `modernc.org/sqlite`
10. **http-sqlite** — Full REST API + SQLite (Go 1.22+ method routing, handler files split by HTTP method)

**Key pattern:** Examples 1-5 build HTTP from scratch using `net` and `bufio` to teach protocol internals. Example 6 then shows the same thing with `net/http` to demonstrate why frameworks exist.

## Build and Run Commands

Each example is a standalone module. Always `cd` into the example directory first:

```bash
cd <example-dir>
go run .              # run (single-file examples)
go run *.go           # run (multi-file examples like http-multi, http-sqlite)
go build -o <name> .  # build binary
```

No tests exist in the repository. No linter configuration is present.

## Key Conventions

- **Raw TCP examples** (telnet-*, http-multi, http-rest): Use `net.Listen("tcp", ":port")` with manual HTTP parsing via `bufio.Reader`. HTTP responses are hand-built strings with `\r\n` line endings.
- **Database examples** use `modernc.org/sqlite` (pure Go, no CGO) via standard `database/sql` interface.
- **http-sqlite** uses Go 1.22+ routing syntax: `GET /series/{id}`, `POST /series`, etc. Handlers are split into `get_handlers.go`, `write_handlers.go`, `delete_handler.go` with a shared `model.go` for the `Serie` struct.
- Go versions vary across modules (1.23.4, 1.24.0, 1.26.1).
