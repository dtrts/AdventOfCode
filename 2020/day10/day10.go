package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// input = inputTest

	inputSlice := []int{}

	for _, line := range strings.Split(input, "\n") {
		lineInt, _ := strconv.Atoi(line)
		inputSlice = append(inputSlice, lineInt)
	}

	// Add your own adaptor
	sort.Ints(inputSlice)
	inputSlice = append(inputSlice, inputSlice[len(inputSlice)-1]+3)
	fmt.Println(inputSlice)

	oneJolt := 0
	threeJolt := 0
	ip := 0

	for _, jolt := range inputSlice {
		diff := jolt - ip

		if diff == 3 {
			threeJolt++
		} else if diff == 1 {
			oneJolt++
		} else {
			fmt.Println("WARN: Other", jolt, ip, diff)
		}

		ip = jolt
	}

	fmt.Println(oneJolt, threeJolt, oneJolt*threeJolt)

	memory := make(map[int]int)
	fmt.Println(countArrangements(len(inputSlice)-1, inputSlice, memory))

}

func countArrangements(position int, values []int, memory map[int]int) int {
	fmt.Println("call", position)
	arrangements := 0
	jolt := values[position]

	// loop thrrough values below this, assuming in array, and the difference in
	// values is <= 3 (its sorted already)

	_, ok := memory[position]
	if ok {
		fmt.Println("returning precalc", memory[position])
		return memory[position]
	}

	for positionCheck := position - 1; positionCheck >= 0 && jolt-values[positionCheck] <= 3; positionCheck-- {
		// fmt.Println("forloop", positionCheck, jolt, jolt-values[positionCheck])
		arrangements += countArrangements(positionCheck, values, memory)
		// fmt.Println("forloopOut", positionCheck, jolt, arrangements)
	}

	if jolt <= 3 {
		// fmt.Println(position, jolt, "endValue")
		arrangements++
	}
	memory[position] = arrangements
	return arrangements

}
