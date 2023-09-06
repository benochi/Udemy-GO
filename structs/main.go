package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//can declare key or it will go by struct order.
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	//another way to assign a struct go will default 0 value.
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
