package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The play function plays a game of rock-paper-scissors.
// params c chan string: a channel to send the result of the game to
func play(c chan string) {
	// Seed the random number generator with the current time in nanoseconds to ensure randomness in the game play each time the program is run: it's very important to seed the random number generator before generating random numbers.
	rand.Seed(time.Now().UnixNano())

	// Create a slice of strings representing the possible choices in the game
	compouterChoices := []string{"rock", "paper", "scissors"}

	// Generate a random index to select a choice from the slice
	i := rand.Intn(len(compouterChoices))

	// Send the randomly selected choice to the channel
	c <- compouterChoices[i]
}

func main() {
	// Create a channel of strings to communicate the result of the game
	c := make(chan string)

	// Start a goroutine that plays the game and sends the result to the channel
	go play(c)

	// Receive the result of the game from the channel
	result := <-c

	// Prompt the user to choose between rock, paper, or scissors
	fmt.Print("Choose between rock: 1, paper: 2, e scissors: 3  ==>   ")

	// Create a slice of strings representing the possible choices in the game
	myChoises := []string{"rock", "paper", "scissors"}

	var i int      // Variable to store the user's choice
	fmt.Scanln(&i) // Read the user's choice from the standard input

	// Check if the user's choice is valid
	if i < 1 || i > 3 {
		fmt.Println("Invalid choice. Please choose between 1, 2, or 3.")
		return
	}

	// Print the result of the game
	fmt.Println("You chose: ", myChoises[i-1])     // Print the user's choice
	fmt.Printf("The computer chose %s.\n", result) // Print the computer's choice

	// Determine the winner of the game
	switch result {
	case "rock":
		if myChoises[i-1] == "rock" {
			fmt.Println("It's a tie!")
		} else if myChoises[i-1] == "paper" {
			fmt.Println("You win!")
		} else {
			fmt.Println("You lose!")
		}
	case "paper":
		if myChoises[i-1] == "rock" {
			fmt.Println("You lose!")
		} else if myChoises[i-1] == "paper" {
			fmt.Println("It's a tie!")
		} else {
			fmt.Println("You win!")
		}
	case "scissors":
		if myChoises[i-1] == "rock" {
			fmt.Println("You win!")
		} else if myChoises[i-1] == "paper" {
			fmt.Println("You lose!")
		} else {
			fmt.Println("It's a tie!")
		}
	}
}

// Output: You chose rock. The computer chose rock. It's a tie!
