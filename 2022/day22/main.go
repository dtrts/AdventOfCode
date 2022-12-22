package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	inputString := strings.TrimRight(string(inputBytes), " \n\t")
	// BOILER PLATE --------------------------------------------------------------------

	boardMap, instructions := parseInput(inputString)

	for _, boardTile := range boardMap {
		p(boardTile)
	}
	p(instructions)

	position := findStartingPosition(boardMap)
	facing := Coord{-1, 0} // Start facing up, then turn right (added by me)

	for _, instruction := range instructions {

		facing = facing.turn(instruction.turn)

		for i := 0; i < instruction.distance; i++ {
			newPosition := boardMap[position].neighbours[facing]
			if boardMap[newPosition].isWall {
				break
			}
			position = newPosition
		}

	}

	p(position, facing)

	part1 := (1000 * position.row) + (4 * position.column) + facingScore(facing)

	// p(generateEdgePoints([2]int{151, 50}, [2]int{151, 1}))

	position = findStartingPosition(boardMap)
	facing = Coord{-1, 0} // Start facing up, then turn right (added by me)

	for _, instruction := range instructions {

		facing = facing.turn(instruction.turn)

		for i := 0; i < instruction.distance; i++ {
			newPosition := boardMap[position].cubeNeighbours[facing]
			if boardMap[newPosition.position].isWall {
				break
			}
			position = newPosition.position
			facing = newPosition.direction
		}

	}
	part2 := (1000 * position.row) + (4 * position.column) + facingScore(facing)

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	p("Part2:", part2) // 55244 is WRONG
	// BOILER PLATE --------------------------------------------------------------------
}

func (facing Coord) turn(turn string) Coord {

	if turn == "R" {
		return Coord{
			row:    facing.column,
			column: -1 * facing.row,
		}
	}

	return Coord{
		row:    -1 * facing.column,
		column: facing.row,
	}
}

func findStartingPosition(boardMap map[Coord]*BoardTile) Coord {
	startingPosition := Coord{1, 0}
	for startingPositionFound := false; !startingPositionFound; {
		startingPosition.column++
		_, startingPositionFound = boardMap[startingPosition]
	}
	return startingPosition
}

func parseInput(input string) (map[Coord]*BoardTile, []Instruction) {

	initialSplit := strings.Split(input, "\n\n")

	return parseBoard(initialSplit[0]), parseInstructions(initialSplit[1])
}

func parseBoard(input string) map[Coord]*BoardTile {

	rows := strings.Split(input, "\n")
	boardMap := make(map[Coord]*BoardTile)
	// Get Initial Coordinates
	for rowI, row := range rows {
		for columnI, tileB := range row {
			tile := string(tileB)
			if strings.TrimSpace(tile) == "" {
				continue
			}

			isWall := tile == "#"

			coord := Coord{row: rowI + 1, column: columnI + 1}

			boardMap[coord] = &BoardTile{
				position: coord,
				isWall:   isWall,
			}
		}
	}

	// Update Up Down Left Right
	for _, boardTile := range boardMap {
		boardTile.neighbours = make(map[Coord]Coord, 0)
		boardTile.neighbours[Coord{-1, 0}] = findAdjacent(boardMap, boardTile.position, Coord{-1, 0})

		boardTile.neighbours[Coord{1, 0}] = findAdjacent(boardMap, boardTile.position, Coord{1, 0})

		boardTile.neighbours[Coord{0, -1}] = findAdjacent(boardMap, boardTile.position, Coord{0, -1})

		boardTile.neighbours[Coord{0, 1}] = findAdjacent(boardMap, boardTile.position, Coord{0, 1})
	}

	cubeEdgeMapping := generateAllEdgeMaps()

	for _, boardTile := range boardMap {
		boardTile.cubeNeighbours = make(map[Coord]EdgeInput, 0)
		boardTile.cubeNeighbours[Coord{-1, 0}] = findCubeAdjacent(boardMap, cubeEdgeMapping, boardTile.position, Coord{-1, 0})
		boardTile.cubeNeighbours[Coord{1, 0}] = findCubeAdjacent(boardMap, cubeEdgeMapping, boardTile.position, Coord{1, 0})
		boardTile.cubeNeighbours[Coord{0, -1}] = findCubeAdjacent(boardMap, cubeEdgeMapping, boardTile.position, Coord{0, -1})
		boardTile.cubeNeighbours[Coord{0, 1}] = findCubeAdjacent(boardMap, cubeEdgeMapping, boardTile.position, Coord{0, 1})
	}

	return boardMap
}

func facingScore(facing Coord) int {

	switch facing {
	case Coord{0, 1}: // right
		return 0
	case Coord{1, 0}: // down
		return 1
	case Coord{0, -1}: //left
		return 2
	case Coord{-1, 0}: // up
		return 3
	}

	panic("Shouldn't be here")

}

func findCubeAdjacent(boardMap map[Coord]*BoardTile, edgeMapping map[EdgeInput]EdgeInput, position Coord, direction Coord) EdgeInput {

	adjCoord := EdgeInput{
		direction: direction,
		position: Coord{
			row:    position.row + direction.row,
			column: position.column + direction.column,
		},
	}

	// If the coordinate doesn't exist,check edgeMapping
	if _, ok := boardMap[adjCoord.position]; !ok {
		adjCoord = edgeMapping[EdgeInput{
			direction: direction,
			position:  position,
		}]
	}

	return adjCoord
}
func findAdjacent(boardMap map[Coord]*BoardTile, position Coord, direction Coord) Coord {

	adjCoord := Coord{
		row:    position.row + direction.row,
		column: position.column + direction.column,
	}

	// If the coordinate doesn't exist, go to the bottom
	if _, ok := boardMap[adjCoord]; !ok {

		for tileFound := true; tileFound; {
			adjCoord.row += -1 * direction.row
			adjCoord.column += -1 * direction.column
			_, tileFound = boardMap[adjCoord]
		}
		adjCoord.row += direction.row
		adjCoord.column += direction.column
	}

	return adjCoord
}

func parseInstructions(input string) []Instruction {
	// Start facing up, then turn right first :)
	input = "R" + input

	instructionsRaw := regexp.MustCompile(`[LR]\d+`).FindAllString(input, -1)

	instructions := make([]Instruction, 0, len(instructionsRaw))

	for _, instructionRaw := range instructionsRaw {

		turn := instructionRaw[:1]
		distanceRaw := string(instructionRaw[1:])

		distance, err := strconv.Atoi(distanceRaw)
		if err != nil {
			panic("Unable to parse number")
		}

		instructions = append(instructions, Instruction{
			turn:     turn,
			distance: distance,
		})
	}

	return instructions
}

// Pretend you start facing up and you turn right first
type Instruction struct {
	turn     string
	distance int
}

type Coord struct {
	row, column int
}

type BoardTile struct {
	position Coord
	isWall   bool

	neighbours     map[Coord]Coord
	cubeNeighbours map[Coord]EdgeInput
}

func generateAllEdgeMaps() map[EdgeInput]EdgeInput {

	edges := []map[EdgeInput]EdgeInput{}

	edges = append(edges, generateEdgeMap(Coord{0, -1}, [2]int{1, 51}, [2]int{50, 51}, Coord{0, 1}, [2]int{150, 1}, [2]int{101, 1}))
	edges = append(edges, generateEdgeMap(Coord{0, -1}, [2]int{51, 51}, [2]int{100, 51}, Coord{1, 0}, [2]int{101, 1}, [2]int{101, 50}))
	edges = append(edges, generateEdgeMap(Coord{-1, 0}, [2]int{101, 1}, [2]int{101, 50}, Coord{0, 1}, [2]int{51, 51}, [2]int{100, 51}))
	edges = append(edges, generateEdgeMap(Coord{0, -1}, [2]int{101, 1}, [2]int{150, 1}, Coord{0, 1}, [2]int{50, 51}, [2]int{1, 51}))
	edges = append(edges, generateEdgeMap(Coord{0, -1}, [2]int{151, 1}, [2]int{200, 1}, Coord{1, 0}, [2]int{1, 51}, [2]int{1, 100}))
	edges = append(edges, generateEdgeMap(Coord{1, 0}, [2]int{200, 1}, [2]int{200, 50}, Coord{1, 0}, [2]int{1, 101}, [2]int{1, 150}))
	edges = append(edges, generateEdgeMap(Coord{0, 1}, [2]int{151, 50}, [2]int{200, 50}, Coord{-1, 0}, [2]int{150, 51}, [2]int{150, 100}))
	edges = append(edges, generateEdgeMap(Coord{1, 0}, [2]int{150, 51}, [2]int{150, 100}, Coord{0, -1}, [2]int{151, 50}, [2]int{200, 50}))
	edges = append(edges, generateEdgeMap(Coord{0, 1}, [2]int{101, 100}, [2]int{150, 100}, Coord{0, -1}, [2]int{50, 150}, [2]int{1, 150}))
	edges = append(edges, generateEdgeMap(Coord{0, 1}, [2]int{51, 100}, [2]int{100, 100}, Coord{-1, 0}, [2]int{50, 101}, [2]int{50, 150}))
	edges = append(edges, generateEdgeMap(Coord{1, 0}, [2]int{50, 101}, [2]int{50, 150}, Coord{0, -1}, [2]int{51, 100}, [2]int{100, 100}))
	edges = append(edges, generateEdgeMap(Coord{0, 1}, [2]int{1, 150}, [2]int{50, 150}, Coord{0, -1}, [2]int{150, 100}, [2]int{101, 100}))
	edges = append(edges, generateEdgeMap(Coord{-1, 0}, [2]int{1, 101}, [2]int{1, 150}, Coord{-1, 0}, [2]int{200, 1}, [2]int{200, 50}))
	edges = append(edges, generateEdgeMap(Coord{-1, 0}, [2]int{1, 51}, [2]int{1, 100}, Coord{0, 1}, [2]int{151, 1}, [2]int{200, 1}))

	mapping := map[EdgeInput]EdgeInput{}

	for _, edge := range edges {

		for k, v := range edge {
			mapping[k] = v
		}

	}

	return mapping

}
func generateEdgeMap(direction Coord, start, end [2]int, resultDirection Coord, resultStart, resultEnd [2]int) map[EdgeInput]EdgeInput {

	// 01 L [1,51] -> [50,51] maps to [150,1] -> [101,1]

	p("Generating edge map", direction, start, end, resultStart, resultEnd)

	inputPoints := generateEdgePoints(start, end)
	resultPoints := generateEdgePoints(resultStart, resultEnd)

	edgeMapping := make(map[EdgeInput]EdgeInput)

	for i, inputPoint := range inputPoints {

		edgeMapping[EdgeInput{
			direction: direction,
			position:  Coord{row: inputPoint[0], column: inputPoint[1]},
		}] = EdgeInput{
			direction: resultDirection,
			position: Coord{
				row:    resultPoints[i][0],
				column: resultPoints[i][1],
			}}
	}
	return edgeMapping
}

func generateEdgePoints(start, end [2]int) [][2]int {

	direction := [2]int{normal(end[0] - start[0]), normal(end[1] - start[1])}

	edgePoints := [][2]int{}

	diff := max(abs(end[0]-start[0]), abs(end[1]-start[1]))

	for i := 0; i <= diff; i++ {
		edgePoints = append(edgePoints, [2]int{start[0] + (direction[0] * i), start[1] + (direction[1] * i)})
	}

	return edgePoints
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func normal(a int) int {
	if a == 0 {
		return a
	}

	return a / abs(a)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

type EdgeInput struct {
	direction Coord // L,R,U,D
	position  Coord
}

/*

01 L [1,51] -> [50,51] maps to [150,1] -> [101,1] now facing R
02 L [51,51] -> [100,51] maps to [101,1] -> [101,50] now facing D
03 U [101,1] -> [101,50] maps to [51,51] -> [100,51] now facing R
04 L [101,1] -> [150,1] maps to [50,51] -> [1,51] now facing R
05 L [151,1] -> [200,1] maps to [1,51] -> [1,100] now facing D
06 D [200,1] -> [200,50] maps to [1,101] -> [1,150] now facing D
07 R [151,50] -> [200,50] maps to [150,51] -> [150,100] now facing U
08 D [150,51] -> [150,100] maps to [151,50] -> [200,50] now facing L
09 R [101,100] -> [150,100] maps to [50,150] -> [1,150] now facing L
10 R [51,100] -> [100,100] maps to  [50,101] -> [50,150] now facing U
11 D [50,101] -> [50,150] maps to [51,100] -> [100,100] now facing L
12 R [1,150] -> [50,150] maps to [150,100] -> [101,100] now facing L
13 U [1,101] -> [1,150] maps to [200,1] -> [200,50] now facing U
14 U [1.51] -> [1,100] maps to [151,1] -> [200,1] now facing R

*/
