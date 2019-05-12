package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type dim struct {
	l, w, h int
}

func areaOfSmallestSide(d dim) (area int) {
	a1 := d.l * d.h
	a2 := d.h * d.w
	a3 := d.w * d.l

	if a1 <= a2 && a1 <= a3 {
		return a1
	}
	if a2 <= a3 {
		return a2
	}
	return a3

}

func wrappingPaperNeeded(d dim) (area int) {
	a := (2 * d.l * d.h) + (2 * d.h * d.w) + (2 * d.w * d.l) + areaOfSmallestSide(d)
	return a
}

func ribbonForWrap(d dim) (len int) {

	//lwh
	//l
	if d.l >= d.w && d.l >= d.h {
		return 2 * (d.w + d.h)
	}
	//w
	if d.w >= d.l && d.w >= d.h {
		return 2 * (d.l + d.h)
	}
	//h
	if d.h >= d.l && d.h >= d.w {
		return 2 * (d.l + d.w)
	}
	//default return
	return 2 * (d.w + d.h)
}

func main() {
	//  Pull in input file
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("file not read")
	}

	//  Split file into a slice of dimensions
	dimensionsRaw := strings.Split(string(input), "\n")
	dimensions := make([]dim, len(dimensionsRaw))
	var l, w, h int
	for _, oneDim := range dimensionsRaw {
		fmt.Sscanf(oneDim, "%dx%dx%d", &l, &w, &h)
		dimensions = append(dimensions, dim{l, w, h})
	}

	//  Calc total wrapping paper needed for all boxes
	totalWrappingPaperNeeded := 0
	for _, oneDim := range dimensions {
		totalWrappingPaperNeeded += wrappingPaperNeeded(oneDim)
	}
	//Answer to part 1
	fmt.Println(totalWrappingPaperNeeded)

	//calc total ribbon needed
	totalRibbonNeeded := 0
	for _, oneDim := range dimensions {
		totalRibbonNeeded += (oneDim.l * oneDim.w * oneDim.h) + ribbonForWrap(oneDim)
	}

	//Answer part2
	fmt.Println(totalRibbonNeeded)

}
