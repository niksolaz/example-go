package main

import "fmt"

func Pointers() {
	var pnt *int // pnt is a pointer to an integer

	valueForPnt := 120 // valueForPnt is an integer with value 120
	pnt = &valueForPnt // pnt is now pointing to the memory address of valueForPnt

	fmt.Println("Value of valueForPnt:", valueForPnt) // valueForPnt is an integer with value 120
	fmt.Println("Value of *pnt:", *pnt)               // *pnt is the value at the memory address pnt is pointing to
	fmt.Println("Memory Address:", pnt)               // pnt is the memory address of valueForPnt

	*pnt = 130 // changing the value at the memory address pnt is pointing to

	fmt.Println("New Value of valueForPnt:", valueForPnt) // valueForPnt is an integer with value 130
	fmt.Println("New Value of *pnt:", *pnt)               // *pnt is the value at the memory address pnt is pointing to
	fmt.Println("Memory Address:", pnt)                   // pnt is the memory address of valueForPnt

	valueForPnt = 140                                     // changing the value of valueForPnt
	fmt.Println("New Value of valueForPnt:", valueForPnt) // valueForPnt is an integer with value 130
	fmt.Println("New Value of *pnt:", *pnt)               // *pnt is the value at the memory address pnt is pointing to
	fmt.Println("Memory Address:", pnt)                   // pnt is the memory address of valueForPnt
}

func main() {
	Pointers()
}
