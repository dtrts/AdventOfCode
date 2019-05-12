/*
--- Day 13: Knights of the Dinner Table ---
In years past, the holiday feast with your family hasn't gone so well. Not
everyone gets along! This year, you resolve, will be different. You're going to
find the optimal seating arrangement and avoid all those awkward conversations.

You start by writing up a list of everyone invited and the amount their
happiness would increase or decrease if they were to find themselves sitting
next to each other person. You have a circular table that will be just big
enough to fit everyone comfortably, and so each person will have exactly two
neighbors.

For example, suppose you have only four attendees planned, and you calculate
their potential happiness as follows:

Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
Then, if you seat Alice next to David, Alice would lose 2 happiness units
(because David talks so much), but David would gain 46 happiness units (because
Alice is such a good listener), for a total change of 44.

If you continue around the table, you could then seat Bob next to Alice (Bob
gains 83, Alice gains 54). Finally, seat Carol, who sits next to Bob (Carol
gains 60, Bob loses 7) and David (Carol gains 55, David gains 41). The
arrangement looks like this:

     +41 +46
+55   David    -2
Carol       Alice
+60    Bob    +54
     -7  +83
After trying every other seating arrangement in this hypothetical scenario, you
find that this one is the most optimal, with a total change in happiness of 330.

What is the total change in happiness for the optimal seating arrangement of the
actual guest list?

--- Part Two ---
In all the commotion, you realize that you forgot to seat yourself. At this
point, you're pretty apathetic toward the whole thing, and your happiness
wouldn't really go up or down regardless of who you sit next to. You assume
everyone else would be just as ambivalent about sitting next to you, too.

So, add yourself to the list, and give all happiness relationships that involve
you a score of 0.

What is the total change in happiness for the optimal seating arrangement that
actually includes yourself?

*/
package main

import (
	// "AdventOfCode/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

type pairing struct {
	guest0    string
	guest1    string
	happiness int
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("file unable to be read")
	}

	inputSplit := strings.Split(string(input), "\n")

	allPairings, guests := []pairing{}, []string{}

	guest0, affect, guest1, happiness := "", "", "", 0

	for _, v := range inputSplit {
		fmt.Sscanf(v, "%s would %s %d happiness units by sitting next to %s.", &guest0, &affect, &happiness, &guest1)
		if affect == "lose" {
			happiness = happiness * -1
		}
		guest1 = strings.TrimSuffix(guest1, ".")
		allPairings = append(allPairings, pairing{guest0, guest1, happiness})

		//list all guests
		if len(guests) == 0 {
			guests = append(guests, guest0)
		}
		//append missing new guests
		guestAppeared0 := false
		guestAppeared1 := false
		for _, guestTemp := range guests {
			if guestTemp == guest0 {
				guestAppeared0 = true
			}
			if guestTemp == guest1 {
				guestAppeared1 = true
			}
		}
		if guestAppeared0 == false {
			guests = append(guests, guest0)
		}
		if guestAppeared1 == false {
			guests = append(guests, guest1)
		}

	}

	// fmt.Println(allPairings)
	// fmt.Println(guests)

	maxHappiness := calcMaxHappiness(guests, allPairings)
	//answer to pt1
	fmt.Println(maxHappiness)

	//add self to guests
	guests = append(guests, "me")

	maxHappiness = calcMaxHappiness(guests, allPairings)
	//answer to pt1
	fmt.Println(maxHappiness)

}

func calcMaxHappiness(guests []string, allPairings []pairing) int {
	arrangements := stringSlicePermutate(guests)
	maxHappiness, minHappiness := 0, 0
	totalHappiness := 0
	for x, arrangement := range arrangements {
		// fmt.Println(arrangement, "----------")
		// get happiness
		totalHappiness = 0
		for i := range arrangement {

			indexPairGuest0 := i - 1
			if indexPairGuest0 < 0 {
				indexPairGuest0 = len(arrangement) - 1
			}
			indexPairGuest1 := i + 1
			if indexPairGuest1 >= len(arrangement) {
				indexPairGuest1 = 0
			}
			pairGuest0 := arrangement[indexPairGuest0]
			pairGuest1 := arrangement[indexPairGuest1]

			//need to get distance between 0 and current and current and 1
			// fmt.Print(pairGuest0, arrangement[i], pairGuest1)
			for i2 := range allPairings {
				if allPairings[i2].guest0 == arrangement[i] && allPairings[i2].guest1 == pairGuest0 {
					// fmt.Print(allPairings[i2])
					totalHappiness += allPairings[i2].happiness
				}

				// if allPairings[i2].guest1 == arrangement[i] && allPairings[i2].guest0 == pairGuest0 {
				// 	fmt.Print(allPairings[i2])
				// 	totalHappiness += allPairings[i2].happiness
				// }

				if allPairings[i2].guest0 == arrangement[i] && allPairings[i2].guest1 == pairGuest1 {
					// fmt.Print(allPairings[i2])
					totalHappiness += allPairings[i2].happiness
				}
				// if allPairings[i2].guest1 == arrangement[i] && allPairings[i2].guest0 == pairGuest1 {
				// 	fmt.Print(allPairings[i2])
				// 	totalHappiness += allPairings[i2].happiness
				// }
			}

			// fmt.Print("\n")

		}
		if x == 0 {
			maxHappiness = totalHappiness
			minHappiness = totalHappiness
		} else {
			if totalHappiness > maxHappiness {
				maxHappiness = totalHappiness
			}
			if totalHappiness < minHappiness {
				minHappiness = totalHappiness
			}
		}

		// fmt.Println(arrangement, "----------", totalHappiness)

	}

	return maxHappiness
}

//StringSlicePermutate will return a slice of all possible permutations of the
//original slice. Does not check for duplicates
func stringSlicePermutate(s []string) (permutations [][]string) {

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
			for _, v2 := range stringSlicePermutate(removeElementFromSlice(s, i)) {
				//permutationsElement := append([][]string{}, append([]string{v}, v2...))
				permutations = append(permutations, append([][]string{}, append([]string{v}, v2...))...)
			}
		}
		return

	}
}

//RemoveElementFromSlice removes the element from a string slice.
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
