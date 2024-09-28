package game

import (
	"fmt"
	"number-guessing-game/pkg/utils"
)

func Start() {
	printWelcomeMessage()
	chances, err := getChances()
	if err != nil {
		fmt.Errorf("error getting chances: %v", err)
	}
}

func printWelcomeMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")
}

func getChances() (int, error) {
	prompt := "Enter your choice: "
	input, err := utils.ReadInput(prompt)
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)

	}

	chances, err := utils.ParseInt(input)
	if err != nil {
		return 0, fmt.Errorf("error parsing int: %v", err)
	}

	return chances, nil
}
