package main

import "fmt"

// builder facats
// builder facets are the sub-builders that are used to build the final object in a step-by-step manner. The builder facade is the main builder that is used to create the final object.

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// employment
	CompanyName, Position string
	AnnualIncome          int
}

// builder
type PersonBuilder struct {
	person *Person
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// methods for PersonAddressBuilder
func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

// methods for PersonJobBuilder

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (pjb *PersonJobBuilder) Build() *Person {
	return pjb.person
}

func (pab *PersonAddressBuilder) Build() *Person {
	return pab.person
}

func main() {
	fmt.Println("-------------Builder Facets-------------")
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Engineer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(person)

}
