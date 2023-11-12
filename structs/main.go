package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var matt person
	matt.firstName = "matt"
	matt.lastName = "Anderson"

	fmt.Println(matt)
	fmt.Printf("%+v", matt)
}
