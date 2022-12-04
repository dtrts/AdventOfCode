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

	// BOILER PLATE --------------------------------------------------------------------

	pairs := strings.Fields(inputString)

	min1, max1, min2, max2 := 0, 0, 0, 0
	count := 0
	count2 := 0
	for _, pair := range pairs {

		fmt.Sscanf(pair, "%d-%d,%d-%d", &min1, &max1, &min2, &max2)

		if (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1) {
			count++
		}

		if between(min1, min2, max2) || between(max1, min2, max2) || between(min2, min1, max1) || between(max2, min1, max1) {
			fmt.Println((pair))
			count2++
		}

	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", count)
	fmt.Println("Part2:", count2)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

}

func between(val, min, max int) bool {
	if val >= min && val <= max {
		return true
	}
	return false
}
