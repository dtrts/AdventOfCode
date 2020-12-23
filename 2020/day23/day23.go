package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"time"
)

var input = []int{1, 3, 7, 8, 2, 6, 4, 9, 5}
var inputTest = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

func main() {

	// Part 1
	start := time.Now()

	currentCup := initRing(input, len(input))

	nodeMap := initNodeMap(currentCup)

	maxValue := len(input)

	for moveNumber := 1; moveNumber <= 100; moveNumber++ {

		currentCup = move(currentCup, moveNumber, maxValue, nodeMap)

	}

	fmt.Printf("\n-- Part 1 --\n%v\nDuration: %v", part1Score(currentCup), time.Since(start))

	// Part 2
	start = time.Now()

	currentCup = initRing(input, 1e6)

	nodeMap = initNodeMap(currentCup)

	maxValue = 1e6

	for moveNumber := 1; moveNumber <= 1e7; moveNumber++ {

		currentCup = move(currentCup, moveNumber, maxValue, nodeMap)

	}

	fmt.Printf("\n-- Part 2 --\n%v\nDuration: %v", part2Score(currentCup), time.Since(start))

}

func initRing(input []int, expandedInputLength int) *ring.Ring {

	currentCup := ring.New(expandedInputLength)

	for _, v := range input {

		currentCup.Value = v

		currentCup = currentCup.Next()

	}

	for i := len(input); i < expandedInputLength; i++ {

		currentCup.Value = i + 1

		currentCup = currentCup.Next()

	}

	return currentCup

}

func initNodeMap(currentCup *ring.Ring) map[int]*ring.Ring {

	nodeMap := make(map[int]*ring.Ring)

	nodeMap[currentCup.Value.(int)] = currentCup

	for travellingCup := currentCup.Next(); travellingCup != currentCup; travellingCup = travellingCup.Next() {

		nodeMap[travellingCup.Value.(int)] = travellingCup

	}

	return nodeMap

}

func move(currentCup *ring.Ring, moveNumber int, maxValue int, nodeMap map[int]*ring.Ring) *ring.Ring {

	removedCups := currentCup.Link(currentCup.Move(4))

	destinationValue := destination(currentCup.Value.(int), maxValue, removedCups)

	nodeMap[destinationValue].Link(removedCups)

	return currentCup.Next()

}

func part1Score(currentCup *ring.Ring) (resultString string) {

	for currentCup.Value.(int) != 1 {

		currentCup = currentCup.Next()

	}

	currentCup = currentCup.Next()

	for currentCup.Value.(int) != 1 {

		resultString += strconv.Itoa(currentCup.Value.(int))

		currentCup = currentCup.Next()

	}

	return resultString

}

func part2Score(currentCup *ring.Ring) int {

	ans := 1

	for currentCup.Value.(int) != 1 {

		currentCup = currentCup.Next()

	}

	for i := 0; i < 2; i++ {
		currentCup = currentCup.Next()

		ans *= currentCup.Value.(int)

	}

	return ans

}

func destination(value int, maxValue int, removedCups *ring.Ring) int {

	value = subtractOne(value, maxValue)

	for ringContains(removedCups, value) {

		value = subtractOne(value, maxValue)

	}

	return value

}

func ringContains(r *ring.Ring, value int) bool {

	ringLength := r.Len

	for i := 0; i < ringLength(); i++ {

		if r.Move(i).Value.(int) == value {

			return true

		}

	}

	return false

}

func subtractOne(value int, maxValue int) int {

	if value == 1 {

		return maxValue

	}

	return value - 1
}
