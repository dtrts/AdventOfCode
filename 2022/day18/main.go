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
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))

	// BOILER PLATE --------------------------------------------------------------------

	points := parseInput(inputString)
	part1 := getExposedSides(points)

	boundingBox := fillBoundingBox(points)

	boundingBoxSurfaceArea := getExposedSides(boundingBox)

	min, max, _ := getBoundingPoints(points)

	dx, dy, dz := (1 + max.x - min.x), (1 + max.y - min.y), (1 + max.z - min.z)

	externalSurfaceArea := (2 * dx * dy) + (2 * dx * dz) + (2 * dy * dz)

	part2 := boundingBoxSurfaceArea - externalSurfaceArea

	// Print slice map of lava droplet
	for x := min.x; x <= max.x; x++ {
		p("== X Slice", x, "==")
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				point := Coord{x, y, z}

				if _, ok := boundingBox[point]; ok {
					fmt.Printf("#")
				} else if _, ok := points[point]; ok {
					fmt.Printf("X")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
}

func fillBoundingBox(points map[Coord]bool) map[Coord]bool {

	min, max, newPoints := getBoundingPoints(points)

	p("Min", min, "max", max)

	boundingBox := make(map[Coord]bool)

	for _, corner := range newPoints {
		boundingBox[corner] = true
	}

	boundingBoxLen := len(boundingBox)

	for increaseBoundingBox := true; increaseBoundingBox; {

		newNewPoints := []Coord{}

		for _, newPoint := range newPoints {

			newNeighbours := neighbours(newPoint)
			for _, newNeighbour := range newNeighbours {
				inBoxRange := inBox(newNeighbour, min, max)

				_, inExistingPoints := points[newNeighbour]

				_, inBoundingBox := boundingBox[newNeighbour]

				if inBoxRange && !inExistingPoints && !inBoundingBox {
					boundingBox[newNeighbour] = true
					newNewPoints = append(newNewPoints, newNeighbour)
				}
			}
		}

		newPoints = newNewPoints

		if boundingBoxLen == len(boundingBox) {
			increaseBoundingBox = false
		}
		boundingBoxLen = len(boundingBox)
	}

	return boundingBox

}

func inBox(point Coord, min Coord, max Coord) bool {

	if between(point.x, min.x, max.x) && between(point.y, min.y, max.y) && between(point.z, min.z, max.z) {
		return true
	}
	return false

}

func between(val, min, max int) bool {
	if val >= min && val <= max {
		return true
	}
	return false
}

func getExposedSides(points map[Coord]bool) int {
	exposedSides := 0
	for point := range points {
		exposedSides += 6

		neighbours := neighbours(point)

		for _, neighbour := range neighbours {

			if _, ok := points[neighbour]; ok {
				exposedSides--
			}

		}
	}
	return exposedSides
}

func getBoundingPoints(points map[Coord]bool) (Coord, Coord, []Coord) {

	avgPoint := getAvgPoint(points)

	min, max := avgPoint, avgPoint

	for point := range points {
		if point.x < min.x {
			min.x = point.x
		}
		if point.y < min.y {
			min.y = point.y
		}
		if point.z < min.z {
			min.z = point.z
		}

		if point.x > max.x {
			max.x = point.x
		}
		if point.y > max.y {
			max.y = point.y
		}
		if point.z > max.z {
			max.z = point.z
		}
	}

	min.x -= 1
	min.y -= 1
	min.z -= 1

	max.x += 1
	max.y += 1
	max.z += 1

	ret := []Coord{}

	ret = append(ret, Coord{min.x, min.y, min.z})
	ret = append(ret, Coord{min.x, min.y, max.z})
	ret = append(ret, Coord{min.x, max.y, min.z})
	ret = append(ret, Coord{min.x, max.y, max.z})

	ret = append(ret, Coord{max.x, min.y, min.z})
	ret = append(ret, Coord{max.x, min.y, max.z})
	ret = append(ret, Coord{max.x, max.y, min.z})
	ret = append(ret, Coord{max.x, max.y, max.z})

	return min, max, ret
}

func getAvgPoint(points map[Coord]bool) Coord {

	avgPoint := Coord{}

	for point := range points {

		avgPoint.x += point.x
		avgPoint.y += point.y
		avgPoint.z += point.z

	}

	avgPoint.x /= len(points)
	avgPoint.y /= len(points)
	avgPoint.z /= len(points)

	return avgPoint

}

func neighbours(c Coord) []Coord {

	ret := []Coord{}

	ret = append(ret, Coord{x: c.x + 1, y: c.y, z: c.z})
	ret = append(ret, Coord{x: c.x - 1, y: c.y, z: c.z})
	ret = append(ret, Coord{x: c.x, y: c.y + 1, z: c.z})
	ret = append(ret, Coord{x: c.x, y: c.y - 1, z: c.z})
	ret = append(ret, Coord{x: c.x, y: c.y, z: c.z + 1})
	ret = append(ret, Coord{x: c.x, y: c.y, z: c.z - 1})

	return ret
}

func parseInput(input string) map[Coord]bool {

	lines := strings.Split(input, "\n")

	points := make(map[Coord]bool, 0)

	for _, line := range lines {

		numbers := strings.Split(line, ",")

		numList := [3]int{}

		for i, numString := range numbers {
			num, err := strconv.Atoi(numString)
			if err != nil {
				p("Whattt", num)
				panic("uh oh")
			}
			numList[i] = num
		}

		points[Coord{x: numList[0], y: numList[1], z: numList[2]}] = true

	}
	return points
}

type Coord struct {
	x, y, z int
}
