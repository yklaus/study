package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.reddit.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}

	c := make(chan requestResult)

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for index := 0; index < len(urls); index++ {
		result := <-c
		fmt.Println(result.url, result.status)
	}
}

func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || (resp != nil && resp.StatusCode >= 400) {
		status = "FAILED"
	}

	c <- requestResult{url: url, status: status}
}
