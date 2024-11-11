package main

import (
	"fmt"
)

// Interface Segregation Principle
// A client should never be forced to implement an interface that it doesn't use or clients shouldn't be forced to depend on methods they do not use.

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {
	fmt.Println("Print")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	fmt.Println("Fax")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	fmt.Println("Scan")
}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	fmt.Println("Print OldFashionedPrinter")
}

// below is a violation of the Interface Segregation Principle
// The OldFashionedPrinter is forced to implement the Fax and Scan methods which it does not use.
// but we can use panic to prevent the OldFashionedPrinter from implementing the Fax and Scan methods

func (o OldFashionedPrinter) Fax(d Document) {
	panic("Operation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("Operation not supported")
}

// ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {
	fmt.Println("Print MyPrinter")
}

type Photocopier struct {
}

func (p Photocopier) Print(d Document) {
	fmt.Println("Print Photocopier")
}

func (p Photocopier) Scan(d Document) {
	fmt.Println("Scan Photocopier")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// Decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	fmt.Println("Interface Segregation Principle")

	ofp := OldFashionedPrinter{}
	ofp.Print(Document{})
	// ofp.Fax(Document{})  // panic: Operation not supported
	// ofp.Scan(Document{}) // panic: Operation not supported

	fmt.Println("ISP")
	mfp := MultiFunctionPrinter{}
	mfp.Print(Document{})
	mfp.Fax(Document{})
	mfp.Scan(Document{})

	fmt.Println("ISP")
	mp := MyPrinter{}
	mp.Print(Document{})
	// mp.Scan(Document{}) // panic: Operation not supported

	pc := Photocopier{}
	pc.Print(Document{})
	pc.Scan(Document{})
	// pc.Fax(Document{}) // panic: Operation not supported

	// MutliFunctionDevice
	mfd := MultiFunctionMachine{printer: MyPrinter{}, scanner: Photocopier{}}
	mfd.Print(Document{}) // Print MyPrinter

}
