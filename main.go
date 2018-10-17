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

	p := person{
		firstName: "Jim",
		lastName:  "Random",
		contactInfo: contactInfo{
			email:   "jim.random@example.com",
			zipCode: 55555,
		},
	}

	p.print()
	p.updateName("Chris", "Roberts")
	p.print()
	p.updateContactInfo("croberts@gmail.com", 11111)
	p.print()
}

/**
Takes a struct of type person as a receiver and updates the person's first name
and last name to the values passed as parameters.
**/
func (p *person) updateName(fn string, ln string) {
	p.firstName = fn
	p.lastName = ln
}

/**
Takes a struct of type contactInfo as a receiver and updates the person's email
and zip code to the values passed as parameters.
**/
func (c *contactInfo) updateContactInfo(e string, zc int) {
	c.email = e
	c.zipCode = zc
}

/**
Takes a struct of type contactInfo as a receiver and prints the contents to the console.
**/
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
