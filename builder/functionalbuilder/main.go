package main

import "fmt"

type Person struct {
	name, position string
}

type PersonMod func(*Person)

type PersonBuilder struct {
	actions []PersonMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, action := range b.actions {
		action(&p)
	}
	return &p
}

// extend PersonBuilder
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func main() {
	fmt.Println("Functional Builder")
	b := PersonBuilder{}
	p := b.Called("John").
		WorksAsA("developer").
		Build()
	name := p.name
	job := p.position
	fmt.Println(name, job)
	fmt.Println(*p)
}
