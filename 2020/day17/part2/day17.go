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
		// printActiveCubes(activeCubes)
	}

	fmt.Printf("Part2. Number of Active Cubes: %v\n", len(activeCubes))

}

func updateActiveCubes(activeCubes [][4]int, neighbourCount map[[4]int]int) [][4]int {

	newActiveCubes := [][4]int{}

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

func neighbourCounter(activeCubes [][4]int) map[[4]int]int {
	directions := directions()
	// Take each cube location, go in all directions (apart from zero)
	neighbourCount := make(map[[4]int]int)
	for _, activeCube := range activeCubes {

		for _, dir := range directions {
			neighbourLoc := moveLoc(activeCube, dir)
			neighbourCount[neighbourLoc]++
		}
	}
	return neighbourCount
}

func isActiveCubeHere(activeCubes [][4]int, loc [4]int) bool {
	for _, cubeLoc := range activeCubes {
		if cubeLoc == loc {
			return true
		}
	}
	return false
}

func moveLoc(loc [4]int, direction [4]int) [4]int {
	newLoc := [4]int{}
	newLoc[0] = loc[0] + direction[0]
	newLoc[1] = loc[1] + direction[1]
	newLoc[2] = loc[2] + direction[2]
	newLoc[3] = loc[3] + direction[3]
	return newLoc
}

func parseInput(input string) [][4]int {
	activeCubes, z, w := [][4]int{}, 0, 0
	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				activeCubes = append(activeCubes, [4]int{x - 1, y - 1, z, w})
			}
		}
	}
	return activeCubes
}

func directions() [][4]int {
	dirs := [][4]int{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					dirs = append(dirs, [4]int{x, y, z, w})
				}
			}
		}
	}
	return dirs
}

func printActiveCubes(activeCubes [][4]int) {
	min, max := minMaxLoc(activeCubes)
	for z := min[2]; z <= max[2]; z++ {
		for w := min[3]; w <= max[3]; w++ {
			printSlice(activeCubes, min, max, z, w)
		}
	}
	return
}

func printSlice(activeCubes [][4]int, min [4]int, max [4]int, z int, w int) {
	fmt.Printf("x:%v y:%v z:%v w:%v\n", min[0], min[1], z, w)
	for y := min[1]; y <= max[1]; y++ {
		for x := min[0]; x <= max[0]; x++ {
			if isActiveCubeHere(activeCubes, [4]int{x, y, z, w}) {
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

func minMaxLoc(activeCubes [][4]int) (min [4]int, max [4]int) {
	min = [4]int{activeCubes[0][0], activeCubes[0][1], activeCubes[0][2], activeCubes[0][3]}
	max = [4]int{activeCubes[0][0], activeCubes[0][1], activeCubes[0][2], activeCubes[0][3]}

	for _, activeCube := range activeCubes {

		for i := 0; i < 4; i++ {
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
