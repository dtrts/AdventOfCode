package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	// BOILER PLATE --------------------------------------------------------------------
	start := time.Now()
	log.Printf("Starting... %s", start.Format("Jan 2 15:04:05 2006 MST"))

	var inputFileName string
	var part string
	flag.StringVar(&inputFileName, "inputFileName", "input.txt", "Name of the input file")
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.StringVar(&part, "p", "0", "Which part to Solve")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))
	fmt.Println("Input:", inputString)
	// BOILER PLATE --------------------------------------------------------------------
	part1, part2 := 0, 0
	if part == "0" || part == "1" {
		part1 = solvePart1(parseInput(inputString))
	}

	if part == "0" || part == "2" {
		instructions, ghostNodes, ghostNodesWithEnding := parseInput2(inputString)

		fmt.Println("Ghost Nodes")
		for _, ghostNode := range ghostNodes {
			fmt.Println(ghostNode)
		}

		fmt.Println("With Ending")
		for ending, ghostNodes := range ghostNodesWithEnding {

			fmt.Println("Ending:", ending)
			for _, ghostNode := range ghostNodes {
				fmt.Println(ghostNode)
			}
			fmt.Println("")

		}
		// fmt.Printf("%v\n%v\n\n", ghostNodes, ghostNodesWithEnding)

		part2 = solvePart2(instructions, ghostNodesWithEnding, ghostNodes)
	}

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePart1(instructions string, nodes map[string]Node) int {
	currentNode := "AAA"

	instructionIdx := 0

	for currentNode != "ZZZ" {

		instruction := instructions[instructionIdx%len(instructions)]

		switch string(instruction) {
		case "L":
			currentNode = nodes[currentNode].left
		case "R":
			currentNode = nodes[currentNode].right
		default:
			panic("aaaaa")
		}

		instructionIdx++

		if instructionIdx > 1e9 {
			panic("way too long")
		}
	}

	return instructionIdx
}

func numStepsEndings(instructions string, ghostNodes map[string]*GhostNode, startNode *GhostNode) int {
	currentNode := startNode

	instructionIdx := 0

	for currentNode.endsWith != "Z" {

		instruction := instructions[instructionIdx%len(instructions)]

		switch string(instruction) {
		case "L":
			currentNode = ghostNodes[currentNode.left]
		case "R":
			currentNode = ghostNodes[currentNode.right]
		default:
			panic("aaaaa")
		}

		instructionIdx++

		if instructionIdx > 1e9 {
			panic("way too long")
		}
	}

	return instructionIdx
}

func parseInput(input string) (string, map[string]Node) {
	instructions, nodesInput, _ := strings.Cut(input, "\n\n")

	nodes := map[string]Node{}
	for _, nodeInput := range strings.Split(nodesInput, "\n") {
		matches := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`).FindAllStringSubmatch(nodeInput, -1)
		nodes[matches[0][1]] = Node{matches[0][2], matches[0][3]}
	}

	return instructions, nodes
}

func solvePart2(instructions string, ghostNodesWithEnding map[string][]*GhostNode, ghostNodes map[string]*GhostNode) int {
	fmt.Println("Part2")

	startingGhosts := ghostNodesWithEnding["A"]

	pathLengths := make([]int, 0, len(startingGhosts))
	for _, ghost := range startingGhosts {

		pathLength := numStepsEndings(instructions, ghostNodes, ghost)
		pathLengths = append(pathLengths, pathLength)

	}

	if len(pathLengths) == 2 {
		return LCM(pathLengths[0], pathLengths[1])
	} else {
		return LCM(pathLengths[0], pathLengths[1], pathLengths[2:]...)
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func parseInput2(input string) (string, map[string]*GhostNode, map[string][]*GhostNode) {
	instructions, nodesInput, _ := strings.Cut(input, "\n\n")
	fmt.Println(instructions)

	ghostNodes := map[string]*GhostNode{}
	ghostNodesWithEnding := map[string][]*GhostNode{}

	for _, nodeInput := range strings.Split(nodesInput, "\n") {
		matches := regexp.MustCompile(`^(\w\w(\w)) = \((\w\w(\w)), (\w\w(\w))\)$`).FindAllStringSubmatch(nodeInput, -1)

		// [[11A = (11B, XXX) 11A A 11B B XXX X]]

		ghostNode := &GhostNode{
			matches[0][1],
			matches[0][2],
			matches[0][3],
			matches[0][4],
			matches[0][5],
			matches[0][6],
		}

		ghostNodes[ghostNode.id] = ghostNode

		if _, ok := ghostNodesWithEnding[ghostNode.endsWith]; !ok {
			ghostNodesWithEnding[ghostNode.endsWith] = make([]*GhostNode, 0)
		}
		ghostNodesWithEnding[ghostNode.endsWith] = append(ghostNodesWithEnding[ghostNode.endsWith], ghostNode)
	}

	return instructions, ghostNodes, ghostNodesWithEnding
}

type GhostNode struct {
	id            string
	endsWith      string
	left          string
	leftEndsWith  string
	right         string
	rightEndsWith string
}

type NodeId struct {
	id     string
	endsIn string
}

type Node struct {
	left, right string
}
