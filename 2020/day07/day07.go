package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// input = inputTest // 32 126

	remove := []*regexp.Regexp{}
	remove = append(remove, regexp.MustCompile(`[.,]`))
	remove = append(remove, regexp.MustCompile(` bags?`))
	remove = append(remove, regexp.MustCompile(` contain( no other)?`))
	for _, re := range remove {
		fmt.Println("Removing:", re)
		input = re.ReplaceAllString(input, "")
		// fmt.Println(input)
	}

	inputLines := strings.Split(input, "\n")

	// fmt.Println(inputLines)

	// Parsing input
	holds := make(map[string]map[string]int)
	heldIn := make(map[string]map[string]struct{})

	for _, ruleLine := range inputLines {

		words := strings.Fields(ruleLine)
		bag := strings.Join(words[:2], " ")
		holds[bag] = make(map[string]int)
		// fmt.Println(bag)

		for i := 2; i < len(words); i += 3 {

			innerBag := strings.Join(words[i+1:i+3], " ")
			amount, err := strconv.Atoi(words[i])
			if err != nil {
				fmt.Println(err)
			}

			holds[bag][innerBag] += amount

			_, ok := heldIn[innerBag]
			if !ok {
				heldIn[innerBag] = make(map[string]struct{})
			}
			heldIn[innerBag][bag] = struct{}{}
		}

	}
	fmt.Println(holds)
	fmt.Println(heldIn)

	// bag A contains:BagB appearsIn:BagC

	fmt.Println(outerBag("shiny gold", heldIn))
	fmt.Println(len(outerBag("shiny gold", heldIn)))

	fmt.Println(numInnerBags("shiny gold", holds, 1) - 1)

}

func numInnerBags(startingBag string, holds map[string]map[string]int, multiplier int) int {

	heldBags := holds[startingBag]
	bags := 1

	if len(heldBags) == 0 {
		return multiplier * bags
	}

	for bag, amount := range heldBags {
		bags += numInnerBags(bag, holds, amount)
	}
	return (bags * multiplier)
}

func outerBag(startingBag string, heldIn map[string]map[string]struct{}) map[string]struct{} {
	output := make(map[string]struct{})

	bagContainedIn := heldIn[startingBag]

	for k, v := range bagContainedIn {
		output[k] = v
	}

	if len(bagContainedIn) == 0 {
		output[startingBag] = struct{}{}

		return output
	}

	for biggerBag := range bagContainedIn {
		topLevel := outerBag(biggerBag, heldIn)

		for k, v := range topLevel {
			output[k] = v
		}
	}

	return output
}
