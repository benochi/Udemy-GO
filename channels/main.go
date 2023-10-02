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

	for _, link := range links {
		go checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		return
	}
	fmt.Println(link, "is up and running.")
}
