package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact contactinfo
}

type contactinfo struct {
	email string
	zipCode int
}

func main() {
	// var clyne person
	// clyne.firstName = "clyne"
	// clyne.lastName = "fuentes"
	clyne := person{
		firstName: "clyne test", 
		lastName: "fuentes",
		contact: contactinfo{
			email: "clyne@email",
			zipCode: 123,
		},
	}

	clyne.updateName("clippe")
	fmt.Println(clyne)
	clyne.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}