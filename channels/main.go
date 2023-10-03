package main

import (
	"fmt"
	"net/http"
)

// App should make reqs to multiple websites to see status.
// channels handle communication between go routines and handle a type.
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://cloudflare.com",
		"http://ebay.com",
		"http://rumble.com",
		"http://udemy.com",
		"http://youtube.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		// fmt.Println(<-c) this works but seems slower.
	}

	for { //infinite loop.
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		c <- link
		return
	}

	fmt.Println(link, "is up and running.")
	c <- link
}
