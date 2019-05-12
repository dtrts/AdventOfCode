/*
--- Day 17: No Such Thing as Too Much ---
The elves bought too much eggnog again - 150 liters this time. To fit it all
into your refrigerator, you'll need to move it into smaller containers. You take
an inventory of the capacities of the available containers.

For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If
you need to store 25 liters, there are four ways to do it:

15 and 10
20 and 5 (the first 5)
20 and 5 (the second 5)
15, 5, and 5
Filling all containers entirely, how many different combinations of containers
can exactly fit all 150 liters of eggnog?

Puzzle Input
50
44
11
49
42
46
18
32
26
40
21
7
18
43
10
47
36
24
22
40

--- Part Two ---
While playing with all the containers in the kitchen, another load of eggnog
arrives! The shipping and receiving department is requesting as many containers
as you can spare.

Find the minimum number of containers that can exactly fit all 150 liters of
eggnog. How many different ways can you fill that number of containers and still
hold exactly 150 litres?

In the example above, the minimum number of containers was two. There were three
ways to use that many containers, and so the answer there would be 3.


*/

package main

import (
	"fmt"
)

// All combinations of all lengths.
// Need a
// loop through all emlements. pass through the reamining slice AFTER that element

func genContainerSets(containers []int, maxVal int, currentContainerSet []int) (validContainerSets [][]int) {
	// Calc num of litres currently held
	currentCapacity := 0
	for _, currentContainer := range currentContainerSet {
		currentCapacity += currentContainer
	}

	// For each element check if it is sufficient to be returned, otherwise recursively act on remaining slice
	for i, container := range containers {

		remainingContainers := []int{}
		if i < len(containers)-1 {
			remainingContainers = append([]int{}, containers[i+1:]...)
		}

		newCurrentContainerSet := append(currentContainerSet, container)
		if currentCapacity+container == maxVal {
			validContainerSets = append(validContainerSets, newCurrentContainerSet)
		}

		if currentCapacity+container <= maxVal {
			validContainerSets = append(validContainerSets, genContainerSets(remainingContainers, maxVal, newCurrentContainerSet)...)
		}

	}

	return

}

func main() {
	fmt.Println("Hello World!")

	totalLitres := 150
	containers := []int{50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40}
	// containers := []int{20, 15, 10, 5, 5}
	// fmt.Println(containers)
	// fmt.Println(removeFromSlice(containers, 0))
	// fmt.Println(removeFromSlice(containers, 1))
	// fmt.Println(removeFromSlice(containers, len(containers)-1))

	allContainerSets := genContainerSets(containers, totalLitres, []int{})

	// fmt.Println(allContainerSets)
	fmt.Println("Part 1:", len(allContainerSets))

	//Get min length
	lenFewestContainerSet := 0
	for i, containerSet := range allContainerSets {
		if i == 0 {
			lenFewestContainerSet = len(containerSet)
		}

		if len(containerSet) < lenFewestContainerSet {
			lenFewestContainerSet = len(containerSet)
		}
	}

	// Get num containerSet with that length
	numSimilarSet := 0
	for _, containerSet := range allContainerSets {
		if len(containerSet) == lenFewestContainerSet {
			numSimilarSet++
		}
	}

	//Part 2 ans
	fmt.Println("Len shortest container set", lenFewestContainerSet)
	fmt.Println("Part 2:", numSimilarSet)
}
