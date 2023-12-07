package _go

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CubeGame struct {
	Id        int
	ShownSets []CubeSet
}

// NewCubeGamesFromString creates a list of CubeGame's from the full input string.
func NewCubeGamesFromString(input string) []*CubeGame {
	cgs := []*CubeGame{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		cg := NewCubeGameFromString(line)
		cgs = append(cgs, &cg)
	}

	return cgs
}

// NewCubeGameFromString creates a CubeGame from a given line input string.
// Example: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func NewCubeGameFromString(input string) CubeGame {
	cg := CubeGame{
		ShownSets: []CubeSet{},
	}

	// Split into two sections (before and after colon)
	sections := strings.Split(input, ": ")

	// Extract ID from left section
	id, err := strconv.Atoi(strings.Split(sections[0], " ")[1])
	if err != nil {
		log.Fatalf("failed to extract cube game ID: %w (original input: '%s')", err, input)
	}
	cg.Id = id

	// Extract the CubeSets
	cubeSetSections := strings.Split(sections[1], "; ")
	for _, cubeSetSection := range cubeSetSections {
		cg.ShownSets = append(
			cg.ShownSets,
			NewCubeSet(cubeSetSection),
		)
	}

	return cg
}

// Highest computes a CubeSet containing the highest values attained in the CubeGame.
func (cg CubeGame) Highest() CubeSet {
	res := CubeSet{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, set := range cg.ShownSets {
		if set.Red > res.Red {
			res.Red = set.Red
		}
		if set.Green > res.Green {
			res.Green = set.Green
		}
		if set.Blue > res.Blue {
			res.Blue = set.Blue
		}
	}

	return res
}

// IsSetPossible gets whether the given comparison set is possible with this CubeSet.
func (cg CubeGame) IsSetPossible(cmpSet CubeSet) bool {
	highest := cg.Highest()

	return highest.Red <= cmpSet.Red && highest.Green <= cmpSet.Green && highest.Blue <= cmpSet.Blue
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

// NewCubeSet creates a new CubeSet from the given input string describing it.
// Example: "3 blue, 4 red"
func NewCubeSet(input string) CubeSet {
	cs := CubeSet{}

	sections := strings.Split(input, ", ")
	for _, section := range sections {
		parts := strings.Split(section, " ")
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("failed to extract cube count: %w (original input: '%s')", err, input)
		}

		cubeType := parts[1]
		switch cubeType {
		case "red":
			cs.Red = count
		case "green":
			cs.Green = count
		case "blue":
			cs.Blue = count
		}
	}

	return cs
}

// Power gets the power of the set of CubeSet.
// > The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
func (cs CubeSet) Power() int {
	return cs.Red * cs.Green * cs.Blue
}

func (Runner) Y23_2_1(input string) {
	cgs := NewCubeGamesFromString(input)

	// Which games would have been possible if the bag contained only:
	// - 12 red cubes
	// - 13 green cubes
	// - 14 blue cubes
	cmpSet := CubeSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sum := 0

	for _, cg := range cgs {
		if cg.IsSetPossible(cmpSet) {
			sum += cg.Id
		}
	}

	fmt.Printf("%d\n", sum)
}

func (Runner) Y23_2_2(input string) {
	cgs := NewCubeGamesFromString(input)
	powerSum := 0

	for _, cg := range cgs {
		powerSum += cg.Highest().Power()
	}

	fmt.Printf("%d\n", powerSum)
}
