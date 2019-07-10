package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("file not read")
	}

	currentFloor := 0

	//  to flag the first basement appearance
	event := false

	upOne := byte('(')
	downOne := byte(')')

	for i, instruction := range input {

		switch instruction {
		case upOne:
			currentFloor++
		case downOne:
			currentFloor--
		default:
			panic("unexpected shits")
		}

		if currentFloor < 0 && event == false {
			fmt.Printf("The instruction which causes the current floor to be -1 is the %dth instruction\n", i+1) //  answer to part 2
			event = true
		}
	}

	fmt.Printf("The final floor is %d\n", currentFloor) //  answer to part one

}
