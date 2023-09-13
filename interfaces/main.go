package main

import "fmt"

//Interfaces are trying to make reusing code easier.

type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("Duh")
}

//could put these on structs
func (eb englishBot) getGreeting() string {
	//Very custom logic for generating English greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	//Very custom logic for generating spanish greeting
	return "Hola!"
}
