package main

import (
	"container/ring"

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

	originalOrder, zeroNode := parseInput(inputString, 1)

	mix(originalOrder)

	part1 := getCoordinates(zeroNode)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	// BOILER PLATE --------------------------------------------------------------------

	decryptionKey := 811589153
	originalOrder, zeroNode = parseInput(inputString, decryptionKey)

	for i := 0; i < 10; i++ {
		mix(originalOrder)
	}

	part2 := getCoordinates(zeroNode)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
}

func parseInput(input string, decryptionKey int) ([]*ring.Ring, *ring.Ring) {

	inputLines := strings.Split(input, "\n")

	ringLen := len(inputLines)

	r := ring.New(ringLen)
	originalOrder := []*ring.Ring{}

	var zeroNode *ring.Ring

	for _, line := range inputLines {
		num, err := strconv.Atoi(line)

		if err != nil {
			panic("Unable to convert number")
		}

		r.Value = nodeValue{
			value: num * decryptionKey,
			move:  (num * decryptionKey) % (ringLen - 1),
		}

		originalOrder = append(originalOrder, r)

		if num == 0 {
			zeroNode = r
		}

		r = r.Next()
	}

	return originalOrder, zeroNode

}

type nodeValue struct {
	value int
	move  int
}

func mix(originalOrder []*ring.Ring) {
	for _, curR := range originalOrder {

		num := getMove(curR)

		q := curR.Prev()
		splitR := q.Unlink(1)

		s := q.Move(num)
		s.Link(splitR)

		// printRing(zeroNode)
	}
}

func getCoordinates(zeroNode *ring.Ring) int {

	return getVal(zeroNode.Move(1000)) + getVal(zeroNode.Move(2000)) + getVal(zeroNode.Move(3000))

}

func getVal(r *ring.Ring) int {
	if val, ok := r.Value.(nodeValue); ok {
		return val.value
	}

	panic("Oh no")
}

func getMove(r *ring.Ring) int {
	if val, ok := r.Value.(nodeValue); ok {
		return val.move
	}

	panic("Oh no")
}

func printRing(zeroNode *ring.Ring) {
	printR := zeroNode
	for i := 0; i < printR.Len(); i++ {
		fmt.Printf("%+3d", printR.Value)
		printR = printR.Next()
	}
	p("")

}

// func moveNode(r *ring.Ring) {
// 	if num, ok := r.Value.(int); !ok {
// 		panic("unable to assert value")
// 	}

// 	s := r.Move(num)

// 	r.Prev().

// }
