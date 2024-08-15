package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The play function plays a game of rock-paper-scissors.
// params c chan string: a channel to send the result of the game to
func play(c chan string) {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Create a slice of strings representing the possible choices in the game
	choices := []string{"rock", "paper", "scissors"}

	// Generate a random index to select a choice from the slice
	i := rand.Intn(len(choices))

	// Send the randomly selected choice to the channel
	c <- choices[i]
}

func main() {
	// Create a channel of strings to communicate the result of the game
	c := make(chan string)

	// Start a goroutine that plays the game and sends the result to the channel
	go play(c)

	// Receive the result of the game from the channel
	result := <-c

	// Print the result of the game
	fmt.Println("You chose rock.")
	fmt.Printf("The computer chose %s.\n", result)

	// Determine the winner of the game
	switch result {
	case "rock":
		fmt.Println("It's a tie!")
	case "paper":
		fmt.Println("You lose!")
	case "scissors":
		fmt.Println("You win!")
	}
}

// Output: You chose rock. The computer chose rock. It's a tie!
