package hello

import (
	"fmt"

	"example-go/niksolaz/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Nik")
	fmt.Println(message)
}
