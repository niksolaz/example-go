package main

import "fmt"

// The sum function sums the elements of a slice and sends the result to a channel.
// params s []int: a slice of integers to sum up and send to the channel
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s { // iterate over the slice
		sum += v // add the value to sum
	}
	c <- sum // send sum to c (channel)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0} // create a slice of integers

	c := make(chan int)     // create a channel of integers
	go sum(s[:len(s)/2], c) // start a goroutine that calls sum(s[:len(s)/2], c), passing the first half of the slice
	go sum(s[len(s)/2:], c) // start another goroutine that calls sum(s[len(s)/2:], c), passing the second half of the slice
	x, y := <-c, <-c        // receive from c (channel) and assign to x and y respectively (blocking until the channel is closed)

	fmt.Println(x, y, x+y) // print x, y, and the sum of x and y
}
