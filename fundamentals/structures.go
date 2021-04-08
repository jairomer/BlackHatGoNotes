package main

import (
	"fmt"
)


// Define functionalities in terms of interfaces if different kind of
// structures need to be processed on those functionalities.
// Then define the particulars on the methods of thos structures.
// Go will automatically infer if such structure has a defined interface.
type Friend interface {
	SayHello()
}

// This is a function.
// No structure as dependency.
// Coupled to an interface, we can then use the double dispatch method
// to reuse different structures that implement common functionalities.
func Greet (f Friend) {
	f.SayHello()
}

type Person struct {
	Name string
	Age int
}

// This is a method.
// Coupled to a specific type of Structure.
func (p* Person) SayHello() {
	fmt.Println("Hello, ", p.Name)
}


type Dog struct {}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)

	var dog = new(Dog)
	Greet(dog)
}
