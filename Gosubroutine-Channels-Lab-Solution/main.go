package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := [5]string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel for communication b/w subroutines.
	channel := make(chan string)

	for _, link := range links {
		go makeRequest(link, channel)
	}

	for index := 0; index < len(links); index++ {
		fmt.Println(<-channel)
	}
}

func makeRequest(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		message := link + " " + "Website is down"
		channel <- message;
		return
	}
	message := link + " " + "Website is up"
	channel <- message;
}
