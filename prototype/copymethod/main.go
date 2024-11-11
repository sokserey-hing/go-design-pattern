package main

// Prototype : A partially or fully initialized object that you copy (clone) and make use of..
// Copy method is a method that creates a new object and copies all the fields of the current object to the new object.

import "fmt"

type Address struct {
	StreetAddress string
	City          string
	Country       string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p // copies Name
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

func main() {
	fmt.Println("----------------- Deep Copying -----------------")

	john := Person{
		"John",
		&Address{
			StreetAddress: "123 Oak St",
			City:          "Toronto",
			Country:       "Canada",
		},
		[]string{"Chris", "Matt"},
	}

	fmt.Println("John: ", john, john.Address)

	jane := john.DeepCopy()
	jane.Address.StreetAddress = "456 Elm St"
	jane.Address.City = "Vancouver"
	jane.Name = "Jane" // okay
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println("Jane: ", jane, jane.Address)

}
