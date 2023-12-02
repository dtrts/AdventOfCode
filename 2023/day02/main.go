package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	// fmt.Println("Input bytes:", inputBytes)
	// fmt.Println("Input string", inputString)
	inputString = strings.TrimSpace(inputString)
	// BOILER PLATE --------------------------------------------------------------------

	lines := strings.Split(inputString, "\n")
	part1, part2 := 0, 0


	games := make([]Game,0,len(lines))

	for _, line := range lines {
		games = append(games, parseLine(line))
	}

	part1 = solvePart1(games)
	part2 = solvePart2(games)


	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:",part1)
	fmt.Println("Part2:",part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePart2(games []Game) int {
	ret := 0
	for _, game := range games {

		maxRed, maxGreen, maxBlue := 0, 0, 0

		for _, pull := range game.pulls {
			if pull.red > maxRed {
				maxRed = pull.red
			}
			if pull.green > maxGreen {
				maxGreen = pull.green
			}
			if pull.blue > maxBlue {
				maxBlue = pull.blue
			}
		}
		gamePower := maxRed * maxGreen * maxBlue

		ret += gamePower
	}

	return ret
}

func solvePart1(games []Game) int {
	//only 12 red cubes, 13 green cubes, and 14 blue cubes
	redLimit := 12
	greenLimit := 13
	blueLimit := 14
	ret := 0
	for _, game := range games {
		validGame := true

		for _, pull := range game.pulls {
			if pull.red > redLimit {
				validGame = false
			}
			if pull.green > greenLimit {
				validGame = false
			}
			if pull.blue > blueLimit {
				validGame = false
			}
		}

		if validGame {
			ret += game.id
		}

	}

	return ret
}

func parseLine(line string) Game {

	game := Game{}
	gameIdSplit := strings.Split(strings.TrimSpace(line), ":")
	gameId, _ := strconv.Atoi(gameIdSplit[0][5:])

	game.id = gameId

	pulls := strings.Split(gameIdSplit[1], ";")

	game.pulls = make([]Pull, 0, len(pulls))

	for _, pull := range pulls {
			game.pulls = append(game.pulls, parsePull(pull))
	}

	return game

}

func parsePull(line string) Pull {
	colourPulls := strings.Split(line, ",")

	pull := Pull{}


	for _, colourPull := range colourPulls {

		count, colour := 0, ""
		fmt.Sscanf(colourPull, "%d %s", &count, &colour)

		switch colour {
		case "red":
			pull.red = count
		case "green":
			pull.green = count
		case "blue":
			pull.blue = count
		}
	}

	return pull
}

type Game struct {
	id int
	pulls []Pull
}

type Pull struct {
	red int
	green int
	blue int
}
