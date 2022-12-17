package main

import (
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
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
	var part1, part2 int
	// p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	valvesRaw := parseInput(inputString)

	startNodes := []string{"AA"}
	flowNodes := []string{}
	for _, valve := range valvesRaw {
		if valve.flowRate > 0 {
			startNodes = append(startNodes, valve.name)
			flowNodes = append(flowNodes, valve.name)
		}
	}

	valves := generateCondensedMap(inputString, startNodes)

	// Okay I now have a condensed map, and a list of remaining valves to open.
	//

	p(valves)
	for k, v := range valves {
		p(k, v)
	}

	p(startNodes, flowNodes)

	cache := make(map[string]int, 0)
	part1 = pressureRelease(valves, flowNodes, []string{}, 30, cache)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1) // 1651 (Test) // 1751
	// BOILER PLATE --------------------------------------------------------------------
	cache = make(map[string]int, 0)
	cachePR := make(map[string]int, 0)

	// part2 = pressureReleases(valves, flowNodes, [2][]string{make([]string, 0), make([]string, 0)})
	part2 = hmm(valves, flowNodes, make([]string, 0), 26, cache, cachePR)
	// 10 -> 4.7

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part2:", part2) // 1707 (Test) //
	// BOILER PLATE --------------------------------------------------------------------

}

func remove(s []string, i int) []string {
	s2 := make([]string, len(s))
	copy(s2, s)

	// return append(s[:i], s[i+1:]...)
	s2[i] = s2[len(s2)-1]
	return s2[:len(s2)-1]
}

// func getPermutations(remainingValues []string) [][2][]string {

// }

func hmm(condensedMap map[string]*Valve, remainingValves []string, currentRoute []string, timeLimit int, cache map[string]int, cachePR map[string]int) int {

	key := strings.Join(remainingValves, ",") + "|" + strings.Join(currentRoute, ",")
	if val, ok := cache[key]; ok {
		return val
	}

	currentScore := totalScore(condensedMap, currentRoute, timeLimit) + pressureRelease(condensedMap, remainingValves, make([]string, 0), timeLimit, cachePR)

	for i, nextValve := range remainingValves {

		// If the sum of the times +
		newRoute := append(currentRoute, nextValve)
		if totalTime(condensedMap, newRoute) > timeLimit {
			continue
		}

		newRemaining := remove(remainingValves, i)
		newScore := hmm(condensedMap, newRemaining, newRoute, timeLimit, cache, cachePR)

		if newScore > currentScore {
			currentScore = newScore
		}
	}

	cache[key] = currentScore
	return currentScore

}

func pressureRelease(condensedMap map[string]*Valve, remainingValves []string, currentRoute []string, timeLimit int, cache map[string]int) int {

	key := strings.Join(remainingValves, ",") + "|" + strings.Join(currentRoute, ",")
	if val, ok := cache[key]; ok {
		return val
	}

	currentScore := totalScore(condensedMap, currentRoute, timeLimit)

	for i, nextValve := range remainingValves {

		// If the sum of the times +
		newRoute := append(currentRoute, nextValve)
		if totalTime(condensedMap, newRoute) > timeLimit {
			continue
		}

		newRemaining := make([]string, len(remainingValves[:i]))
		copy(newRemaining, remainingValves[:i])
		newRemaining = append(newRemaining, remainingValves[i+1:]...)

		newScore := pressureRelease(condensedMap, newRemaining, newRoute, timeLimit, cache)
		if newScore > currentScore {
			currentScore = newScore
		}
	}

	cache[key] = currentScore
	return currentScore

}

func totalTime(condensedMap map[string]*Valve, route []string) int {

	time := 0
	currValve := "AA"
	for _, valve := range route {
		time += condensedMap[currValve].tunnels[valve] + 1
		currValve = valve
	}

	return time
}

func totalScore(condensedMap map[string]*Valve, route []string, timeLimit int) int {

	totalScore := 0
	timeRemaining := timeLimit
	currValve := "AA"
	for _, valve := range route {

		timeRemaining -= (condensedMap[currValve].tunnels[valve] + 1)

		if timeRemaining <= 0 {
			break
		}
		totalScore += (timeRemaining * condensedMap[valve].flowRate)
		currValve = valve
	}
	return totalScore
}

func generateCondensedMap(input string, startNodes []string) map[string]*Valve {

	p("Generating condensed map")

	condensedMap := parseInput(input)

	// For each key, check if in start nodes. Otherwise delete
	for condensedKey, _ := range condensedMap {

		contains := false
		for _, startNode := range startNodes {
			if startNode == condensedKey {
				contains = true
			}
		}

		if !contains {
			delete(condensedMap, condensedKey)
		}
	}

	for _, startNode := range startNodes {
		freshValves := parseInput(input)

		condensedMap[startNode].tunnels = generateTunnels(startNode, freshValves)
	}

	p("Condensed map")
	for _, node := range condensedMap {
		p(node)
	}

	return condensedMap
}

func generateTunnels(startNode string, valves map[string]*Valve) map[string]int {

	// p("Generate Tunnels", startNode)

	pq := make(PriorityQueue, len(valves))
	i := 0
	for _, valve := range valves {
		pq[i] = valve
		pq[i].index = i

		if valve.name == startNode {
			pq[i].priority = 0
		} else {
			pq[i].priority = math.MaxInt
		}
		i++
	}

	heap.Init(&pq)

	for len(pq) > 0 {
		currentValve := heap.Pop(&pq).(*Valve)

		for tunnelName, distance := range currentValve.tunnels {

			destinationPriority := currentValve.priority + distance

			if destinationPriority < valves[tunnelName].priority {
				pq.update(valves[tunnelName], destinationPriority)
			}
		}
	}

	tunnels := make(map[string]int, 0)

	for _, valve := range valves {
		if valve.flowRate > 0 && valve.name != startNode {
			tunnels[valve.name] = valve.priority
		}
	}

	return tunnels
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

		tunnelNames := strings.Split(inputLineSplit[1], ", ")

		tunnels := make(map[string]int, 0)

		for _, tunnelName := range tunnelNames {
			tunnels[tunnelName] = 1
		}

		// p(inputLine, valveName, flowRate, destNames)

		valves[valveName] = &Valve{
			name:     valveName,
			flowRate: flowRate,
			tunnels:  tunnels,
		}

	}

	return valves
}

type Valve struct {
	name     string
	flowRate int
	tunnels  map[string]int

	// For managing in a Priority Queue
	priority int
	index    int
}

// Implementing Dijkstra for finding new graph of valves only

type PriorityQueue []*Valve

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Valve)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Valve, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
