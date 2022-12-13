package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"unicode"
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

	inputString := strings.TrimSpace(string(inputBytes))
	p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------
	inputPairs := strings.Split(inputString, "\n\n")

	// re := regexp.MustCompile(`\[(\[[\d,]+\]|\d)*(,((\[[\d,]+\]|\d)))+\]`)
	part1 := 0
	for index, inputPair := range inputPairs {
		inputPairSplit := strings.Split(inputPair, "\n")
		left, right := inputPairSplit[0], inputPairSplit[1]

		inCorrectOrder := compareGroups([]string{left}, []string{right}, 0)

		if inCorrectOrder {
			p("Groups in correct order. Index:", index)
			part1 += 1 + index
		}

	}

	p("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1) // 2136 is too low
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}

// func recursiveParseGroups(s string) (ret []string) {

// 	ret = append(ret, parseGroups(s)...)

// 	for _, group := range ret {

// 		if len(group) > 1 {
// 			ret = append(ret, recursiveParseGroups(group)...)
// 		}
// 	}

// 	return ret
// }

func compareGroups(left, right []string, nested int) bool {
	p("")
	nested++
	p(strings.Repeat("  ", nested), "Comparing", left, "and", right)

	for i, _ := range left {

		leftItem := left[i]
		rightItem := right[i]

		if len(leftItem) == 1 && len(rightItem) == 1 {
			leftNum, _ := strconv.Atoi(leftItem)
			rightNum, _ := strconv.Atoi(rightItem)

			p(strings.Repeat("  ", nested), "both numbers, comparing", leftNum, "and", rightNum)
			if leftNum > rightNum {
				return false
			}

		}

		if len(leftItem) == 1 && len(rightItem) > 1 {
			p(strings.Repeat("  ", nested), "Turning", leftItem, "into an array")
			leftItem = "[" + leftItem + "]"
			p(strings.Repeat("  ", nested), "Post conversion", leftItem)

		}
		if len(leftItem) > 1 && len(rightItem) == 1 {
			p(strings.Repeat("  ", nested), "Turning", rightItem, "into an array")
			rightItem = "[" + rightItem + "]"
			p(strings.Repeat("  ", nested), "Post conversion", rightItem)
		}

		if len(leftItem) > 1 && len(rightItem) > 1 {
			p(strings.Repeat("  ", nested), "Comparing two groups:", leftItem, "and", rightItem)
			leftGroups := parseGroups(leftItem)
			rightGroups := parseGroups(rightItem)

			p(strings.Repeat("  ", nested), "New Groups:", leftGroups, "and", rightGroups)
			if len(rightGroups) < len(leftGroups) {
				return false
			}

			return compareGroups(leftGroups, rightGroups, nested)
		}

	}

	return true
	// If both values are integers, the lower integer should come first. If the left integer is lower than the right integer, the
	// inputs are in the right order. If the left integer is higher than the right integer, the inputs are not in the right order.
	// Otherwise, the inputs are the same integer; continue checking the next part of the input.

	// If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out
	// of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the
	// right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the
	// next part of the input.

	// If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry
	// the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2); the result
	// is then found by instead comparing [0,0,0] and [2].

}

func parseGroups(s string) (ret []string) {

	// p("Start parse", s, len(s))
	nestedCount := -1
	group := ""
	for _, byte := range s {

		char := string(byte)

		if nestedCount == 0 && unicode.IsDigit(rune(char[0])) {
			group += char
			ret = append(ret, group)
			group = ""
			continue
		}

		if nestedCount == 0 && char == "," {
			group = ""
			continue
		}

		if char == "[" {
			nestedCount++
			if nestedCount > 0 {
				group += char
			}
			continue
		}

		if char == "]" {
			if nestedCount > 0 {
				group += char
			}
			nestedCount--
			if nestedCount == 0 {
				ret = append(ret, group)
				group = ""
			}
			continue
		}

		if nestedCount > 0 {
			group += char
		}

	}

	return ret
}
