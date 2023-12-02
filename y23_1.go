package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func (Runner) Y23_1_1(input string) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	// Loop all lines
	for scanner.Scan() {
		line := scanner.Text()

		var firstDigit rune
		var secondDigit rune

		// Loop from back to front, store first digit
		for _, char := range line {
			if unicode.IsDigit(char) {
				firstDigit = char
				break
			}
		}

		// Loop from front to back, store second digit
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if unicode.IsDigit(rune(char)) {
				secondDigit = rune(char)
				break
			}
		}

		sum += int(((firstDigit - '0') * 10) + (secondDigit - '0'))
	}

	fmt.Printf("%d\n", sum)
}

func (Runner) Y23_1_2(input string) {
	fmt.Println("Part 2")
}
