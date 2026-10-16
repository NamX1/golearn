package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// catalogueResponse is the JSON body the endpoints send back. The struct tags
// decide the keys on the wire, so the JSON reads {"status": ..., "message": ...}.
type catalogueResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// handleRoot answers any path with a plain-text greeting. The request is unused
// here, so it is discarded as _.
func handleRoot(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hey there!")
}

// handleStatus reports whether the catalogue is reachable, as JSON.
func handleStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := catalogueResponse{
		Status:  "OK",
		Message: "Library catalogue is open",
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/status", handleStatus)

	// Timeouts are not optional on a public server: without them one slow
	// client can hold a connection open indefinitely.
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Library catalogue listening on http://localhost:8080")

	// ListenAndServe always returns a non-nil error. A clean Shutdown reports
	// http.ErrServerClosed, which is not a failure. Nothing calls Shutdown yet.
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Library catalogue failed: %s", err)
	}
}
