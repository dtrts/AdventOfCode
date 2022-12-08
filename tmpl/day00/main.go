package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func p(s ...interface{}) {
	fmt.Println(s...)
}

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
	p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	p("Calculating Part 1....")

	p("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	p("Part1:")
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}
