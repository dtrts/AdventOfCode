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
	p("Input string", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	p("Calculating Part 1....")

	nodes := parseMap(inputString)

	lowPoints := []Coord{}
	for _, node := range nodes {
		if node.height == 1 {
			lowPoints = append(lowPoints, node.coordinate)
		}
	}

	p(lowPoints)
	// for _, coord := range lowPoints {

	// 	testNodes := parseMap(inputString)
	// 	shortestPath := shortestPath(testNodes, testNodes[coord])

	// 	p(shortestPath)

	// }

	// startNode, endNode := &Node{}, &Node{}
	// for _, node := range nodes {
	// 	if node.start {
	// 		startNode = node
	// 	}
	// 	if node.end {
	// 		endNode = node
	// 	}
	// }
	// p(startNode, endNode)

	// pq := make(PriorityQueue, len(nodes))
	// i := 0
	// for _, node := range nodes {
	// 	pq[i] = &Item{
	// 		value:    node,
	// 		priority: node.distance,
	// 		index:    i,
	// 	}
	// 	i++
	// }
	// heap.Init(&pq)

	// p(len(pq))

	// for len(pq) > 0 {
	// 	// for i := 0; i < 5; i++ {
	// 	currentItem := heap.Pop(&pq).(*Item)
	// 	p("---")
	// 	p(currentItem)
	// 	p(currentItem.value)
	// 	p(currentItem.value.dest)

	// 	for _, destNode := range currentItem.value.dest {

	// 		var destNodeItem *Item
	// 		for _, item := range pq {
	// 			if item.value == destNode {
	// 				destNodeItem = item
	// 				break
	// 			}
	// 		}
	// 		if destNodeItem == nil {
	// 			continue
	// 		}

	// 		newDistance := currentItem.value.distance + 1

	// 		if newDistance < destNode.distance {
	// 			destNode.distance = newDistance
	// 			destNode.via = currentItem.value
	// 			pq.update(destNodeItem, destNode, newDistance)
	// 		}

	// 	}

	// }

	// p(len(pq))

	p("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	// p("Part1:", endNode.distance)
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}

func parseMap(inputString string) map[Coord]*Node {
	mapLines := strings.Split(inputString, "\n")

	aHeightAdjustment := int(string('a')[0]) - 1

	nodes := make(map[Coord]*Node)

	for y, mapLine := range mapLines {

		for x, char := range mapLine {

			height := int(string(char)[0]) - aHeightAdjustment
			distance := math.MaxInt
			var isStart, isEnd bool
			if string(char) == "S" {
				height = 1 // 'a'
				// distance = 0
				isStart = true
			}
			if string(char) == "E" {
				height = 26 // 'z'
				isEnd = true
			}

			coordinate := Coord{x: x, y: y}

			newNode := &Node{
				coordinate: coordinate,
				char:       string(char),
				height:     height,
				start:      isStart,
				end:        isEnd,
				dest:       make([]*Node, 0),
				distance:   distance,
				via:        nil,
			}
			nodes[coordinate] = newNode

		}

	}

	for coord, node := range nodes {

		up := Coord{x: coord.x, y: coord.y - 1}
		down := Coord{x: coord.x, y: coord.y + 1}
		left := Coord{x: coord.x - 1, y: coord.y}
		right := Coord{x: coord.x + 1, y: coord.y}

		checkDirs := []Coord{up, down, left, right}

		for _, checkDir := range checkDirs {
			if neighbour, ok := nodes[checkDir]; ok {
				if neighbour.height <= node.height+1 {
					node.dest = append(node.dest, neighbour)
				}
			}
		}
	}

	return nodes
}

func shortestPath(nodes map[Coord]*Node, startNode *Node) int {

	endNode := &Node{}
	for _, node := range nodes {
		if node.end {
			endNode = node
		}
	}

	startNode.distance = 0

	pq := make(PriorityQueue, len(nodes))
	i := 0
	for _, node := range nodes {
		pq[i] = &Item{
			value:    node,
			priority: node.distance,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		// for i := 0; i < 5; i++ {
		currentItem := heap.Pop(&pq).(*Item)

		for _, destNode := range currentItem.value.dest {

			var destNodeItem *Item
			for _, item := range pq {
				if item.value == destNode {
					destNodeItem = item
					break
				}
			}
			if destNodeItem == nil {
				continue
			}

			newDistance := currentItem.value.distance + 1

			if newDistance < destNode.distance {
				destNode.distance = newDistance
				destNode.via = currentItem.value
				pq.update(destNodeItem, destNode, newDistance)
			}

		}

	}

	return endNode.distance

}

type Node struct {
	coordinate Coord
	char       string
	height     int
	start      bool
	end        bool
	dest       []*Node
	distance   int
	via        *Node
}

type Coord struct {
	x int
	y int
}

type Item struct {
	value    *Node // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	index    int   // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// actually the otherway around, shorter is better
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
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
func (pq *PriorityQueue) update(item *Item, value *Node, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
