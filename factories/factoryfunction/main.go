package main

// Factory function is a function that returns an instance of a struct

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		panic("Person is too young")
	}
	return &Person{name, age, 2}
}

func main() {
	fmt.Println("-----factory function-----")
	p := NewPerson("John", 10)

	fmt.Println(p)
}
