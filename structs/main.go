package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Roberts",
		contact: contactInfo{
			email:   "123@12334.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", jim)
}
