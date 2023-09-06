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
//the pointer to type pointer of a person will automatically handle if a 'person' is passed in instead of a pointer explicitly.
//Slices are a reference type, so they point to the same addess in memory when copied by GO and passed into a function
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
