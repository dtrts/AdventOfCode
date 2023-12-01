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
	fmt.Println("Input bytes:", inputBytes)
	fmt.Println("Input string", inputString)
	inputString = strings.TrimSpace(inputString)
	// BOILER PLATE --------------------------------------------------------------------

	lines := strings.Split(inputString, "\n")

	fmt.Println(lines)

	digitsRe := regexp.MustCompile(`^(1|2|3|4|5|6|7|8|9|0)`)
	numsRe := regexp.MustCompile(`^(1|2|3|4|5|6|7|8|9|0|one|two|three|four|five|six|seven|eight|nine)`)

	part1, part2 := 0, 0

	for _, line := range lines {
		digitsStrings := []string{}
		numsStrings, numsStringsConv := []string{}, []string{}

		// Search the start of ever decreasing line to include overlapping numbers (FindAll doesn't
		// match overlaps :( )
		for i := 0; i < len(line); i++ {
			shortLine := line[i:]

			digitString := digitsRe.FindString(shortLine)
			if len(digitString) > 0 {
				digitsStrings = append(digitsStrings, digitString)
			}

			numString := numsRe.FindString(shortLine)
			if len(numString) > 0 {
				numsStrings = append(numsStrings, numString)
			}
		}

		numsStringsConv = convertNumNames(numsStrings)
		fmt.Println("Line:     ", line)
		fmt.Println("Digits:   ", digitsStrings)
		fmt.Println("Nums:     ", numsStrings)
		fmt.Println("NumsConv: ", numsStringsConv)
		fmt.Println("")

		if len(digitsStrings) > 0 {
			digits := convertToIntSlice(digitsStrings)
			part1 += (10*digits[0]) + digits[len(digits)-1]
		}
		if len(numsStringsConv) > 0 {
			nums := convertToIntSlice(numsStringsConv)
			part2 += (10*nums[0]) + nums[len(nums)-1]
		}

	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:",part1)
	fmt.Println("Part2:",part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func convertToIntSlice(s []string) []int {
	i := make([]int, 0, len(s))

	for _, v := range s {
		vi, _ := strconv.Atoi(v)
		i = append(i, vi)
	}

	return i
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

func convertNumNames(s []string) []string {

	ret := make([]string,0, len(s))

	for _, v := range s {
		if digit, ok := digitNameMap[v]; ok {
			ret = append(ret, digit)
		} else {
			ret = append(ret, v)
		}
	}
	return ret
}
