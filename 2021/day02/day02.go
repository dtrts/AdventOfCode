package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("Starting... %s", start.String())

	testInputPtr := flag.Bool("testInput", false, "a bool")
	flag.Parse()

	if *testInputPtr {
		fmt.Println("Using test Input")
		input = inputTest
	}

	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	// Position and depth start at 0
	aim := 0
	position := 0
	depth := 0
	depth2 := 0
	action := ""
	distance := 0
	for _, line := range inputLines {
		fmt.Sscanf(line, "%s %d", &action, &distance)

		switch action {
		case "forward":
			position += distance
			depth2 += aim * distance
		case "up":
			depth -= distance
			aim -= distance
		case "down":
			depth += distance
			aim += distance
		}
	}

	fmt.Println(position * depth)
	fmt.Println(position * depth2)
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
}
