package main

import "fmt"

var inputTest1 = []int{0, 3, 6} // Given the starting numbers 0,3,6, the 2020th number spoken is 436.
var inputTest2 = []int{1, 3, 2} // Given the starting numbers 1,3,2, the 2020th number spoken is 1.
var inputTest3 = []int{2, 1, 3} // Given the starting numbers 2,1,3, the 2020th number spoken is 10.
var inputTest4 = []int{1, 2, 3} // Given the starting numbers 1,2,3, the 2020th number spoken is 27.
var inputTest5 = []int{2, 3, 1} // Given the starting numbers 2,3,1, the 2020th number spoken is 78.
var inputTest6 = []int{3, 2, 1} // Given the starting numbers 3,2,1, the 2020th number spoken is 438.
var inputTest7 = []int{3, 1, 2} // Given the starting numbers 3,1,2, the 2020th number spoken is 1836

var input = []int{19, 0, 5, 1, 10, 13}

func main() {
	// input = inputTest
	fmt.Println(input)
	origLength := len(input) + 1

	// Create game map
	currGameMap := make(map[int][]int)
	for i, v := range input {
		_, ok := currGameMap[v]
		if !ok {
			currGameMap[v] = []int{}
		}
		currGameMap[v] = append(currGameMap[v], i)
	}
	fmt.Println(currGameMap)

	currGameForMapFunc := make([]int, len(input))
	copy(currGameForMapFunc, input)

	// Run through part one with dumb algo
	for i := 0; i <= 2020-origLength; i++ {
		input = append(input, nextNum(input))
	}
	fmt.Println(input[len(input)-1])

	origLength = len(currGameForMapFunc) + 1

	for i := 0; i <= 30000000-origLength; i++ {

		nextNumFromMap, turnNumFromMap := nextNumMap(currGameForMapFunc, currGameMap)
		currGameForMapFunc = append(currGameForMapFunc, nextNumFromMap)
		currGameMap[nextNumFromMap] = append(currGameMap[nextNumFromMap], turnNumFromMap)

	}
	fmt.Println(currGameForMapFunc[len(currGameForMapFunc)-1])
	// fmt.Println(currGameMap, currGameForMapFunc)
}

// This has an array of how the game has played, along with a map which keeps
// track of which turns each number has been played. The two have to be updated
// together!
// I should really look into passing in pointers to the map and slice and
// affecting them within the function. This would ensure that they are updated
// together.
// ALso the map could easily be map[int][2]int since we only need the last two
// turns played, and the current game only needs to be the total number of turns
// played AND what the last num played was (otherwise you'd have to search the
// map for the biggest turn number which would take too long).
// Maybe i'll come back and optimise later. Or you can do it instead of cheating ;p
func nextNumMap(currGame []int, currGameMap map[int][]int) (nextNum int, turnNum int) {
	turnNum = len(currGame)
	prevTurnNum := len(currGame) - 1
	prevNumSpoken := currGame[prevTurnNum]

	prevTurnNums := currGameMap[prevNumSpoken]

	if len(prevTurnNums) == 1 {
		return 0, turnNum
	}

	return prevTurnNums[len(prevTurnNums)-1] - prevTurnNums[len(prevTurnNums)-2], turnNum
}

func nextNum(currGame []int) int {
	prevTurnNum := len(currGame) - 1
	prevNumSpoken := currGame[prevTurnNum]

	mostRecentBeforePrevTurnNum := prevTurnNum

	// Go through game backwards until out of range or mostRecentBeforePrevTurnNum
	// changes
	for i := len(currGame) - 2; i >= 0 && mostRecentBeforePrevTurnNum == prevTurnNum; i-- {
		if currGame[i] == prevNumSpoken {
			mostRecentBeforePrevTurnNum = i
		}
	}

	// This will return 0 if the value doesn't change
	return prevTurnNum - mostRecentBeforePrevTurnNum

}
