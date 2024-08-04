package main

import (
	"example-go/niksolaz/greetings"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("New Mod Nik")
	fmt.Println(message)
}
