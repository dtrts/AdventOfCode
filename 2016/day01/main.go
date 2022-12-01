package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
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
	// BOILER PLATE --------------------------------------------------------------------

	instructions := strings.Split(strings.TrimSpace(inputString), ", ")

	facing, position := [2]int{1, 0}, [2]int{0, 0}

	rotations := make([]rune, 0, len(instructions))
	distances := make([]int, 0, len(instructions))

	locationsVisited := 0

	for _, instruction := range instructions {
		rotation := rune(instruction[0])
		rotations = append(rotations, rotation)

		distance, _ := strconv.Atoi(instruction[1:])
		distances = append(distances, distance)

		locationsVisited += distance
	}

	locations := make([][2]int, 0, locationsVisited)

	for instructionIndex := 0; instructionIndex < len(instructions); instructionIndex++ {
		facing = turn(facing, rotations[instructionIndex])

		for i2 := 0; i2 < distances[instructionIndex]; i2++ {
			position[0] += facing[0] * 1
			position[1] += facing[1] * 1

			locations = append(locations, position)
		}
	}

	visitTwice := [2]int{0, 0}

	for i, location := range locations {
		index2 := slices.Index(locations[i+1:], location)

		if index2 > 0 {
			visitTwice = locations[i+index2+1]
			break
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", intAbs(position[0])+intAbs(position[1]))
	fmt.Println("Part2:", intAbs(visitTwice[0])+intAbs(visitTwice[1]))
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------
}

var N [2]int = [2]int{1, 0}
var S [2]int = [2]int{-1, 0}
var E [2]int = [2]int{0, 1}
var W [2]int = [2]int{0, -1}

func intAbs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func turn(facing [2]int, rotation rune) [2]int {

	strRotation := string(rotation)

	if strRotation != "L" && strRotation != "R" {
		panic("Only L or R")
	}

	lefts := [][2]int{N, W, S, E}
	rights := [][2]int{N, E, S, W}

	if strRotation == "L" {
		i := slices.Index(lefts, facing)
		return lefts[(i+1)%4]
	}

	i := slices.Index(rights, facing)

	return rights[(i+1)%4]

}
