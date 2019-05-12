/*
--- Day 7: Some Assembly Required ---
This year, Santa brought little Bobby Tables a set of wires and bitwise logic
gates! Unfortunately, little Bobby is a little under the recommended age range,
and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit
signal (a number from 0 to 65535). A signal is provided to each wire by a gate,
another wire, or some specific value. Each wire can only get a signal from one
source, but can provide its signal to multiple destinations. A gate provides no
signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together:
x AND y -> z means to connect wires x and y to an AND gate, and then connect its
 output to wire z.

For example:

123 -> x means that the signal 123 is provided to wire x.
x AND y -> z means that the bitwise AND of wire x and wire y is provided to
wire z.
p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then
provided to wire q.
NOT e -> f means that the bitwise complement of the value from wire e is
provided to wire f.
Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for
some reason, you'd like to emulate the circuit instead, almost all programming
languages (for example, C, JavaScript, or Python) provide operators for these
gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
In little Bobby's kit's instructions booklet (provided as your puzzle input),
what signal is ultimately provided to wire a?

--- Part Two ---
Now, take the signal you got on wire a, override wire b to that signal, and
reset the other wires (including wire a). What new signal is ultimately provided
to wire a?


*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//  dependancies, instruction ,

type wireAndInstruction struct {
	wire string
	//instruction           string
	instructionType   string
	instructionInputs []string
	value             int
	processed         bool
}

//will return type, references wires (including values?) value if it has one and
//processed true if it is just a value in.
func processInstruction(instruction string) (instructionType string, instructionInputs []string, value int, processed bool) {
	instructionInputs0, instructionInputs1 := "", ""
	//NOT AND OR RSHIFT LSHIFT
	if strings.Contains(instruction, "NOT") {

		fmt.Sscanf(instruction, "%s %s", &instructionType, &instructionInputs0)
		instructionInputs = append(instructionInputs, instructionInputs0)

		return
	} else if strings.Contains(instruction, "AND") || strings.Contains(instruction, "OR") || strings.Contains(instruction, "RSHIFT") || strings.Contains(instruction, "LSHIFT") {

		fmt.Sscanf(instruction, "%s %s %s", &instructionInputs0, &instructionType, &instructionInputs1)
		instructionInputs = append(instructionInputs, instructionInputs0, instructionInputs1)
		return
	} else {
		instructionType = "DIRECT"
		instructionInputs = append(instructionInputs, instruction)
		if i, err := strconv.Atoi(instructionInputs[0]); err == nil {
			value = i
			processed = true
		}
		return
	}

	panic("AAAHHH")
	return
}

//Done badly as shoudl really pass in a pointer to the set?
func processInstructions(wireAndInstructions []wireAndInstruction) ([]wireAndInstruction, int) {
	updatesMade := 0
	for i := range wireAndInstructions {
		//if not processed loop through and
		if wireAndInstructions[i].processed == false {
			// check if this set can be processed and collect full set of inputs on way in
			canBeProcessed := true
			instructionInputsValues := []int{} //  this will be filled in the correct order

			for _, processingInputs := range wireAndInstructions[i].instructionInputs {
				//two behavious on if input is int or a wire name
				if processingInputInt, err := strconv.Atoi(processingInputs); err == nil {
					instructionInputsValues = append(instructionInputsValues, processingInputInt)
				} else {
					//loop though wires and see if the wire has been processed.
					//retrieve value if possible (ie. wire has been processed)
					//set as unable to be processed if not
					for i2 := range wireAndInstructions {
						if canBeProcessed == true && processingInputs == wireAndInstructions[i2].wire {
							//Check if retrieved wire has been processed
							if wireAndInstructions[i2].processed == true {
								instructionInputsValues = append(instructionInputsValues, wireAndInstructions[i2].value)
							} else {
								canBeProcessed = false
							}
						}
					}
				}
			}

			if canBeProcessed == true {
				// should now know if wire can be processed and have the right values to go with it
				newValue := 0
				aByte := (2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2) - 1

				if wireAndInstructions[i].instructionType == "NOT" {
					if len(instructionInputsValues) != 1 {
						panic("wrong number of arguments not")
					}
					newValue = aByte &^ instructionInputsValues[0]
				} else if wireAndInstructions[i].instructionType == "DIRECT" {
					newValue = instructionInputsValues[0]
					if len(instructionInputsValues) != 1 {
						panic("wrong number of arguments direct")
					}
				} else {
					if len(instructionInputsValues) != 2 {
						panic("wrong number of arguments the rest")
					}
					switch wireAndInstructions[i].instructionType {
					case "AND":
						newValue = instructionInputsValues[0] & instructionInputsValues[1]
					case "OR":
						newValue = instructionInputsValues[0] | instructionInputsValues[1]
					case "RSHIFT":
						newValue = instructionInputsValues[0] >> uint(instructionInputsValues[1])
					case "LSHIFT":
						newValue = instructionInputsValues[0] << uint(instructionInputsValues[1])
					}
				}
				//fmt.Printf("%v ", wireAndInstructions[i])

				wireAndInstructions[i].processed = true
				wireAndInstructions[i].value = newValue
				updatesMade++
				//fmt.Println(wireAndInstructions[i])
				// if wireAndInstructions[i].wire == "a" {
				// 	fmt.Println("-----")
				// }
				// fmt.Println(wireAndInstructions[i])
			}
			// pull out all values by appending s.wireInputs with
		}

		// fmt.Println(inputStrings)
	}
	return wireAndInstructions, updatesMade
}

func main() {
	fmt.Println("Hello World!")
	//import input
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Input file could not be read.")
	}
	inputStrings := strings.Split(string(input), "\n")

	//parse input

	//need a map of wires and their signals.
	//need a slice of wires which have been processed
	//need a wire and it's instruction.
	//for each string populate the struct
	wireAndInstructions := []wireAndInstruction{}
	for _, s := range inputStrings {
		wireAndInstructionRaw := strings.Split(s, " -> ")
		wire := wireAndInstructionRaw[1]
		instruction := wireAndInstructionRaw[0]
		instructionType, instructionInputs, value, processed := processInstruction(instruction)
		wireAndInstructionNew := wireAndInstruction{wire, instructionType, instructionInputs, value, processed}
		//add to slice of structs
		wireAndInstructions = append(wireAndInstructions, wireAndInstructionNew)

	}

	// fmt.Println("--------------------------------------------------------------------------------")

	// for i := range wireAndInstructions {
	// 	if wireAndInstructions[i].processed == true {
	// 		fmt.Println(wireAndInstructions[i])
	// 	}
	// }
	// fmt.Println("--------------------------------------------------------------------------------")

	// loop through set
	// for i := range wireAndInstructions {
	// 	if wireAndInstructions[i].processed == true {
	// 		fmt.Println(wireAndInstructions[i])
	// 	}
	// }
	updatesMade := 1

	for updatesMade > 0 {
		wireAndInstructions, updatesMade = processInstructions(wireAndInstructions)
	}

	fmt.Println("--------------------------------------------------------------------------------")
	// for i := range wireAndInstructions {
	// 	if wireAndInstructions[i].processed == true {
	// 		fmt.Println(wireAndInstructions[i])
	// 	}
	// }
	// for i := range wireAndInstructions {
	// 	if wireAndInstructions[i].processed == false {
	// 		fmt.Println(wireAndInstructions[i])
	// 	}
	// }
	//part1 answer
	wireAResult := 0
	for i := range wireAndInstructions {
		if wireAndInstructions[i].wire == "a" {
			wireAResult = wireAndInstructions[i].value
			fmt.Println(wireAndInstructions[i].value)
			break
		}
		fmt.Printf(".")
	}
	fmt.Println("--------------------------------------------------------------------------------")

	// 	--- Part Two ---
	// Now, take the signal you got on wire a, override wire b to that signal, and
	// reset the other wires (including wire a). What new signal is ultimately provided
	// to wire a?
	// Signal in a to b
	for i := range wireAndInstructions {
		if wireAndInstructions[i].wire == "b" {
			wireAndInstructions[i].value = wireAResult
		} else {
			wireAndInstructions[i].processed = false
		}
	}
	//  process again
	updatesMade = 1
	for updatesMade > 0 {
		wireAndInstructions, updatesMade = processInstructions(wireAndInstructions)
	}
	//part2 answer
	for i := range wireAndInstructions {
		if wireAndInstructions[i].wire == "a" {

			fmt.Println(wireAndInstructions[i].value)
			break
		}
		fmt.Printf(".")
	}

}
