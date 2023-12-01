package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	digitsRe     := regexp.MustCompile(`1|2|3|4|5|6|7|8|9|0`)
	numsRe       := regexp.MustCompile(`1|2|3|4|5|6|7|8|9|0|one|two|three|four|five|six|seven|eight|nine`)
	numsReversRe := regexp.MustCompile(`1|2|3|4|5|6|7|8|9|0|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)

	part1, part2 := 0, 0

	for _, line := range lines {
		// Part 1
		digit1, digit2 := findFirstAndLast(line,digitsRe,digitsRe)
		part1 += sumDigits(digit1,digit2)

		// Part 2
		digit1, digit2 = findFirstAndLast(line, numsRe, numsReversRe)
		part2 += sumDigits(digit1,digit2)
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:",part1)
	fmt.Println("Part2:",part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

var digitNameMap = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

func findFirstAndLast(line string, forwardRE, backwardRE *regexp.Regexp) (string,string) {
		lineRev := reverseString(line)
		return convertNumberToDigit(forwardRE.FindString(line)), convertNumberToDigit(reverseString(backwardRE.FindString(lineRev)))
}

func convertNumberToDigit(s string) string {
	if digit, ok := digitNameMap[s]; ok {
		return digit
	}
  return s
}

func reverseString(s string) string {
	ret := ""
	for i := len(s)-1; i>= 0 ; i-- {
		ret += string(s[i])
	}
	return ret
}

func sumDigits(s1,s2 string) int {
	n1, _ := strconv.Atoi(s1)
	n2, _ := strconv.Atoi(s2)
	return (10 * n1) + n2
}
