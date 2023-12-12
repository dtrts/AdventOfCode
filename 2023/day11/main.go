package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
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

	inputString := strings.TrimSpace(string(inputBytes))
	// fmt.Println("Input:", inputString)
	// BOILER PLATE --------------------------------------------------------------------
	part1, part2 := 0, 0

	galaxies := parseInput(inputString)
	expandAllGalaxy(galaxies, 2-1) // each line replaced with 2
	solvePaths(galaxies)

	galaxies2 := parseInput(inputString)
	expandAllGalaxy(galaxies2, 1e6-1) // each line replaced with 1e6
	solvePaths(galaxies2)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePaths(galaxies []*Galaxy) int {
	ret := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			ret += galaxies[i].pathLen(galaxies[j])
		}
	}
	return ret
}

func expandAllGalaxy2(galaxies [][2]int, times int) {

	for coord := 0; coord < 2; coord++ {

		slices.SortFunc(galaxies, func(a, b [2]int) int {
			return a[coord] - b[coord]
		})

		oldXIdx := 0
		for i := 0; i < len(galaxies); i++ {
			newXIdx := galaxies[i][coord]
			if i > 0 {
				oldXIdx = galaxies[i-1][coord]
			}
			gap := newXIdx - oldXIdx - 1
			if gap < 1 {
				oldXIdx = newXIdx
				continue
			}
			for j := i; j < len(galaxies); j++ {
				galaxies[j][coord] += (gap * times)
			}
		}
	}
}

func expandAllGalaxy(galaxies []*Galaxy, times int) {

	// sort by X, then for all galaxies add the differences to all following Xs from the current X

	slices.SortFunc(galaxies, func(a, b *Galaxy) int {
		return a.x - b.x
	})

	// fmt.Println("X-sort")
	// printGalaxies(galaxies)

	oldXIdx := 0
	for i := 0; i < len(galaxies); i++ {
		// fmt.Println("X: processing ", galaxies[i])
		newXIdx := galaxies[i].x
		if i > 0 {
			oldXIdx = galaxies[i-1].x
		}
		gap := newXIdx - oldXIdx - 1
		// fmt.Println("Gap ", gap)

		if gap < 1 {
			oldXIdx = newXIdx
			continue
		}

		for j := i; j < len(galaxies); j++ {
			// fmt.Println("Adding gap onto ", galaxies[j], gap, i, j)
			galaxies[j].x += (gap * times)
		}
	}

	// fmt.Println("X-pansion")
	// printGalaxies(galaxies)

	slices.SortFunc(galaxies, func(a, b *Galaxy) int {
		return a.y - b.y
	})

	// fmt.Println("Y-sort")
	// printGalaxies(galaxies)

	oldYIdx := 0
	for i := 0; i < len(galaxies); i++ {
		newYIdx := galaxies[i].y
		if i > 0 {
			oldYIdx = galaxies[i-1].y
		}
		gap := newYIdx - oldYIdx - 1

		if gap < 1 {
			oldYIdx = newYIdx
			continue
		}

		for j := i; j < len(galaxies); j++ {
			galaxies[j].y += (gap * times)
		}
	}
	// fmt.Println("yPansion")
	// printGalaxies(galaxies)
}

func parseInput(input string) []*Galaxy {
	lines := strings.Split(input, "\n")

	galaxies := make([]*Galaxy, 0, 0)

	for y, line := range lines {
		for x, c := range line {
			if string(c) == "#" {
				galaxies = append(galaxies, &Galaxy{x: x, y: y})
			}
		}
	}

	return galaxies
}

type Galaxy struct {
	x, y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func printGalaxies(g []*Galaxy) {
	for _, gal := range g {
		fmt.Println(gal)
	}
}

func diff(a, b int) int {
	d := b - a
	if d > 0 {
		return d
	}
	return -1 * d
}

func (g1 *Galaxy) pathLen(g2 *Galaxy) int {
	return diff(g1.x, g2.x) + diff(g1.y, g2.y)
}
