package main

import (
	"fmt"
	"strings"
)

func main() {
	// input = inputTest

	var inputLines []string

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	treesHit(inputLines, 3, 1)

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	treesHitLoop := 1
	fmt.Println(slopes)

	for _, slope := range slopes {
		treesHitLoop *= treesHit(inputLines, slope[0], slope[1])
	}

	fmt.Println(treesHitLoop)

}

func treesHit(inputLines []string, slopeX int, slopeY int) int {

	// slopeX := 3
	// slopeY := 1
	position := [2]int{0, 0}
	mapWidth := len(inputLines[0])
	locations := []string{}
	treesHit := 0

	fmt.Println(len(inputLines), slopeX, slopeY, position, mapWidth, treesHit)
	// fmt.Println(len(inputLines), slopeX, slopeY, position, mapWidth, treesHit, inputLines)

	for position[1] < len(inputLines) {
		// fmt.Println("Start loop:", position, treesHit, locations)

		thing := string(inputLines[position[1]][position[0]%mapWidth])

		locations = append(locations, thing)
		if thing == "#" {
			treesHit++
		}

		position[0] += slopeX
		position[1] += slopeY

		fmt.Print(thing, " ")
		// fmt.Println("End Loop Position", position, treesHit, locations)
	}

	fmt.Print("\n")
	fmt.Println("Results", position, treesHit, locations)
	return treesHit
}
