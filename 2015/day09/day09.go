/*
--- Day 9: All in a Single Night ---
Every year, Santa manages to deliver all of his presents in a single night.

This year, however, he has some new locations to visit; his elves have provided
him the distances between every pair of locations. He can start and end at any
two (different) locations he wants, but he must visit each location exactly
once. What is the shortest distance he can travel to achieve this?

For example, given the following distances:

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
The possible routes are therefore:

Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982
The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is
605 in this example.

What is the distance of the shortest route?

--- Part Two ---
The next year, just to show off, Santa decides to take the route with the
longest distance instead.

He can still start and end at any two (different) locations he wants, and he
still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for
example) Dublin -> London -> Belfast.

What is the distance of the longest route?
*/
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cityDistance struct {
	city0    string
	city1    string
	distance int
}

//Assigning a variable to a slice range gives the underlying slice, it DOES NOT make a copy!!!!
func removeElementFromSlice(sliceIncoming []string, i int) (alteredSlice []string) {
	// assuming len(s) > 0 and i < len(s)
	if i == 0 {

		alteredSlice = append(alteredSlice, sliceIncoming[1:]...)
		return alteredSlice

	} else if i == len(sliceIncoming)-1 {

		alteredSlice = append(alteredSlice, sliceIncoming[:len(sliceIncoming)-1]...)
		return alteredSlice

	} else {

		alteredSlice = append(append(alteredSlice, sliceIncoming[:i]...), sliceIncoming[i+1:]...)
		return alteredSlice

	}
}

//Recursive call
func permutate(s []string) (permutations [][]string) {

	if len(s) == 0 {

		return

	} else if len(s) == 1 {

		permutations = [][]string{[]string{s[0]}}
		return

	} else if len(s) == 2 {

		permutations = [][]string{[]string{s[0], s[1]}, []string{s[1], s[0]}} //  two permutations
		return

	} else {

		for i, v := range s {
			// append onto permutations each element and the permutations it has
			// append sub permutations onto element
			for _, v2 := range permutate(removeElementFromSlice(s, i)) {
				//permutationsElement := append([][]string{}, append([]string{v}, v2...))
				permutations = append(permutations, append([][]string{}, append([]string{v}, v2...))...)
			}
		}
		return

	}
}

func main() {

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(strings.Repeat("-", 80))
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("file unable to be read")
	}

	inputSplit := strings.Split(string(input), "\n")

	//Generate slice of distances and cities
	cityDistances := []cityDistance{}
	cities := []string{}

	//Get slice of distances and slice of all cities
	var city0, city1 string
	distance := 0

	for _, v := range inputSplit {
		fmt.Sscanf(v, "%s to %s = %d", &city0, &city1, &distance)
		cityDistances = append(cityDistances, cityDistance{city0, city1, distance})

		//list all cities
		if len(cities) == 0 {
			cities = append(cities, city0)
		}
		//append missing new cities
		cityAppeared0 := false
		cityAppeared1 := false
		for _, cityTmp := range cities {
			if cityTmp == city0 {
				cityAppeared0 = true
			}
			if cityTmp == city1 {
				cityAppeared1 = true
			}
		}
		if cityAppeared0 == false {
			cities = append(cities, city0)
		}
		if cityAppeared1 == false {
			cities = append(cities, city1)
		}

	}

	fmt.Println(len(permutate(cities)))
	fmt.Println(cityDistances)
	fmt.Println(strings.Repeat("-", 80))

	//permutate city list and find distance of cityPath
	cityPermutations := permutate(cities)
	minDistance := 0
	maxDistance := 0

	for i, cityPath := range cityPermutations {

		pathDistance := 0
		//Get distance of path
		for i := 1; i < len(cityPath); i++ {
			cityStart := cityPath[i-1]
			cityFinish := cityPath[i]

			//Get distance for city pair
			cityDistanceThis := 0
			for _, cityDistanceRetreval := range cityDistances {
				if (cityStart == cityDistanceRetreval.city0 || cityStart == cityDistanceRetreval.city1) && (cityFinish == cityDistanceRetreval.city0 || cityFinish == cityDistanceRetreval.city1) {
					cityDistanceThis = cityDistanceRetreval.distance
				}
			}

			pathDistance += cityDistanceThis

		}

		//Decide if new min/max
		if i == 0 {
			minDistance = pathDistance
			maxDistance = pathDistance
		} else {
			if pathDistance < minDistance {
				minDistance = pathDistance
			}
			if pathDistance > maxDistance {
				maxDistance = pathDistance
			}
		}

		//fmt.Println(minDistance, maxDistance, pathDistance, cityPath)

	}

	//Print answers to both parts
	fmt.Println(minDistance, maxDistance)

}
