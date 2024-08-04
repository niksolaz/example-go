package main

import (
	"fmt"

	"example-go/niksolaz/greetings"
	"example-go/niksolaz/messages"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Nik")
	fmt.Println(message)

	listenMessage := messages.WriteMessage("Thsi is new message")
	fmt.Println(listenMessage)
}
