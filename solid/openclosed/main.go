package main

import (
	"fmt"
)

// OCP states that a class should be open for extension but closed for modification.
// Specification pattern is a way to implement OCP.

// Product is a struct

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// Filter is a struct

}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {

	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {

	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// specification pattern

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type OrSpecification struct {
	first, second Specification
}

func (o OrSpecification) IsSatisfied(p *Product) bool {
	return o.first.IsSatisfied(p) || o.second.IsSatisfied(p)
}

type BetterFilter struct {
	// BetterFilter is a struct
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	f := Filter{}
	greenProducts := f.FilterByColor(products, green)
	fmt.Println("Green products are(old):")
	for _, v := range greenProducts {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeProducts := f.FilterBySize(products, large)
	fmt.Println("Large products are(old):")
	for _, v := range largeProducts {
		fmt.Printf(" - %s is large\n", v.name)
	}

	largeGreenProducts := f.FilterBySizeAndColor(products, large, green)
	fmt.Println("Large green products are(old):")
	for _, v := range largeGreenProducts {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

	// OCP with specification pattern
	fmt.Println("OCP with specification pattern")
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}

	largeGreenSpec := greenSpec.IsSatisfied(&tree)
	fmt.Println("Is tree large and green? ", largeGreenSpec)

	largeGreenSpec = largeSpec.IsSatisfied(&tree)
	fmt.Println("Is tree large? ", largeGreenSpec)

	bf := BetterFilter{}
	greenProductsNew := bf.Filter(products, greenSpec)
	fmt.Println("Green products are(new):")
	for _, v := range greenProductsNew {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// and specification
	largeGreenSpec = AndSpecification{greenSpec, largeSpec}.IsSatisfied(&tree)
	fmt.Println("Is tree large and green? ", largeGreenSpec)

	// or specification
	largeGreenSpec = OrSpecification{greenSpec, largeSpec}.IsSatisfied(&tree)
	fmt.Println("Is tree large or green? ", largeGreenSpec)

	fmt.Println("Large green products are(new):")
	for _, v := range bf.Filter(products, AndSpecification{greenSpec, largeSpec}) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

}
