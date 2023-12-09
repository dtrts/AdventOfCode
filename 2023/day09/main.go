package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
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

	inputString := strings.TrimSpace(string(inputBytes))
	// fmt.Println("Input:", inputString)
	// BOILER PLATE --------------------------------------------------------------------
	part1, part2 := 0, 0

	histories := parseInput(inputString)

	for _, history := range histories {
		back, front := extrapolate(history)
		part1 += front
		part2 += back
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func extrapolate(history []int) (int, int) {
	allZeros := true
	for _, i := range history {
		if i != 0 {
			allZeros = false
		}
	}
	if allZeros {
		return 0, 0
	}

	diffs := make([]int, 0, len(history)-1)

	for i := 0; i < len(history)-1; i++ {
		diffs = append(diffs, history[i+1]-history[i])
	}

	back, front := extrapolate(diffs)
	return history[0] - back, history[len(history)-1] + front
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	histories := make([][]int, 0, len(lines))

	for _, line := range lines {
		history := convertNumList(line, " ")
		histories = append(histories, history)
	}

	return histories
}

func convertNumList(s string, sep string) []int {
	stringList := strings.Split(s, sep)
	stringList = slices.DeleteFunc(stringList, func(e string) bool { return e == "" })

	stringInt := make([]int, 0, len(stringList))

	for _, e := range stringList {
		n, _ := strconv.Atoi(e)
		stringInt = append(stringInt, n)
	}

	return stringInt
}
