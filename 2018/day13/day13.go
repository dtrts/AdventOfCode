package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFilePath13 = "input.txt"

func importInput(inputFilePath string) string {
	input, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	return string(input)
}

var mapRaw = importInput(inputFilePath13)
var mapGlobal [][]track
var cartGlobal []cart

type track struct {
	N bool
	S bool
	E bool
	W bool
}

var trackV = track{true, true, false, false}
var trackH = track{false, false, true, true}
var trackNW = track{true, false, false, true}
var trackNE = track{true, false, true, false}
var trackSW = track{false, true, false, true}
var trackSE = track{false, true, true, false}
var trackInt = track{true, true, true, true}
var trackN = track{true, false, false, false}
var trackS = track{false, true, false, false}
var trackE = track{false, false, true, false}
var trackW = track{false, false, false, true}
var trackNil = track{false, false, false, false}

type cart struct {
	lineNum   int
	charNum   int
	direction track //  direction of travel 1 2 3 4 =>  E N W S
	prevTurn  int8  //  -1 (left) 1 (right) -2 (left straight) 2 (right straight)
}

// track type  0 is vertical 1 is horizontal

//  processes all characters into mapGlobal, and appends carts onto a list
func populateMap(mapRaw string) {

	mapRawLen := len(mapRaw)
	if string(mapRaw[mapRawLen-1]) == "\n" {
		mapRaw = mapRaw[:mapRawLen-1]
		mapRawLen--
	}

	numNewLineChars := strings.Count(mapRaw, "\n")
	numLines := numNewLineChars + 1
	lineLength := (mapRawLen - numNewLineChars) / numLines

	fmt.Printf("number of new lines: %v. num new line chars: %v. line length: %v\n", numLines, numNewLineChars, lineLength)

	mapGlobal = make([][]track, numLines)

	for i := range mapGlobal {
		mapGlobal[i] = make([]track, lineLength)
	}

	mapRaw = strings.Replace(mapRaw, "\n", "", -1)

	for i, v := range mapRaw {

		lineNum := i / lineLength
		charNum := i % lineLength

		var newTrack track

		switch string(v) {
		case "-":
			newTrack = trackH
		case ">":
			newTrack = trackH
		case "<":
			newTrack = trackH
		case "|":
			newTrack = trackV
		case "v":
			newTrack = trackV
		case "^":
			newTrack = trackV
		case "+":
			newTrack = trackInt
		}

		if charNum == 0 && lineNum == 0 && string(v) == "/" {
			newTrack = trackSE
		} else if lineNum == 0 {
			switch string(v) {
			case "/":
				newTrack = trackSE
			case "\\":
				newTrack = trackSW
			}
		} else if charNum == 0 {
			switch string(v) {
			case "/":
				newTrack = trackSE
			case "\\":
				newTrack = trackNE
			}
		} else {

			if string(v) == "/" {
				if mapGlobal[lineNum][charNum-1].E {
					newTrack = trackNW
				} else {
					newTrack = trackSE
				}
			}

			if string(v) == "\\" {
				if mapGlobal[lineNum][charNum-1].E {
					newTrack = trackSW
				} else {
					newTrack = trackNE
				}
			}
		}

		mapGlobal[lineNum][charNum] = newTrack

		//  adding carts onto cart global
		if string(v) == ">" || string(v) == "^" || string(v) == "<" || string(v) == "v" {

			var dir track
			switch string(v) {
			case ">":
				dir = trackE
			case "^":
				dir = trackN
			case "<":
				dir = trackW
			case "v":
				dir = trackS
			}
			newCart := cart{lineNum, charNum, dir, 0}
			cartGlobal = append(cartGlobal, newCart)

		}

	}

}

//  Used for drawing the map, will just return the first cart in the list at that spot
func retrieveCart(lineNum int, charNum int) cart {
	for _, v := range cartGlobal {
		if v.lineNum == lineNum && v.charNum == charNum {
			return v
		}

	}
	return cart{}
}

// Prints map with the carts on it
func printMap() {
	for i := range mapGlobal {
		for i2, v := range mapGlobal[i] {
			var pic string
			switch v {
			case trackV:
				pic = "|"
			case trackH:
				pic = "-"
			case trackNW:
				pic = "/"
			case trackNE:
				pic = "\\"
			case trackSW:
				pic = "\\"
			case trackSE:
				pic = "/"
			case trackInt:
				pic = "+"
			default:
				pic = " "
			}
			cart := retrieveCart(i, i2)
			if cart.direction != trackNil {
				switch cart.direction {
				case trackE:
					pic = ">"
				case trackN:
					pic = "^"
				case trackW:
					pic = "<"
				case trackS:
					pic = "v"
				}
			}
			fmt.Printf(pic)

		}
		fmt.Println("")
	}
}

func tickCart(c *cart) {
	//fmt.Println(*c)
	//  Move one space
	switch c.direction {
	case trackN:
		c.lineNum-- //"^"
	case trackS:
		c.lineNum++ //"v"
	case trackE:
		c.charNum++ //">"
	case trackW:
		c.charNum-- //"<"
	}

	newGround := mapGlobal[c.lineNum][c.charNum]
	// reaching a corner, direction in matches the opposite of one direction, direction out is then set as the other value
	if c.direction == trackN && newGround == trackSE {
		c.direction = trackE
	}
	if c.direction == trackN && newGround == trackSW {
		c.direction = trackW
	}
	if c.direction == trackS && newGround == trackNE {
		c.direction = trackE
	}
	if c.direction == trackS && newGround == trackNW {
		c.direction = trackW
	}
	if c.direction == trackE && newGround == trackNW {
		c.direction = trackN
	}
	if c.direction == trackE && newGround == trackSW {
		c.direction = trackS
	}
	if c.direction == trackW && newGround == trackNE {
		c.direction = trackN
	}
	if c.direction == trackW && newGround == trackSE {
		c.direction = trackS
	}

	if newGround == trackInt {

		if c.prevTurn == 0 { //  going stright, has turned right, now  turning left
			c.direction = leftTurn(c.direction)
			c.prevTurn = 1
		} else if c.prevTurn == 1 {
			c.prevTurn = 2
		} else if c.prevTurn == 2 {
			c.direction = rightTurn(c.direction)
			c.prevTurn = 0
		}

	}
	//fmt.Println(*c)

}

func leftTurn(dir track) track {
	newTrack := trackNil
	switch dir {
	case trackN:
		newTrack = trackW
	case trackW:
		newTrack = trackS
	case trackS:
		newTrack = trackE
	case trackE:
		newTrack = trackN
	}
	return newTrack
}
func rightTurn(dir track) track {
	newTrack := trackNil
	switch dir {
	case trackN:
		newTrack = trackE
	case trackE:
		newTrack = trackS
	case trackS:
		newTrack = trackW
	case trackW:
		newTrack = trackN

	}
	return newTrack
}

func isCollision(carts []cart) bool {

	isCollision := false

	for i, v := range carts {
		for _, v2 := range carts[i+1:] {
			if v.lineNum == v2.lineNum && v.charNum == v2.charNum {
				fmt.Println("Collision! Cart ", v, " and ", v2, " have collided at ", v.lineNum, v.charNum)
				isCollision = true
			}
		}
	}
	return isCollision
}

func deleteCart(c cart) {
	for i, v := range cartGlobal {

		if v == c && i == len(cartGlobal)-1 {
			cartGlobal = cartGlobal[:i]
		} else if v == c && i == 0 {
			cartGlobal = cartGlobal[1:]
		} else if v == c {
			cartGlobal = append(cartGlobal[:i], cartGlobal[i+1:]...)
		}

	}
}

var tickNum = 0

func isCollisionCheck(carts []cart) {

	for i, v := range carts {
		for _, v2 := range carts[i+1:] {
			if v.lineNum == v2.lineNum && v.charNum == v2.charNum {
				fmt.Println("Collision! Cart ", v, " and ", v2, " have collided at ", v.lineNum, v.charNum)
				deleteCart(v)
				deleteCart(v2)
			}
		}
	}
}

func moveCartOnSpot(lineNum int, charNum int) {
	for _, c := range cartGlobal {
		if c.lineNum == lineNum && c.charNum == charNum {
			tickCart(&c)
		}
	}
}

//D13 choo choo
func main() {
	//fmt.Println(mapRaw)

	populateMap(mapRaw)
	fmt.Println(cartGlobal)
	printMap()

	for len(cartGlobal) > 1 {
		for lineNum, line := range mapGlobal {
			for colNum := range line {
				for i3 := range cartGlobal {

					if cartGlobal[i3].lineNum == lineNum && cartGlobal[i3].charNum == colNum {

					}
				}
			}
		}
	}
	i0 := 0
	for len(cartGlobal) > 1 {
		cartCopy := cartGlobal
		for i := range cartGlobal {
			tickCart(&cartGlobal[i])
		}
		for i := range cartCopy {
			tickCart(&cartCopy[i])
			isCollisionCheck(cartCopy)
		}

		/*
			for i := range cartGlobal {
				tickCart(&cartGlobal[i])
			}
			isCollisionDelCarts(cartGlobal)
		*/
		i0++
		//printMap()
		//fmt.Println(cartGlobal)

		fmt.Println("tick number : ", tickNum, "num carts:", len(cartGlobal))
		tickNum++
	}
	//printMap()
	for i := range cartGlobal {

		tickCart(&cartGlobal[i])
	}

	fmt.Println(cartGlobal)
}

/*

func isCollisionDelCarts(carts []cart) bool {

	isCollision := false

	var collidedCarts []cart

	for i, v := range carts {
		for _, v2 := range carts[i+1:] {
			if v.lineNum == v2.lineNum && v.charNum == v2.charNum {
				fmt.Println("Collision! Cart ", v, " and ", v2, " have collided at ", v.lineNum, v.charNum)
				collidedCarts = append(collidedCarts, v, v2)
				isCollision = true
			}
		}
	}
	if isCollision {
		fmt.Println(cartGlobal)

	}
	for _, c := range collidedCarts {
		deleteCart(c)
	}
	if isCollision {

		fmt.Println(cartGlobal)
		fmt.Println("Num reaining carts:", len(cartGlobal))
		fmt.Println("...")
	}
	return isCollision
}

*/
