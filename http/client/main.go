package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

var url string

func main() {
	//Asking user for website url
	fmt.Print("Website url: ")
	fmt.Scan(&url)
	fmt.Println("\nRequest url: ", url)

	//http client function
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//http request get
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error request: ", err)
		return
	}
	defer resp.Body.Close()

	scan := bufio.NewScanner(resp.Body)
	for scan.Scan() {
		fmt.Println(scan.Text())
	}
	if err := scan.Err(); err != nil {
		fmt.Println("Scanner error: ", err)
	}
}
