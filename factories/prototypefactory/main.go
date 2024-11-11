package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	fmt.Println("----------------- Prototype Factory -----------------")
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)

	d := NewEmployee(Developer)
	d.Name = "John"
	fmt.Println(d)
}
