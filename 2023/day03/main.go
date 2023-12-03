package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// BOILER PLATE --------------------------------------------------------------------
	start := time.Now()
	log.Printf("Starting... %s", start.Format("Jan 2 15:04:05 2006 MST"))

	var inputFileName string
	flag.StringVar(&inputFileName, "inputFileName", "input.txt", "Name of the input file")
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := string(inputBytes)
	// fmt.Println("Input bytes:", inputBytes)
	// fmt.Println("Input string", inputString)
	inputString = strings.TrimSpace(inputString)
	// BOILER PLATE --------------------------------------------------------------------

	numbers, symbols := parseInput(inputString)

	part1 := solvePart1(numbers, symbols)
	part2 := solvePart2(numbers, symbols)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePart2(numbers []Number, symbols []Symbol) int {
	ret := 0

	for _, symbol := range symbols {
		if symbol.value != "*" {
			fmt.Println("Skipping non gear:", symbol.value)
			continue
		}

		symSurroundingCoords := surroundingCoordinatesSym(symbol)
		surroundingNumbers := []Number{}

		for _, number := range numbers {
			for _, coord := range symSurroundingCoords {
				if coordInNum(coord, number) {
					surroundingNumbers = append(surroundingNumbers, number)
					break
				}
			}
		}

		if len(surroundingNumbers) == 2 {
			ret += surroundingNumbers[0].value * surroundingNumbers[1].value
		}
	}
	return ret

}

func surroundingCoordinatesSym(symbol Symbol) [][2]int {
	coords := make([][2]int, 0, 8)

	for x := symbol.x - 1; x <= symbol.x+1; x++ {
		for y := symbol.y - 1; y <= symbol.y+1; y++ {
			if x == symbol.x && y == symbol.y {
				continue
			}
			coords = append(coords, [2]int{x, y})
		}
	}

	return coords
}

func coordInNum(coord [2]int, num Number) bool {
	if num.minX <= coord[0] && coord[0] <= num.maxX && coord[1] == num.y {
		return true
	}
	return false
}

func solvePart1(numbers []Number, symbols []Symbol) int {
	ret := 0
NumLoop:
	for _, number := range numbers {

		checkCoords := surroundingCoordinates(number)

		for _, symbol := range symbols {
			for _, checkCoord := range checkCoords {
				if checkCoord[0] == symbol.x && checkCoord[1] == symbol.y {
					ret += number.value

					continue NumLoop
				}
			}
		}
	}

	return ret
}

func surroundingCoordinates(num Number) [][2]int {
	ret := [][2]int{
		{num.minX - 1, num.y},
		{num.maxX + 1, num.y},
	}

	for x := num.minX - 1; x <= num.maxX+1; x++ {
		ret = append(
			ret,
			[2]int{x, num.y + 1},
			[2]int{x, num.y - 1},
		)
	}
	return ret
}

func parseInput(input string) ([]Number, []Symbol) {
	digitRegexp := regexp.MustCompile(`\d+`)
	symbolRegexp := regexp.MustCompile(`[^\d.]`)

	lines := strings.Split(input, "\n")

	numbers := []Number{}
	symbols := []Symbol{}

	for y, line := range lines {
		digitIndexes := digitRegexp.FindAllStringIndex(line, -1)
		digitStrings := digitRegexp.FindAllString(line, -1)

		for i := 0; i < len(digitIndexes); i++ {
			newNumValue, _ := strconv.Atoi(digitStrings[i])
			newNum := Number{
				value: newNumValue,
				y:     y,
				minX:  digitIndexes[i][0],
				maxX:  digitIndexes[i][1] - 1,
			}

			numbers = append(numbers, newNum)
		}

		symbolIndexes := symbolRegexp.FindAllStringIndex(line, -1)
		symbolStrings := symbolRegexp.FindAllString(line, -1)

		for i := 0; i < len(symbolIndexes); i++ {

			newSymbol := Symbol{
				value: symbolStrings[i],
				y:     y,
				x:     symbolIndexes[i][0],
			}

			symbols = append(symbols, newSymbol)
		}
	}

	return numbers, symbols
}

type Number struct {
	value int
	minX  int
	maxX  int
	y     int
}

type Symbol struct {
	value string
	x     int
	y     int
}
