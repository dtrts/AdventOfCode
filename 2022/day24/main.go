package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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

	valley := parseInput(inputString)
	valley.populateSpaces()

	// startingPosition := Coord{0, 0}
	// startingIteration := 1
	// goalPosition := Coord{valley.width - 1, valley.height - 1}
	// cache := make(map[CacheKey]int, 0)
	// visited := make(map[CacheKey]int, 0)
	// globalMinimum := math.MaxInt
	// shortestPathLength := valley.findShortestPath(startingPosition, startingIteration, goalPosition, visited, cache, &globalMinimum)
	entrance := Coord{x: 0, y: -1}
	exit := Coord{x: valley.width - 1, y: valley.height}
	shortestPathLength := valley.findShortestPathBFS([]Coord{entrance}, 0, exit)
	part1 := shortestPathLength
	p("Shortest Path Length 1", shortestPathLength)
	shortestPathLength2 := valley.findShortestPathBFS([]Coord{exit}, part1, entrance)
	p("Shortest Path Length 2", shortestPathLength2)
	shortestPathLength3 := valley.findShortestPathBFS([]Coord{entrance}, shortestPathLength2, exit)
	p("Shortest Path Length 3", shortestPathLength3)

	// iteration := part1
	// for tripFound := false; !tripFound; {

	// 	short := valley.findShortestPathBFS([]Coord{goalPosition}, iteration, startingPosition)

	// 	if short > 0 {
	// 		p("Short found with no blockages", iteration)
	// 		tripFound = true
	// 		p("length", short-iteration+2)
	// 	} else {
	// 		p("iteration", iteration, "found blockage")
	// 	}
	// 	iteration++

	// }

	// valley.printBlizzards(part1)
	/*
	   #.######
	   #>2.<.<#
	   #.2v^2<#
	   #>..>2>#
	   #<....>#
	   ######E#
	*/
	// for i := 0; i <= valley.loopLength; i++ {
	// 	short := valley.findShortestPathBFS([]Coord{goalPosition}, i, startingPosition)
	// 	p(i, short, short-i+2)
	// }

	// shortestPathLength2 := valley.findShortestPathBFS([]Coord{goalPosition}, part1+1, startingPosition)

	// shortestPathLength3 := valley.findShortestPathBFS([]Coord{startingPosition}, shortestPathLength2+3, goalPosition)

	// p(shortestPathLength3 + 1)
	// shortestPathLength3 := valley.findShortestPathBFS([]Coord{goalPosition}, shortestPathLength2+1, startingPosition)
	// p(shortestPathLength3)
	// In the above example, the first trip to the goal takes 18 minutes, the trip back to the start takes 23 minutes, and the
	// trip back to the goal again takes 13 minutes, for a total time of 54 minutes.

	// part2 := shortestPathLength2 + 2

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1) // 368 is too high??
	// p("Part2:", part2) // 834 is too hight? 836 also too hgih

	// BOILER PLATE --------------------------------------------------------------------
}

// Need a map of coordinates ?
// the blizzards arrive in cycles, so by keeping track of which move we have made we can determin if the square has a blizzard on it just by storing the repeat.
// I.e we don' track the positions of the blizzards, we just track when they would arrive at each square.
// Since x and y can be different, the cycles are different.

// So we will have the GRID:

// the grid has coordinates.
// the grid has a width and height
// on each coordinate we have:
// height repeats and length repeats. This is a slice of numbers. st.
// (iteration - NUM)  % height == 0 => A BLIZZARD IS THERE
// the map repeats every lcm(height,width)

// (We can use this information and invert it, then we will have a GRID which is easy to track next steps.)

// Is presense of bliz better than preseense of not bliz?

// not bliz i think, can compare sizes?

type CacheKey struct {
	currentPosition Coord
	iteration       int
}

func (valley *Valley) findShortestPathBFS(currentPositions []Coord, iteration int, goalPosition Coord) int {

	// breadth first. Take each point. search all neighbours. Keep track of path?
	if len(currentPositions) == 1 {

		if !contains(valley.tiles[currentPositions[0]].noBlizzard, (iteration)%valley.loopLength) {
			p("Trying to start in a place with a blizzard. Wait one sec.", currentPositions, iteration)
			return valley.findShortestPathBFS(currentPositions, iteration+1, goalPosition)
		}
	}
	// If current position

	// Go through current positions and find possible moves
	//
	possibleMoves := []Coord{}
	for _, currentPosition := range currentPositions {
		subPossibleMoves := currentPosition.moves(valley.width, valley.height)
		for _, subPossibleMove := range subPossibleMoves {
			if !contains(possibleMoves, subPossibleMove) {
				possibleMoves = append(possibleMoves, subPossibleMove)
			}
		}
	}

	viableMoves := []Coord{}

	for _, possibleMove := range possibleMoves {
		if contains(valley.tiles[possibleMove].noBlizzard, (iteration+1)%valley.loopLength) {
			viableMoves = append(viableMoves, possibleMove)
		}
	}

	if len(viableMoves) == 0 {
		// p(currentPositions)
		// p("No Viable Moves")
		return -1 //valley.findShortestPathBFS(currentPositions, iteration+1, goalPosition)

		// return valley.findShortestPathBFS(currentPositions, iteration+1, goalPosition)
	}

	if contains(viableMoves, goalPosition) {
		return iteration + 1
	}

	return valley.findShortestPathBFS(viableMoves, iteration+1, goalPosition)
}

func (valley *Valley) findShortestPath(currentPosition Coord, iteration int, goalPosition Coord, visited, cache map[CacheKey]int, globalMinimum *int) int {
	// p("Finding Shortest Path", currentPosition, iteration, globalMinimum)
	cacheKey := CacheKey{
		currentPosition: currentPosition,
		iteration:       iteration % valley.loopLength,
	}

	if _, ok := visited[cacheKey]; !ok {
		visited[cacheKey] = 0
	}
	visited[cacheKey]++

	// if at goalPosition return iteration
	if currentPosition == goalPosition {
		// p("Made it to goal position")
		globalMinimum = &iteration
		cache[cacheKey] = iteration
		return iteration
	}

	// If already visited

	// If global minimum is less than current iteration + shortest path to goal path return max int
	crowFliesLength := abs(goalPosition.x-currentPosition.x) + abs(goalPosition.y-currentPosition.y)
	if iteration+crowFliesLength > *globalMinimum {
		return math.MaxInt
	}

	// Check cache for minimum path
	// Cache key: Current Position, Iteration Mod Loop Length.

	// if shortestPathLength, ok := cache[cacheKey]; ok {
	// 	// p("Cache Hit", cacheKey, currentPosition, iteration, valley.loopLength, valley.width, valley.height)
	// 	return shortestPathLength
	// }

	// If cache miss, and already visited, then return and set to max.
	if visited[cacheKey] > 1 {
		// p("Already visited within loop length", cacheKey, currentPosition, iteration, valley.loopLength, valley.width, valley.height)
		cache[cacheKey] = math.MaxInt
		return math.MaxInt
	}

	// Find neighboring positions.

	possibleMoves := currentPosition.moves(valley.width, valley.height)
	// Sort possible moves so we search closer to end position?
	// sort.Slice(possibleMoves, func(i, j int) bool {
	// 	return distance(possibleMoves[i], goalPosition) < distance(possibleMoves[j], goalPosition)
	// })

	// p("Checking positions:", possibleMoves)
	shortestPathLength := math.MaxInt

	for _, possibleMove := range possibleMoves {
		// Check to see if the slice contains

		if contains(valley.tiles[possibleMove].noBlizzard, (iteration+1)%valley.loopLength) {
			newShortestPathLength := valley.findShortestPath(possibleMove, iteration+1, goalPosition, visited, cache, globalMinimum)
			shortestPathLength = min(shortestPathLength, newShortestPathLength)
		}
	}

	cache[cacheKey] = shortestPathLength
	return shortestPathLength
}

func (pos Coord) moves(width, height int) []Coord {

	entrance := Coord{x: 0, y: -1}
	exit := Coord{x: width - 1, y: height}

	possibleMoves := []Coord{}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 || dy == 0 {
				newCoord := Coord{x: pos.x + dx, y: pos.y + dy}

				if (betweenIncl(newCoord.x, 0, width-1) && betweenIncl(newCoord.y, 0, height-1)) || contains([]Coord{entrance, exit}, newCoord) {
					possibleMoves = append(possibleMoves, newCoord)
				}
			}
		}
	}
	return possibleMoves
}

type Valley struct {
	width      int
	height     int
	loopLength int

	tiles map[Coord]*Tile
}

type Tile struct {
	movingLeft  []int
	movingRight []int
	movingUp    []int
	movingDown  []int
	noBlizzard  []int
}

type Coord struct {
	x, y int
}

func parseInput(input string) *Valley {

	// Assuming there is a wall all around. But don't car about that TBH
	// Initialise the valley
	valley := &Valley{tiles: make(map[Coord]*Tile, 0)}

	lines := strings.Split(input, "\n")

	valley.height = len(lines) - 2

	for y, line := range lines[1 : valley.height+1] {
		valley.width = len(line) - 2

		for x := range line[1 : valley.width+1] {
			valley.tiles[Coord{x: x, y: y}] = &Tile{
				movingLeft:  make([]int, 0),
				movingRight: make([]int, 0),
				movingUp:    make([]int, 0),
				movingDown:  make([]int, 0),
				noBlizzard:  make([]int, 0),
			}
		}
	}

	for y, line := range lines[1 : valley.height+1] {
		for x, charB := range line[1 : valley.width+1] {
			char := string(charB)
			valley.addBlizzard(char, Coord{x: x, y: y})
		}
	}

	p("")
	p("")

	valley.loopLength = lcm(valley.width, valley.height)

	return valley

}

func (valley *Valley) addBlizzard(direction string, pos Coord) {
	switch direction {
	case ">":
		for dx := 0; dx < valley.width; dx++ {
			c := Coord{x: dx, y: pos.y}
			valley.tiles[c].movingRight = append(valley.tiles[c].movingRight, dx-pos.x)
		}
	case "<":
		for dx := 0; dx < valley.width; dx++ {
			c := Coord{x: dx, y: pos.y}
			valley.tiles[c].movingLeft = append(valley.tiles[c].movingLeft, pos.x-dx)
		}
	case "v":
		for dy := 0; dy < valley.height; dy++ {
			c := Coord{x: pos.x, y: dy}
			valley.tiles[c].movingDown = append(valley.tiles[c].movingDown, dy-pos.y)
		}
	case "^":
		for dy := 0; dy < valley.height; dy++ {
			c := Coord{x: pos.x, y: dy}
			valley.tiles[c].movingUp = append(valley.tiles[c].movingUp, pos.y-dy)
		}
	}
}

func (valley *Valley) populateSpaces() {
	entrance := Coord{x: 0, y: -1}
	exit := Coord{x: valley.width - 1, y: valley.height}
	valley.tiles[entrance] = &Tile{
		noBlizzard: []int{},
	}
	valley.tiles[exit] = &Tile{
		noBlizzard: []int{},
	}
	// iterations:
	for iteration := 0; iteration < valley.loopLength; iteration++ {

		valley.tiles[entrance].noBlizzard = append(valley.tiles[entrance].noBlizzard, iteration)
		valley.tiles[exit].noBlizzard = append(valley.tiles[entrance].noBlizzard, iteration)

		for y := 0; y < valley.height; y++ {
		tiles:
			for x := 0; x < valley.width; x++ {
				c := Coord{x: x, y: y}
				for _, blizzard := range valley.tiles[c].movingRight {
					if (iteration-blizzard)%valley.width == 0 {
						continue tiles
					}
				}
				for _, blizzard := range valley.tiles[c].movingLeft {
					if (iteration-blizzard)%valley.width == 0 {
						continue tiles
					}
				}
				for _, blizzard := range valley.tiles[c].movingUp {
					if (iteration-blizzard)%valley.height == 0 {
						continue tiles
					}
				}
				for _, blizzard := range valley.tiles[c].movingDown {
					if (iteration-blizzard)%valley.height == 0 {
						continue tiles
					}
				}
				valley.tiles[c].noBlizzard = append(valley.tiles[c].noBlizzard, iteration)
			}
		}
	}
}

// When parsing the input we want to make a blizzardGrid

func (valley *Valley) printSpaces(iteration int) {
	for y := 0; y < valley.height; y++ {
		for x := 0; x < valley.width; x++ {
			c := Coord{x: x, y: y}

			isSpace := false

			for _, space := range valley.tiles[c].noBlizzard {
				if (iteration-space)%valley.loopLength == 0 {
					isSpace = true
					break
				}
			}

			if isSpace {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
		p("")
	}
	p("")
}

func (valley *Valley) printBlizzards(iteration int) {
	for y := 0; y < valley.height; y++ {
		for x := 0; x < valley.width; x++ {
			c := Coord{x: x, y: y}

			blizzards := []string{}

			for _, blizzard := range valley.tiles[c].movingLeft {
				if (iteration-blizzard)%valley.width == 0 {
					blizzards = append(blizzards, "<")
				}
			}

			for _, blizzard := range valley.tiles[c].movingRight {
				if (iteration-blizzard)%valley.width == 0 {
					blizzards = append(blizzards, ">")
				}
			}

			for _, blizzard := range valley.tiles[c].movingUp {
				if (iteration-blizzard)%valley.height == 0 {
					blizzards = append(blizzards, "^")
				}
			}
			for _, blizzard := range valley.tiles[c].movingDown {
				if (iteration-blizzard)%valley.height == 0 {
					blizzards = append(blizzards, "v")
				}
			}

			if len(blizzards) == 0 {
				fmt.Print(".")
			} else if len(blizzards) == 1 {
				fmt.Print(blizzards[0])
			} else {
				fmt.Print(len(blizzards))
			}
		}
		p("")
	}
	p("")
}

// Maths stolen from stack overflow
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func betweenIncl(a, min, max int) bool {
	if min <= a && a <= max {
		return true
	}

	return false
}

func contains[T comparable](s []T, v T) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func distance(a, b Coord) int {

	return abs(b.x-a.x) + abs(b.y-a.y)

}
