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
	flag.StringVar(&inputFileName, "inputFileName", "input.txt", "Name of the input file")
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := string(inputBytes)
	// BOILER PLATE --------------------------------------------------------------------

	rows := strings.Fields(inputString)
	height, width := len(rows), len(rows[0])

	treeMap := make([][]int, height)
	for y, row := range rows {
		treeMap[y] = make([]int, width)
		for x, c := range row {
			num, _ := strconv.Atoi(string(c))
			treeMap[y][x] = num
		}
	}

	visible := 0
	maxVisibility := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				visible++
				continue
			}

			visibleCheckLeft, visibleCheckRight, visibleCheckUp, visibleCheckDown := true, true, true, true
			visibleLeft, visibleRight, visibleUp, visibleDown := 0, 0, 0, 0

			for x2 := x - 1; x2 >= 0; x2-- {
				visibleLeft++
				if treeMap[y][x2] >= treeMap[y][x] {
					visibleCheckLeft = false
					break
				}
			}
			for x2 := x + 1; x2 < width; x2++ {
				visibleRight++
				if treeMap[y][x2] >= treeMap[y][x] {
					visibleCheckRight = false
					break
				}
			}
			for y2 := y - 1; y2 >= 0; y2-- {
				visibleUp++
				if treeMap[y2][x] >= treeMap[y][x] {
					visibleCheckUp = false
					break
				}
			}
			for y2 := y + 1; y2 < height; y2++ {
				visibleDown++
				if treeMap[y2][x] >= treeMap[y][x] {
					visibleCheckDown = false
					break
				}
			}

			if visibleCheckLeft || visibleCheckRight || visibleCheckUp || visibleCheckDown {
				visible++
			}

			visibility := visibleLeft * visibleRight * visibleUp * visibleDown
			if visibility > maxVisibility {
				maxVisibility = visibility
			}
		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
	p("Part1:", visible)
	p("Part2:", maxVisibility)
	// BOILER PLATE --------------------------------------------------------------------
}
