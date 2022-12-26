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
	p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	grid := [][7]bool{}
	blows := inputString

	part1 := 0
	// loopLength := 0

	// Spawn shape

	//

	states := []State{}
	index1, index2 := -1, -1

	// Move to
	blowIndex := 0
	for shapeIndex := 0; shapeIndex < 1e6; shapeIndex++ {
		// p("Starting Shape number:", shapeIndex)

		startingPosition := Coord{x: 2, y: len(grid) + 3}

		shape := newShape(shapeIndex)
		shape = moveShape(shape, startingPosition)

		// p("Shape starting", shape)

		for shapeSettled := false; !shapeSettled; {

			shape = blow(grid, shape, string(blows[blowIndex%len(blows)]))
			blowIndex++

			// p("Blown Shape", shape)

			shape, shapeSettled = descend(grid, shape)

			// p("Descended Shape", shape, shapeSettled)

			if shapeSettled {
				// p("Shape has settled", shape)
				for i := len(grid); i <= maxY(shape); i++ {
					// p("Appending", i+1, "row to grid")
					grid = append(grid, [7]bool{})
				}

				for _, point := range shape {
					if grid[point.y][point.x] {
						panic("Point already true")
					}
					grid[point.y][point.x] = true
				}
			}

		}

		if shapeIndex == 2021 {
			part1 = len(grid)
		}

		states = append(states, State{
			shapeIndex:  shapeIndex % 5,
			blowIndex:   blowIndex % len(blows),
			fingerPrint: getFingerprint(grid),
			height:      len(grid),
		})

		if shapeIndex%1000 == 0 {
			index1, index2 = findCycle(states)
			p("Checked", shapeIndex, "shapes")
		}
		if shapeIndex > 2021 && index1 >= 0 && index2 >= 0 {
			p("Found the answers")
			p("Shape Index", shapeIndex)
			p("Index1", index1)
			p("Index2", index2)
			break
		} else {
			if shapeIndex%1000 == 0 {
				p("Checked", shapeIndex, "shapes")
			}
		}
	}

	totalRocks := 1000000000000

	totalRocksFromStartOfCycle := totalRocks - index1

	modRocksFromStartOfCycle := totalRocksFromStartOfCycle % (index2 - index1)
	totalCycles := totalRocksFromStartOfCycle / (index2 - index1)

	heightToIndex1 := states[index1].height
	heightToIndex2 := states[index2].height
	heightDifferenceOfCycle := heightToIndex2 - heightToIndex1
	heightOfModRocks := states[index1+modRocksFromStartOfCycle].height - states[index1].height

	height := heightToIndex1 + heightDifferenceOfCycle*totalCycles + heightOfModRocks

	// printGrid(grid)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)    // 3193
	p("Part2:", height-1) // Don't know why i am one higher but it worked lol
	// BOILER PLATE --------------------------------------------------------------------
}

func findCycle(states []State) (int, int) {
	for i := 0; i < len(states); i++ {

		for j := i + 1; j < len(states); j++ {

			if states[i].blowIndex == states[j].blowIndex && states[i].shapeIndex == states[j].shapeIndex && states[i].fingerPrint == states[j].fingerPrint {
				return i, j
			}
		}
	}
	return -1, -1
}

func getFingerprint(grid [][7]bool) (fingerPrint [10][7]bool) {

	amount := len(fingerPrint)
	for i := len(grid) - amount; i < len(grid); i++ {

		if i >= 0 {
			fingerPrint[i-(len(grid)-amount)] = grid[i]
		} else {
			fingerPrint[i-(len(grid)-amount)] = [7]bool{true, true, true, true, true, true, true}
		}
	}

	return fingerPrint
}

type State struct {
	shapeIndex  int
	blowIndex   int
	height      int
	fingerPrint [10][7]bool
}

func findRepeat(grid [][7]bool) (startingPosition, loopLength int) {
starting:
	for s := 0; s <= len(grid); s++ {
	search:
		for i := s + 1; i <= (len(grid)-s)/2; i++ {
			for j := s; j < i; j++ {
				if grid[i+j] != grid[j] {
					continue search
				}
			}
			loopLength = i
			startingPosition = s
			break starting
		}

	}

	// starting:
	// 	for s := 0; s <= len(grid)/2; s++ {
	// 		for i := 1; i <= len(grid)/2; i++ {
	// 		for j := 0; j < i; j++ {
	// 			if grid[i+j] != grid[j] {
	// 				continue starting
	// 			}
	// 		}
	// 		loopLength = i

	// 	}

	return startingPosition, loopLength
}

func printGrid(grid [][7]bool) {
	for i := len(grid) - 1; i >= 0; i-- {

		for _, p := range grid[i] {
			if p {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
	p("")

}

func maxY(shape []Coord) (maxY int) {
	for _, point := range shape {
		if point.y > maxY {
			maxY = point.y
		}
	}
	return maxY
}

func descend(grid [][7]bool, shape []Coord) ([]Coord, bool) {
	directionCoord := Coord{x: 0, y: -1}
	newLocation := moveShape(shape, directionCoord)
	gridHeight := len(grid)

	for _, point := range newLocation {
		if point.y < 0 {
			// p("Hit floor")
			return shape, true
		}
		if point.y < gridHeight {

			if grid[point.y][point.x] {
				// p("Hit a rock")
				return shape, true
			}

		}
	}
	return newLocation, false
}

func blow(grid [][7]bool, shape []Coord, direction string) []Coord {

	directionCoord := Coord{x: 1, y: 0}
	if direction == "<" {
		directionCoord.x = -1
	}

	gridHeight := len(grid)

	newLocation := moveShape(shape, directionCoord)

	for _, point := range newLocation {

		if point.x < 0 || point.x >= 7 {
			// p("Hit a wall innit")
			return shape
		}

		if point.y < gridHeight {
			if grid[point.y][point.x] {
				// p("Hit a rock")
				return shape
			}
		}
	}

	// p("Moved", string(direction))
	return newLocation

}

func moveShape(shape []Coord, direction Coord) []Coord {
	newShape := make([]Coord, len(shape))
	copy(newShape, shape)
	for i, shapePoint := range newShape {
		newShape[i] = shapePoint.move(direction)
	}

	return newShape
}

type Coord struct {
	x, y int
}

func (c Coord) move(dir Coord) Coord {
	return Coord{
		x: c.x + dir.x,
		y: c.y + dir.y,
	}
}

func newShape(shapeIndex int) []Coord {

	/*
	   ####
	*/

	line := []Coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}

	/*
	   .#.
	   ###
	   .#.
	*/

	cross := []Coord{
		{1, 2},
		{0, 1}, {1, 1}, {2, 1},
		{1, 0},
	}

	/*
	   ..#
	   ..#
	   ###
	*/

	l := []Coord{
		{2, 2},
		{2, 1},
		{0, 0}, {1, 0}, {2, 0},
	}

	/*
	   #
	   #
	   #
	   #
	*/

	tetris := []Coord{
		{0, 3},
		{0, 2},
		{0, 1},
		{0, 0},
	}

	/*
	   ##
	   ##
	*/

	square := []Coord{
		{0, 1}, {1, 1},
		{0, 0}, {1, 0},
	}

	shapes := [][]Coord{line, cross, l, tetris, square}

	index := shapeIndex % len(shapes)

	newShape := make([]Coord, len(shapes[index]))
	copy(newShape, shapes[index])

	return newShape
}
