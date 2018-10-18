package main

import "testing"

func TestUpdateName(t *testing.T) {
	p := person{
		firstName: "FirstName",
		lastName:  "LastName",
		contactInfo: contactInfo{
			email:   "FirstName.LastName@test.com",
			zipCode: 55555,
		},
	}

	p.updateName("UpdatedFirstName", "UpdatedLastName")

	if p.firstName != "UpdatedFirstName" {
		t.Errorf("EXPECTED First Name = UpdatedFirstName; GOT First Name = %+v", p.firstName)
	}

	if p.lastName != "UpdatedLastName" {
		t.Errorf("EXPECTED Last Name = UpdatedLastName; GOT Last Name = %+v", p.lastName)
	}
}

func TestUpdateContactInfo(t *testing.T) {
	p := person{
		firstName: "FirstName",
		lastName:  "LastName",
		contactInfo: contactInfo{
			email:   "FirstName.LastName@test.com",
			zipCode: 55555,
		},
	}

	p.updateContactInfo("TotallyNewAndLegitEmail@test.com", 12345)

	if p.email != "TotallyNewAndLegitEmail@test.com" {
		t.Errorf("EXPECTED Email = TotallyNewAndLegitEmail@test.com; GOT First Name = %+v", p.email)
	}

	if p.zipCode != 12345 {
		t.Errorf("EXPECTED Zip Code = 12345; GOT Last Name = %+v", p.zipCode)
	}
}
