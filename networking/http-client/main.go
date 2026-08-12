package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// run the http-server first
const statusURL = "http://localhost:8080/status"

type statusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func fetchStatus(client *http.Client, url string) (statusResponse, error) {
	var result statusResponse

	resp, err := client.Get(url)
	if err != nil {
		return result, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code: %w", resp.StatusCode)
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

	result, err := fetchStatus(client, statusURL)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Status %s\nMessage: %s\n", result.Status, result.Message)
}
