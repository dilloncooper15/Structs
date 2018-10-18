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
	p.updateAllDemographics("updatedFirstName", "updatedLastName", "updatedEmail@example.com", 99999)
	p.print()
}

/*
[p] takes a struct of type [person] as a receiver and overwrites the value of
[person.firstName] to [fn], [person.lastName] to [ln], [person.email] to [e],
and [person.zipCode] to [zc].
*/
func (p *person) updateAllDemographics(fn string, ln string, e string, zc int) {
	p.firstName = fn
	p.lastName = ln
	p.email = e
	p.zipCode = zc
}

/*
[p] takes a struct of type [person] as a receiver and overwrites the value of
[person.firstName] to [fn] and the value for [person.lastName] to [ln].
*/
func (p *person) updateName(fn string, ln string) {
	p.firstName = fn
	p.lastName = ln
}

/*
[p] takes a struct of type [person] as a receiver and overwrites the value of
[person.firstName] to [fn].
*/
func (p *person) updateFirstName(fn string) {
	p.firstName = fn
}

/*
[p] takes a struct of type [person] as a receiver and overwrites the value of
[person.firstName] to [fn].
*/
func (p *person) updateLastName(ln string) {
	p.lastName = ln
}

/*
[c] takes a struct of type [contactInfo] as a receiver and overwrites the value
of [contactInfo.email] to [e] and the value of [contactInfo.zipCode] to [zc].
*/
func (c *contactInfo) updateContactInfo(e string, zc int) {
	c.email = e
	c.zipCode = zc
}

/*
[c] takes a struct of type [contactInfo] as a receiver and overwrites the value
of [contactInfo.email] to [e].
*/
func (c *contactInfo) updateContactEmail(e string) {
	c.email = e
}

/*
[c] takes a struct of type [contactInfo] as a receiver and overwrites the value
of [contactInfo.zipCode] to [zc].
*/
func (c *contactInfo) updateContactZipCode(zc int) {
	c.zipCode = zc
}

/*
Takes a struct of type contactInfo as a receiver and prints the contents to the console.
*/
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
