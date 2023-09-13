package main

import "fmt"

//Interfaces are trying to make reusing code easier.
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type russianBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	rb := russianBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(rb)
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//could put these on structs
//can remove receiver value and only leave type if not using variable IE 'eb/sb'
func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (russianBot) getGreeting() string {
	return "привет!"
}
