package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input = inputTest
	// preamble := 5
	preamble := 25

	inputLines := []int{}
	for _, line := range strings.Split(input, "\n") {
		inputInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		inputLines = append(inputLines, inputInt)
	}

	fmt.Println(preamble, inputLines)

	positionInvalidNum := 0

	for position, value := range inputLines {
		if position >= preamble {
			if sumOfPrevious(position, preamble, inputLines) == false {
				positionInvalidNum = position
				fmt.Println(position, value, sumOfPrevious(position, preamble, inputLines))
				break
			}

		}
	}

	valueInvalidNum := inputLines[positionInvalidNum]
	fmt.Println(positionInvalidNum, valueInvalidNum)
	// for _, v := range inputLines {
	start, length := findValue(valueInvalidNum, inputLines)
	fmt.Println(valueInvalidNum, start, length, smallest(inputLines[start:start+length]), biggest(inputLines[start:start+length]), smallest(inputLines[start:start+length])+biggest(inputLines[start:start+length]))
	// }
}

func smallest(values []int) int {
	o := 2147483647
	for _, v := range values {
		if v < o {
			o = v
		}
	}
	return o
}

func biggest(values []int) int {
	o := 0
	for _, v := range values {
		if v > o {
			o = v
		}
	}
	return o
}
func findValue(value int, values []int) (int, int) {
	for i := range values {
		for length := 0; length+i <= len(values) && sumOfRange(i, length, values) <= value; length++ {
			if sumOfRange(i, length, values) == value {
				return i, length
			}
		}
	}
	return -1, -1
}

func sumOfRange(position int, length int, values []int) int {
	sum := 0
	for _, v := range values[position : position+length] {
		sum += v

	}
	return sum
}

func sumOfPrevious(position int, preamble int, values []int) bool {

	// preambleSlice := values[position-preamble : position]

	for i, v := range values[position-preamble : position] {
		for _, v2 := range values[position-preamble+i : position] {
			if v+v2 == values[position] {
				// fmt.Println(position, preamble, i, v, v2, values[position-preamble:position], values[position-preamble+i:position], values[position])
				return true
			}
		}
	}
	return false
}
