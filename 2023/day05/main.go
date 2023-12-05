package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

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

	inputString := strings.TrimSpace(string(inputBytes))
	fmt.Println("Input:", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	entities, entityCharts, transforms := parseInput(inputString)

	part1, part2 := solvePart1(entities, transforms), solvePart2(entityCharts, transforms)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func solvePart2(entityCharts []*EntityChart, transforms []Transform) int {

	fmt.Println("iter")

	for _, eh := range entityCharts {
		fmt.Printf("%v\n", eh)
	}

	for _, transform := range transforms {

		newEc := make([]*EntityChart, 0)
		for _, ec := range entityCharts {
			newEc = append(newEc, transformEC(ec, transform)...)
		}
		entityCharts = newEc
	}

	for _, eh := range entityCharts {
		fmt.Printf("%v\n", eh)
	}

	minLoc := entityCharts[0].start
	for _, eh := range entityCharts {
		if eh.start < minLoc {
			minLoc = eh.start
		}
	}
	return minLoc
}

func transformEC(entityChart *EntityChart, transform Transform) []*EntityChart {

	if entityChart.kind != transform.from {
		panic("shouldn't happen")
	}

	transformedEntities := make([]*EntityChart, 0)

	for _, chart := range transform.charts {
		fmt.Printf("Processing EC: %v against %v\n", entityChart, chart)

		entityChartEnd := entityChart.start + entityChart.size - 1
		chartFromEnd := chart.fromStart + chart.size - 1
		newEntityStart, newEntityEnd := entityChart.start, entityChartEnd

		if entityChartEnd < chart.fromStart || chartFromEnd < entityChart.start {
			fmt.Println("No intersection, continue")
			continue
		}

		// Intersection

		if entityChart.start <= chart.fromStart && chart.fromStart <= entityChartEnd {
			newEntityStart = chart.fromStart
		}

		if entityChart.start <= chartFromEnd && chartFromEnd <= entityChartEnd {
			newEntityEnd = chartFromEnd
		}

		if newEntityEnd < newEntityStart {
			panic("New Entity wrong way around")
		}

		intersectionEntity := &EntityChart{
			kind:  transform.to,
			start: chart.toStart + newEntityStart - chart.fromStart, // this needs to be transformed!!
			size:  newEntityEnd - newEntityStart + 1,                // size is the same. 5  to 6
		}

		fmt.Printf("Interseciotn: %v\n", intersectionEntity)

		transformedEntities = append(transformedEntities, intersectionEntity)

		if entityChart.start < newEntityStart {

			startEntity := &EntityChart{
				kind:  entityChart.kind,
				start: entityChart.start,
				size:  newEntityStart - entityChart.start, // AAA
			}

			transformedEntities = append(transformedEntities, transformEC(startEntity, transform)...)
			fmt.Printf("New entity at start: %v returns %v\n", startEntity, transformedEntities)
		}

		if newEntityEnd < entityChartEnd {
			endEntity := &EntityChart{
				kind:  entityChart.kind,
				start: newEntityEnd + 1,
				size:  entityChartEnd - newEntityEnd, // BBBBAAAA 4, 8 -> 5 r:4
			}
			transformedEntities = append(transformedEntities, transformEC(endEntity, transform)...)

			fmt.Printf("New entity at end: %v returns %v\n", endEntity, transformedEntities)

		}

		break

	}

	if len(transformedEntities) == 0 {
		fmt.Println("no intersections found, tranforming in place")
		transformedEntities = append(transformedEntities, &EntityChart{
			kind:  transform.to,
			start: entityChart.start,
			size:  entityChart.size,
		})
	}

	// Get non intersecting parts, run through same function to get transformed entities out.
	// If no intersects then return transformed entity

	return transformedEntities
}

// func splitEntities(entityCharts []*EntityChart, transform Transform) []*EntityChart {

// 	// get entity.
// 	// see if start intersects with a transform

// 	// // take the entity range. See if intersects by a single transform, return two entities.

// 	// for _, c := range transform.charts {

// 	// 	tmpEC := make([]*EntityChart, 0, len(entityCharts))

// 	// 	for _, e := range entityCharts {
// 	// 		eStart, eEnd := e.start, e.start+e.size-1
// 	// 		cStart, cEnd := c.fromStart, c.fromStart+c.size-1

// 	// 		switch {
// 	// 		case cStart == eStart && cEnd == eEnd:
// 	// 			tmpEC = append(tmpEC, e)
// 	// 		// Ceee|eee
// 	// 		case cStart <= eStart && cEnd < eEnd:
// 	// 			tmpEC = append(tmpEC,
// 	// 				&EntityChart{
// 	// 					e.kind,
// 	// 					eStart,
// 	// 					eStart - cEnd,
// 	// 				},
// 	// 				&EntityChart{
// 	// 					e.kind,
// 	// 					cEnd + 1,
// 	// 					cEnd + 1 - eEnd,
// 	// 				},
// 	// 			)
// 	// 		case: cStart < eStart && eEnd
// 	// 			// ee|ee|eee
// 	// 			// ee|eeetttt
// 	// 			// tttteeeeetttt
// 	// 			// eetttteee
// 	// 		}

// 	// 	}

// 	// 	entityCharts = tmpEC
// 	// }
// 	// // make temp entity slice
// 	// // loop through transforms
// 	// // loop through entities
// 	// // if transform intersects entity then split and generate tmpEntity slice
// 	// // if transform doesn't intersect then put onto tmp Entity slice

// }

// func (e *EntityChart) transform(t Transform) []*EntityChart {

// }

func solvePart1(entities []*Entity, transforms []Transform) int {

	for _, e := range entities {
		for _, t := range transforms {
			e.transform(t)
		}
	}

	minLoc := entities[0].value
	for _, e := range entities[1:] {
		if e.value < minLoc {
			minLoc = e.value
		}
	}

	return minLoc
}

func (e *Entity) transform(t Transform) {

	if e.kind != t.from {
		panic("Using wrong transform on entity")
	}

	// Run through all charts finding a valid one and returning if successful
	for _, chart := range t.charts {
		if chart.fromStart <= e.value && e.value < chart.fromStart+chart.size {
			diff := e.value - chart.fromStart
			e.value = chart.toStart + diff
			e.kind = t.to
			return
		}
	}

	// If it hasn't been transformed then the value is the same
	e.kind = t.to
	return
}

func parseInput(input string) ([]*Entity, []*EntityChart, []Transform) {

	sections := strings.Split(input, "\n\n")

	// seeds
	seeds := convertNumList(strings.Split(sections[0], ": ")[1], " ")
	entities := make([]*Entity, 0, len(seeds))
	for _, seed := range seeds {
		entities = append(entities,
			&Entity{
				kind:  "seed",
				value: seed,
			},
		)
	}

	entityCharts := make([]*EntityChart, 0)
	for i := 0; i < len(seeds); i += 2 {
		entityCharts = append(entityCharts, &EntityChart{
			kind:  "seed",
			start: seeds[i],
			size:  seeds[i+1],
		})
	}

	transforms := make([]Transform, 0, len(sections)-1)
	for _, transformInput := range sections[1:] {
		transforms = append(transforms, parseTransform(transformInput))
	}

	return entities, entityCharts, transforms
}

func parseTransform(input string) Transform {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`(\w+)-to-(\w+) map:`).FindAllStringSubmatch(lines[0], -1)
	transform := Transform{
		re[0][1],
		re[0][2],
		make([]Chart, 0, len(lines)-1),
	}

	for _, line := range lines[1:] {
		numbers := convertNumList(line, " ")

		chart := Chart{
			numbers[1],
			numbers[0],
			numbers[2],
		}

		transform.charts = append(transform.charts, chart)
	}

	return transform
}

type Entity struct {
	kind  string
	value int
}

type EntityChart struct {
	kind  string
	start int
	size  int
}

type Transform struct {
	from   string
	to     string
	charts []Chart
}

type Chart struct {
	fromStart int
	toStart   int
	size      int
}

func convertNumList(s string, sep string) []int {
	stringList := strings.Split(s, sep)
	stringList = slices.DeleteFunc(stringList, func(e string) bool { return e == "" })

	stringInt := make([]int, 0, len(stringList))

	for _, e := range stringList {
		n, _ := strconv.Atoi(e)
		stringInt = append(stringInt, n)

	}

	return stringInt
}
