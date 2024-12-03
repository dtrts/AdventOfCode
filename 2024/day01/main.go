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

	inputString := string(inputBytes)
	// fmt.Println("Input bytes:", inputBytes)
	// fmt.Println("Input string", inputString)
	inputString = strings.TrimSpace(inputString)
	// BOILER PLATE --------------------------------------------------------------------

	part1, part2 := 0, 0

	lines := strings.Split(inputString, "\n")

	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))

	for _, line := range lines {
		chars := strings.Split(line, " ")
		leftInt, _ := strconv.Atoi(chars[0])
		rightInt, _ := strconv.Atoi(chars[3])
		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	for i, v := range leftList {
		part1 += abs(leftList[i] - rightList[i])

		part2 += v * count(rightList, v)
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func count(list []int, val int) int {
	count := 0

	for _, v := range list {
		if v == val {
			count++
		}
	}
	return count
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}
