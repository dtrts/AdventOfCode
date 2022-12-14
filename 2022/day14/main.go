package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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

	part1 := part1(inputString)
	part2 := part2(inputString)
	// sandStartPointInput := "500,0"

	// Get lowest point  from all rocks and sand? Once a sand particle has passed

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
}

func part2(input string) int {

	p("")
	p("Starting part 2")

	settledLocations := parseRocks(input)

	floorLayer := 0
	for location := range settledLocations {
		if location.y > floorLayer {
			floorLayer = location.y
		}
	}
	floorLayer += 2

	infinityBreak := 0

	source := Coord{
		x: 500,
		y: 0,
	}

	sourceReached := false
	for !sourceReached {
		sand := Coord{
			x: 500,
			y: 0,
		}

		sandSettled := false

		for !sandSettled {
			sandNew := descend(sand, settledLocations, floorLayer)

			if sandNew == sand {
				settledLocations[sandNew] = 2
				p("Sand has settled", sandNew)
				sandSettled = true
				if sandNew == source {
					p("Sand has settled on source")
					sourceReached = true
				}
			}

			sand = sandNew
		}

		infinityBreak++
		if infinityBreak >= 1e6 {
			p("Get outta here")
			break
		}
	}

	sandCount := 0
	for _, typeOfSettled := range settledLocations {
		if typeOfSettled == 2 {
			sandCount++
		}
	}

	return sandCount

}
func part1(input string) int {

	settledLocations := parseRocks(input)

	p(settledLocations)

	abyssLayer := 0
	for location := range settledLocations {
		if location.y > abyssLayer {
			abyssLayer = location.y
		}
	}

	abyssFound := false
	infinityBreak := 0

	for !abyssFound {
		sand := Coord{
			x: 500,
			y: 0,
		}

		for sand.y <= abyssLayer {
			sandNew := descend(sand, settledLocations, math.MaxInt)

			if sandNew == sand {
				settledLocations[sandNew] = 2
				p("Sand has settled", sandNew)
				break
			}
			sand = sandNew
		}

		if sand.y > abyssLayer {
			p("Abyss found")
			break
		}

		infinityBreak++
		if infinityBreak >= 1e6 {
			p("Get outta here")
			break
		}
	}

	sandCount := 0
	for _, typeOfSettled := range settledLocations {
		if typeOfSettled == 2 {
			sandCount++
		}
	}

	return sandCount

}

func descend(sand Coord, settled map[Coord]int, floor int) Coord {

	direction1 := Coord{
		x: 0,
		y: 1,
	}
	direction2 := Coord{
		x: -1,
		y: 1,
	}
	direction3 := Coord{
		x: 1,
		y: 1,
	}

	sand1 := move(sand, direction1)
	sand2 := move(sand, direction2)
	sand3 := move(sand, direction3)

	if _, ok := settled[sand1]; !ok && sand1.y < floor {
		return sand1
	}
	if _, ok := settled[sand2]; !ok && sand1.y < floor {
		return sand2
	}
	if _, ok := settled[sand3]; !ok && sand1.y < floor {
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
