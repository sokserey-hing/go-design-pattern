package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry I am too tired")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}

	if age < 16 {
		panic("Person is too young")
	}
	return &person{name, age}
}

func main() {
	// person := person{name: "John", age: 30}
	// person.SayHello()

	fmt.Println("\n-----interface factory-----")
	p := NewPerson("John", 30)
	p.SayHello() //Hello, my name is John and I am 30 years old

	fmt.Println("\n-----interface factory-----")
	p1 := NewPerson("Jim", 120)
	p1.SayHello() //Sorry I am too tired

	// p and p1 will return different values because they are different types even though they implement the same interface

}
