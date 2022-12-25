package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	input := strings.TrimSpace(string(inputBytes))
	p("Input string", input)
	// BOILER PLATE --------------------------------------------------------------------

	p("Calculating Part 1....")

	p("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:")
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}
