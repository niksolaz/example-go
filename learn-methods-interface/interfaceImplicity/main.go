package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string // S field of type string in T struct type
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

type P struct {
	N int // N field of type int in P struct type
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (p P) M() {
	fmt.Println(p.N) // prints the value of N field in P struct type
}

func main() {
	// The interface value i can hold any value that implements the M method.
	var i I = T{"hello"}
	i.M()

	// The interface value j can hold any value that implements the M method.
	var j I = P{10}
	j.M()
}
