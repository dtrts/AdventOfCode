package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
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
	// BOILER PLATE --------------------------------------------------------------------

	fmt.Println("Calculating....")

	elves := strings.Split(inputString, "\n\n")

	totalByElf := make([]int, 0, len(elves))

	for _, elf := range elves {
		carriedCalories := strings.Split(elf, "\n")
		totalCaloriesByElf := 0

		for _, calorieStr := range carriedCalories {
			calorie, _ := strconv.Atoi(calorieStr)
			totalCaloriesByElf += calorie
		}

		totalByElf = append(totalByElf, totalCaloriesByElf)

	}

	slices.Delete()

	sort.Slice(totalByElf, func(i, j int) bool {
		return totalByElf[i] > totalByElf[j]
	})

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", totalByElf[0])
	fmt.Println("Part2:", totalByElf[0]+totalByElf[1]+totalByElf[2])
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}
