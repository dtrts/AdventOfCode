package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	inputString := strings.TrimSpace(string(inputBytes))
	p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------
	inputPairs := strings.Split(inputString, "\n\n")

	part1 := 0

	for index, inputPair := range inputPairs {
		inputPairSplit := strings.Split(inputPair, "\n")
		left, right := inputPairSplit[0], inputPairSplit[1]

		pair := index + 1

		p("")
		p("== Pair", pair, "==")
		inCorrectOrder := compare(left, right, 0, true)

		if inCorrectOrder == 1 {
			p("Groups in correct order. Pair:", pair)
			part1 += pair
		}

	}

	p("Calculating Part 2....")

	allPackets := strings.Fields(strings.Replace(inputString, "\n\n", "\n", -1))

	allPackets = append(allPackets, "[[2]]", "[[6]]")
	sort.Slice(allPackets, func(i, j int) bool {
		return compare(allPackets[i], allPackets[j], 0, false) > 0
	})

	twoI, sixI := 0, 0
	for i, packet := range allPackets {
		if packet == "[[2]]" {
			twoI = i + 1
		}
		if packet == "[[6]]" {
			sixI = i + 1
		}
		if twoI != 0 && sixI != 0 {
			break
		}
	}

	part2 := twoI * sixI

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
}

func compare(left, right string, nested int, debug bool) int {
	if debug {
		p(strings.Repeat("  ", nested), "- Compare", left, "vs", right)
	}

	if isDigit(left) && isDigit(right) {
		leftNum, _ := strconv.Atoi(left)
		rightNum, _ := strconv.Atoi(right)

		if leftNum < rightNum {
			if debug {

				p(strings.Repeat("  ", nested), "- Left side is smaller, so inputs are in the right order")
			}
			return 1
		}

		if rightNum < leftNum {
			if debug {
				p(strings.Repeat("  ", nested), "- Right side is smaller, so inputs are not in the right order")
			}
			return -1
		}

		return 0
	}

	if isList(left) && isList(right) {

		leftGroups, rightGroups := parseGroups(left), parseGroups(right)
		if debug {
			p(strings.Repeat("  ", nested), "- left groups", leftGroups, "rightGroups", rightGroups)
		}
		leftLen, rightLen := len(leftGroups), len(rightGroups)

		for i := 0; i < min(leftLen, rightLen); i++ {
			itemCompare := compare(leftGroups[i], rightGroups[i], nested+1, debug)
			if itemCompare != 0 {
				return itemCompare
			}
		}

		if leftLen < rightLen {
			if debug {
				p(strings.Repeat("  ", nested), "- Left side ran out of items, so inputs are in the right order")
			}
			return 1
		}
		if rightLen < leftLen {
			if debug {
				p(strings.Repeat("  ", nested), "- Right side ran out of items, so inputs are not in the right order")
			}
			return -1
		}

		return 0
	}

	if isDigit(left) && isList(right) {
		left = "[" + left + "]"
		if debug {
			p(strings.Repeat("  ", nested), "  - Mixed types; convert left to", left, "and retry comparison")
		}
		return compare(left, right, nested+1, debug)
	}

	if isDigit(right) && isList(left) {
		right = "[" + right + "]"
		if debug {
			p(strings.Repeat("  ", nested), "  - Mixed types; convert right to", right, "and retry comparison")
		}
		return compare(left, right, nested+1, debug)
	}

	panic("ASdwdasd")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isList(s string) bool {
	return string(s[0]) == "["
}

func isDigit(s string) bool {
	return !isList(s)
}

func parseGroups(s string) (ret []string) {

	// p("Start parse", s, len(s))
	nestedCount := -1
	group := ""
	for _, byte := range s {

		char := string(byte)

		if char == "[" {
			nestedCount++
			if nestedCount > 0 {
				group += char
			}
			continue
		}

		if char == "]" { // [[]] -> ['[]']
			nestedCount--

			if nestedCount == -1 {
				if len(group) > 0 {
					ret = append(ret, group)
				}
				group = ""
				continue
			}

			group += char

			if nestedCount == 0 {
				if len(group) > 0 {
					ret = append(ret, group)
				}
				group = ""
				continue
			}
		}

		if char == "," && nestedCount == 0 {
			if len(group) > 0 {
				ret = append(ret, group)
			}
			group = ""
			continue
		}

		group += char

	}

	return ret
}
