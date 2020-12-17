package main

import (
	"fmt"
	"strings"
)

func main() {
	// input = inputTest
	fmt.Printf("State:%v\n\n", input)
	// start with space size, i.e
	//Array values are comparable if values of the array element type are comparable. Two array values are equal if their
	//corresponding elements are equal.

	// Could try [][][]bool for [x][y][z]isActive however the map size is changing so maybe not.
	// I could have [][3]int representing actives then map[[3]int]int counting neighbours, then comparing that with the
	// thing. Yeah I think i'll do that.

	activeCubes := parseInput(input)
	printActiveCubes(activeCubes)

	for cycle := 0; cycle < 6; cycle++ {
		neighbourCount := neighbourCounter(activeCubes)
		newActiveCubes := updateActiveCubes(activeCubes, neighbourCount)
		activeCubes = newActiveCubes

		fmt.Printf("CYCLE %v --------------------------------------------------------------------------------------------\n", cycle+1)
		printActiveCubes(activeCubes)
	}

	fmt.Printf("Part1. Number of Active Cubes: %v\n", len(activeCubes))

}

func updateActiveCubes(activeCubes [][3]int, neighbourCount map[[3]int]int) [][3]int {

	newActiveCubes := [][3]int{}

	for neighbourLoc, neighbourNum := range neighbourCount {
		currentlyActive := isActiveCubeHere(activeCubes, neighbourLoc)
		if currentlyActive && (neighbourNum == 2 || neighbourNum == 3) {
			newActiveCubes = append(newActiveCubes, neighbourLoc)
		} else if !currentlyActive && neighbourNum == 3 {
			newActiveCubes = append(newActiveCubes, neighbourLoc)
		}
	}

	return newActiveCubes

}

func neighbourCounter(activeCubes [][3]int) map[[3]int]int {
	directions := directions()
	// Take each cube location, go in all directions (apart from zero)
	neighbourCount := make(map[[3]int]int)
	for _, activeCube := range activeCubes {

		for _, dir := range directions {
			neighbourLoc := moveLoc(activeCube, dir)
			neighbourCount[neighbourLoc]++
		}
	}
	return neighbourCount
}

func isActiveCubeHere(activeCubes [][3]int, loc [3]int) bool {
	for _, cubeLoc := range activeCubes {
		if cubeLoc == loc {
			return true
		}
	}
	return false
}

func moveLoc(loc [3]int, direction [3]int) [3]int {
	newLoc := [3]int{}
	newLoc[0] = loc[0] + direction[0]
	newLoc[1] = loc[1] + direction[1]
	newLoc[2] = loc[2] + direction[2]
	return newLoc
}

func parseInput(input string) [][3]int {
	activeCubes, z := [][3]int{}, 0
	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				activeCubes = append(activeCubes, [3]int{x - 1, y - 1, z})
			}
		}
	}
	return activeCubes
}

func directions() [][3]int {
	dirs := [][3]int{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				dirs = append(dirs, [3]int{x, y, z})
			}
		}
	}
	return dirs
}

func printActiveCubes(activeCubes [][3]int) {
	min, max := minMaxLoc(activeCubes)
	for z := min[2]; z <= max[2]; z++ {
		printSlice(activeCubes, min, max, z)
	}
	return
}

func printSlice(activeCubes [][3]int, min [3]int, max [3]int, z int) {
	fmt.Printf("x:%v y:%v z:%v\n", z, min[0], min[1])
	for y := min[1]; y <= max[1]; y++ {
		for x := min[0]; x <= max[0]; x++ {
			if isActiveCubeHere(activeCubes, [3]int{x, y, z}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	return
}

func minMaxLoc(activeCubes [][3]int) (min [3]int, max [3]int) {
	min = [3]int{activeCubes[0][0], activeCubes[0][1], activeCubes[0][2]}
	max = [3]int{activeCubes[0][0], activeCubes[0][1], activeCubes[0][2]}

	for _, activeCube := range activeCubes {

		for i := 0; i < 3; i++ {
			if activeCube[i] < min[i] {
				min[i] = activeCube[i]
			}
			if activeCube[i] > max[i] {
				max[i] = activeCube[i]
			}
		}

	}

	return min, max
}

var input = `..##.##.
#.#..###
##.#.#.#
#.#.##.#
###..#..
.#.#..##
#.##.###
#.#..##.`

var inputTest = `.#.
..#
###`
