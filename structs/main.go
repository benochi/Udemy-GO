package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Roberts",
		contactInfo: contactInfo{
			email:   "123@12334.com",
			zipCode: 12345,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.print()

	jim.updateName("Jimmy")
	jim.print()
}

//Was able to do this by only adding the asterisk, not sure why we need to assign the pointer inside main, as it worked without it.
//the pointer to type pointer of a person was able to fix the bug and I need to confirm why.
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
