package main

import (
	"fmt"
	"strings"
)

func main() {
	// input = inputTest
	//Vars
	seatingArrangement := [][]string{} // So seatingArrangement[row][column]

	// Parse input
	for _, line := range strings.Split(input, "\n") {
		seatingArrangement = append(seatingArrangement, strings.Split(line, ""))
	}

	//Print initial state
	fmt.Println(seatingArrangement, seatingArrangementCounter(seatingArrangement, true), numOccupiedSeats(seatingArrangement))

	// Update the seating arangement until the number of occupied seats stabalises
	// (Potential for a false positives where seats change but number occupied
	// stays the same?)
	occupiedSeats, tick := numOccupiedSeats(seatingArrangement), 0
	for tick < 10000 {
		// Update the seating arrangement using the counter map
		update(seatingArrangement, seatingArrangementCounter(seatingArrangement, true), 4)

		newOccupiedSeats := numOccupiedSeats(seatingArrangement)
		if occupiedSeats == newOccupiedSeats {
			fmt.Println("Num Occupied Seat stabalized")
			break
		}

		occupiedSeats = newOccupiedSeats
		tick++
	}

	// Print out final arrangement and number of occupied seats
	fmt.Println(seatingArrangement, seatingArrangementCounter(seatingArrangement, true), numOccupiedSeats(seatingArrangement))

	//Part 2
	fmt.Println("Part2")
	seatingArrangement = nil // So seatingArrangement[row][column]

	// Parse input
	for _, line := range strings.Split(input, "\n") {
		seatingArrangement = append(seatingArrangement, strings.Split(line, ""))
	}

	occupiedSeats, tick = numOccupiedSeats(seatingArrangement), 0
	for tick < 10000 {
		update(seatingArrangement, seatingArrangementCounter(seatingArrangement, false), 5)

		newOccupiedSeats := numOccupiedSeats(seatingArrangement)
		if occupiedSeats == newOccupiedSeats {
			fmt.Println("Num Occupied Seat stabalized")
			break
		}
		occupiedSeats = newOccupiedSeats
		tick++
	}
	fmt.Println(seatingArrangement, seatingArrangementCounter(seatingArrangement, false), numOccupiedSeats(seatingArrangement))

}

func update(seatingArrangement [][]string, seatingArrangementCount [][]int, fussyness int) {
	for row, rowLine := range seatingArrangementCount {
		for column, numSurroundingSeats := range rowLine {
			seat := seatingArrangement[row][column]
			if seat == "L" && numSurroundingSeats == 0 {
				seatingArrangement[row][column] = "#"
			} else if seat == "#" && numSurroundingSeats >= fussyness {
				seatingArrangement[row][column] = "L"
			}
		}
	}
	return
}

// Produces a corressonding arrangement to the seats, showing how many
// surrounding seats each position has. -1 is used for the floor
func seatingArrangementCounter(seatingArrangement [][]string, adjacent bool) (seatingArrangementCount [][]int) {
	// Go through all positions and find the number of occupied seats next door.
	for row, rowLine := range seatingArrangement {
		rowArrangementCount := []int{}

		for column, seat := range rowLine {

			if seat == "." {
				rowArrangementCount = append(rowArrangementCount, -1)
			} else {
				// Theres two methods of counting surrounding seats
				if adjacent {
					rowArrangementCount = append(rowArrangementCount, numSurroundingSeats3(seatingArrangement, row, column))
				} else {
					rowArrangementCount = append(rowArrangementCount, numSurroundingSeats2(seatingArrangement, row, column))
				}
			}

		}

		seatingArrangementCount = append(seatingArrangementCount, rowArrangementCount)

	}
	return seatingArrangementCount
}

// Count all occupied seats in an arrangment
func numOccupiedSeats(seatingArrangement [][]string) (num int) {
	for _, v := range seatingArrangement {
		for _, v2 := range v {
			if v2 == "#" {
				num++
			}
		}
	}
	return num
}

// Check if a 2d coordinate is valid based on max indexes
func inPlane(position []int, maxRow int, maxCol int) bool {
	if position[0] >= 0 && position[0] <= maxRow && position[1] >= 0 && position[1] <= maxCol {
		return true
	}
	return false
}

// Multiply a coordinate by a scalar.
func multCoord(coord []int, mult int) []int {
	newCoord := []int{}
	for _, v := range coord {
		newCoord = append(newCoord, v*mult)
	}
	return newCoord
}

// Count occupied seats in all 8 directions, skipping over floors
func numSurroundingSeats2(seatingArrangement [][]string, row int, column int) (num int) {
	maxRow := len(seatingArrangement) - 1
	maxCol := len(seatingArrangement[0]) - 1

	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

	// for each direction
	// increasing distance
	// check if in plane otherwise break
	// check if a seat count and break

	for _, direction := range directions {
		// fmt.Println("Direction", direction)

		for distance := 1; distance <= maxCol || distance <= maxRow; distance++ {

			deltaCoord := multCoord(direction, distance)
			newPosition := []int{row + deltaCoord[0], column + deltaCoord[1]}
			// fmt.Println("Distance", distance, deltaCoord, newPosition)

			if inPlane(newPosition, maxRow, maxCol) {
				seatCheck := seatingArrangement[newPosition[0]][newPosition[1]]
				// fmt.Println("in place", seatCheck)
				if seatCheck == "#" {
					num++
					// fmt.Println("Found occupied seat", num)
					break
				} else if seatCheck == "L" {
					// fmt.Println("Found empty seat")
					break
				}
			} else {
				// fmt.Println("No longer in plane")

				break
			}
		}
	}
	// fmt.Println("numSurroundingSeats2()", row, column, num)

	return num
}

func numSurroundingSeats3(seatingArrangement [][]string, row int, column int) (num int) {
	maxRow := len(seatingArrangement) - 1
	maxCol := len(seatingArrangement[0]) - 1

	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

	for _, direction := range directions {

		newPosition := []int{row + direction[0], column + direction[1]}

		if inPlane(newPosition, maxRow, maxCol) {
			seatCheck := seatingArrangement[newPosition[0]][newPosition[1]]
			if seatCheck == "#" {
				num++
			}
		}

	}

	return num
}
