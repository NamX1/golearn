package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// catalogueURL is the address this client expects. Start the server first:
//
//	go run ./stdlib/net/http/server
const catalogueURL = "http://localhost:8080/status"

// catalogueResponse mirrors the JSON body the server sends.
type catalogueResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// fetchStatus does the work and returns a decoded reply or an error. It never
// prints, so it stays easy to test. Failures are wrapped with %w so callers can
// unwrap them with errors.Is.
func fetchStatus(client *http.Client, url string) (catalogueResponse, error) {
	var result catalogueResponse

	resp, err := client.Get(url)
	if err != nil {
		return result, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// A non-200 is not an error to net/http; it has to be checked by hand.
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("decoding response: %w", err)
	}

	return result, nil
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	result, err := fetchStatus(client, catalogueURL)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Status %s\nMessage: %s\n", result.Status, result.Message)
}
