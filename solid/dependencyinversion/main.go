package main

import (
	"fmt"
)

// Dependency Inversion Principle(DIP)
// High-level modules should not depend on low-level modules. Both should depend on abstractions.
// Abstractions should not depend on details. Details should depend on abstractions.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// DOB
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module
type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for _, rel := range r.relations {
		if rel.from.name == name && rel.relationship == Parent {
			result = append(result, rel.to)
		}
	}
	return result

}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// break DIP
	// relationships Relationships

	// follow DIP
	browser RelationshipBrowser // abstraction

}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

// main is the entry point of the program. It demonstrates the Dependency Inversion Principle
// by creating instances of Person and establishing parent-child relationships using the
// Relationships struct. It then performs research on these relationships.
func main() {
	fmt.Println("Dependency Inversion Principle")
	fmt.Println("High-level modules should not depend on low-level modules. Both should depend on abstractions.")

	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	fmt.Println("Investigate")
	r.Investigate() // John has a child called Chris
	// John has a child called Matt

}
