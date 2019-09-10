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
	// basic struct
	// comrade := person{"archie", "isdiningrat"}
	comrade := person{
		firstName: "archie",
		lastName:  "atrie",
		contact: contactInfo{
			email:   "archie@mail.com",
			zipCode: 201020,
		},
	}
	comrade.lastName = "isdiningrat"

	// turn into memory address
	// comradePointer := &comrade
	comrade.updateName("angga")
	comrade.print()
}

// struct funciton receiver
// pointers, memory address to value
func (pointer *person) updateName(name string) {
	(*pointer).firstName = name
}

func (pointer person) print() {
	fmt.Println(pointer)
}
