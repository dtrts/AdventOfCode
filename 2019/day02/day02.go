package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	linesStr := strings.Split(string(file), "\n")
	codesStr := strings.Split(linesStr[0], ",")

	codes := make([]int, len(codesStr), len(codesStr))
	for i, code := range codesStr {
		codes[i], err = strconv.Atoi(code)
		if err != nil {
			panic("Wooaoahah")
		}
	}

	fmt.Println(len(codes), codes)

	// codes[1] = 12
	// codes[2] = 2
	fmt.Println(codes)

	loc := 0
	for loc < len(codes) {
		if codes[loc] == 99 {
			fmt.Println("hit 99")
			break
		}
		if codes[loc] != 1 && codes[loc] != 2 {
			panic("bad op code")
		}
		if codes[loc] == 1 {
			codes[codes[loc+3]] = codes[loc+1] + codes[loc+2]
		}
		if codes[loc] == 1 {
			codes[codes[loc+3]] = codes[loc+1] * codes[loc+2]
		}
		fmt.Println(codes)

		loc += 4
	}

	fmt.Println(codes)

}
