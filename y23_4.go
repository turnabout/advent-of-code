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
	card := Card{Id: id}

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
	ogCards := NewCardsFromInput(input)
	cardsLeft := make([]Card, len(ogCards))
	copy(cardsLeft, ogCards)

	cardsEncountered := 0

	for i := 0; i < len(cardsLeft); i++ {
		card := cardsLeft[i]
		matchingNumsCount := card.GetMatchingNumbers()
		cardsEncountered++

		if matchingNumsCount == 0 {
			continue
		}

		originalI := card.Id - 1

		wonCardsIndexStart := originalI + 1
		wonCardsIndexEnd := originalI + matchingNumsCount + 1

		// Make sure we don't go OOB (can't get cards beyond the limit)
		if wonCardsIndexEnd > len(cardsLeft) {
			wonCardsIndexEnd = len(cardsLeft)
		}

		wonCards := ogCards[wonCardsIndexStart:wonCardsIndexEnd]
		wonCardsIds := []string{}
		for _, wonCard := range wonCards {
			wonCardsIds = append(wonCardsIds, strconv.Itoa(wonCard.Id))
		}

		cardsLeft = append(cardsLeft, wonCards...)

		// This approach is quite inefficient and slow, especially with the prints
		// TODO: Look at a different approach.
		// Instead, we could probably add a "CardCount" field to each cards.
		// Then whenever there is a winning card, it adds '+n' to the "CardCount" of cards ahead of it,
		// where 'n' = the "CardCount" of the current card.
		/*
			fmt.Printf(
				"Card %d has %d matching numbers. Winning: {%s}\n",
				card.Id,
				matchingNumsCount,
				strings.Join(wonCardsIds, ","),
			)
		*/
	}

	fmt.Println("---------------------")
	fmt.Printf("Total cards: %d\n", cardsEncountered)
	fmt.Println("---------------------")
	/*
		---------------------
		Total cards: 8570000
		---------------------
	*/
}
