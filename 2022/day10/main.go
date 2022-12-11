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
	flag.StringVar(&inputFileName, "inputFileName", "input.txt", "Name of the input file")
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))
	p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	instructions := strings.Split(inputString, "\n")

	cycles := []int{}
	cycle := 1
	register := 1
	value := 0
	cycles = append(cycles, register) // just to make the matsh easier belowe
	pixels := []string{}

	for _, instruction := range instructions {

		// Cycle 1
		cycles = append(cycles, register)
		if register-1 <= (cycle-1)%40 && (cycle-1)%40 <= register+1 {
			pixels = append(pixels, "#")
		} else {
			pixels = append(pixels, ".")
		}
		cycle++

		// Cycle 2
		if instruction != "noop" {
			cycles = append(cycles, register)
			if register-1 <= (cycle-1)%40 && (cycle-1)%40 <= register+1 {
				pixels = append(pixels, "#")
			} else {
				pixels = append(pixels, ".")
			}
			cycle++

			fmt.Sscanf(instruction, "addx %d", &value)
			register += value
		}

	}
	cycles = append(cycles, register)
	if register-1 <= (cycle-1)%40 && (cycle-1)%40 <= register+1 {
		pixels = append(pixels, "#")
	} else {
		pixels = append(pixels, ".")
	}

	p(cycles)

	part1 := 0
	for _, v := range []int{20, 60, 100, 140, 180, 220} {
		part1 += (cycles[v] * v)
	}

	p(cycles[20], cycles[60], cycles[100], cycles[140], cycles[180], cycles[220])

	p(pixels[40*0 : 40*1])
	p(pixels[40*1 : 40*2])
	p(pixels[40*2 : 40*3])
	p(pixels[40*3 : 40*4])
	p(pixels[40*4 : 40*5])
	p(pixels[40*5 : 40*6])

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	p("Part1:", part1)
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}
