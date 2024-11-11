package main

// Copy Through Serialization is a technique that allows you to copy complex objects by serializing them to a stream of bytes and then deserializing them back into a new object. This technique is useful when you need to copy objects that are not serializable or when you need to create deep copies of objects that contain circular references or other complex structures.

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

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
	b := bytes.Buffer{}     // A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	e := gob.NewEncoder(&b) // NewEncoder returns a new encoder that writes to w.
	_ = e.Encode(p)

	// fmt.Println(b.String())

	d := gob.NewDecoder(&b) // NewDecoder returns a new decoder that reads from r.
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

// main demonstrates the prototype design pattern using deep copy through serialization.
// It creates a Person object, makes a deep copy of it, modifies the copy, and prints both the original and the copied objects.
func main() {
	fmt.Println("Copy Through Serialization")

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
