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
	galaxies, galaxies2 := parseInput(inputString), parseInput(inputString)

	expandAllGalaxy(galaxies, 2-1)    // each line replaced with 2
	expandAllGalaxy(galaxies2, 1e6-1) // each line replaced with 1e6

	part1, part2 := solvePaths(galaxies), solvePaths(galaxies2)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePaths(galaxies [][2]int) int {
	ret := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			ret += pathLen(galaxies[i], galaxies[j])
		}
	}
	return ret
}

func expandAllGalaxy(galaxies [][2]int, times int) {

	for coord := 0; coord < 2; coord++ {
		slices.SortFunc(galaxies, func(a, b [2]int) int {
			return a[coord] - b[coord]
		})

		for i := 0; i < len(galaxies); i++ {
			previousCoord, currentCoord := 0, galaxies[i][coord]
			if i > 0 {
				previousCoord = galaxies[i-1][coord]
			}

			emptyLines := currentCoord - previousCoord - 1
			if emptyLines < 1 {
				continue
			}

			for j := i; j < len(galaxies); j++ {
				galaxies[j][coord] += (emptyLines * times)
			}
		}
	}
}

func parseInput(input string) [][2]int {
	galaxies := make([][2]int, 0, 0)

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			if string(c) == "#" {
				galaxies = append(galaxies, [2]int{x, y})
			}
		}
	}

	return galaxies
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printGalaxies(galaxies [][2]int) {
	for _, galaxy := range galaxies {
		fmt.Println(galaxy)
	}
}

func diff(a, b int) int {
	d := b - a
	if d > 0 {
		return d
	}
	return -1 * d
}

func pathLen(g1, g2 [2]int) int {
	return diff(g1[0], g2[0]) + diff(g1[1], g2[1])
}
