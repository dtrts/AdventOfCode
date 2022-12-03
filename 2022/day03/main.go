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

	rucksackPairs := strings.Fields(inputString)

	score := 0

	for _, rucksackPair := range rucksackPairs {
		ruckSackA, ruckSackB := splitString(rucksackPair)

		for _, c := range ruckSackA {
			if strings.Contains(ruckSackB, string(c)) {
				score += scoreChar(string(c))
				break
			}
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	p("Part1:", score)
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

	score = 0

	for i := 0; i < len(rucksackPairs); i += 3 {
		group := rucksackPairs[i : i+3]

		for _, c := range group[0] {
			if strings.Contains(group[1], string(c)) && strings.Contains(group[2], string(c)) {
				score += scoreChar(string(c))
				break
			}
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	p("Part2:", score)
	elapsed = time.Since(start)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

}

func p(s ...interface{}) {
	fmt.Println(s...)
}

func scoreChar(c string) int {
	A, a := int(string('A')[0]), int(string('a')[0])

	cI := int(c[0])

	if cI >= a {
		return cI - a + 1
	}

	return cI - A + 1 + 26
}

func splitString(s string) (string, string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}
