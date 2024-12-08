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

	inputString := string(inputBytes)
	// fmt.Println("Input bytes:", inputBytes)
	// fmt.Println("Input string", inputString)
	inputString = strings.TrimSpace(inputString)
	// BOILER PLATE --------------------------------------------------------------------

	part1, part2 := 0, 0

	updates, rules := parseInput(inputString)

	for _, update := range updates {
		reducedRules := reduceRules(update, rules)
		if validateUpdates(update, reducedRules) {
			part1 += update[len(update)/2]
		} else {
			constructedUpdate := constructUpdate([]int{}, update, reducedRules)
			part2 += constructedUpdate[len(constructedUpdate)/2]
		}
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func constructUpdate(update []int, remainingPages []int, rules [][2]int) []int {
	// No need to duplicate slice when popping off the page, can reduce length as recursion deepens
	insertingPage := remainingPages[len(remainingPages)-1]
	newRemainingPages := remainingPages[:len(remainingPages)-1]

	for i := 0; i <= len(update); i++ {
		// Must duplicate slice here since we will be inserting at many points. slices.Insert alters the underlying slice. Safer to copy
		newUpdate := sliceInsert(update, i, insertingPage)
		newUpdateIsValid := validateUpdates(newUpdate, rules)

		if !newUpdateIsValid {
			continue
		}

		if newUpdateIsValid {
			if len(newRemainingPages) == 0 {
				return newUpdate
			}
			return constructUpdate(newUpdate, newRemainingPages, rules)
		}
	}

	panic("Shouldn't be here. Suggest bad input")
}

// func sliceDelete(s []int, i int) []int {
// 	newSlice := make([]int, i, len(s)-1)
// 	copy(newSlice, s[:i])
// 	newSlice = append(newSlice, s[i+1:]...)
// 	return newSlice
// }

// func slicePop(s []int, i int) ([]int, int) {
// 	val := s[i]
// 	return sliceDelete(s, i), val
// }

func sliceInsert(s []int, i int, v int) []int {
	newSlice := make([]int, i, len(s)+1)
	copy(newSlice, s[:i])
	newSlice = append(newSlice, v)
	if i < len(s) {
		newSlice = append(newSlice, s[i:]...)
	}
	return newSlice
}

func validateUpdates(update []int, rules [][2]int) bool {
	for _, rule := range rules {
		indexBefore := slices.Index(update, rule[0])
		if indexBefore < 0 {
			continue
		}
		indexAfter := slices.Index(update, rule[1])
		if indexAfter < 0 {
			continue
		}
		if indexBefore > indexAfter {
			return false
		}
	}
	return true
}

func reduceRules(update []int, rules [][2]int) [][2]int {
	reducedRules := [][2]int{}

	for _, rule := range rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
			reducedRules = append(reducedRules, rule)
		}
	}

	return reducedRules
}

func parseInput(inputString string) ([][]int, [][2]int) {

	parts := strings.Split(inputString, "\n\n")

	rulesRaw := strings.Split(parts[0], "\n")

	rules := make([][2]int, 0, len(rulesRaw))
	for _, ruleRaw := range rulesRaw {
		rule := stringsToInts(strings.Split(ruleRaw, "|"))
		rules = append(rules, [2]int{rule[0], rule[1]})
	}

	updatesRaw := strings.Split(parts[1], "\n")
	updates := make([][]int, 0, len(updatesRaw))
	for _, update := range updatesRaw {
		splitUpdates := strings.Split(update, ",")

		updates = append(updates, stringsToInts(splitUpdates))
	}

	return updates, rules
}

func stringsToInts(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}
