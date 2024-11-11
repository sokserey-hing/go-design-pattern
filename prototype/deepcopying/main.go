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
	}

	fmt.Println("John: ", john, john.Address)

	jane := john

	// jane.Address.StreetAddress = "456 Elm St"
	jane.Address = &Address{
		StreetAddress: "456 Elm St",
		City:          "Vancouver",
		Country:       "Canada",
	}
	jane.Name = "Jane" // okay

	fmt.Println("Jane: ", jane, jane.Address)

}
