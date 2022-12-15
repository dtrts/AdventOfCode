package main

import (
	"flag"
	"fmt"
	"log"
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

	sensors := parseInput(inputString)

	var avgX, avgY int
	for _, sensor := range sensors {
		avgX += sensor.location.x + sensor.beacon.x
		avgY += sensor.location.y + sensor.beacon.y
	}

	minX, maxX := avgX, avgX
	minY, maxY := avgY, avgY
	for _, sensor := range sensors {

		minX = min(sensor.location.x-sensor.distance, minX)
		minY = min(sensor.location.y-sensor.distance, minY)

		maxX = max(sensor.location.x+sensor.distance, maxX)
		maxY = max(sensor.location.y+sensor.distance, maxY)

	}
	avgX /= len(sensors) / 2
	avgY /= len(sensors) / 2

	p("AvgX:", avgX, "AvgY:", avgY, "MinX:", minX, "MinY:", minY, "MaxX:", maxX, "MaxY:", maxY)

	checkY := 2000000
	if inputFileName == "input-test.txt" {
		checkY = 10
	}

	// Can definitely optimise part 1 by finding ranges at the Y level.
	noBeaconCount := 0
	for x := minX - 1; x <= maxX+1; x++ {

		for _, sensor := range sensors {

			currentLoc := Coord{x, checkY}

			if calcDistance(sensor.location, currentLoc) <= sensor.distance {

				// Check to make sure a beacon isn't already there:
				beaconPresent := false
				for _, sensor := range sensors {
					if currentLoc == sensor.beacon {
						beaconPresent = true
					}
				}

				// If no beacon present, count space.
				if !beaconPresent {
					noBeaconCount++
				}

				// We are within range of sensor so move onto next location
				break
			}
		}
	}

	checkBox := 4000000
	if inputFileName == "input-test.txt" {
		checkBox = 20
	}

	emptySpace := Coord{}

doubleLoop:
	for x := 0; x <= checkBox; x++ {
		if x < 20 && x%10000 == 0 {
			p("Checked starting points", x, "out of", checkBox, "%:", x/checkBox)
		}
		for y := 0; y <= checkBox; {

			sensorsInRange := intersectingSensors(Coord{x, y}, sensors, -1)

			if len(sensorsInRange) == 0 {
				p("Empty space found!")
				p(x, y)
				emptySpace = Coord{x, y}
				break doubleLoop
			}

			biggestDy := 0
			for _, sensorInRange := range sensorsInRange {

				dy := sensorInRange.distance - calcDistance(sensorInRange.location, Coord{x, y})
				biggestDy = max(biggestDy, dy)
			}

			y += biggestDy + 1

		}
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", noBeaconCount)
	p("Part2:", emptySpace.x*4000000+emptySpace.y)
	// BOILER PLATE --------------------------------------------------------------------
}

func parseInput(input string) (sensors []Sensor) {

	for _, inputLine := range strings.Split(input, "\n") {
		var sensorX, sensorY, beaconX, beaconY int

		fmt.Sscanf(inputLine, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

		location, beacon := Coord{x: sensorX, y: sensorY}, Coord{x: beaconX, y: beaconY}

		distance := calcDistance(location, beacon)

		sensors = append(sensors, Sensor{
			location: location,
			beacon:   beacon,
			distance: distance,
		})
	}

	return sensors
}

type Coord struct {
	x, y int
}

type Sensor struct {
	location Coord
	beacon   Coord
	distance int
}

func calcDistance(s Coord, b Coord) int {
	return abs(b.x-s.x) + abs(b.y-s.y)
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersectingSensors(location Coord, sensors []Sensor, n int) (intersectingSensors []Sensor) {
	for _, sensor := range sensors {
		if calcDistance(sensor.location, location) <= sensor.distance {
			intersectingSensors = append(intersectingSensors, sensor)
		}

		if n > 0 && len(intersectingSensors) == n {
			return intersectingSensors
		}
	}

	return intersectingSensors
}

// for each position on the y axis
