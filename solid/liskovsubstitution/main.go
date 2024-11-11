package main

import (
	"fmt"
)

// Liskov Substitution Principle

// If S is a subtype of T, then objects of type T may be replaced with objects of type S without altering any of the desirable properties of the program.
// In other words, the objects of a superclass shall be replaceable with objects of its subclasses without affecting the functionality of the program.

type Sized interface {
	GetWidth() int
	GetHeight() int

	SetWidth(width int)
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width

}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle // Embedding
}

func NewSoquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

func UseIt(size Sized) {
	width := size.GetWidth()
	size.SetHeight(10)
	expectedArea := width * 10
	actualArea := size.GetWidth() * size.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

type Square2 struct {
	size int // width and height
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	fmt.Println("Lisko Substitution Principle")
	rc := &Rectangle{2, 3}
	UseIt(rc) // Expected an area of 20, but got 20

	sq := NewSoquare(5)
	UseIt(sq) // Expected an area of 50, but got 100 which is wrong and violates the LSP principle because the square is not a rectangle and it does not have the same behavior as a rectangle

	// sq2 := &Square2{5}
	// UseIt(&sq2.Rectangle()) // Expected an area of 50, but got 50

}
