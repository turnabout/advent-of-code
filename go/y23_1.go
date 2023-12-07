package _go

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
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

	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	// Loop all lines
	for scanner.Scan() {
		line := scanner.Text()
		line = preProcessLine(line)
		// fmt.Println(line)

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

		lineNum := int(((firstDigit - '0') * 10) + (secondDigit - '0'))
		// fmt.Printf("Line num: %d\n", lineNum)
		sum += lineNum
	}

	fmt.Printf("%d\n", sum)
}

// findAllIndices finds all indices where the substring appears in the given string.
func findAllIndices(str, subStr string) []int {
	indices := []int{}

	startIndex := 0
	for {
		index := strings.Index(str[startIndex:], subStr)
		if index == -1 {
			break
		}

		indices = append(indices, startIndex+index)
		startIndex += index + len(subStr)
	}

	return indices
}

// lineWithDigitAdditions applies lineWithDigitAddition to the given line, for all numbers from 1 to 9.
func lineWithDigitAdditions(line string) string {
	for i := 1; i <= 9; i++ {
		line = lineWithDigitAddition(i, line)
	}

	return line
}

// lineWithDigitAddition returns the line with "number additions" applied to it.
//
// Every time a string representation of a number is encountered, its corresponding digit is added right before it.
// For example: "one4ffdsaffoury" becomes "1one4ffdsaf4foury"
func lineWithDigitAddition(num int, line string) string {
	numStr := strNums[num-1]
	digitStr := strconv.Itoa(num)

	// Find all indices with this number string
	indices := findAllIndices(line, numStr)

	// For each index, add the digit equivalent to the string right before it
	alreadyProcessed := 0
	for _, idx := range indices {
		line = withSubstring(line, digitStr, idx+alreadyProcessed)
		alreadyProcessed++
	}

	return line
}

func preProcessLine(line string) string {

	var accumStr string
	var accumStrIndexStart = 0

	// Loop from back to front, store first digit

	for i := 0; i < len(line); i++ {
		if accumStr == "" {
			accumStrIndexStart = i
		}

		accumStr += string(line[i])
		numVal := getStrNumVal(accumStr)

		// No match
		if numVal == -1 {
			// If accumulated string length higher than 1, we need to reset to right after where it started
			i = accumStrIndexStart
			accumStr = ""
			continue
		}

		// Complete match - insert number where accum string is
		if numVal > 0 {
			line = withSubstring(line, strconv.Itoa(numVal), accumStrIndexStart)

			// Reset accum string
			accumStr = ""

			// Put index to after where accum string started.
			// We need to set the new "i" value to right after where the accum string match started.
			//
			// This means we need to add +2 to its value, to also account for the newly added digit character.
			//
			// Here we add + 1, because the loop will automatically add +1 during its next loop
			i = accumStrIndexStart + 1
			accumStrIndexStart = 0
		}
	}

	// return lineWithDigitAdditions(line)

	// The below approach doesn't work, as it doesn't account for strings with numbers that overlap, like "oneight".
	/*
		return strings.NewReplacer(
			"one", "1",
			"two", "2",
			"three", "3",
			"four", "4",
			"five", "5",
			"six", "6",
			"seven", "7",
			"eight", "8",
			"nine", "9",
		).Replace(line)
	*/
	return line
}

// withSubstring returns an original string with a substring inserted at the given index
func withSubstring(original, substring string, index int) string {
	return original[:index] + substring + original[index:]
}

var strNums = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// getStrNumVal gets, from the given string, the associated integer value (excluding 0).
// - returns > 0: complete match with one of the number string. Returned value corresponds to matched integer value.
// - returns 0:   partial match with at least one of the number string.
// - returns -1:  no matches (not partial nor complete).
//
// Ended up not using this - was misguided in my approach.
// Using this to solve numbers doesn't work when we go backwards to get the last digit.
// Instead, it makes a lot more sense to just reuse the same exact algorithm as before, but pre-process the line by
// replacing all "one"s with "1"s, "two"s with "2"s, etc.
func getStrNumVal(sourceStr string) int {
	partialMatchFound := false

	for i, strNum := range strNums {
		match, err := stringsMatch(strNum, sourceStr)

		// No match - go to next one
		if err != nil {
			continue
		}

		// Complete match found - return equivalent value
		if match {
			return i + 1
		}

		// At least a partial match found - set
		partialMatchFound = true
	}

	if partialMatchFound {
		return 0
	} else {
		return -1
	}
}

// stringsMatch gets whether the two strings match.
//
// The "parent" string may be longer than the "child" string, which may partially match it.
// But the "child" string may NOT be longer than the "parent" string.
// For example: child = "sev" , parent = "seven"
//
// - returns true:  perfect match
// - returns false: partial match
// - non-nil error: no match
func stringsMatch(parent, child string) (bool, error) {
	childLen := len(child)

	for i := range parent {
		// Unable to get next child letter - child is too short
		// This means a partial match
		if childLen <= i {
			return false, nil
		}

		// If any letters mismatch, we have a bad match
		if child[i] != parent[i] {
			return false, errors.New("no match")
		}
	}

	// Complete match
	return true, nil
}
