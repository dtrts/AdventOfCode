package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// input = inputTest

	inputLines := strings.Split(input, "\n")

	fmt.Println(inputLines)

	instructions := []instruction{}
	for _, line := range inputLines {
		instructions = append(instructions, parseInstruction(line))
	}

	fmt.Println(instructions)

	// Part 1
	route, loop := runProgram(instructions)
	fmt.Println(route, loop)

	// Part 2. Loop through instructions, if jmp or nop then change. then run
	// function to run the program,

	route2, loop2 := []computer{}, true
	for i, instruct := range instructions {

		if instruct.function == "jmp" {
			instructions[i].function = "nop"

			route2, loop2 = runProgram(instructions)

			if loop2 {
				instructions[i].function = "jmp"
			} else {
				fmt.Println("Program finishes!")
				break
			}
		}

		if instruct.function == "nop" {
			instructions[i].function = "jmp"

			route2, loop2 = runProgram(instructions)

			if loop2 {
				instructions[i].function = "nop"
			} else {
				fmt.Println("Program finishes!")
				break
			}
		}

	}

	fmt.Println(route2, loop2)
}

func runProgram(instructions []instruction) ([]computer, bool) {

	current, route, loopFound := computer{}, []computer{}, false

	for loopCheck(route) == false && inSlice(instructions, current) && len(route) <= len(instructions) {
		// fmt.Print(" : ", current)

		current = processInstruction(current, instructions)
		route = append(route, current)

		// fmt.Println("End:", current, route)
	}

	if loopCheck(route) {
		fmt.Println("Program found infinite loop")
		loopFound = true
	}

	if inSlice(instructions, current) == false {
		fmt.Println("Program outside of instructions")
	}
	return route, loopFound

}

func inSlice(instructions []instruction, current computer) bool {
	if current.position >= 0 && current.position < len(instructions) {
		return true
	}
	return false
}

func processInstruction(current computer, instructions []instruction) computer {

	value := instructions[current.position].value

	switch function := instructions[current.position].function; function {

	case "nop":
		current.position++
		return current

	case "acc":
		current.accumulator += value
		current.position++
		return current

	case "jmp":
		current.position += value
		return current

	default:
		current.position++
		return current
	}
	return current

}

func parseInstruction(strInstruction string) instruction {
	output := instruction{}
	tmpFunction, tmpValue := "", ""
	fmt.Sscanf(strInstruction, "%3s %s", &tmpFunction, &tmpValue)

	output.function = tmpFunction
	intValue, err := strconv.Atoi(tmpValue)
	if err != nil {
		fmt.Println(err)
	}
	output.value = intValue
	return output

}

func loopCheck(route []computer) bool {

	if len(route) <= 1 {
		return false
	}

	currentPosition := route[len(route)-1].position
	route = route[:len(route)-1]

	for _, visitedProcess := range route {
		if visitedProcess.position == currentPosition {
			return true
		}
	}
	return false
}

type instruction struct {
	function string
	value    int
}

type computer struct {
	accumulator int
	position    int
}
