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
	Id             int
	WinningNumbers Numbers
	Numbers        Numbers

	// CardCount is the count of this card we have.
	CardCount int
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

// GetMatchingNumbers gets the amount of Numbers that are matching with winning Numbers.
func (c Card) GetMatchingNumbers() int {
	return c.WinningNumbers.CountSame(c.Numbers)
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

func NewCardFromLineInput(line string, id int) Card {
	card := Card{Id: id, CardCount: 1}

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
	i := 0
	for scanner.Scan() {
		i++
		cards = append(
			cards,
			NewCardFromLineInput(scanner.Text(), i),
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
	input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	cards := NewCardsFromInput(input)
	cardsEncountered := 0

	for i, card := range cards {
		matchingNumsCount := card.GetMatchingNumbers()
		cardsEncountered += card.CardCount

		if matchingNumsCount == 0 {
			continue
		}

		wonCardsIds := []string{}

		// Update following cards' CardCount
		for j := 1; j <= matchingNumsCount; j++ {
			nextCardIdx := i + j

			// OOB, break out
			if nextCardIdx >= len(cards) {
				break
			}

			wonCardsIds = append(wonCardsIds, strconv.Itoa(cards[nextCardIdx].Id))
			cards[nextCardIdx].CardCount += card.CardCount
		}

		/*
			fmt.Printf(
				"Card %d (count: %d) has %d matching numbers. Winning %d of: {%s}\n",
				card.Id,
				card.CardCount,
				matchingNumsCount,
				card.CardCount,
				strings.Join(wonCardsIds, ","),
			)
		*/
	}

	fmt.Println("---------------------")
	fmt.Printf("Total cards: %d\n", cardsEncountered)
	fmt.Println("---------------------")
	// Answer with full input:
	/*
		---------------------
		Total cards: 8570000
		---------------------
	*/
}
