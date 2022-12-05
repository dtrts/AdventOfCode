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

	// Parsing the input. Split into diagrams + instructions.
	parse1 := strings.Split(inputString, "\n\n")

	stacks, stacks2 := parseDiagram(parse1[0]), parseDiagram(parse1[0])

	instructions := strings.Split(strings.TrimSpace(parse1[1]), "\n")

	moveN, from, to := 0, 0, 0
	for _, instruction := range instructions {
		fmt.Sscanf(instruction, "move %d from %d to %d", &moveN, &from, &to)

		// Move values to 0-index
		from--
		to--

		// For part 1 move boxes one at a time
		for i := 0; i < moveN; i++ {
			stacks[to] = append(stacks[to], popN(&stacks[from], 1)...)
		}

		// For part 2 move boxes together
		stacks2[to] = append(stacks2[to], popN(&stacks2[from], moveN)...)
	}

	// BOILER PLATE --------------------------------------------------------------------
	fmt.Println("Part1:", topValues(stacks))
	fmt.Println("Part2:", topValues(stacks2))
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------
}
