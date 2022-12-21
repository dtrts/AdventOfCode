package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func p(s ...interface{}) {
	fmt.Println(s...)
}

func main() {
	// BOILER PLATE --------------------------------------------------------------------
	start := time.Now()
	log.Printf("Starting... %s", start.Format("Jan 2 15:04:05 2006 MST"))

	var inputFileName string
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))

	// BOILER PLATE --------------------------------------------------------------------
	// Monkey names are unique
	monkeys := parseInput(inputString)

	mathMonkeysToCalc := getMathMonkeysToCalc(monkeys)

	iteration := 0
	for len(mathMonkeysToCalc) > 0 {
		p("Starting iteration", iteration, "with", len(mathMonkeysToCalc), "monkey(s)")

		for _, mathMonkey := range mathMonkeysToCalc {

			mathMonkey.calculateValue(iteration + 1)

		}

		mathMonkeysToCalc = getMathMonkeysToCalc(monkeys)
		iteration++
	}

	p("Iterations", iteration)
	part1 := 0
	for _, monkey := range monkeys {
		if monkey.iteration == iteration {
			part1 = monkey.value
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	// BOILER PLATE --------------------------------------------------------------------
	p("Calculating Part 2....")

	monkeys = parseInput(inputString)
	monkeys["humn"].isMath, monkeys["humn"].isNumber = true, false

	for i := range [2]int{0, 1} {
		monkeys["root"].parents[i].collapseMonkey()
	}

	// p(monkeys["root"])
	// p(monkeys["root"].parents[0])
	// p(monkeys["root"].parents[1])
	// p(equalityMonkey)
	// p(humanMonkeyRoot)

	monkeys["humn"].isMath, monkeys["humn"].isNumber = false, true

	// curMonk := monkeys["humn"]
	// for curMonk != nil {
	// 	p(curMonk)
	// 	curMonk = curMonk.child
	// }

	currentMonkey := monkeys["root"]
	humanMonkey, valueMonkey := currentMonkey.parentSplit()
	matchingValue := valueMonkey.value
	currentMonkey = humanMonkey

	for currentMonkey.name != "humn" {

		humanMonkey, valueMonkey = currentMonkey.parentSplit()
		p("Current Monkey", currentMonkey)
		p("Current Value", matchingValue)
		p("Human Monkey", humanMonkey)
		p("Value Monkey", valueMonkey)

		switch currentMonkey.operation {
		case "*":
			matchingValue = matchingValue / valueMonkey.value
		case "+":
			matchingValue = matchingValue - valueMonkey.value
		case "-":
			if currentMonkey.parents[0] == humanMonkey {
				matchingValue = matchingValue + valueMonkey.value
			} else {
				matchingValue = valueMonkey.value - matchingValue
			}

		case "/":
			if currentMonkey.parents[0] == humanMonkey {
				matchingValue = matchingValue * valueMonkey.value
			} else {
				matchingValue = valueMonkey.value / matchingValue
			}
		default:
			panic("Unable to handle operation case")
		}
		currentMonkey = humanMonkey
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part2:", matchingValue)
	// BOILER PLATE --------------------------------------------------------------------
}

func (monkey *Monkey) parentSplit() (*Monkey, *Monkey) {

	if monkey.parents[0].isMath || monkey.parents[0].name == "humn" {
		return monkey.parents[0], monkey.parents[1]
	}

	return monkey.parents[1], monkey.parents[0]
}

func calcValue(monkey *Monkey) int {
	valueRet := 0

	switch monkey.operation {
	case "*":
		valueRet = calcValue(monkey.parents[0]) * calcValue(monkey.parents[1])
	case "+":
		valueRet = calcValue(monkey.parents[0]) + calcValue(monkey.parents[1])
	case "-":
		valueRet = calcValue(monkey.parents[0]) - calcValue(monkey.parents[1])
	case "/":
		valueRet = calcValue(monkey.parents[0]) / calcValue(monkey.parents[1])
	case "":
		valueRet = monkey.value
	default:
		panic("Unable to handle operation case")
	}

	return valueRet
}

func (monkey *Monkey) collapseMonkey() *Monkey {

	if monkey.isMath {
		parentsNowBothValue := true
		for i := range [2]int{0, 1} {

			if monkey.parents[i].name == "humn" {
				parentsNowBothValue = false
				continue
			}

			if monkey.parents[i].isMath {
				monkey.parents[i].collapseMonkey()
			}

			if monkey.parents[i].isMath {
				parentsNowBothValue = false
			}
		}

		if parentsNowBothValue {
			switch monkey.operation {
			case "*":
				monkey.value = monkey.parents[0].value * monkey.parents[1].value
			case "+":
				monkey.value = monkey.parents[0].value + monkey.parents[1].value
			case "-":
				monkey.value = monkey.parents[0].value - monkey.parents[1].value
			case "/":
				monkey.value = monkey.parents[0].value / monkey.parents[1].value
			default:
				panic("Unable to handle operation case")
			}

			monkey.isMath = false
			monkey.isNumber = true
		}
	}

	return monkey

}

func parseInput(input string) map[string]*Monkey {

	monkeys := make(map[string]*Monkey)
	for _, line := range strings.Split(input, "\n") {

		lineSplit := strings.Split(line, ":")
		name := lineSplit[0]
		operationFields := strings.Fields(lineSplit[1])

		newMonkey := &Monkey{
			name:      name,
			iteration: -1,
		}

		if len(operationFields) == 1 {

			newMonkey.isNumber = true

			newMonkey.value, _ = strconv.Atoi(operationFields[0])

		} else if len(operationFields) == 3 {

			newMonkey.isMath = true

			newMonkey.operation = operationFields[1]

			newMonkey.parentNames[0] = operationFields[0]
			newMonkey.parentNames[1] = operationFields[2]

		} else {
			panic("Parsing error")
		}
		monkeys[name] = newMonkey
	}

	// Update parents
	for _, monkey := range monkeys {
		if monkey.isMath {
			for i, parentName := range monkey.parentNames {
				monkey.parents[i] = monkeys[parentName]

				if monkeys[parentName].child != nil {
					panic("Parsing input child fail")
				}
				monkeys[parentName].child = monkey

			}

		}
	}

	return monkeys
}

type Monkey struct {
	name        string
	operation   string
	parentNames [2]string
	parents     [2]*Monkey
	child       *Monkey
	isNumber    bool
	value       int
	isMath      bool
	iteration   int
}

func getMathMonkeysToCalc(monkeys map[string]*Monkey) []*Monkey {
	retMonkeys := []*Monkey{}

	for _, monkey := range monkeys {
		// Need it to be a math monkey
		// Need parents to be a number or calculated
		if monkey.isNumber || monkey.iteration >= 0 {
			continue
		}

		parent0Ready := monkey.parents[0].isNumber || monkey.parents[0].iteration >= 0
		parent1Ready := monkey.parents[1].isNumber || monkey.parents[1].iteration >= 0

		if monkey.isMath && parent0Ready && parent1Ready {

			retMonkeys = append(retMonkeys, monkey)

		}

	}
	return retMonkeys
}

func (monkey *Monkey) calculateValue(iteration int) {

	monkey.iteration = iteration

	if monkey.isMath {

		for i := range [2]int{0, 1} {
			if monkey.parents[i].iteration < 0 {
				monkey.parents[i].iteration = iteration
			}
		}

		switch monkey.operation {
		case "*":
			monkey.value = monkey.parents[0].value * monkey.parents[1].value
		case "+":
			monkey.value = monkey.parents[0].value + monkey.parents[1].value
		case "-":
			monkey.value = monkey.parents[0].value - monkey.parents[1].value
		case "/":
			monkey.value = monkey.parents[0].value / monkey.parents[1].value
		default:
			panic("Unable to handle operation case")

		}
	}

}

// Generate the monkey tree.
//
