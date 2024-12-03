package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	fmt.Println(inputString)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllString(inputString, -1)

	for _, match := range matches {
		re := regexp.MustCompile(`\d+`)
		numbers := stringsToInts(re.FindAllString(match, -1))
		part1 += numbers[0] * numbers[1]
	}

	re2 := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches2 := re2.FindAllString(inputString, -1)

	enabled := true

	for _, match := range matches2 {
		if match == "do()" {
			enabled = true
			continue
		} else if match == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			re := regexp.MustCompile(`\d+`)
			numbers := stringsToInts(re.FindAllString(match, -1))
			part2 += numbers[0] * numbers[1]
		}
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func stringsToInts(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}
