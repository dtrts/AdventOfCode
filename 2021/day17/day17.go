package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("Starting...")

	testInputPtr := flag.Bool("testInput", false, "a bool")
	flag.Parse()

	// target area: x=185..221, y=-122..-74
	rangeX := [2]int{185, 221}
	rangeY := [2]int{-122, -74}

	if *testInputPtr {
		fmt.Println("Using test Input")
		rangeX[0], rangeX[1] = 20, 30
		rangeY[0], rangeY[1] = -10, -5
		// target area: x=20..30, y=-10..-5
	}

	// Working on the assumption the initial target area is completely to the lower right of the source.

	// initial Y is abs(minY) - 1. So maxY is sum of 1..N where N is initial Y
	absY := Abs(rangeY[0])

	fmt.Println("Part 1:", (((absY - 1) * (absY)) / 2))

	// Part 2 is going to be bruteforce.

	// X has to at least get to the target area, so choose lowerBoundX such that
	// SUM(1..lbX) >= min(rangeX)
	// => (lbX(lbX+1))/2 >= minX
	// => lbX^2 + lbX >= minX*2 => (lbX + 0.5)^2 - 0.25 >= minX*2

	lbX := int(math.Sqrt(float64(rangeX[0]*2)+0.25) - 0.5)
	ubX := rangeX[1]
	fmt.Println("Checking X ranges", lbX, ubX)

	lbY := rangeY[0]
	ubY := (-1 * rangeY[0]) - 1
	fmt.Println("Checking Y ranges", lbY, ubY)

	goodProbes := [][2]int{}

	for ivx := lbX; ivx <= ubX; ivx++ {
		for ivy := lbY; ivy <= ubY; ivy++ {

			testProbe := probe{x: 0, y: 0, vx: ivx, vy: ivy}

			if testProbe.hitTarget(rangeX, rangeY) {
				goodProbes = append(goodProbes, [2]int{ivx, ivy})
			}

		}
	}

	// fmt.Println(goodProbes)
	fmt.Println("Part 2:", len(goodProbes))
	log.Printf("Duration: %s", time.Since(start))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *probe) hitTarget(rangeX, rangeY [2]int) bool {
	for !p.overShot(rangeX, rangeY) {
		if p.inTarget(rangeX, rangeY) {
			return true
		}
		p.move()
	}
	return false
}

func (p *probe) inTarget(rangeX, rangeY [2]int) bool {
	if between(p.x, rangeX) && between(p.y, rangeY) {
		return true
	}

	return false
}

func between(val int, intRange [2]int) bool {

	if val >= intRange[0] && val <= intRange[1] {
		return true
	}

	return false
}

func (p *probe) overShot(rangeX, rangeY [2]int) bool {
	if p.y < rangeY[0] || rangeX[1] < p.x {
		return true
	}

	if p.vx == 0 && p.x < rangeX[0] {
		return true
	}

	return false
}

func (p *probe) move() {
	p.x += p.vx

	p.y += p.vy

	if p.vx < 0 {
		p.vx++
	}
	if p.vx > 0 {
		p.vx--
	}

	p.vy--
}

type probe struct {
	x  int
	y  int
	vx int
	vy int
}
