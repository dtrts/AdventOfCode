package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	fmt.Println("Input bytes:", inputBytes)
	inputString := string(inputBytes)
	fmt.Println("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	fmt.Println("Calculating Part 1....")

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:")
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

	fmt.Println("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	elapsed = time.Since(start)
	fmt.Println("Part2:")
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

}
