package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // ====> contactInfo   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newfirstName string) {
	(*pointerToPerson).firstName = newfirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
