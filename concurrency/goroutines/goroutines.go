package main

import (
	"fmt"
	"time"
)

// The say function prints a string 5 times with a 100ms delay.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // Sleep for 100ms
		fmt.Println(s, i)                  // Print the string and the loop counter
	}
}

func main() {
	go say("world") // Start a goroutine that calls say("world")
	say("hello")    // Call say("hello") in the main goroutine
}
