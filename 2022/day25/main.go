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

	// BOILER PLATE --------------------------------------------------------------------

	lines := strings.Split(inputString, "\n")

	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
		// sum += convertToDecimal(line)
	}

	// p(sum)

	numbers := make([]map[int]int, 0)

	for _, line := range lines {
		number := make(map[int]int, 0)
		pos := 0
		for i := len(line) - 1; i >= 0; i-- {
			char := string(line[i])
			if char == "-" {
				char = "-1"
			}
			if char == "=" {
				char = "-2"
			}

			num, err := strconv.Atoi(char)
			if err != nil {
				panic("asd")
			}

			number[pos] = num
			pos++
		}
		p(line, number)
		numbers = append(numbers, number)
	}

	sum := make(map[int]int, 0)
	for i := 0; i <= maxLength; i++ {
		sum[i] = 0
	}

	for _, number := range numbers {
		for i := 0; i <= maxLength; i++ {
			if digit, ok := number[i]; ok {
				sum[i] += digit
			}
		}
	}

	p("Sum", sum)

	decimalSum := 0
	for digit, amount := range sum {
		decimalSum += amount * pow(5, digit)
	}
	p(decimalSum)

	digit := 0
	for areWeDoneYet := false; !areWeDoneYet; {
		parity := 1
		if sum[digit] < 0 {
			parity = -1
		}

		num := sum[digit]
		numAbs := num * parity

		p("digit", digit, "num", num, "numAbs", numAbs, "Parity", parity)
		// Take out any multiples of five and add onto the next digit up.
		if numAbs >= 5 {
			nextDigitAdd := numAbs / 5

			if _, ok := sum[digit+1]; !ok {
				sum[digit+1] = 0
			}

			sum[digit] -= parity * nextDigitAdd * 5
			sum[digit+1] += parity * nextDigitAdd

			p("  nextDA", nextDigitAdd, "new Current Digit", sum[digit], "next current digit", sum[digit+1])

			num = sum[digit]
			parity = 1
			if sum[digit] < 0 {
				parity = -1
			}
			numAbs = num * parity
		}

		// Now we need to handle -4, -3, 3 and 4
		// Same parity one above
		// Then handle distance to five.

		if numAbs > 2 {
			p("  Dealing with 3s and 4s")

			diff := 5 - numAbs

			p("    digit", digit, "num", num, "numAbs", numAbs, "Parity", parity, "diff", diff)

			if _, ok := sum[digit+1]; !ok {
				sum[digit+1] = 0
			}

			sum[digit] = -1 * parity * diff
			sum[digit+1] += parity

			p("  nextDA", diff, "new Current Digit", sum[digit], "next current digit", sum[digit+1])

			num = sum[digit]
			parity = 1
			if sum[digit] < 0 {
				parity = -1
			}
			numAbs = num * parity
		}

		maxSumDigit := 0
		for sumDigit := range sum {
			if sumDigit > maxSumDigit {
				maxSumDigit = sumDigit
			}
		}

		digit++
		if digit > maxSumDigit {
			areWeDoneYet = true
		}
	}

	p("Sum", sum)

	decimalSum = 0
	for digit, amount := range sum {

		decimalSum += amount * pow(5, digit)

	}

	//33448434171005
	p(decimalSum)

	maxSumDigit := 0
	for sumDigit := range sum {
		if sumDigit > maxSumDigit {
			maxSumDigit = sumDigit
		}
	}

	SNAFU := ""
	for digit := maxSumDigit; digit >= 0; digit-- {

		if _, ok := sum[digit]; !ok {
			sum[digit] = 0
		}

		switch sum[digit] {
		case -2:
			SNAFU = SNAFU + "="
		case -1:
			SNAFU = SNAFU + "-"
		case 0:
			SNAFU = SNAFU + "0"
		case 1:
			SNAFU = SNAFU + "1"
		case 2:
			SNAFU = SNAFU + "2"
		default:
			panic("Uh oh ")

		}
	}

	p(strings.TrimLeft(SNAFU, "0"))

	// p(pow(5, 0))
	// p(pow(5, 1))
	// p(pow(5, 2))
	// p(pow(5, 3))
	// p(pow(5, 4))
	// p(pow(5, 5))
	// p(pow(5, 6))
	// p(pow(5, 20))

	// mult := 1
	// base := 25
	// for

	// SNAFtoDecimal := map[string]int

	// for len <

	// sum := 0
	// for _, line := range strings.Split(inputString, "\n") {

	// 	sum += convertToDecimal(line)
	// 	p("")
	// }

	// p(sum)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:")
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}

// Generate

func pow(num, exp int) int {

	if exp == 0 {
		return 1
	}

	ret := num

	for i := 2; i <= exp; i++ {
		ret *= num
	}
	return ret
}

func convertToDecimal(s string) int {

	multiples := []int{}

	for sI := len(s) - 1; sI >= 0; sI-- {
		val := string(s[sI])

		if val == "-" {
			val = "-1"
		}
		if val == "=" {
			val = "-2"
		}

		numVal, err := strconv.Atoi(val)
		if err != nil {
			p(s, val, sI)
			panic("Unable to parse val")
		}

		multiples = append(multiples, numVal)
	}

	base := 5
	ret := 0
	for i, m := range multiples {
		ret += m * pow(base, i)
	}

	return ret

}
