package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

	inputString := strings.TrimSpace(string(inputBytes))
	fmt.Println("Input:", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	races, singleRace := parseInput(inputString)

	part1, part2 := solvePart1(races), solveRace(singleRace)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solveRace(race Race) int {
	lower, upper := solveQuadRatic(race)
	// We want all ints within the bounds solved by the quadratic.
	// So move the solution 'inward' and then find closest int outward.
	// Edge case if  | upper - lower | <= 1 ?
	return int(math.Ceil(upper-1)) - int(math.Floor(lower+1)) + 1
}

func solvePart1(races []Race) int {
	part1 := 1

	for _, race := range races {
		part1 *= solveRace(race)
	}

	return part1
}

func solveQuadRatic(race Race) (float64, float64) {

	// timeHolding = (( totalTime  +/- sqrRoot( totalTime^2 - 4 * distance) ) / 2)

	b2m4ac := math.Sqrt(math.Pow(float64(race.Time), 2) - 4*float64(race.Record))

	sol1 := (float64(race.Time) - b2m4ac) / 2
	sol2 := (float64(race.Time) + b2m4ac) / 2

	// return such that the values are in order
	if sol1 > sol2 {
		return sol2, sol1
	}

	return sol1, sol2
}

func parseInput(input string) ([]Race, Race) {

	timeString, distString, _ := strings.Cut(input, "\n")

	times := convertNumList(strings.Split(timeString, ":")[1], " ")
	dists := convertNumList(strings.Split(distString, ":")[1], " ")

	if len(times) != len(dists) {
		panic("o no")
	}

	races := make([]Race, 0, len(times))
	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			times[i],
			dists[i],
		})
	}

	singleTime, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(timeString, ":")[1], " ", ""))
	singleDist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(distString, ":")[1], " ", ""))

	singleRace := Race{singleTime, singleDist}

	return races, singleRace
}

func convertNumList(s string, sep string) []int {
	stringList := strings.Split(s, sep)
	stringList = slices.DeleteFunc(stringList, func(e string) bool { return e == "" })

	stringInt := make([]int, 0, len(stringList))

	for _, e := range stringList {
		n, _ := strconv.Atoi(e)
		stringInt = append(stringInt, n)
	}

	return stringInt
}

type Race struct {
	Time   int
	Record int
}

/*
speed (mm/ms) = timeHolding (ms)
distance (mm) = speed (mm/ms) * timeRacing (ms)
totalTime (ms) = timeHolding (ms) + timeRacing (ms)

----

distance = speed * (totalTime - timeHolding)

distance = timeHolding * (totalTime - timeHolding)

find timeHolding so that:

timeHolding^2 - totalTime * timeHolding + distance = 0

timeHolding = -b +- sqroot( b^2 - 4ac) / 2a

timeHolding = (( totalTime  +/- sqroot( totalTime^2 - 4 * distance) ) / 2)

distance = (speed * totalTime) - speed^2
*/
