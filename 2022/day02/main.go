package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

var scoreResults map[string]int = map[string]int{
	"win":  6,
	"draw": 3,
	"loss": 0,
}
var scoreChoice map[string]int = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var moveLeft map[string]string = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}
var moveRight map[string]string = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var wins []string = []string{"A Y", "B Z", "C X"}
var losses []string = []string{"A Z", "B X", "C Y"}

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
	// BOILER PLATE --------------------------------------------------------------------

	// Values
	rounds := strings.Split(inputString, "\n")
	score := 0
	for _, round := range rounds {
		yourChoice := strings.Fields(round)[1]
		yourShape := moveRight[yourChoice]

		score += scoreChoice[yourShape]

		if slices.Contains(wins, round) {
			score += scoreResults["win"]
		} else if slices.Contains(losses, round) {
			score += scoreResults["loss"]
		} else {
			score += scoreResults["draw"]
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", score)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

	fmt.Println("Calculating Part 2....")

	score = 0

	for _, round := range rounds {
		choices := strings.Fields(round)
		if choices[1] == "X" {
			// Need to lose
			move := findMove(choices[0], losses)
			score += scoreChoice[moveRight[move]]
			score += scoreResults["loss"]
		} else if choices[1] == "Y" {
			/// draw
			score += scoreChoice[moveLeft[choices[0]]]
			score += scoreResults["draw"]
		} else {
			// win
			move := findMove(choices[0], wins)
			score += scoreChoice[moveRight[move]]
			score += scoreResults["win"]
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed = time.Since(start)
	fmt.Println("Part2:", score)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

}

func findMove(opponent string, options []string) string {
	winningRound := slices.IndexFunc(options, func(s string) bool {
		return string(s[0]) == opponent
	})

	return string(options[winningRound][2])
}
