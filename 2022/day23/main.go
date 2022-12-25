package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func p(s ...interface{}) {
	fmt.Println(s...)
}

func main() {
	// BOILER PLATE --------------------------------------------------------------------
	start := time.Now()
	log.Printf("Starting... %s", start.Format("Jan 2 15:04:05 2006 MST"))

	var inputFileName string
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))
	// BOILER PLATE --------------------------------------------------------------------

	elves := parseInput(inputString)

	// p("=== Initial State ===")
	// printElves(elves, 2)
	// p("")

	part1, part2 := 0, 0
	iteration := 0
	elvesMoved := len(elves)
	for elvesMoved > 0 && iteration < 1e5 {

		updateSurroundingElves(elves)

		proposedPositions := updateProposedPositions(elves, iteration)

		elvesMoved = updatePositions(elves, proposedPositions)
		if elvesMoved == 0 {
			part2 = iteration + 1
		}

		// fmt.Printf("=== End of Round %d ===\n", iteration+1)
		// p("Elves Moved", elvesMoved)
		// p("len elves", len(elves))
		// // printElves(elves, 0)
		// p("")

		if iteration+1 == 10 {
			minRect, maxRect := findMinMax(elves)
			part1 = ((maxRect.x - minRect.x + 1) * (maxRect.y - minRect.y + 1)) - len(elves)
		}

		iteration++
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1) // 4138 // 110 test
	p("Part2:", part2) // 519 is tooooo low // 20 test
	// BOILER PLATE --------------------------------------------------------------------
}

type Coord struct {
	x, y int
}

type Elf struct {
	position         Coord
	proposedPosition Coord
	surroundingElves map[string]int
	canMove          bool
}

func updatePositions(elves map[Coord]*Elf, proposedPositions map[Coord]int) int {
	elvesMoved := 0

	for pos, elf := range elves {
		if elf.surroundingElves["A"] >= 1 && elf.canMove {
			if proposedPositions[elf.proposedPosition] == 1 {
				elvesMoved++
				delete(elves, pos)
				elf.position = elf.proposedPosition
				elves[elf.position] = elf
			}
		}
	}
	return elvesMoved
}

func updateProposedPositions(elves map[Coord]*Elf, iteration int) map[Coord]int {

	searchOrder := [4]string{"N", "S", "W", "E"}

	proposedPositions := make(map[Coord]int, 0)

	for _, elf := range elves {
		elf.proposedPosition = Coord{0, 0}
		elf.canMove = false
		if elf.surroundingElves["A"] >= 1 {
			for i := 0; i < 4; i++ {

				direction := searchOrder[(iteration+i)%4]
				if elf.surroundingElves[direction] == 0 {

					elf.proposedPosition = directionProposedPosition(elf.position, direction)

					if _, ok := proposedPositions[elf.proposedPosition]; !ok {
						proposedPositions[elf.proposedPosition] = 0
					}

					proposedPositions[elf.proposedPosition]++
					elf.canMove = true
					break
				}
			}
		}
	}
	return proposedPositions
}

func updateSurroundingElves(elves map[Coord]*Elf) {
	directions := []string{"A", "N", "S", "E", "W"}
	for position, elf := range elves {
		for _, direction := range directions {
			elf.surroundingElves[direction] = 0

			neighbours := neighbours(position, direction)

			for _, neighbour := range neighbours {
				if _, ok := elves[neighbour]; ok {
					elf.surroundingElves[direction]++
				}
			}
		}
	}
}

func parseInput(input string) map[Coord]*Elf {
	lines := strings.Split(input, "\n")

	elves := make(map[Coord]*Elf)

	for y, line := range lines {
		for x, charB := range line {
			char := string(charB)

			if char == "#" {
				position := Coord{x: x, y: y}
				elves[position] = &Elf{
					position:         position,
					surroundingElves: make(map[string]int),
				}
			}
		}
	}
	return elves
}

func directionProposedPosition(pos Coord, direction string) Coord {

	switch direction {
	case "N":
		return Coord{pos.x, pos.y - 1}
	case "S":
		return Coord{pos.x, pos.y + 1}
	case "E":
		return Coord{pos.x + 1, pos.y}
	case "W":
		return Coord{pos.x - 1, pos.y}

	default:
		panic("Please provide a direction")
	}

}

func neighbours(pos Coord, direction string) []Coord {

	dxRange, dyRange := []int{-1, 0, 1}, []int{-1, 0, 1}
	switch direction {
	case "N":
		dyRange = []int{-1}
	case "S":
		dyRange = []int{1}
	case "E":
		dxRange = []int{1}
	case "W":
		dxRange = []int{-1}
	case "A":
	default:
		panic("Please provide a direction")
	}

	neighbours := make([]Coord, 0, 8)

	for _, dx := range dxRange {
		for _, dy := range dyRange {
			if dx == 0 && dy == 0 {
				continue
			}
			neighbours = append(neighbours, Coord{x: pos.x + dx, y: pos.y + dy})
		}
	}
	return neighbours
}

func printElves(elves map[Coord]*Elf, padding int) {
	minPos, maxPos := findMinMax(elves)

	for y := minPos.y - padding; y <= maxPos.y+padding; y++ {
		for x := minPos.x - padding; x <= maxPos.x+padding; x++ {

			_, ok := elves[Coord{x: x, y: y}]

			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		p("")
	}
}

func listElves(elves map[Coord]*Elf) {
	for _, elf := range elves {
		p(elf)
	}
}

func findMinMax(elves map[Coord]*Elf) (Coord, Coord) {
	var avg Coord

	for pos := range elves {
		avg.x += pos.x
		avg.y += pos.y
	}
	avg.x /= len(elves)
	avg.y /= len(elves)

	minPos, maxPos := avg, avg

	for pos := range elves {
		minPos.x = min(pos.x, minPos.x)
		minPos.y = min(pos.y, minPos.y)
		maxPos.x = max(pos.x, maxPos.x)
		maxPos.y = max(pos.y, maxPos.y)
	}

	return minPos, maxPos
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
N dy -1 dx -1,0,1
S dy 1 dx -1,0,1
E dx 1 dy -1,0,1
W dx -1 dy -1,0,1
*/
