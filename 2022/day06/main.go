package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	// BOILER PLATE --------------------------------------------------------------------

	charactersProcessed, charactersProcessed2 := 0, 0

	for i := 4; i < len(inputBytes); i++ {
		if allDiffSlice(inputBytes[i-4 : i]) {
			charactersProcessed = i
			break
		}
	}

	for i := 14; i < len(inputBytes); i++ {
		if allDiffSlice(inputBytes[i-14 : i]) {
			charactersProcessed2 = i
			break
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	fmt.Println("Part1:", charactersProcessed)
	fmt.Println("Part2:", charactersProcessed2)
	// BOILER PLATE --------------------------------------------------------------------
}

func allDiffSlice[T comparable](s []T) bool {
	for i, v := range s {
		if slices.Contains(s[i+1:], v) {
			return false
		}
	}
	return true
}
