package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	settledLocations := parseRocks(inputString)

	lowestLevel, floorLayer := 0, 0
	for location := range settledLocations {
		if location.y > lowestLevel {
			lowestLevel = location.y
		}
	}
	floorLayer = lowestLevel + 2

	source := Coord{
		x: 500,
		y: 0,
	}

	part1 := -1
	part2 := -1
	part1Found := false

	for sourceReached := false; !sourceReached; {
		sand := Coord{x: 500, y: 0}

		for sandSettled := false; !sandSettled; {
			sandNew := descend(sand, settledLocations, floorLayer)

			if sandNew == sand {
				settledLocations[sandNew] = 2
				sandSettled = true
				if sandNew == source {
					sourceReached = true
				}
			}

			if sandNew.y > lowestLevel && !part1Found {
				for _, typeOfSettled := range settledLocations {
					if typeOfSettled == 2 {
						part1++
					}
				}
				part1Found = true
			}

			sand = sandNew
		}
	}

	part2 = 0
	for _, typeOfSettled := range settledLocations {
		if typeOfSettled == 2 {
			part2++
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
}

func descend(sand Coord, settled map[Coord]int, floor int) Coord {

	sand1 := move(sand, Coord{x: 0, y: 1})
	if _, ok := settled[sand1]; !ok && sand1.y < floor {
		return sand1
	}

	sand2 := move(sand, Coord{x: -1, y: 1})
	if _, ok := settled[sand2]; !ok && sand2.y < floor {
		return sand2
	}

	sand3 := move(sand, Coord{x: 1, y: 1})
	if _, ok := settled[sand3]; !ok && sand3.y < floor {
		return sand3
	}

	return sand

}

func parseRocks(input string) map[Coord]int {
	inputLines := strings.Split(input, "\n")

	pairs := [][2]string{}

	for _, inputLine := range inputLines {
		coordsInput := strings.Split(inputLine, " -> ")
		for i := 1; i < len(coordsInput); i++ {
			pairs = append(pairs, [2]string{coordsInput[i-1], coordsInput[i]})
		}
	}

	rocks := map[Coord]int{}

	for _, pair := range pairs {
		pairPath := pathBetweenCoord(covertCoord(pair[0]), covertCoord(pair[1]))

		for _, step := range pairPath {
			rocks[step] = 1
		}
	}

	return rocks
}

func pathBetweenCoord(start, end Coord) []Coord {

	path := []Coord{start}

	if start.x != end.x && start.y != end.y {
		panic("Not a line of rocks")
	}

	direction := Coord{
		x: normal(end.x - start.x),
		y: normal(end.y - start.y),
	}

	for start != end {
		start = move(start, direction)
		path = append(path, start)
	}

	return path
}

func normal(i int) int {
	if i == 0 {
		return 0
	}

	return i / abs(i)

}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func move(c Coord, d Coord) Coord {
	return Coord{
		x: c.x + d.x,
		y: c.y + d.y,
	}
}

func covertCoord(s string) Coord {

	coordsInput := strings.Split(s, ",")

	x, err := strconv.Atoi(coordsInput[0])
	if err != nil {
		panic("Unable to convert X coord")
	}
	y, err := strconv.Atoi(coordsInput[1])
	if err != nil {
		panic("Unable to convert Y coord")
	}

	return Coord{
		x,
		y,
	}
}

type Coord struct {
	x int
	y int
}
