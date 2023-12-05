package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Numbers []int

// CountSame gets the amount of numbers in the given Numbers are present in this one.
func (nums Numbers) CountSame(cmpNums Numbers) int {
	count := 0

	for _, n1 := range nums {
		for _, n2 := range cmpNums {
			if n1 == n2 {
				count++
			}
		}
	}

	return count
}

type Card struct {
	WinningNumbers Numbers
	Numbers        Numbers
}

func (c Card) GetScore() int {
	sameCount := c.WinningNumbers.CountSame(c.Numbers)

	if sameCount == 1 {
		return 1
	}

	sameCount--

	return int(
		math.Pow(2, float64(sameCount)),
	)
}

func PrintCardsList(cards []Card) {
	for i, c := range cards {
		fmt.Printf(
			"Card %d: %s | %s\n",
			i+1,
			strings.Join(intsToFields(c.WinningNumbers), " "),
			strings.Join(intsToFields(c.Numbers), " "),
		)
	}
}

func fieldsToInts(fields []string) []int {
	ints := []int{}
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			log.Fatalf("failed to convert field '%s' to int: %w", f, err)
		}

		ints = append(ints, n)
	}
	return ints
}

func intsToFields(ints []int) []string {
	fields := []string{}
	for _, i := range ints {
		fields = append(fields, strconv.Itoa(i))
	}

	return fields
}

func NewCardFromLineInput(line string) Card {
	card := Card{}

	numbersSectionStr := strings.Split(line, ": ")[1]
	numbersStrArr := strings.Split(numbersSectionStr, " | ")

	// Get string for both winning nums (left of '|') and actual nums (right of '|')
	winningNumsStr := numbersStrArr[0]
	numsStr := numbersStrArr[1]

	card.WinningNumbers = fieldsToInts(
		strings.Fields(winningNumsStr),
	)

	card.Numbers = fieldsToInts(
		strings.Fields(numsStr),
	)

	return card
}

func NewCardsFromInput(input string) []Card {
	cards := []Card{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		cards = append(
			cards,
			NewCardFromLineInput(scanner.Text()),
		)
	}

	return cards
}

func (Runner) Y23_4_1(input string) {
	cards := NewCardsFromInput(input)

	accumScore := 0
	for i, card := range cards {
		score := card.GetScore()
		fmt.Printf("Card %d has score: %d\n", i+1, score)
		accumScore += score
	}

	fmt.Println("---------------------")
	fmt.Printf("Total score: %d\n", accumScore)
	fmt.Println("---------------------")
}

func (Runner) Y23_4_2(input string) {
}
