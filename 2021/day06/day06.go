package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("Starting...")

	testInputPtr := flag.Bool("testInput", false, "a bool")
	flag.Parse()

	if *testInputPtr {
		fmt.Println("Using test Input")
		input = inputTest
	}

	lanternFish := parseInput(input)
	lanternMap := genMap(lanternFish)

	numDays := 256
	for day := 1; day <= numDays; day++ {
		lanternMap = tickMap(lanternMap)

		if day == 18 {
			fmt.Println("Quick Test:", sumMap(lanternMap))
			fmt.Println(lanternMap)

		}

		if day == 80 {
			fmt.Println("Part 1:", sumMap(lanternMap))
			fmt.Println(lanternMap)

		}

		if day == 256 {
			fmt.Println("Part 2:", sumMap(lanternMap))
			fmt.Println(lanternMap)

		}
	}

	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
}

func genMap(lanternFish []int) map[int]int {
	lanternMap := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, v := range lanternFish {
		lanternMap[v]++
	}

	return lanternMap
}

func tickMap(lanternMap map[int]int) map[int]int {
	zeros := lanternMap[0]
	sevens := lanternMap[7]

	for i := 0; i <= 8; i++ {
		if i != 0 && i != 7 {
			lanternMap[i-1] = lanternMap[i]
		}
	}
	lanternMap[8] = zeros
	lanternMap[6] = zeros + sevens
	return lanternMap
}

func sumMap(s map[int]int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func parseInput(input string) []int {
	splits := strings.Split(input, ",")
	ret := make([]int, len(splits))
	for i, v := range splits {
		num, err := strconv.Atoi(v)
		ret[i] = num
		if err != nil {
			panic(err)
		}
	}
	return ret
}

var inputTest string = `3,4,3,1,2`

var input string = `1,5,5,1,5,1,5,3,1,3,2,4,3,4,1,1,3,5,4,4,2,1,2,1,2,1,2,1,5,2,1,5,1,2,2,1,5,5,5,1,1,1,5,1,3,4,5,1,2,2,5,5,3,4,5,4,4,1,4,5,3,4,4,5,2,4,2,2,1,3,4,3,2,3,4,1,4,4,4,5,1,3,4,2,5,4,5,3,1,4,1,1,1,2,4,2,1,5,1,4,5,3,3,4,1,1,4,3,4,1,1,1,5,4,3,5,2,4,1,1,2,3,2,4,4,3,3,5,3,1,4,5,5,4,3,3,5,1,5,3,5,2,5,1,5,5,2,3,3,1,1,2,2,4,3,1,5,1,1,3,1,4,1,2,3,5,5,1,2,3,4,3,4,1,1,5,5,3,3,4,5,1,1,4,1,4,1,3,5,5,1,4,3,1,3,5,5,5,5,5,2,2,1,2,4,1,5,3,3,5,4,5,4,1,5,1,5,1,2,5,4,5,5,3,2,2,2,5,4,4,3,3,1,4,1,2,3,1,5,4,5,3,4,1,1,2,2,1,2,5,1,1,1,5,4,5,2,1,4,4,1,1,3,3,1,3,2,1,5,2,3,4,5,3,5,4,3,1,3,5,5,5,5,2,1,1,4,2,5,1,5,1,3,4,3,5,5,1,4,3`
