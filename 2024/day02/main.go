package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	reports := make([][]int, 0, len(lines))
	for _, line := range lines {
		reports = append(reports, stringsToInts(strings.Split(line, " ")))
	}

	for _, report := range reports {
		if isSafe(report) {
			part1++
		}
		if isSafe(report) {
			part2++
		} else {
			for i := range report {
				if isSafe(del(report, i)) {
					part2++
					break
				}
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

func isSafe(report []int) bool {
	direction := 0
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff == 0 {
			return false
		}
		if abs(diff) > 3 {
			return false
		}
		if direction == 0 {
			direction = norm(diff)
		}
		if direction != norm(diff) {
			return false
		}
	}
	return true
}

func stringsToInts(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func norm(i int) int {
	if i < 0 {
		return -1
	}
	if i == 0 {
		return 0
	}
	return 1
}

func del(report []int, idx int) []int {
	newSlice := make([]int, 0, len(report)-1)
	newSlice = append(newSlice, report[0:idx]...)
	newSlice = append(newSlice, report[idx+1:]...)
	return newSlice
}
