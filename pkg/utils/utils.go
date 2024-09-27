package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func ParseInt(input string) (int, error) {
	return strconv.Atoi(input)
}
