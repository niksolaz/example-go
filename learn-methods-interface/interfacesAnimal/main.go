package main

import "fmt"

// Animal interface with Speak method that returns a string value
type Animal interface {
	Speak() string // Speak method returns a string value
}

type Dog struct{} // Dog struct with no fields or methods defined on it yet

// Speak method for Dog struct that returns a string value
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{} // Cat struct with no fields or methods defined on it yet

// Speak method for Cat struct that returns a string value
func (c Cat) Speak() string {
	return "Meow!"
}

type AnimalWithPaws interface {
	Animal             // Animal interface is embedded in AnimalWithPaws interface
	NumberOfPaws() int // NumberOfPaws method returns a int value
}

type Parrot struct {
	paws int // paws field of type int
}

// NumberOfPaws method for Parrot struct that returns a int value
func (p Parrot) NumberOfPaws() int {
	return p.paws
}

// Speak method for Parrot struct that returns a string value
func (p Parrot) Speak() string {
	return "Squawk!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Parrot{}} // slice of Animal interface with Dog and Cat structs as values in it (Dog{} and Cat{} are struct literals)
	for _, animal := range animals {            // iterate over the slice of Animal interface
		fmt.Println(animal.Speak())
	}

	p := Parrot{paws: 2}          // Parrot struct literal with paws field set to 2
	fmt.Println(p.NumberOfPaws()) // prints 2

}
