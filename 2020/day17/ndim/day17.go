package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input = inputTest
	fmt.Printf("State:\n%v\n\n", input)
	// start with space size, i.e
	//Array values are comparable if values of the array element type are comparable. Two array values are equal if their
	//corresponding elements are equal.

	// Could try [][][]bool for [x][y][z]isActive however the map size is changing so maybe not.
	// I could have [][3]int representing actives then map[[3]int]int counting neighbours, then comparing that with the
	// thing. Yeah I think i'll do that.

	dimensions := 4

	// fmt.Println(len(directions(4)))

	activeCubes := parseInput(input, dimensions)
	printActiveCubes(activeCubes)

	numCycles := 7

	for cycle := 0; cycle < numCycles; cycle++ {
		neighbourCount := neighbourCounter(activeCubes, dimensions)
		newActiveCubes := updateActiveCubes(activeCubes, neighbourCount)
		activeCubes = newActiveCubes

		fmt.Printf("CYCLE %v --------------------------------------------------------------------------------------------\n", cycle+1)
		printActiveCubes(activeCubes)
		fmt.Printf(" Number of Active Cubes: %v\n", len(activeCubes))

	}

	fmt.Printf("Part2. Number of Active Cubes: %v\n", len(activeCubes))

}

func updateActiveCubes(activeCubes [][]int, neighbourCount map[string]int) [][]int {

	newActiveCubes := [][]int{}

	for neighbourLoc, neighbourNum := range neighbourCount {
		currentlyActive := isActiveCubeHere(activeCubes, locToSlice(neighbourLoc))
		if currentlyActive && (neighbourNum == 2 || neighbourNum == 3) {
			newActiveCubes = append(newActiveCubes, locToSlice(neighbourLoc))
		} else if !currentlyActive && neighbourNum == 3 {
			newActiveCubes = append(newActiveCubes, locToSlice(neighbourLoc))
		}
	}

	return newActiveCubes

}

func neighbourCounter(activeCubes [][]int, dimensions int) map[string]int {
	directions := directions(dimensions)
	// Take each cube location, go in all directions (apart from zero)
	neighbourCount := make(map[string]int)
	for _, activeCube := range activeCubes {

		for _, dir := range directions {
			neighbourLoc := moveLoc(activeCube, dir)
			// fmt.Println(activeCube, dir, neighbourLoc, locSliceToString(neighbourLoc), neighbourCount[locSliceToString(neighbourLoc)])
			neighbourCount[locSliceToString(neighbourLoc)]++
		}
	}
	return neighbourCount
}

func locToSlice(locs string) (loc []int) {
	for _, point := range strings.Split(locs, ",") {
		val, _ := strconv.Atoi(point)
		loc = append(loc, val)
	}
	return loc
}

func locSliceToString(loc []int) (locs string) {
	for i, point := range loc {
		if i == 0 {
			locs += strconv.Itoa(point)
		} else {
			locs += "," + strconv.Itoa(point)
		}
	}
	return locs
}

func isActiveCubeHere(activeCubes [][]int, loc []int) bool {
	for _, cubeLoc := range activeCubes {
		if compareLoc(cubeLoc, loc) {
			return true
		}
	}
	return false
}

func moveLoc(loc []int, direction []int) []int {
	newLoc := make([]int, len(loc))
	for i := range loc {
		newLoc[i] = loc[i] + direction[i]
	}
	return newLoc
}

func parseInput(input string, dimensions int) [][]int {
	activeCubes := [][]int{}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				activeCubeLoc := make([]int, dimensions)
				activeCubeLoc[0] = x - 1
				activeCubeLoc[1] = y - 1
				activeCubes = append(activeCubes, activeCubeLoc)
			}
		}
	}
	return activeCubes
}

func compareLoc(loc0 []int, loc1 []int) bool {
	if len(loc0) != len(loc1) {
		return false
	}
	for i := range loc0 {
		if loc0[i] != loc1[i] {
			return false
		}
	}
	return true
}

func directions(dimensions int) [][]int {
	min := make([]int, dimensions)
	max := make([]int, dimensions)
	for i := range min {
		min[i] = -1
		max[i] = 1
	}

	combos := highDimCombos(min, max)

	//remove origin
	zeros := make([]int, dimensions)
	comboZeroIndex := 0
	for i, v := range combos {
		if compareLoc(zeros, v) {
			comboZeroIndex = i
			break
		}
	}
	combos[comboZeroIndex] = combos[len(combos)-1]
	return combos[:len(combos)-1]
}

func printActiveCubes(activeCubes [][]int) {
	min, max := minMaxLoc(activeCubes)
	highDimCombos := highDimCombos(min[2:], max[2:])

	for _, highDimCombo := range highDimCombos {
		printSlice(activeCubes, min, max, highDimCombo)
	}
	return
}

func highDimCombos(min []int, max []int) [][]int {
	start, grower := [][]int{}, [][]int{}
	for i := min[0]; i <= max[0]; i++ {
		start = append(start, []int{i})
	}
	for dimNum := 1; dimNum < len(min); dimNum++ {
		for i := min[dimNum]; i <= max[dimNum]; i++ {
			for _, s := range start {
				grower = append(grower, append(s, i))
			}
		}
		start = grower
		grower = [][]int{}
	}
	return start
}

func printSlice(activeCubes [][]int, min []int, max []int, higherDims []int) {
	// Display coordinate of top left cube displayed
	fmt.Printf("x:%v y:%v", min[0], min[1])
	for i, v := range higherDims {
		fmt.Printf(" %v:%v", string(rune(((i+25)%26)+97)), v)
	}
	fmt.Printf("\n")

	for y := min[1]; y <= max[1]; y++ {
		for x := min[0]; x <= max[0]; x++ {
			loc := append([]int{x, y}, higherDims...)
			if isActiveCubeHere(activeCubes, loc) {
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

func minMaxLoc(activeCubes [][]int) (min []int, max []int) {
	min, max = make([]int, len(activeCubes[0])), make([]int, len(activeCubes[0]))
	copy(min, activeCubes[0])
	copy(max, activeCubes[0])

	for _, activeCube := range activeCubes {

		for i := 0; i < len(min); i++ {
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
