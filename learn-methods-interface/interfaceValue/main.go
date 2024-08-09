package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

// T struct type with S field of type string
type T struct {
	S string
}

// M method for T struct type that prints the value of S field
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64 // F type with underlying type float64

// M method for F type that prints the value of F type
func (f F) M() {
	fmt.Println(f)
}

type P struct {
	K int // K field of type int
}

// M method for P struct type that prints the value of K field
func (p *P) M() {
	if p == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(p.K)
}

func main() {
	var i I // i is an interface type

	i = &T{"Hello"} // i is assigned a pointer to T struct type with S field set to "Hello"
	describe(i)
	i.M() // prints "Hello"

	i = F(math.Pi) // i is assigned a F type with value of math.Pi
	describe(i)
	i.M() // prints 3.141592653589793

	var p *P // p is a pointer to P struct type
	i = p    // i is assigned a pointer to P struct type
	describe(i)
	i.M() // prints nil
}

// describe function that takes an interface type as an argument and prints the value and type of the interface type argument passed to it
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i) // prints the value and type of the interface type argument passed to it
}
