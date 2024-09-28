# Number Guessing Game

A simple CLI-based number guessing game built in Go. Test your luck by guessing the number the computer has randomly selected!

## Features

- **Difficulty Levels**: Choose from Easy, Medium, or Hard, each with different number of chances.
- **Hints**: Receive helpful hints after each incorrect guess.
- **Timer**: Track how long it takes to guess the number.
- **High Scores**: Compete against your best scores for each difficulty level.
- **Multiple Rounds**: Play multiple rounds until you decide to quit.

## Requirements

- Go 1.20 or higher

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/yourusername/number-guessing-game.git
    cd number-guessing-game
    ```

2. **Build the application**:

    ```bash
    go build -o guess ./cmd/guess
    ```

3. **Run the game**:

    ```bash
    ./guess
    ```

## Usage

Upon running the game, you'll be greeted with a welcome message and prompted to select a difficulty level. Enter the corresponding number to choose your difficulty:

1. **Easy**: 10 chances
2. **Medium**: 5 chances
3. **Hard**: 3 chances

After selecting the difficulty, start guessing the number between 1 and 100. After each guess, you'll receive feedback indicating whether the target number is higher or lower than your guess, along with a hint to help you.

The game ends when you either guess the correct number or run out of chances. After each round, you can choose to play again or exit.

## High Scores

High scores are tracked based on the fewest number of attempts and the shortest time taken to guess the number for each difficulty level. High scores are stored in `highscores.json` in the project directory.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## Contact

For any questions or suggestions, feel free to open an issue or contact [adwaithrajeev@gmail.com].

Happy Guessing!
