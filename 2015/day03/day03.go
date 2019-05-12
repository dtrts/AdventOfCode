/*
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and
then an elf at the North Pole calls him via radio and tells him where to move
next. Moves are always exactly one house to the north (^), south (v), east (>),
or west (<). After each move, he delivers another present to the house at his
new location.

However, the elf back at the north pole has had a little too much eggnog, and
so his directions are a little off, and Santa ends up visiting some houses more
than once.

How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the
east.
^>v< delivers presents to 4 houses in a square, including twice to the house at
his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2
houses.

--- Part Two ---

The next year, to speed up the process, Santa creates a robot version of
himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the
same starting house), then take turns moving based on instructions from the elf,
who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa
goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back
where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction
and Robo-Santa going the other.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

func coordSort(c [][2]int) [][2]int {
	sort.SliceStable(c, func(i, j int) bool {
		return c[i][0] < c[j][0]
	})
	sort.SliceStable(c, func(i, j int) bool {
		return c[i][0] == c[j][0] && c[i][1] < c[j][1]
	})
	return c
}

func moveCoord(c [2]int, d string) [2]int {
	switch d {
	case "^":
		c[1]++
	case ">":
		c[0]++
	case "v":
		c[1]--
	case "<":
		c[0]--
	}
	return c
}

//assumes slice has been ordered using coordSort
func housesVisited(s [][2]int) int {
	s = coordSort(s)
	housesVisited := 1
	coordTemp := s[0]

	for _, coord := range s[1:] {

		if coord != coordTemp {
			housesVisited++
		}
		coordTemp = coord
	}
	return housesVisited
}

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Input file unable to be read.")
	}
	//part 1
	//generate list of coordinates visited
	currentCoord := [2]int{0, 0}
	visitedCoords := make([][2]int, 0, len(input))
	visitedCoords = append(visitedCoords, currentCoord)

	for _, direction := range input {
		currentCoord = moveCoord(currentCoord, string(direction))
		visitedCoords = append(visitedCoords, currentCoord)
	}

	//part1 answer
	fmt.Println(housesVisited(visitedCoords))

	//part2
	currentCoordSanta := [2]int{0, 0}
	visitedCoordsSanta := make([][2]int, 0, len(input)/2+len(input)%2)
	visitedCoordsSanta = append(visitedCoordsSanta, currentCoordSanta)

	currentCoordRoboSanta := [2]int{0, 0}
	visitedCoordsRoboSanta := make([][2]int, 0, len(input)/2+len(input)%2)
	visitedCoordsRoboSanta = append(visitedCoordsRoboSanta, currentCoordRoboSanta)

	//generate list of coordinates visited by the two present deliverer's
	for i, direction := range input {
		if i%2 == 0 {
			currentCoordSanta = moveCoord(currentCoordSanta, string(direction))
			visitedCoordsSanta = append(visitedCoordsSanta, currentCoordSanta)
		}
		if i%2 == 1 {
			currentCoordRoboSanta = moveCoord(currentCoordRoboSanta, string(direction))
			visitedCoordsRoboSanta = append(visitedCoordsRoboSanta, currentCoordRoboSanta)
		}
	}
	//join list of coordinates visited
	visitedCoords2 := append(visitedCoordsSanta, visitedCoordsRoboSanta...)
	//part2 answer
	fmt.Println(housesVisited(visitedCoords2))
}
