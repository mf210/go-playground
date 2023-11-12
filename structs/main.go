package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	matt := person{firstName: "matt", lastName: "Faresi"}
	fmt.Println(matt)
}
