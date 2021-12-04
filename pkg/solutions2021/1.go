package solutions2021

import (
	"log"
	"strconv"
	"strings"
)

func (s Solution) S1(input string) string {
	inputs := strings.Split(input, "\n")

	increaseCount := 0
	previousNum := -1

	for _, inputLine := range inputs {
		inputNum, err := strconv.Atoi(inputLine)
		if err != nil {
			log.Fatal(err)
		}

		// Edge case (first iteration)
		if previousNum == -1 {
			previousNum = inputNum
			continue
		}

		if inputNum > previousNum {
			increaseCount++
		}

		previousNum = inputNum
	}

	return strconv.Itoa(increaseCount)
}
