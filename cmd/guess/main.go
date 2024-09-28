package main

import (
	"fmt"
	"number-guessing-game/internal/game"
	"number-guessing-game/pkg/utils"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")

	for {
		game.Start()

		var playAgain string
		for {
			prompt := "Do you want to play again? (y/n): "
			playAgain, err := utils.ReadInput(prompt)
			if err != nil {
				fmt.Println("Error reading input. Please try again.")
				continue
			}
			playAgain = strings.ToLower(playAgain)
			if playAgain == "y" || playAgain == "n" {
				break
			}
			fmt.Println("Invalid input. Please enter 'y' or 'n")
		}

		if playAgain != "y" {
			fmt.Println("Thank you for playing! Goodbye!")
			break
		}
	}
}
