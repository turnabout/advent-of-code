package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var (
	colorCyan *color.Color
	colorRed  *color.Color
)

func init() {
	colorCyan = color.New(color.FgCyan)
	colorRed = color.New(color.FgRed)
}

type CBoard [][]rune

func (c CBoard) IsOob(x, y int) bool {
	return x < 0 || y < 0 || y >= len(c) || x >= len(c[0])
}

func (c CBoard) CharAtCoordsIsSymbolish(x, y int) bool {
	if x < 0 || y < 0 || y >= len(c) || x >= len(c[0]) {
		return false
	}

	return charIsSymbol(c[y][x])
}

func (b CBoard) PrintAllSymbolInBoundaries(topLeftX, topLeftY, botRightX, botRightY int) {
	for x := topLeftX; x <= botRightX; x++ {
		y := topLeftY

		if !b.IsOob(x, y) {
			fmt.Printf("%c", b[y][x])
		}
	}

	for x := topLeftX; x <= botRightX; x++ {
		y := botRightY

		if !b.IsOob(x, y) {
			fmt.Printf("%c", b[y][x])
		}
	}

	for y := topLeftY; y <= botRightY; y++ {
		x := topLeftX
		if !b.IsOob(x, y) {
			fmt.Printf("%c", b[y][x])
		}
	}

	for y := topLeftY; y <= botRightY; y++ {
		x := botRightX
		if !b.IsOob(x, y) {
			fmt.Printf("%c", b[y][x])
		}
	}
}

func (b CBoard) HasSymbolInBoundaries(topLeftX, topLeftY, botRightX, botRightY int) bool {

	for x := topLeftX; x <= botRightX; x++ {
		y := topLeftY

		if b.CharAtCoordsIsSymbolish(x, y) {
			return true
		}
	}

	for x := topLeftX; x <= botRightX; x++ {
		y := botRightY

		if b.CharAtCoordsIsSymbolish(x, y) {
			return true
		}
	}

	for y := topLeftY; y <= botRightY; y++ {
		x := topLeftX
		if b.CharAtCoordsIsSymbolish(x, y) {
			return true
		}
	}

	for y := topLeftY; y <= botRightY; y++ {
		x := botRightX
		if b.CharAtCoordsIsSymbolish(x, y) {
			return true
		}
	}

	/*
		for x := topLeftX; x <= botRightX; x++ {
			// Check OOB
			if x < 0 {
				continue
			}

			for y := topLeftY; y <= botRightY; y++ {
				// Check OOB
				if y < 0 || y >= len(b) || x >= len(b[y]) {
					continue
				}

				if charIsSymbol(b[y][x]) {
					return true
				}
			}
		}
	*/

	return false
}

func NewCBoard(input string) CBoard {
	cb := CBoard{}

	y := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		cb = append(
			cb,
			[]rune{},
		)

		line := scanner.Text()

		for _, c := range line {
			cb[y] = append(
				cb[y],
				c,
			)
		}

		y++
	}

	return cb
}

func printCBoard(cb CBoard) {
	for _, row := range cb {
		for _, col := range row {
			colorCyan.Printf("%c", col)
		}
		fmt.Printf("\n")
	}
}

func charIsSymbolish(char rune) bool {
	return char != '.'
}

func charIsSymbol(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

type Board struct {
	SymbolCoords []Coords       `json:"symbol_coords"`
	NumberCoords []NumberCoords `json:"number_coords"`
}

type Coords struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var (
	xLen = 0
	yLen = 0
)

func (c Coords) TouchesAny(coordsList []Coords) bool {
	topLeft := Coords{
		X: c.X - 1,
		Y: c.Y - 1,
	}

	botRight := Coords{
		X: c.X + 1,
		Y: c.Y + 1,
	}

	for _, listC := range coordsList {
		if listC.X >= topLeft.X && listC.Y >= topLeft.Y &&
			listC.X <= botRight.X && listC.Y <= botRight.Y {
			return true
		}
	}

	return false
}

type NumberCoords struct {
	Num       int      `json:"num"`
	AllCoords []Coords `json:"all_coords"`
}

func (nc NumberCoords) TouchesAny(coordsList []Coords) bool {
	for _, coord := range nc.AllCoords {
		if coord.TouchesAny(coordsList) {
			return true
		}
	}

	return false
}

func (Runner) Y23_3_1(input string) {

	/*
			input = `467..114..
		...*......
		..35..633.
		......#...
		617*......
		.....+.58.
		..592.....
		......755.
		...$.*....
		.664.598..`
	*/

	/*
			input = `...............................930...................................283...................453.34.............................867....282....
		....=.........370...........................48..456......424...-.341*.....554...*807.571............971..958............166......*..........
		..159.........../..........539*.....73......-...*.......+....954.........*.....7.......*........*.....*....*.....405$..*.......31.........15
		...............................873..*............726.............94.......126.........699....253....584..750................................
		.660.................................336.....391.................*....860......76..................................435....576.....-.........
		.................................888............*924...55......308.......*91.........446...535......87...136/........*...*........793.=351..
		...........826...949...120...985..&....................*.......................462.../......*.........*.......358..932..599.479*............
		............../.....%..*......%...............151.304..931..471.......601.....*............765........805....%..................149...345...
		........................216..........................+......*............#..906...-......................................105...........&....
		.......&..827*327.375-.................923.......*..........630......851..........459..656.......340.432........915.288....#.865*...........
		.....693......................866......*......575.970...........201...................%........%...*.=...........+....*..........305.....666
		.........%536......345..............166........................*....@905....863.&...........916..212.....386*963.....183....................`
	*/

	b := NewCBoard(input)
	// printCBoard(b)

	total := 0

	for y, row := range b {
		accumNum := ""

		for x, char := range row {

			xFactor := 0

			// Add to accumulated num
			if unicode.IsDigit(char) {
				accumNum += string(char)

				// Continue to next one if NOT last character
				if x != len(row)-1 {
					continue
				} else {
					// Else, keep processing
					xFactor = 1
				}
			}

			// Close off accumulated num
			if accumNum != "" {
				// Convert accumulated num to int
				accumNumInt, err := strconv.Atoi(accumNum)
				if err != nil {
					log.Fatalf("failed to convert '%s' to integer", accumNum)
				}

				// fmt.Printf("Encountered %d\n", accumNumInt)

				xCalc := x + xFactor

				// Get whether the number collides with a symbol
				topLeftX := xCalc - len(accumNum) - 1
				topLeftY := y - 1

				botRightX := topLeftX + len(accumNum) + 1
				botRightY := y + 1

				// fmt.Printf("encountered %d\n", accumNumInt)

				if b.HasSymbolInBoundaries(topLeftX, topLeftY, botRightX, botRightY) {
					// fmt.Printf("\tAdding number %d\n", accumNumInt)
					total += accumNumInt

					// TODO
					colorCyan.Printf(accumNum)
				} else {
					// TODO
					colorRed.Printf(accumNum)

					// fmt.Printf("%d is not in. Chars:\n", accumNumInt)
					// b.PrintAllSymbolInBoundaries(topLeftX, topLeftY, botRightX, botRightY)
					// fmt.Println()
					// fmt.Println()
				}

				// fmt.Printf("checking from %d,%d to %d,%d for num %d\n", topLeftX, topLeftY, botRightX, botRightY, accumNumInt)

				// Print accum num

				// Reset
				accumNum = ""

				// Print number

				// charIsSymbol(char)
			}

			// TODO
			if !unicode.IsDigit(char) {
				fmt.Printf("%c", char)
			}
		}

		// TODO
		fmt.Println()
	}

	fmt.Printf("Result: %d\n", total)

	return

	/*
		symbolsCords := []Coords{}
		numbersCoords := []NumberCoords{}

		// Load symbol coords
		scanner := bufio.NewScanner(strings.NewReader(input))
		y := 0
		for scanner.Scan() {
			line := scanner.Text()

			if y == 0 {
				xLen = len(line)
			}

			accumNum := ""

			for x, char := range line {
				// If is number, add to number accumulation
				if unicode.IsDigit(char) {
					accumNum += string(char)
					continue
				}

				// If accumNum was present and reached non-num, reset accumNum & record its coordinates
				if accumNum != "" {
					// Convert accumulated num to int
					accumNumInt, err := strconv.Atoi(accumNum)
					if err != nil {
						log.Fatalf("failed to convert '%s' to integer", accumNum)
					}

					digitLen := len(accumNum)
					digitStartX := x - digitLen

					allCoords := []Coords{}

					// Record all individual coords making up the number
					for i := 0; i < digitLen; i++ {
						allCoords = append(
							allCoords,
							Coords{
								X: digitStartX + i,
								Y: y,
							},
						)
					}

					// Record coordinates of number
					numbersCoords = append(
						numbersCoords,
						NumberCoords{
							Num:       accumNumInt,
							AllCoords: allCoords,
						},
					)

					// Reset string
					accumNum = ""
				}

				// If character is dot, ignore
				if char == '.' {
					continue
				}

				fmt.Printf("%c at %d,%d\n", char, x, y)

				// Encountered symbol!
				symbolsCords = append(
					symbolsCords,
					Coords{
						X: x,
						Y: y,
					},
				)
			}

			y++
			yLen++
		}

		accumRes := 0

		// For each number coords, get whether they collide with all the symbol coords. If so, record them
		for _, numbersCoord := range numbersCoords {
			if numbersCoord.TouchesAny(symbolsCords) {
				fmt.Printf("adding %d\n", numbersCoord.Num)
				accumRes += numbersCoord.Num
			}
		}

		fmt.Printf("%d\n", accumRes)
	*/

	/*
		printBoard(Board{
			SymbolCoords: symbolsCords,
			NumberCoords: numbersCoords,
		})

		fmt.Printf("X len: %d\n", xLen)
		fmt.Printf("Y len: %d\n", yLen)
	*/
}

func (Runner) Y23_3_2(input string) {
}

func printBoard(b Board) {
	bs, err := json.MarshalIndent(&b, "\t", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
