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
	comrade.updateName("angga")
	comrade.print()
}

// struct funciton receiver

func (p person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Println(p)
}
