/*
--- Day 18: Like a GIF For Your Yard ---
After the million lights incident, the fire code has gotten stricter: now, at
most ten thousand lights are allowed. You arrange them in a 100x100 grid.

Never one to let you down, Santa again mails you instructions on the ideal
lighting configuration. With so few lights, he says, you'll have to resort to
animation.

Start by setting your lights to the included initial configuration (your puzzle
input). A # means "on", and a . means "off".

Then, animate your grid in steps, where each step decides the next configuration
based on the current one. Each light's next state (either on or off) depends on
its current state and the current states of the eight lights adjacent to it
(including diagonals). Lights on the edge of the grid might have fewer than
eight neighbors; the missing ones always count as "off".

For example, in a simplified 6x6 grid, the light marked A has the neighbors
numbered 1 through 8, and the light marked B, which is on an edge, only has the
neighbors marked 1 through 5:

1B5...
234...
......
..123.
..8A4.
..765.
The state a light should have next is based on its current state (on or off)
plus the number of neighbors that are on:

A light which is on stays on when 2 or 3 neighbors are on, and turns off
otherwise.
A light which is off turns on if exactly 3 neighbors are on, and stays off
otherwise.
All of the lights update simultaneously; they all consider the same current
state before moving to the next.

Here's a few steps from an example configuration of another 6x6 grid:

Initial state:
.#.#.#
...##.
#....#
..#...
#.#..#
####..

After 1 step:
..##..
..##.#
...##.
......
#.....
#.##..

After 2 steps:
..###.
......
..###.
......
.#....
.#....

After 3 steps:
...#..
......
...#..
..##..
......
......

After 4 steps:
......
......
..##..
..##..
......
......
After 4 steps, this example has four lights on.

In your grid of 100x100 lights, given your initial configuration, how many
lights are on after 100 steps?

--- Part Two ---
You flip the instructions over; Santa goes on to point out that this is all just
an implementation of Conway's Game of Life. At least, it was, until you notice
that something's wrong with the grid of lights you bought: four lights, one in
each corner, are stuck on and can't be turned off.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("panic")
	}

	inputSplit := strings.Split(string(input), "\n")

	mapHeight := len(inputSplit)   // assumes no new line at end of file
	mapWidth := len(inputSplit[0]) //  assuming all runes are on byte

	fmt.Printf("Height: %d. Width: %d.\n", mapHeight, mapWidth)

	//  make map map with the right level of capacity.
	mapMap := make([][]bool, mapHeight)
	for i := range mapMap {
		mapMap[i] = make([]bool, mapWidth)
	}

	// Populate mapMap with input calues
	for iInput, inputLine := range inputSplit {
		for iInputLine, inputChar := range inputLine {
			if string(inputChar) == "#" {
				mapMap[iInput][iInputLine] = true
			} else {
				mapMap[iInput][iInputLine] = false
			}
		}
	}

	mapMap2 := make([][]bool, mapHeight)
	for i := range mapMap2 {
		mapMap2[i] = make([]bool, mapWidth)
	}
	copy(mapMap2, mapMap)
	// printMap(mapMap2)

	//  Tick Map
	for tick := 0; tick < 100; tick++ {

		mapMap = tickLights(mapMap, false)
		//mapMap = newMapMap

		// fmt.Println("-------------------")

	}
	// printMap(mapMap)

	// Count number of trues
	numLightsOn := 0
	for _, v := range mapMap {
		for _, v2 := range v {
			if v2 {
				numLightsOn++
			}
		}
	}
	fmt.Println("Part 1:", numLightsOn)

	//PART2
	//Run ticks in part 2 mode
	//  Tick Map

	mapMap2[0][0] = true
	mapMap2[len(mapMap2)-1][len(mapMap2[0])-1] = true
	mapMap2[len(mapMap2)-1][0] = true
	mapMap2[0][len(mapMap2[0])-1] = true

	// printMap(mapMap2)

	for tick := 0; tick < 100; tick++ {
		// fmt.Println("tick:", tick, "-------------------")

		mapMap2 = tickLights(mapMap2, true)
		//mapMap = newMapMap
		// printMap(mapMap2)

		// fmt.Println("-------------------")

	}
	printMap(mapMap2)
	fmt.Println("-------------------")

	numLightsOn2 := 0
	for _, v := range mapMap2 {
		for _, v2 := range v {
			if v2 {
				numLightsOn2++
			}
		}
	}
	printMap(mapMap2)

	fmt.Println("Part 2:", numLightsOn2)

}

// Tick mapMap
func tickLights(mapMap [][]bool, part2 bool) (newMapMap [][]bool) {

	// How slices interact is really confusing atm
	newMapMap = make([][]bool, len(mapMap))
	for i := range mapMap {
		newMapMap[i] = make([]bool, len(mapMap[0]))
	}

	// populate newMapMap by going through all lements in current mapMap
	for i, v := range mapMap {
		for i2, v2 := range v {

			neighbours := countSurroundingLights(mapMap, i2, i)

			if v2 == true {
				if neighbours == 2 || neighbours == 3 {
					newMapMap[i][i2] = true
				} else {
					newMapMap[i][i2] = false
				}
			} else if v2 == false {
				if neighbours == 3 {
					newMapMap[i][i2] = true
				} else {
					newMapMap[i][i2] = false
				}
			}

		}
	}

	// for second part make the corners always on
	if part2 {
		newMapMap[0][0] = true
		newMapMap[len(newMapMap)-1][len(newMapMap[0])-1] = true
		newMapMap[len(newMapMap)-1][0] = true
		newMapMap[0][len(newMapMap[0])-1] = true

	}

	return

}

func printMap(mapMap [][]bool) {
	for _, v := range mapMap {
		for _, v2 := range v {
			if v2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	// Print neighbours as each point
	for i, v := range mapMap {
		for i2 := range v {
			neighbours := strconv.Itoa(countSurroundingLights(mapMap, i2, i))
			if neighbours == "0" {
				neighbours = "."
			}
			fmt.Print(neighbours)
		}
		fmt.Print("\n")
	}

}

func countSurroundingLights(mapMap [][]bool, charIndex, lineIndex int) (numLightsOn int) {

	mapWidth := len(mapMap[0])
	mapHeight := len(mapMap)
	if mapWidth <= 1 && mapHeight <= 1 {
		return
	}

	// TOP LEFT CORNER
	if charIndex == 0 && lineIndex == 0 {

		//EAST
		if mapWidth > 1 {
			if mapMap[lineIndex][charIndex+1] {
				numLightsOn++
			}
		}
		//SOUTH
		if mapHeight > 1 {
			if mapMap[lineIndex+1][charIndex] {
				numLightsOn++
			}
		}
		//SOUTH EAST
		if mapWidth > 1 && mapHeight > 1 {
			if mapMap[lineIndex+1][charIndex+1] {
				numLightsOn++
			}
		}
		// TOP RIGHT CORNER --  in else otherwise a 1x1 would do this twice.
	} else if charIndex == mapWidth-1 && lineIndex == 0 {

		//WEST
		if mapWidth > 1 {
			if mapMap[lineIndex][charIndex-1] {
				numLightsOn++
			}
		}
		//SOUTH WEST
		if mapWidth > 1 && mapHeight > 1 {
			if mapMap[lineIndex+1][charIndex-1] {
				numLightsOn++
			}
		}
		//SOUTH
		if mapHeight > 1 {
			if mapMap[lineIndex+1][charIndex] {
				numLightsOn++
			}
		}

		// BOTTOM LEFT CORNER
	} else if charIndex == 0 && lineIndex == mapHeight-1 {

		// NORTH
		if mapHeight > 1 {
			if mapMap[lineIndex-1][charIndex] {
				numLightsOn++
			}
		}
		// NORTH EAST
		if mapWidth > 1 && mapHeight > 1 {
			if mapMap[lineIndex-1][charIndex+1] {
				numLightsOn++
			}
		}
		// EAST
		if mapWidth > 1 {
			if mapMap[lineIndex][charIndex+1] {
				numLightsOn++
			}
		}
		// BOTTOM RIGHT CORNER
	} else if charIndex == mapWidth-1 && lineIndex == mapHeight-1 {
		// NORTH WEST
		if mapMap[lineIndex-1][charIndex-1] {
			numLightsOn++
		}
		// NORTH
		if mapMap[lineIndex-1][charIndex] {
			numLightsOn++
		}
		// WEST
		if mapMap[lineIndex][charIndex-1] {
			numLightsOn++
		}

		// TOP ROW
	} else if lineIndex == 0 {
		//W
		//E
		//SW
		//S
		//SE

		// W
		if mapMap[lineIndex][charIndex-1] {
			numLightsOn++
		}
		// E
		if mapMap[lineIndex][charIndex+1] {
			numLightsOn++
		}
		// SW
		if mapHeight > 1 {
			if mapMap[lineIndex+1][charIndex-1] {
				numLightsOn++
			}
			// S
			if mapMap[lineIndex+1][charIndex] {
				numLightsOn++
			}
			// SE
			if mapMap[lineIndex+1][charIndex+1] {
				numLightsOn++
			}
		}

		//LEFT COLUMN
	} else if charIndex == 0 {

		// N
		if mapMap[lineIndex-1][charIndex] {
			numLightsOn++
		}
		// S
		if mapMap[lineIndex+1][charIndex] {
			numLightsOn++
		}

		if mapWidth > 1 {

			// NE
			if mapMap[lineIndex-1][charIndex+1] {
				numLightsOn++
			}

			// E
			if mapMap[lineIndex][charIndex+1] {
				numLightsOn++
			}
			// SE
			if mapMap[lineIndex+1][charIndex+1] {
				numLightsOn++
			}
		}
		//RIGHT COLUMN
	} else if charIndex == mapWidth-1 {

		// NW
		if mapMap[lineIndex-1][charIndex-1] {
			numLightsOn++
		}
		// N
		if mapMap[lineIndex-1][charIndex] {
			numLightsOn++
		}

		// W
		if mapMap[lineIndex][charIndex-1] {
			numLightsOn++
		}

		// SW
		if mapMap[lineIndex+1][charIndex-1] {
			numLightsOn++
		}
		// S
		if mapMap[lineIndex+1][charIndex] {
			numLightsOn++
		}

		// BOTTOM ROW
	} else if lineIndex == mapHeight-1 {

		// NW
		if mapMap[lineIndex-1][charIndex-1] {
			numLightsOn++
		}
		// N
		if mapMap[lineIndex-1][charIndex] {
			numLightsOn++
		}
		// NE
		if mapMap[lineIndex-1][charIndex+1] {
			numLightsOn++
		}
		// W
		if mapMap[lineIndex][charIndex-1] {
			numLightsOn++
		}
		// E
		if mapMap[lineIndex][charIndex+1] {
			numLightsOn++
		}

		// The rest of them
	} else {
		// NW
		if mapMap[lineIndex-1][charIndex-1] {
			numLightsOn++
		}
		// N
		if mapMap[lineIndex-1][charIndex] {
			numLightsOn++
		}
		// NE
		if mapMap[lineIndex-1][charIndex+1] {
			numLightsOn++
		}
		// W
		if mapMap[lineIndex][charIndex-1] {
			numLightsOn++
		}
		// E
		if mapMap[lineIndex][charIndex+1] {
			numLightsOn++
		}
		// SW
		if mapMap[lineIndex+1][charIndex-1] {
			numLightsOn++
		}
		// S
		if mapMap[lineIndex+1][charIndex] {
			numLightsOn++
		}
		// SE
		if mapMap[lineIndex+1][charIndex+1] {
			numLightsOn++
		}
	}

	// TOP ROW - by definition will have
	if charIndex < mapWidth-1 && lineIndex == 0 {

	}

	return
}
