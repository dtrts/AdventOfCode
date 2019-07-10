package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFilePath = "input.txt"

func parseIntArray(inputFilePath string) []int {
	input, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	values := strings.Split(string(input), " ")

	var valuesInts []int

	for _, v := range values {
		appendValue, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		valuesInts = append(valuesInts, appendValue)
	}
	return valuesInts
}

type node struct { //level      int
	numChild   int
	numMeta    int
	meta       []int
	childNodes []node
}

var input = parseIntArray(inputFilePath)

func parseNode(currentPosition int) (node, int) {

	newNode := node{}

	newNode.numChild = input[currentPosition]
	newNode.numMeta = input[currentPosition+1]

	currentPosition += 2

	for i := 0; i < newNode.numChild; i++ {
		childNode, newPosition := parseNode(currentPosition)
		newNode.childNodes = append(newNode.childNodes, childNode)
		currentPosition = newPosition
	}
	newNode.meta = input[currentPosition : currentPosition+newNode.numMeta]

	currentPosition = currentPosition + newNode.numMeta

	return newNode, currentPosition
}

func nodeSum(n node) (sum int) {
	for _, v := range n.meta {
		sum += v
	}
	for _, n := range n.childNodes {
		sum += nodeSum(n)
	}

	return sum
}

func nodeValue(n node) (value int) {
	if len(n.childNodes) == 0 {
		value += nodeSum(n)
	} else {
		for _, m := range n.meta {
			if m <= len(n.childNodes) {
				value += nodeValue(n.childNodes[m-1])
			}
		}
	}
	return value
}

func main() {

	rootNode, _ := parseNode(0)

	sum := nodeSum(rootNode)
	value := nodeValue(rootNode)
	fmt.Printf("Part-1: %v\nPart-2: %v\n", sum, value)
}

//D08pt1 is great
// func D08pt1() {
// 	input, err := ioutil.ReadFile(inputFilePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//fmt.Println(string(input))
// 	values := strings.Split(string(input), " ")

// 	//fmt.Println(values)

// 	var valuesInts []int

// 	for _, v := range values {
// 		appendValue, err := strconv.Atoi(v)
// 		if err != nil {
// 			panic(err)
// 		}
// 		valuesInts = append(valuesInts, appendValue)
// 	}

// 	fmt.Println(valuesInts)

// 	allMeta := []int{}
// 	for len(valuesInts) > 0 {
// 		fmt.Println("loop start:", len(valuesInts))
// 		//for i := 1; i <= 4; i++ {
// 		startIndex, numMeta := getNextBottomNode(valuesInts)

// 		allMeta = append(allMeta, valuesInts[startIndex+2:startIndex+2+numMeta]...)

// 		if startIndex > 1 {
// 			valuesInts[startIndex-2] = valuesInts[startIndex-2] - 1
// 		}

// 		if startIndex > 0 {
// 			valuesInts = append(valuesInts[:startIndex], valuesInts[startIndex+2+numMeta:]...)
// 		} else {
// 			valuesInts = []int{}
// 		}

// 		//fmt.Println(startIndex, numMeta)
// 	}
// 	fmt.Println(valuesInts)
// 	fmt.Println(allMeta)

// 	sumMeta := 0
// 	for _, v := range allMeta {
// 		sumMeta += v
// 	}

// 	fmt.Println(sumMeta)
// }

// func getNextBottomNode(input []int) (startIndex int, numMeta int) {

// 	for i, v := range input {
// 		if i%2 == 0 && v == 0 {
// 			startIndex = i
// 			numMeta := input[i+1]
// 			return startIndex, numMeta
// 		}
// 	}

// 	return 0, 0
// }
