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

	monkeys := parseInput(inputString)

	monkeys["root"].collapseMonkey("")

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", monkeys["root"].value)
	// BOILER PLATE --------------------------------------------------------------------
	p("Calculating Part 2....")

	monkeys = parseInput(inputString)
	humanName := "humn"
	for i := range [2]int{0, 1} {
		monkeys["root"].parents[i].collapseMonkey(humanName)
	}

	currentMonkey := monkeys["root"]
	humanMonkey, valueMonkey := currentMonkey.parentSplit(humanName)
	matchingValue := valueMonkey.value
	currentMonkey = humanMonkey

	for currentMonkey.name != humanName {

		humanMonkey, valueMonkey = currentMonkey.parentSplit(humanName)

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

func (monkey *Monkey) parentSplit(humanName string) (*Monkey, *Monkey) {

	if monkey.parents[0].isMath || monkey.parents[0].name == humanName {
		return monkey.parents[0], monkey.parents[1]
	}

	return monkey.parents[1], monkey.parents[0]
}

func (monkey *Monkey) collapseMonkey(humanName string) {

	if monkey.isMath {
		parentsNowBothValue := true

		// Run collapse on both parents.
		// If both collapse successfully then you can update current monkey to a Value type
		for i := range [2]int{0, 1} {

			if monkey.parents[i].name == humanName {
				parentsNowBothValue = false
				continue
			}

			if monkey.parents[i].isMath {
				monkey.parents[i].collapseMonkey(humanName)
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
		}
	}

}

func parseInput(input string) map[string]*Monkey {

	monkeys := make(map[string]*Monkey)
	for _, line := range strings.Split(input, "\n") {

		lineSplit := strings.Split(line, ":")
		name := lineSplit[0]
		operationFields := strings.Fields(lineSplit[1])

		newMonkey := &Monkey{
			name: name,
		}

		if len(operationFields) == 1 {
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
	value       int
	isMath      bool
}
