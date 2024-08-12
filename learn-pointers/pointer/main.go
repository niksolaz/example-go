package main

import "fmt"

type ninja struct {
	name string
}

func main() {
	tommy := ninja{name: "Tommy"}
	fmt.Println(tommy)
	tommyPointer := &tommy
	fmt.Println(tommyPointer)
	johnnyPointer := &ninja{name: "Johnny"}
	fmt.Println(johnnyPointer)
	var ninjaPointer *ninja = new(ninja)

	fmt.Println(ninjaPointer)
	fmt.Println(*johnnyPointer)
}
