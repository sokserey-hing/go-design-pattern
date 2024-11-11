package main

import (
	"fmt"
)

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
// passing function as a parameter
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach
// passing struct as a parameter
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	fmt.Println("---------factory generator---------")
	fmt.Println("---------functional approach---------")
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer)
	fmt.Println(manager)

	// structural approach
	fmt.Println("---------structural approach---------")
	ceoFactory := EmployeeFactory{"CEO", 100000}
	fmt.Println(ceoFactory)
	ceo := ceoFactory.Create("Sam")
	fmt.Println(ceo)

	bossFactory := NewEmployeeFactory2("CTO", 120000)
	fmt.Println(bossFactory)
	bossFactory.AnnualIncome = 130000 // here we can change the value of the struct
	boss := bossFactory.Create("John")
	fmt.Println(boss)

}
