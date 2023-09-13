package main

import "fmt"

//Interfaces are trying to make reusing code easier.

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//could put these on structs
//can remove receiver value and only leave type if not using variable IE 'eb/sb'
func (englishBot) getGreeting() string {
	//Very custom logic for generating English greeting
	return "Hi There!"
}

// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }
