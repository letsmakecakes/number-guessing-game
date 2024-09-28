package game

import (
	"fmt"
	"math/rand"
	"number-guessing-game/pkg/utils"
)

type Diffculty struct {
	Name    string
	Chances int
}

func Start() {
	diff := getDifficulty()

	target := rand.Intn(100) + 1

	loop(diff.Chances, target)
}

func getDifficulty() *Diffculty {
	var diff Diffculty
	for {
		prompt := "Enter your choice: "
		input, err := utils.ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		choice, err := utils.ParseInt(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter 1, 2, or 3.")
			continue
		}

		switch choice {
		case 1:
			diff = Diffculty{Name: "Easy", Chances: 10}
		case 2:
			diff = Diffculty{Name: "Medium", Chances: 5}
		case 3:
			diff = Diffculty{Name: "Hard", Chances: 3}
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
			continue
		}
		break
	}

	fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", diff.Name)
	fmt.Println("Let's start the game!")

	return &diff
}

func loop(chances int, target int) {
	attempts := 0
	for attempts < chances {
		fmt.Printf("Attempt %d/%d\n", attempts+1, chances)
		prompt := "Enter your guess: "
		input, err := utils.ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		guess, err := utils.ParseInt(input)
		if err != nil || guess < 1 || guess > 100 {
			fmt.Println("Invalid guess. Please enter a number between 1 and 100.")
			continue
		}

		attempts++

		if guess == target {
			fmt.Printf("Congratulations! You guess the correct number in %d attempts.\n", attempts)
			return
		} else if guess < target {
			fmt.Println("Incorret! The number is greater than", guess)
		} else {
			fmt.Println("Incorret! The number is less than", guess)
		}

		fmt.Println()
	}

	fmt.Printf("Sorry, you've run out of chances. The correct number was %d.\n", target)
}
