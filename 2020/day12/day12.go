package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input = inputTest

	inputLines := strings.Split(input, "\n")

	facing := "E" //[1,0]
	position := [2]int{0, 0}
	fmt.Println(facing, position, inputLines)

	for _, line := range inputLines {

		action := string(line[:1])
		amount, _ := strconv.Atoi(line[1:])

		// fmt.Sscanf(line, "%s%d", &action, &amount)
		fmt.Println(position, facing, action, amount)

		if action == "R" || action == "L" {
			facing = rotate(facing, action, amount)
		} else if action == "F" {
			position = forward(position, facing, amount)
		} else {
			position = move(position, action, amount)
		}

		fmt.Println(position, facing)
	}

	fmt.Println("Part1", manDist(position))

	position = [2]int{0, 0}
	waypoint := [2]int{10, 1}

	for _, line := range inputLines {

		action := string(line[:1])
		amount, _ := strconv.Atoi(line[1:])

		// fmt.Sscanf(line, "%s%d", &action, &amount)
		fmt.Println(position, waypoint, action, amount)
		if action == "R" || action == "L" {

			waypoint = rotateWaypoint(waypoint, action, amount)
		} else if action == "F" {
			position[0] += waypoint[0] * amount
			position[1] += waypoint[1] * amount
		} else {
			waypoint = move(waypoint, action, amount)
		}

		fmt.Println(position, waypoint)

	}
	fmt.Println("Part2", manDist(position))

}

func manDist(position [2]int) int {
	for i, v := range position {
		if v < 0 {
			position[i] *= -1
		}
	}

	return position[0] + position[1]

}

func move(position [2]int, action string, amount int) [2]int {

	switch action {
	case "E":
		position[0] += amount
	case "W":
		position[0] -= amount
	case "N":
		position[1] += amount
	case "S":
		position[1] -= amount
	}

	return position
}

func forward(position [2]int, facing string, amount int) [2]int {

	// fmt.Println("FORWARD", position, facing, amount)

	switch facing {
	case "E":
		position[0] += amount
	case "W":
		position[0] -= amount
	case "N":
		position[1] += amount
	case "S":
		position[1] -= amount
	}

	// fmt.Println("FORWARDEND", position, facing, amount)

	return position
}

func rotate(facing string, action string, amount int) string {
	times := amount / 90
	dirs := []string{"N", "E", "S", "W"}
	currIndex := 0

	for i, dir := range dirs {
		if facing == dir {
			currIndex = i
		}
	}

	// fmt.Println("currIndex", currIndex, currIndex%4, "Times", times, amount)

	if action == "R" {
		currIndex += (times)
	} else {
		currIndex += (times * 3) // Three rights make a left, one left made an out of array exepction panic
	}
	return dirs[currIndex%4]
}

func rotateWaypoint(waypoint [2]int, action string, amount int) [2]int {
	times := amount / 90
	newWay := [2]int{0, 0}
	for time := 1; time <= times; time++ {
		fmt.Println("Ratotaing", times, time, action)
		if action == "R" {
			newWay[0] = waypoint[1]
			newWay[1] = -waypoint[0]
		} else {
			newWay[0] = -waypoint[1]
			newWay[1] = waypoint[0]
		}
		waypoint[0] = newWay[0]
		waypoint[1] = newWay[1]
	}
	return newWay
}
