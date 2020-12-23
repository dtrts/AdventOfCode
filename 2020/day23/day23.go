package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"time"
)

// var input = [9]uint8{1, 3, 7, 8, 2, 6, 4, 9, 5}
// var inputTest = [9]uint8{3, 8, 9, 1, 2, 5, 4, 6, 7}
// var minCup = uint8(1)
// var maxCup = uint8(9)
// var moveNumber = 1
// var debug = false
// var gameLength = 9

var input = []int{1, 3, 7, 8, 2, 6, 4, 9, 5}
var inputTest = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
var debug = false

func main() {
	// input = inputTest
	fmt.Println("hw")
	maxValue := len(input)
	// Create game with length of input
	currentCup := ring.New(len(input))
	// Set the value of each element to one of the cup numbers
	for _, v := range input {
		currentCup.Value = v
		currentCup = currentCup.Next()
	}

	nodeMap := make(map[int]*ring.Ring)
	nodeMap[currentCup.Value.(int)] = currentCup
	for travellingCup := currentCup.Next(); travellingCup != currentCup; travellingCup = travellingCup.Next() {
		nodeMap[travellingCup.Value.(int)] = travellingCup
	}

	// Game starts here

	for moveNumber := 1; moveNumber <= 100; moveNumber++ {
		currentCup = move(currentCup, moveNumber, maxValue, nodeMap)
	}

	printf("\n-- final --\n")
	printCups(currentCup)
	fmt.Printf("\n-- Part 1 --\n%v\n", result(currentCup))

	// Part 2.
	// Take input,
	fmt.Printf("Creating part2input...")
	start := time.Now()

	// part2InputTest := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	part2InputTest := []int{1, 3, 7, 8, 2, 6, 4, 9, 5} // actual input
	part2InputTest = append(part2InputTest, make([]int, 1e6-9)...)
	for i := 10; i <= 1e6; i++ {
		part2InputTest[i-1] = i
	}

	fmt.Println("Done", time.Since(start))

	// fmt.Println("len", len(part2InputTest), part2InputTest[:100], part2InputTest[1e6-10:])

	input = part2InputTest
	maxValue = 1e6
	currentCup = ring.New(len(input))

	// Set the value of each element to one of the cup numbers
	for _, v := range input {
		currentCup.Value = v
		currentCup = currentCup.Next()
	}

	nodeMap = make(map[int]*ring.Ring)
	nodeMap[currentCup.Value.(int)] = currentCup
	for travellingCup := currentCup.Next(); travellingCup != currentCup; travellingCup = travellingCup.Next() {
		nodeMap[travellingCup.Value.(int)] = travellingCup
	}

	start = time.Now()

	for moveNumber := 1; moveNumber <= 1e7; moveNumber++ {
		currentCup = move(currentCup, moveNumber, maxValue, nodeMap)
	}

	printf("\n-- final --\n")
	// printCups(currentCup)
	fmt.Printf("\n-- Part 2 --\n%v\n", result2(currentCup))
	fmt.Println("Done", time.Since(start))

	// part2Input :=  []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

}

func result2(currentCup *ring.Ring) int {
	ans := 1

	for currentCup.Value.(int) != 1 {
		currentCup = currentCup.Next()
	}
	currentCup = currentCup.Next()
	ans *= currentCup.Value.(int)
	currentCup = currentCup.Next()
	ans *= currentCup.Value.(int)
	return ans
}

func move(currentCup *ring.Ring, moveNumber int, maxValue int, nodeMap map[int]*ring.Ring) *ring.Ring {
	if moveNumber%10000 == 0 {
		fmt.Printf("\n-- move %v --\nTime:%v", moveNumber, time.Now())
	}
	printf("\n-- move %v --\n", moveNumber)
	printf("cups: ")
	printCups(currentCup)

	removedCups := currentCup.Link(currentCup.Move(4))
	destinationValue := destination(currentCup.Value.(int), maxValue, removedCups)

	printf("pick up: ")
	printCups(removedCups)
	printf("destination: %v\n", destinationValue)

	// Find destination,
	destinationCup := nodeMap[destinationValue]

	destinationCup.Link(removedCups)
	return currentCup.Next()
}

func result(currentCup *ring.Ring) (resultString string) {
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

// Printers

func printCups(r *ring.Ring) {
	if debug {
		r.Do(func(p interface{}) {
			fmt.Printf("%v ", p.(int))
		})
		fmt.Printf("\n")
	}
}

func print(attributes ...interface{}) {
	if debug {
		fmt.Println(attributes...)
	}
}

func printf(printer string, attributes ...interface{}) {
	if debug {
		fmt.Printf(printer, attributes...)
	}
}
