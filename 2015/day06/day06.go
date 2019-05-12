/*
--- Day 6: Probably a Fire Hazard ---
Because your neighbors keep defeating you in the holiday house decorating
contest year after year, you've decided to deploy one million lights in a
1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you
instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at
each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include
whether to turn on, turn off, or toggle various inclusive ranges given as
coordinate pairs. Each coordinate pair represents opposite corners of a
rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to
 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by
doing the instructions Santa sent you in order.

For example:

turn on 0,0 through 999,999 would turn on (or leave on) every light.
toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off
the ones that were on, and turning on the ones that were off.
turn off 499,499 through 500,500 would turn off (or leave off) the middle four
lights.
After following the instructions, how many lights are lit?

--- Part Two ---
You just finish implementing your winning light pattern when you realize you
mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each
light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of
those lights by 1.

The phrase turn off actually means that you should decrease the brightness of
those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of
those lights by 2.

What is the total brightness of all lights combined after following Santa's
instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.

*/
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instruction struct {
	rectStart    [2]int
	rectFin      [2]int
	instructType string
}

func main() {
	//import file
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("File unable to be read")
	}

	// Split instructions into structs
	inputStrings := strings.Split(string(input), "\n")
	//fmt.Println(inputStrings)

	instructType := ""
	rectStart0 := 0
	rectStart1 := 0
	rectEnd0 := 0
	rectEnd1 := 0
	//instructions := make([]instruction, 0, len(inputStrings))
	var instructions []instruction
	prefixes := [3]string{"turn on", "turn off", "toggle"}
	for _, s := range inputStrings {
		//process instruction type
		for _, prefix := range prefixes {

			if strings.HasPrefix(s, prefix) {
				instructType = strings.TrimSpace(prefix)
				s = strings.TrimSpace(strings.TrimPrefix(s, prefix))
				break
			}

		}

		fmt.Sscanf(s, "%d,%d through %d,%d", &rectStart0, &rectStart1, &rectEnd0, &rectEnd1)
		instructions = append(instructions, instruction{[2]int{rectStart0, rectStart1}, [2]int{rectEnd0, rectEnd1}, instructType})

	}

	//fmt.Println(instructions)

	//prt1
	//create a slice of 1000x1000 bools to hold all the lights.
	var allOfTheLights [1e6]bool

	for _, instruct := range instructions {
		for i := instruct.rectStart[0]; i <= instruct.rectFin[0]; i++ {
			for i2 := instruct.rectStart[1]; i2 <= instruct.rectFin[1]; i2++ {
				index := (1000 * i) + i2
				switch instruct.instructType {
				case "turn on":
					allOfTheLights[index] = true
				case "turn off":
					allOfTheLights[index] = false
				case "toggle":
					allOfTheLights[index] = !allOfTheLights[index]
				}
			}
		}
	}

	numLightsOn := 0
	for _, v := range allOfTheLights {
		if v {
			numLightsOn++
		}
	}
	fmt.Println(numLightsOn)

	//part2
	var allOfTheBrights [1e6]int

	for _, instruct := range instructions {
		for i := instruct.rectStart[0]; i <= instruct.rectFin[0]; i++ {
			for i2 := instruct.rectStart[1]; i2 <= instruct.rectFin[1]; i2++ {
				index := (1000 * i) + i2

				switch instruct.instructType {
				case "turn on":
					allOfTheBrights[index]++
				case "toggle":
					allOfTheBrights[index] += 2
				case "turn off":
					if allOfTheBrights[index] > 0 {
						allOfTheBrights[index]--
					}
				}

			}
		}
	}

	totalBrightness := 0
	for _, v := range allOfTheBrights {
		totalBrightness += v
	}
	fmt.Println(totalBrightness)
}
