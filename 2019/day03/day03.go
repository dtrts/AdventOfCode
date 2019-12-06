package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(file))
	wires := strings.Split(string(file), "\n")
	wire1 := strings.Split(wires[0], ",")
	wire2 := strings.Split(wires[1], ",")

	directions1 := parseDirections(wire1)
	directions2 := parseDirections(wire2)

	visitedLocations1 := visitedLocations(directions1)
	visitedLocations2 := visitedLocations(directions2)
	fmt.Println(visitedLocations1[0])
	fmt.Println(visitedLocations2[0])
	// fmt.Println(visitedLocations1, visitedLocations2)

	intersections := make([][2]int, 0)
	for _, loc1 := range visitedLocations1 {
		for _, loc2 := range visitedLocations2 {
			if loc1 == loc2 {
				intersections = append(intersections, loc1)
			}
		}

	}

	minDistance := positive(intersections[0][0]) + positive(intersections[0][1])
	for _, v := range intersections {
		if positive(v[0])+positive(v[1]) < minDistance {
			minDistance = positive(v[0]) + positive(v[1])
		}
	}
	fmt.Println("Part 1:", minDistance)

	minDistSum := 2147483647

	for _, intersection := range intersections {
		fmt.Println(intersection, steps(intersection, visitedLocations1), steps(intersection, visitedLocations2))

		dist := steps(intersection, visitedLocations1) + steps(intersection, visitedLocations2)
		fmt.Println(dist)
		if dist < minDistSum {
			minDistSum = dist
		}
	}
	fmt.Println("Part 2:", minDistSum)

}

func steps(intersection [2]int, visitedLocations [][2]int) int {
	for i, loc := range visitedLocations {
		if loc == intersection {
			return i + 1
		}
	}
	return 0
}

func positive(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func visitedLocations(directions []movement) [][2]int {
	visitedLocs := make([][2]int, 0)
	currentLoc := [2]int{0, 0}
	for _, direction := range directions {
		for i := 0; i < direction.distance; i++ {
			switch direction.direction {
			case "U":
				currentLoc[1]++
			case "D":
				currentLoc[1]--
			case "R":
				currentLoc[0]++
			case "L":
				currentLoc[0]--
			}
			visitedLocs = append(visitedLocs, currentLoc)
		}
	}
	return visitedLocs
}

func parseDirections(wire []string) []movement {
	directions := make([]movement, len(wire))

	for i, v := range wire {
		newM := movement{direction: string(v[0])}
		newM.distance, _ = strconv.Atoi(string(v[1:]))
		directions[i] = newM
	}
	return directions

}

type movement struct {
	direction string
	distance  int
}
