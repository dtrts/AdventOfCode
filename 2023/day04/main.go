package main

import (
	"flag"
	"fmt"
	"log"
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

	cards := parseInput(inputString)

	part1 := solvePart1(cards)
	part2 := solvePart2(cards)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePart2(cards map[int]*Card) int {
	ret := 0

	memo := map[int]int{}
	maxCards := len(cards)

	for _, card := range cards {
		copyCount := countCopies(cards, maxCards, card, memo)
		ret += 1 + copyCount
	}
	return ret
}

// Using memo map to reduce recursion. 60ms -> 2ms for full input
func countCopies(cards map[int]*Card, maxCards int, card *Card, memo map[int]int) int {

	count := 0
	minCopyId := card.id + 1
	maxCopyId := card.id + card.matches
	if maxCopyId > maxCards {
		maxCopyId = maxCards
	}

	for copyCardId := minCopyId; copyCardId <= maxCopyId; copyCardId++ {
		count++
		copyCount := 0

		if _, ok := memo[copyCardId]; ok {
			copyCount = memo[copyCardId]
		} else {
			copyCount = countCopies(cards, maxCards, cards[copyCardId], memo)
			memo[copyCardId] = copyCount
		}

		count += copyCount
	}

	return count
}

func solvePart1(cards map[int]*Card) int {
	ret := 0

	for _, card := range cards {
		ret += card.points
	}

	return ret
}

func (c *Card) processCardNumbs() {
	matches := 0

	for _, haveNum := range c.haveNumbers {
		for _, winningNum := range c.winningNumbers {
			if haveNum == winningNum {
				matches++
			}
		}
	}

	c.matches = matches
}

func (c *Card) processPoints() {
	value := 0
	switch c.matches {
	case 0:
		value = 0
	case 1:
		value = 1
	default:
		for i := 2; i <= c.matches; i++ {
			value *= 2
		}
	}
	c.points = value
}

func parseInput(input string) map[int]*Card {

	lines := strings.Split(input, "\n")

	cards := make(map[int]*Card)

	for _, line := range lines {
		inputCardId, inputNumbers, _ := strings.Cut(line, ":")

		cardId := 0
		fmt.Sscanf(inputCardId, "Card %d", &cardId)

		inputWinningNumbers, inputHaveNumbers, _ := strings.Cut(inputNumbers, "|")

		cards[cardId] = &Card{
			id:             cardId,
			winningNumbers: parseNumbers(inputWinningNumbers),
			haveNumbers:    parseNumbers(inputHaveNumbers),
		}

		cards[cardId].processCardNumbs()
		cards[cardId].processPoints()
	}

	return cards
}

func parseNumbers(s string) []int {
	// Split by one space and remove empty strings
	numberStrings := slices.DeleteFunc(strings.Split(s, " "), func(e string) bool { return e == "" })

	numbers := make([]int, 0, len(numberStrings))

	for _, numberString := range numberStrings {
		num, _ := strconv.Atoi(numberString)
		numbers = append(numbers, num)
	}

	return numbers
}

type Card struct {
	id             int
	winningNumbers []int
	haveNumbers    []int
	matches        int
	points         int
}
