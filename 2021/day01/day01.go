package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func 

func main() {
	start := time.Now()
	log.Printf("Starting... %s", start.String())

	testInputPtr := flag.Bool("testInput", false, "a bool")
	flag.Parse()

	if *testInputPtr {
		input = inputTest
	}

	fmt.Println(len(input))

	increases := 0

	for i := range input {
		if i == 0 {
			continue
		}

		if input[i-1] < input[i] {
			increases++
		}
	}

	fmt.Println(increases)

	// sumArray := []int{}
	increasesPart2 := 0

	for i := range input {
		if i < 3 {
			continue
		}

		if input[i-3] < input[i] {
			increasesPart2++
		}
	}
	// for i := range input {
	// 	if i < 2 {
	// 		continue
	// 	}
	// 	sumArray = append(sumArray, input[i]+input[i-1]+input[i-2])
	// }

	// for i := range sumArray {
	// 	if i == 0 {
	// 		continue
	// 	}

	// 	if sumArray[i-1] < sumArray[i] {
	// 		increasesPart2++
	// 	}

	// }

	fmt.Println(increasesPart2)

	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
}
