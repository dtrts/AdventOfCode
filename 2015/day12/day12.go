/*
--- Day 12: JSAbacusFramework.io ---
Santa's Accounting-Elves need help balancing the books after a recent order.
Unfortunately, their accounting software uses a peculiar storage format. That's
where you come in.

They have a JSON document which contains a variety of things: arrays ([1,2,3]),
objects ({"a":1, "b":2}), numbers, and strings. Your first job is to simply find
all of the numbers throughout the document and add them together.

For example:

[1,2,3] and {"a":2,"b":4} both have a sum of 6.
[[[3]]] and {"a":{"b":4},"c":-1} both have a sum of 3.
{"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0.
[] and {} both have a sum of 0.
You will not encounter any strings containing numbers.

What is the sum of all numbers in the document?

--- Part Two ---
Uh oh - the Accounting-Elves have realized that they double-counted everything
red.

Ignore any object (and all of its children) which has any property with the
value "red". Do this only for objects ({...}), not arrays ([...]).

[1,2,3] still has a sum of 6.
[1,{"c":"red","b":2},3] now has a sum of 4, because the middle object is
ignored.
{"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure
is ignored.
[1,"red",5] has a sum of 6, because "red" in an array has no effect.

*/
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calcSum(input []byte) (sum int) {

	numberStringTmp := ""

	for i, v := range string(input) {
		if strings.ContainsAny(string(v), "-0123456789") {
			numberStringTmp += string(v)

			if i == len(string(input))-1 {
				newVal, err := strconv.Atoi(numberStringTmp)
				if err != nil {
					panic("String conversion doesn't work")
				}
				sum += newVal
			}

		} else {

			if len(numberStringTmp) > 0 {
				newVal, err := strconv.Atoi(numberStringTmp)
				if err != nil {
					panic("String conversion doesn't work")
				}
				sum += newVal
			}

			numberStringTmp = ""

		}
	}
	return sum
}

func main() {
	fmt.Println("Hello world!")

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("File unable to be read.")
	}

	//part1
	fmt.Println(calcSum(input))

	//part2
	// j, err := json.Marshal(input)
	// if err != nil {
	// 	panic("Unable to marshal")
	// }
	// fmt.Println(string(j))
	//Remove all objects with red
	for strings.Index(string(input), ":\"red\"") > 0 {
		redElementIndex := strings.Index(string(input), ":\"red\"")
		//find closing bracket
		closingBracketIndex := redElementIndex
		depth := 1
		for depth > 0 {
			if string(input[closingBracketIndex]) == "}" {
				depth--
				if depth == 0 {
					break
				}
			}
			if string(input[closingBracketIndex]) == "{" {
				depth++
			}
			closingBracketIndex++

		}
		//find opening bracket
		openingBracketIndex := redElementIndex
		depth = 1
		for depth > 0 {
			if string(input[openingBracketIndex]) == "{" {
				depth--
				if depth == 0 {
					break
				}
			}
			if string(input[openingBracketIndex]) == "}" {
				depth++
			}
			openingBracketIndex--

		}
		input = append(input[:openingBracketIndex], input[closingBracketIndex+1:]...)
	}

	fmt.Println(calcSum(input))

}
