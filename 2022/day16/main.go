package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
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
	// p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	valvesRaw := parseInput(inputString)

	startNode := valvesRaw["AA"]
	maxReleased := make(map[State]int, 0)

	possibleOpenValves := ""

	for _, valve := range valvesRaw {
		if valve.flowRate > 0 {
			possibleOpenValves = addOpenValve(valve.name, possibleOpenValves)
		}
	}

	p(possibleOpenValves)

	p(valvesRaw)

	startState := State{
		valve:              startNode,
		time:               0,
		flow:               0,
		flowRate:           0,
		openValves:         "",
		possibleOpenValves: possibleOpenValves,
	}

	part1 := maxRelease2(valvesRaw, startState, maxReleased, 0)

	p("Calculating Part 1....")

	p("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1) // 1651 (Test) // 1751
	// BOILER PLATE --------------------------------------------------------------------

	startStateElephant := ElephantState{
		valve:              startNode,
		elephantValve:      startNode,
		time:               0,
		flow:               0,
		flowRate:           0,
		openValves:         "",
		possibleOpenValves: possibleOpenValves,
	}

	maxReleasedElephant := make(map[ElephantState]int, 0)

	part2 := elephantRelease(valvesRaw, startStateElephant, maxReleasedElephant, 0)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part2:", part2) // 1707 (Test) //
	// BOILER PLATE --------------------------------------------------------------------

}

type State struct {
	valve              *Valve
	time               int
	flow               int
	flowRate           int
	openValves         string
	possibleOpenValves string
}

type ElephantState struct {
	valve              *Valve
	elephantValve      *Valve
	time               int
	flow               int
	flowRate           int
	openValves         string
	possibleOpenValves string
}

func elephantRelease(valves map[string]*Valve, currentState ElephantState, maxReleased map[ElephantState]int, nested int) int {

	if val, ok := maxReleased[currentState]; ok {
		return val
	}

	maxToBeat := currentState.flow + (currentState.flowRate * (26 - currentState.time))

	if currentState.openValves == currentState.possibleOpenValves || currentState.time == 26 {
		maxReleased[currentState] = maxToBeat
		eValve := currentState.elephantValve
		currentState.elephantValve = currentState.valve
		currentState.valve = eValve
		maxReleased[currentState] = maxToBeat

		return maxToBeat
	}

	tunnelsToCheck := currentState.valve.tunnels
	if currentState.valve.flowRate > 0 && !isValveOpen(currentState.valve.name, currentState.openValves) {
		// tunnelsToCheck = append(tunnelsToCheck, Tunnel{dest: currentState.valve, destName: currentState.valve.name, distance: 1})
		tunnelsToCheck = []Tunnel{{dest: currentState.valve, destName: currentState.valve.name, distance: 1}}

	}

	tunnelsToCheckElephant := currentState.elephantValve.tunnels
	if currentState.valve != currentState.elephantValve && currentState.elephantValve.flowRate > 0 && !isValveOpen(currentState.elephantValve.name, currentState.openValves) {
		// tunnelsToCheckElephant = append(tunnelsToCheckElephant, Tunnel{dest: currentState.elephantValve, destName: currentState.elephantValve.name, distance: 1})
		tunnelsToCheckElephant = []Tunnel{{dest: currentState.elephantValve, destName: currentState.elephantValve.name, distance: 1}}
	}

	for _, tunnel := range tunnelsToCheck {
		for _, tunnelElephant := range tunnelsToCheckElephant {

			newState := ElephantState{
				valve:              tunnel.dest,
				elephantValve:      tunnelElephant.dest,
				time:               currentState.time + tunnel.distance,
				flow:               currentState.flow + currentState.flowRate,
				flowRate:           currentState.flowRate,
				openValves:         currentState.openValves,
				possibleOpenValves: currentState.possibleOpenValves,
			}

			if newState.valve == currentState.valve {
				newState.flowRate += newState.valve.flowRate
				newState.openValves = addOpenValve(newState.valve.name, newState.openValves)
			}

			if newState.elephantValve == currentState.elephantValve {
				newState.flowRate += newState.elephantValve.flowRate
				newState.openValves = addOpenValve(newState.elephantValve.name, newState.openValves)
			}

			destinationMax := elephantRelease(valves, newState, maxReleased, nested+1)

			if destinationMax > maxToBeat {
				maxToBeat = destinationMax
			}
		}
	}

	maxReleased[currentState] = maxToBeat
	eValve := currentState.elephantValve
	currentState.elephantValve = currentState.valve
	currentState.valve = eValve
	maxReleased[currentState] = maxToBeat

	return maxToBeat
}

func maxRelease2(valves map[string]*Valve, currentState State, maxReleased map[State]int, nested int) int {

	if val, ok := maxReleased[currentState]; ok {
		// p("Found Cache", currentState)
		return val
	}

	maxToBeat := currentState.flow + (currentState.flowRate * (30 - currentState.time))

	if currentState.openValves == currentState.possibleOpenValves {
		maxReleased[currentState] = maxToBeat
		return maxToBeat
	}

	tunnelsToCheck := currentState.valve.tunnels

	if currentState.valve.flowRate > 0 && !isValveOpen(currentState.valve.name, currentState.openValves) {
		// tunnelsToCheck = append(tunnelsToCheck, Tunnel{dest: currentState.valve, destName: currentState.valve.name, distance: 1})
		tunnelsToCheck = []Tunnel{{dest: currentState.valve, destName: currentState.valve.name, distance: 1}}
	}

	for _, tunnel := range tunnelsToCheck {

		newState := State{
			valve:              tunnel.dest,
			time:               currentState.time + tunnel.distance,
			flow:               currentState.flow + currentState.flowRate,
			flowRate:           currentState.flowRate,
			openValves:         currentState.openValves,
			possibleOpenValves: currentState.possibleOpenValves,
		}

		if newState.valve == currentState.valve {
			newState.flowRate += newState.valve.flowRate
			newState.openValves = addOpenValve(newState.valve.name, currentState.openValves)
		}

		if newState.time > 30 {
			continue
		}

		destinationMax := maxRelease2(valves, newState, maxReleased, nested+1)

		if destinationMax > maxToBeat {
			maxToBeat = destinationMax
		}
	}

	maxReleased[currentState] = maxToBeat
	return maxToBeat
}

func isValveOpen(valve string, openValves string) bool {

	openValvesSplit := strings.Split(openValves, ",")
	for _, openValve := range openValvesSplit {
		if openValve == valve {
			return true
		}
	}
	return false
}

func addOpenValve(valve string, openValves string) string {
	openValvesSplit := strings.Split(openValves, ",")

	for _, openValve := range openValvesSplit {
		if openValve == valve {
			return openValves
		}
	}

	openValvesSplit = append(openValvesSplit, valve)

	sort.Strings(openValvesSplit)

	return strings.Join(openValvesSplit, ",")
}

func parseInput(input string) map[string]*Valve {

	valves := make(map[string]*Valve, 0)

	for _, inputLine := range strings.Split(input, "\n") {

		inputLineSplit := strings.Split(inputLine, "; ")

		var valveName string
		var flowRate int

		fmt.Sscanf(inputLineSplit[0], "Valve %v has flow rate=%d", &valveName, &flowRate)

		inputLineSplit[1] = strings.TrimPrefix(inputLineSplit[1], "tunnels lead to valves ")
		inputLineSplit[1] = strings.TrimPrefix(inputLineSplit[1], "tunnel leads to valve ")

		destNames := strings.Split(inputLineSplit[1], ", ")

		// p(inputLine, valveName, flowRate, destNames)

		valves[valveName] = &Valve{
			name:      valveName,
			flowRate:  flowRate,
			destNames: destNames,
			open:      flowRate == 0, // Act as if open already when flow rate is 0
		}

	}

	for _, valve := range valves {

		for _, destName := range valve.destNames {
			valve.tunnels = append(valve.tunnels, Tunnel{
				dest:     valves[destName],
				destName: destName,
				distance: 1,
			})
		}
	}

	return valves
}

type Tunnel struct {
	dest     *Valve
	destName string
	distance int
}

type Valve struct {
	name      string
	flowRate  int
	destNames []string
	tunnels   []Tunnel
	open      bool

	// For path searching
	timeTaken    int
	flowReleased int

	// For managing in a Priority Queue
	priority int
	index    int
}
