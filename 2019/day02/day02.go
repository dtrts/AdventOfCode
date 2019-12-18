package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	linesStr := strings.Split(string(file), "\n")
	codesStr := strings.Split(linesStr[0], ",")

	codes := make([]int, len(codesStr), len(codesStr))
	for i, code := range codesStr {
		codes[i], err = strconv.Atoi(code)
		if err != nil {
			panic("")
		}
	}

	codes1 := make([]int, len(codes))
	copy(codes1, codes)
	codes1[1] = 12
	codes1[2] = 2

	for opLoc := 0; opLoc < len(codes1); opLoc += 4 {
		if codes1[opLoc] == 99 {
			break
		}
		if codes[opLoc] == 1 {
			codes1[codes1[opLoc+3]] = codes1[codes1[opLoc+1]] + codes1[codes1[opLoc+2]]
		}
		if codes[opLoc] == 2 {
			codes1[codes1[opLoc+3]] = codes1[codes1[opLoc+1]] * codes1[codes1[opLoc+2]]
		}
	}

	fmt.Println("Part 1: ", codes1[0])

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			testCodes := make([]int, len(codes))
			copy(testCodes, codes)

			testCodes[1] = noun
			testCodes[2] = verb

			for opLoc := 0; opLoc < len(testCodes); opLoc += 4 {
				if testCodes[opLoc] == 99 {
					break
				}
				if testCodes[opLoc] == 1 {
					testCodes[testCodes[opLoc+3]] = testCodes[testCodes[opLoc+1]] + testCodes[testCodes[opLoc+2]]
				}
				if codes[opLoc] == 2 {
					testCodes[testCodes[opLoc+3]] = testCodes[testCodes[opLoc+1]] * testCodes[testCodes[opLoc+2]]
				}
			}
			if testCodes[0] == 19690720 {
				fmt.Printf("Noun: %d. Verb: %d. Output: %d.\n", noun, verb, testCodes[0])
				fmt.Println("Part 2: ", (100*noun)+verb)
				break
			}
		}
	}
}
