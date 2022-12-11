package main

import (
	"flag"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"sort"
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
	flag.StringVar(&inputFileName, "inputFileName", "input.txt", "Name of the input file")
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))

	monkeysInput := strings.Split(inputString, "\n\n")
	// BOILER PLATE --------------------------------------------------------------------

	monkeys := make([]Monkey, 0, len(monkeysInput))

	for _, monkeyInput := range monkeysInput {
		// p(monkeyInput)
		monkeyInputLineSplit := strings.Split(monkeyInput, "\n")

		monkeys = append(monkeys, Monkey{
			items:       parseStartingItems(monkeyInputLineSplit[1]),
			operation:   parseOperation(monkeyInputLineSplit[2]),
			test:        parseTest(monkeyInputLineSplit[3]),
			monkeyTrue:  parseResult(monkeyInputLineSplit[4]),
			monkeyFalse: parseResult(monkeyInputLineSplit[5]),
		})

	}

	inspectionCount := make(map[int]int, 0)

	for round := 1; round <= 20; round++ {

		for mIndex, monkey := range monkeys {
			// p("Start", mIndex, monkey)
			inspectionCount[mIndex] += len(monkey.items)
			for _, item := range monkey.items {
				// p("start item:", item)
				item = monkey.operation(item)
				// p("post op:", item)
				item = item / 3
				// p("relief:", item)

				remainder := item % monkey.test
				// p("remainder:", remainder)

				if remainder == 0 {
					monkeys[monkey.monkeyTrue].items = append(monkeys[monkey.monkeyTrue].items, item)
					// p("Test pass, going to monkey:", monkey.monkeyTrue, monkeys[monkey.monkeyTrue].items)
				} else {
					monkeys[monkey.monkeyFalse].items = append(monkeys[monkey.monkeyFalse].items, item)
					// p("Test pass, going to monkey:", monkey.monkeyFalse, monkeys[monkey.monkeyFalse].items)
				}
			}
			monkeys[mIndex].items = make([]int, 0)
			// p("End", mIndex, monkey)

		}

		// p("Round", round)
		// for mI, m := range monkeys {
		// 	p("Monkey ", mI, ":", m.items)
		// }
		// p("")
	}

	// p(inspectionCount)

	inspections := maps.Values(inspectionCount)

	sort.Ints(inspections)
	lenInspections := len(inspections)
	part1 := inspections[lenInspections-1] * inspections[lenInspections-2]

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	p("Part1:", part1)
	// BOILER PLATE --------------------------------------------------------------------

	monkeys2 := make([]Monkey, 0, len(monkeysInput))

	for _, monkeyInput := range monkeysInput {
		p(monkeyInput)

		monkeyInputLineSplit := strings.Split(monkeyInput, "\n")

		monkeys2 = append(monkeys2, Monkey{
			items:       parseStartingItems(monkeyInputLineSplit[1]),
			operation:   parseOperation(monkeyInputLineSplit[2]),
			test:        parseTest(monkeyInputLineSplit[3]),
			monkeyTrue:  parseResult(monkeyInputLineSplit[4]),
			monkeyFalse: parseResult(monkeyInputLineSplit[5]),
		})

	}

	inspectionCount2 := make(map[int]int, 0)

	modulo := 1
	for _, m := range monkeys2 {
		modulo *= m.test
	}

	for round := 1; round <= 10000; round++ {

		for mIndex, monkey := range monkeys2 {
			// p("Start", mIndex, monkey)
			inspectionCount2[mIndex] += len(monkey.items)
			for _, item := range monkey.items {
				// p("start item:", item)
				item = monkey.operation(item)
				// p("post op:", item)
				// item = item / 3

				item = item % modulo
				// p("relief:", item)

				remainder := item % monkey.test
				// p("remainder:", remainder)

				if remainder == 0 {
					monkeys2[monkey.monkeyTrue].items = append(monkeys2[monkey.monkeyTrue].items, item)
					// p("Test pass, going to monkey:", monkey.monkeyTrue, monkeys[monkey.monkeyTrue].items)
				} else {
					monkeys2[monkey.monkeyFalse].items = append(monkeys2[monkey.monkeyFalse].items, item)
					// p("Test pass, going to monkey:", monkey.monkeyFalse, monkeys[monkey.monkeyFalse].items)
				}
			}
			monkeys2[mIndex].items = make([]int, 0)
			// p("End", mIndex, monkey)

		}

		if round == 1 || round == 20 || round%1000 == 0 {
			p("Round", round)
			for k, v := range inspectionCount2 {
				p("Monkey ", k, ":", v)
			}
			p("")
		}
	}

	p(inspectionCount2)

	inspections2 := maps.Values(inspectionCount2)

	sort.Ints(inspections2)

	lenInspections2 := len(inspections2)
	part2 := inspections2[lenInspections2-1] * inspections2[lenInspections2-2]
	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
	// 4788350550 WRONG
}

type Monkey struct {
	items       []int
	operation   func(int) int
	test        int
	monkeyTrue  int
	monkeyFalse int
}

func parseTest(testInput string) int {
	testInputSplit := strings.Fields(testInput)
	divisibleBy, err := strconv.Atoi(testInputSplit[3])
	if err != nil {
		panic("Unable to parse test")
	}
	return divisibleBy
}

func parseResult(resultInput string) int {
	resultInputSplit := strings.Fields(resultInput)
	destination, err := strconv.Atoi(resultInputSplit[5])
	if err != nil {
		panic("Unable to parse result")
	}
	return destination
}

func parseStartingItems(itemsInput string) []int {

	itemsSplit := strings.Fields(itemsInput)
	itemsSplit = itemsSplit[2:]
	items := make([]int, len(itemsSplit))

	for i, item := range itemsSplit {
		item = strings.Trim(item, ",")
		itemNum, err := strconv.Atoi(item)
		if err != nil {
			panic("Unable to convert item")
		}
		items[i] = itemNum
	}

	return items
}

func parseOperation(operationInput string) func(int) int {

	operationInputSplit := strings.Fields(operationInput)
	// a := operationInputSplit[3]
	operation := operationInputSplit[4]
	b := operationInputSplit[5]

	if operation == "*" && b == "old" {
		return func(old int) int {
			return old * old
		}
	}

	bNum, err := strconv.Atoi(b)
	if err != nil {
		panic("Unable to parse operation")
	}

	if operation == "+" {
		return func(old int) int {
			return old + bNum
		}
	}

	if operation == "*" {
		return func(old int) int {
			return old * bNum
		}
	}

	p("No Matches found for", operationInput)
	panic("no matches found")
}
