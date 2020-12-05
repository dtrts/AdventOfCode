package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	// input = inputTest
	inputLines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	maxSeatID := 0

	seatChart := [128][8]string{}

	for _, seatCode := range inputLines {
		seatRow := binarySearch(seatCode[:7])
		seatColumn := binarySearch(seatCode[7:])

		fmt.Println(seatCode, seatRow, seatColumn, seatID(seatCode))

		if seatID(seatCode) > maxSeatID {
			maxSeatID = seatID(seatCode)
		}

		seatChart[seatRow][seatColumn] = "X"
	}

	fmt.Println(maxSeatID)

	printSeatChart(seatChart)
	// fmt.Println(seatChart)

	inFront := true
	mySeatRow := 0
	mySeatColumn := 0
	for i, v := range seatChart {

		for i2, v2 := range v {
			if inFront {
				if v2 == "X" {
					inFront = false
				}
			} else {
				if v2 != "X" {
					mySeatRow = i
					mySeatColumn = i2
					break
				}
			}
		}

		if mySeatRow > 0 || mySeatColumn > 0 {
			break
		}

	}

	fmt.Println(mySeatRow, mySeatColumn, (mySeatRow*8)+mySeatColumn)

}

func printSeatChart(seatChart [128][8]string) {
	for i, v := range seatChart {
		fmt.Print(v, " ", i, "\n")
	}
}

func binarySearch(seatCode string) int {
	minSeat := 0
	maxSeat := int(math.Pow(2, float64(len(seatCode)))) - 1
	// fmt.Println("Start:", minSeat, maxSeat, maxSeat-minSeat, (maxSeat-minSeat)/2)
	// fmt.Print("Start:", minSeat, maxSeat)

	for _, direction := range seatCode {

		if string(direction) == "F" || string(direction) == "L" {
			maxSeat = maxSeat - (maxSeat-minSeat+1)/2
		}
		if string(direction) == "B" || string(direction) == "R" {
			minSeat = minSeat + ((maxSeat - minSeat + 1) / 2)
		}

		// fmt.Println("Loop ", i, string(direction), minSeat, maxSeat, maxSeat-minSeat, (maxSeat-minSeat)/2)
	}
	// fmt.Println(" ", maxSeat)

	return maxSeat
}

func seatID(seatCode string) int {
	return (binarySearch(seatCode[:7]) * 8) + (binarySearch(seatCode[7:]))
}
