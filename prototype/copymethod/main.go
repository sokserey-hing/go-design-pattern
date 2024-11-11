package main

// Deep copying is a process of copying an object and all of its child objects.
// Prototype : A partially or fully initialized object that you copy (clone) and make use of.

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
