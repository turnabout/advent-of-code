package _go

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*

Col 1: Destination range start
Col 2: Source range start
Col 3: Range length

Source seeds:
13 14 55 79

-----------------
seed-to-soil map:
50 98 2
52 50 48
-----------------

50 98 2
==========================
Seed 98 -> 99 (2 values)
Soil 50 -> 51 (2 values)

52 50 48
==========================
Seed 50 -> 97  (48 values)
Soil 52 -> 99  (48 values)

*/

type AlmanacMap struct {
	Source      string          `json:"source"`
	Destination string          `json:"destination"`
	Rows        []AlmanacMapRow `json:"rawRows"`
}

/*
func NewAlmanacMapWithOnlyMatchingRows(inputMap AlmanacMap, destinationRows []AlmanacMapRow) AlmanacMap {
	newMap := AlmanacMap{
		Source:      inputMap.Source,
		Destination: inputMap.Destination,
		Rows:        []AlmanacMapRow{},
	}

	for _, dstRow := range destinationRows {
		for _, srcRow := range inputMap.Rows {
			// Check if overlap between input map (which is the source)'s rows.DST and destination rows (which is destination)
			if dstRow.SrcRangeStart >= srcRow.DstL

		}
	}

	return newMap
}
*/

func (m AlmanacMap) getRowWithLowestDestinationRange() AlmanacMapRow {
	var lowest AlmanacMapRow = AlmanacMapRow{
		SrcRangeStart: -1,
	}

	for _, row := range m.Rows {
		if lowest.SrcRangeStart == -1 {
			lowest = row
			continue
		}

		if lowest.DstRangeStart < row.DstRangeStart {
			lowest = row
		}
	}

	return lowest
}

func (m AlmanacMap) MapInputs(inputs []int) []int {
	res := []int{}

	for _, input := range inputs {
		row, ok := m.FindRowWithSourceInput(input)

		// If no row found, map is 1:1
		if !ok {
			// fmt.Printf("%d -> %d\n", input, input)
			res = append(res, input)
			continue
		}

		// Row found, apply map
		mappedVal := row.MapSourceToDestination(input)
		// fmt.Printf("%d -> %d\n", input, mappedVal)
		res = append(res, mappedVal)
	}

	return res
}

func (m AlmanacMap) FindRowWithSourceInput(input int) (AlmanacMapRow, bool) {
	for _, row := range m.Rows {
		if input >= row.SrcRangeStart && input <= row.SrcRangeStart+row.RangeLength-1 {
			return row, true
		}
	}

	return AlmanacMapRow{}, false
}

func NewAlmanacMap(blockLines []string) AlmanacMap {
	retMap := AlmanacMap{Rows: []AlmanacMapRow{}}

	// Parse first line, which contains the titles
	names := strings.Split(strings.Split(blockLines[0], " ")[0], "-to-")
	retMap.Source = names[0]
	retMap.Destination = names[1]

	for _, line := range blockLines[1:] {
		retMap.Rows = append(retMap.Rows, NewAlmanacMapRow(line))
	}

	return retMap
}

type AlmanacMapRow struct {
	SrcRangeStart int
	DstRangeStart int
	RangeLength   int
}

/*
Col 1: Destination range start
Col 2: Source range start
Col 3: Range length

Source input: 55
Used Map:     52 50
-------------------
Result: 57

*/

func (r AlmanacMapRow) MapSourceToDestination(source int) int {
	return source + (r.DstRangeStart - r.SrcRangeStart)
}

func NewAlmanacMapRow(line string) AlmanacMapRow {
	fields := strings.Fields(line)
	fieldsInt := []int{}
	for _, f := range fields {
		fI, err := strconv.Atoi(f)
		if err != nil {
			log.Fatalf("failed to convert input field '%s' to int: %w", f, err)
		}
		fieldsInt = append(fieldsInt, fI)
	}

	return AlmanacMapRow{
		DstRangeStart: fieldsInt[0],
		SrcRangeStart: fieldsInt[1],
		RangeLength:   fieldsInt[2],
	}
}

const y23_5_shortInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func printAlmanacMaps(maps []AlmanacMap) {
	bs, _ := json.MarshalIndent(maps, "", "\t")
	fmt.Println(string(bs))
}

func resolveAlmanacMaps(scanner *bufio.Scanner) []AlmanacMap {
	// Resolve maps
	var blockLines []string
	maps := []AlmanacMap{}
	for scanner.Scan() {
		line := scanner.Text()

		// Empty line, we gathered a whole block
		if line == "" {
			maps = append(maps, NewAlmanacMap(blockLines))
			blockLines = []string{}
			continue
		}

		// Add to block string
		blockLines = append(blockLines, line)
	}

	return maps
}

func resolveInputVals(line string) []int {
	fields := strings.Fields(strings.Split(line, ": ")[1])
	fieldsInt := []int{}
	for _, f := range fields {
		fI, err := strconv.Atoi(f)
		if err != nil {
			log.Fatalf("failed to convert input field '%s' to int: %w", f, err)
		}
		fieldsInt = append(fieldsInt, fI)
	}

	return fieldsInt
}

func printInputValues(name string, vals []int) {
	fmt.Println(strings.Title(name))
	fmt.Print("{")
	for i, v := range vals {
		fmt.Printf("%d", v)
		if i < len(vals)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("}")
	fmt.Println("-------------------")
}

func getValsLowestValue(vals []int) int {
	lowest := vals[0]

	for _, val := range vals {
		if val < lowest {
			lowest = val
		}
	}

	return lowest
}

func (Runner) Y23_5_1(input string) {
	input = y23_5_shortInput + "\n\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Resolve initial input values from first line
	scanner.Scan()
	inputValues := resolveInputVals(scanner.Text())
	scanner.Scan()

	// Resolve maps from the rest of the lines
	maps := resolveAlmanacMaps(scanner)

	// Loop maps, make the input values go through them as a sort of pipeline
	for _, m := range maps {
		fmt.Println()
		printInputValues(m.Source, inputValues)
		inputValues = m.MapInputs(inputValues)
	}

	printInputValues("Final", inputValues)

	fmt.Printf("Lowest value (answer): %d\n", getValsLowestValue(inputValues))
}

/*
Idea: work your way backwards from the destination ranges in the last map that are lowest, until we find seeds that
correspond to that range.

By looking at the "humidity-to-locations map", we can already see which of the maps will result in the lowest numbers.
*/

// Lowest (answer): 15290096
func (Runner) Y23_5_2(input string) {
	input = y23_5_shortInput + "\n\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Resolve initial input values from first line
	scanner.Scan()
	inputValues := resolveInputVals(scanner.Text())
	scanner.Scan()

	// Resolve maps from the rest of the lines
	maps := resolveAlmanacMaps(scanner)

	fmt.Printf("%d\n", bruteForce(inputValues, maps))
	return

	// Loop maps, make the input values go through them as a sort of pipeline
	for _, m := range maps {
		fmt.Println()
		printInputValues(m.Source, inputValues)
		inputValues = m.MapInputs(inputValues)
	}

	printInputValues("Final", inputValues)

	fmt.Printf("Lowest value (answer): %d\n", getValsLowestValue(inputValues))
}

func resolveSeed(seed int, maps []AlmanacMap) int {
	inputValues := []int{seed}

	// Loop maps, make the input values go through them as a sort of pipeline
	for _, m := range maps {
		inputValues = m.MapInputs(inputValues)
	}

	return inputValues[0]
}

/*
func resolveInputValuePairs(pairs []int, maps []AlmanacMap) []int {
	ints := []int{}

	for i := 0; i < len(pairs); i += 2 {
		numStart := pairs[i]
		numRange := pairs[i+1]

		// Add nums
		for num := numStart; num < numStart+numRange; num++ {
			resolved := resolveSeed(num, maps)
			if resolved < lowest {
				lowest = resolved
			}
			// res = append(res, num)
		}
	}

	return ints
}
*/

func bruteForce(pairs []int, maps []AlmanacMap) int {
	lowest := 9999999999

	for i := 0; i < len(pairs); i += 2 {
		numStart := pairs[i]
		numRange := pairs[i+1]

		// Add nums
		for num := numStart; num < numStart+numRange; num++ {
			resolved := resolveSeed(num, maps)
			if resolved < lowest {
				lowest = resolved
			}
			// res = append(res, num)
		}
	}

	return lowest
}
