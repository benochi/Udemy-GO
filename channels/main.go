package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c { //infinite loop. l = link step over for loop every time channel emits a value
		go func(link string) { //link is from l arg in 2nd parenthesis from parent scope.
			time.Sleep(5 * time.Second)
			checkLink(link, c) //variable referenced in outer scope is bad. pass main routine data.
		}(l) //function literal *Lambda expression* notice extra parenthesis for execution of lambda.
		//args go into the outer parenthesis.
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
