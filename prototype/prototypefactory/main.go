package main

// Prototype Factory is a design pattern that allows you to create new objects by copying an existing object, known as the prototype. The prototype factory is useful when you need to create new objects that are similar to existing objects but have different values or configurations. This pattern is especially useful when creating objects that are expensive to create or when you need to create objects that are similar to existing objects but have different values or configurations.

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress string
	City          string
	Suite         int
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}     // A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	e := gob.NewEncoder(&b) // NewEncoder returns a new encoder that writes to w.
	_ = e.Encode(p)

	// fmt.Println(b.String())

	d := gob.NewDecoder(&b) // NewDecoder returns a new decoder that reads from r.
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{"", Address{"123 East Dr", "London", 0}}
var auxOffice = Employee{"", Address{"66 West Dr", "London", 101}}

// let define a new utility for creating new employees
func NewEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	// john and jane are two different employees with different office addresses

	fmt.Println("John: ", john, john.Office)

	fmt.Println("Jane: ", jane, jane.Office)

}
