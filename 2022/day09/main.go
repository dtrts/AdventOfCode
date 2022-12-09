package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	// "golang.org/x/exp/slices"
)

func p(s ...interface{}) {
	fmt.Println(s...)
}

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

	inputString := strings.TrimSpace(string(inputBytes))
	// BOILER PLATE --------------------------------------------------------------------

	instructions := strings.Split(inputString, "\n")

	visited1, visited2 := map[Coord]bool{Coord{0, 0}: true}, map[Coord]bool{Coord{0, 0}: true}

	rope := make([]Coord, 10)

	instructionSplit, direction, steps := []string{"", ""}, Coord{0, 0}, 0

	for _, instruction := range instructions {
		instructionSplit = strings.Fields(instruction)
		direction = parseDirection(instructionSplit[0])
		steps, _ = strconv.Atoi(instructionSplit[1])

		for stepCount := 1; stepCount <= steps; stepCount++ {
			moveRope(rope, direction)
			visited1[rope[1]] = true
			visited2[rope[9]] = true
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	p("Part1:", len(visited1))
	p("Part2:", len(visited2))
	// BOILER PLATE --------------------------------------------------------------------
}

type Coord struct {
	x, y int
}

func moveRope(rope []Coord, direction Coord) {
	rope[0].move(direction)
	for i := 1; i < len(rope); i++ {
		rope[i].follow(rope[i-1])
	}
}

func (positionTail *Coord) follow(positionHead Coord) {

	dx, dy := positionHead.x-positionTail.x, positionHead.y-positionTail.y

	// on same spot, adjacent, or adjacent diagonal
	if abs(dx)+abs(dy) <= 1 || (abs(dx) == 1 && abs(dy) == 1) {
		return
	}

	positionTail.move(Coord{normal(dx), normal(dy)})
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func normal(i int) int {
	if i == 0 {
		return i
	}

	return i / abs(i)
}

func (position *Coord) move(direction Coord) {
	position.x += direction.x
	position.y += direction.y
}

func parseDirection(s string) Coord {
	up := Coord{1, 0}
	down := Coord{-1, 0}
	left := Coord{0, -1}
	right := Coord{0, 1}

	switch s {
	case "U":
		return up
	case "D":
		return down
	case "L":
		return left
	case "R":
		return right
	default:
		panic("")
	}

}
