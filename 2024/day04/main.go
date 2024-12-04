package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	inputString := string(inputBytes)
	// fmt.Println("Input bytes:", inputBytes)
	// fmt.Println("Input string", inputString)
	inputString = strings.TrimSpace(inputString)
	// BOILER PLATE --------------------------------------------------------------------

	part1, part2 := 0, 0

	lines := strings.Split(inputString, "\n")

	grid := make([][]string, len(lines))

	for y, line := range lines {
		grid[y] = strings.Split(line, "")
	}

	dirs := []int{-1, 0, 1}

	for y, gridLine := range grid {
		for x := range gridLine {
			for _, yDiff := range dirs {
				for _, xDiff := range dirs {
					if spellsXmas(grid, x, y, xDiff, yDiff) {
						part1++
					}
				}
			}
			if isCrossX(grid, x, y) {
				part2++
			}
		}
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func spellsXmas(grid [][]string, x, y, xDiff, yDiff int) bool {
	limY, limX := len(grid), len(grid[0])

	searchChars := []string{"X", "M", "A", "S"}

	for i, c := range searchChars {
		newX := x + (xDiff * i)
		newY := y + (yDiff * i)

		if newX < 0 || newX >= limX || newY < 0 || newY >= limY {
			return false
		}
		if grid[newY][newX] != c {
			return false
		}
	}
	return true
}

func isCrossX(grid [][]string, x, y int) bool {
	limY, limX := len(grid), len(grid[0])

	if grid[y][x] != "A" {
		return false
	}

	if x-1 < 0 || y-1 < 0 || x+1 >= limX || y+1 >= limY {
		return false
	}

	tl := grid[y-1][x-1]
	tr := grid[y-1][x+1]
	bl := grid[y+1][x-1]
	br := grid[y+1][x+1]

	backslash := (tl == "M" && br == "S") || (tl == "S" && br == "M")
	slash := (tr == "M" && bl == "S") || (tr == "S" && bl == "M")

	return backslash && slash
}
