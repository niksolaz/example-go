package messages

import "fmt"

func WriteMessage(msg string) string {
	message := fmt.Sprintf("Hi, %v. You Have new message: ", msg)
	return message
}
