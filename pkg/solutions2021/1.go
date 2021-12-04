package solutions2021

import (
	"log"
	"strconv"
	"strings"
)

func (s Solution) S1_1(input string) string {
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

func (s Solution) S1_2(input string) string {
	inputs := strings.Split(input, "\n")
	inputsLen := len(inputs)

	increaseCount := 0
	previousSum := -1

	for i, inputLine := range inputs {
		inputNum, err := strconv.Atoi(inputLine)
		if err != nil {
			log.Fatal(err)
		}

		// If we don't have two more numbers for sum, stop
		if (i + 2) >= inputsLen {
			break
		}

		// Get next two numbers
		inputNum2, err := strconv.Atoi(inputs[i+1])
		if err != nil {
			log.Fatal(err)
		}

		inputNum3, err := strconv.Atoi(inputs[i+2])
		if err != nil {
			log.Fatal(err)
		}

		// Get the current sum
		sum := inputNum + inputNum2 + inputNum3

		// Edge case (first iteration)
		if previousSum == -1 {
			previousSum = sum
			continue
		}

		if sum > previousSum {
			increaseCount++
		}

		previousSum = sum
	}

	return strconv.Itoa(increaseCount)
}
