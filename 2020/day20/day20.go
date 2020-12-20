package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// The approach:
// Let us assume that ever edge will match with at most one other (otherwise it will be like the tickets with the rules
// where we will have to check every combination, find the valid ones and start reducing from there. However that ends
// up with 4mil+ combos to start with which i think it just SILLY)
// So lets assume once we have found a match for a border we can move on.
// Need to get a tile, part of a set of 8 varations. 4 rotations, 2 flips.
// Lets make a placing of 19 by 19, and put any tile in the middle. Remove all variations of that tile in the queue to compare.
// Then move through all the remaining tiles and see if you can stick it on the palcement. Ie move through all empty
// squares on the placement,. If that square has a neighbour then test the edges. If we find a match then add the tile
// to the placement, remove the variations of the tile from the queue.
// Then we can reduce the placement down.

// The & operator generates a pointer to its operand.

// i := 42
// p = &i

// The * operator denotes the pointer's underlying value.

// fmt.Println(*p) // read i through the pointer p
// *p = 21         // set i through the pointer p

func main() {
	fmt.Println("hw")
	// input = inputTest

	initTiles := parseInput(input)
	numTiles := len(initTiles)

	// Check square
	rootNumTilesFloat := math.Sqrt(float64(numTiles))
	rootNumTiles := int(rootNumTilesFloat)
	if rootNumTilesFloat-float64(rootNumTiles) != float64(0) {
		panic("Number of tiles doesn't form a square :(")
	} else {
		fmt.Println("Square of sides:", rootNumTiles)
	}
	fmt.Println(initTiles)

	// Make a collection of all tile Variations
	allTileVariations := []*Tile{}
	for _, tile := range initTiles {
		tileVariations := tile.genTileVariations()
		allTileVariations = append(allTileVariations, tileVariations...)
	}
	// tile := parseTile(tileInput)
	// tile.print()
	fmt.Println("Num of variations", len(allTileVariations))

	// Get all edges form all variations and count how many they appear
	edges := make(map[[10]bool]int)
	for _, tile := range allTileVariations {
		edges[tile.topBorder()]++
		edges[tile.bottomBorder()]++
		edges[tile.leftBorder()]++
		edges[tile.rightBorder()]++
	}
	// Define edges on init tiles. An edge is unique if it appears only 4 times sure to the flips and such.
	for edge, num := range edges {
		if num == 4 {
			for _, tile := range initTiles {
				if edge == tile.topBorder() {
					tile.topUnique = true
				}
				if edge == tile.bottomBorder() {
					tile.bottomUnique = true
				}
				if edge == tile.leftBorder() {
					tile.leftUnique = true
				}
				if edge == tile.rightBorder() {
					tile.rightUnique = true
				}
			}
		}
	}

	// fmt.Println(edges)
	part1 := 1
	for _, tile := range initTiles {
		uniqueEdges := 0
		if tile.topUnique {
			uniqueEdges++
		}
		if tile.bottomUnique {
			uniqueEdges++
		}
		if tile.leftUnique {
			uniqueEdges++
		}
		if tile.rightUnique {
			uniqueEdges++
		}
		if uniqueEdges > 2 {
			fmt.Println(tile)
			panic("too many edges are unique")
		}
		if uniqueEdges == 2 {
			part1 *= tile.id
		}
	}

	fmt.Println(part1) //17032646100079

	// End Part 1 attempt ------------------------------------------------------------------------

	placement := initPlacement(rootNumTiles) // placement is an array of nil pointers ready to accept matching tiles. One has to place a corner first

	allTileVariations = []*Tile{}
	for _, tile := range initTiles {
		tileVariations := tile.genTileVariations()
		allTileVariations = append(allTileVariations, tileVariations...)
	}

	fmt.Println(len(allTileVariations))

	// Put a tile in top left
	for _, tile := range allTileVariations {
		if tile.topUnique && !tile.rightUnique && !tile.bottomUnique && tile.leftUnique {
			// tile.print()
			placement[0][0] = tile
			removeTileByID(allTileVariations, tile)
			break
		}
	}

	// fill top row
	for i := 1; i < rootNumTiles-1; i++ {
		for _, tile := range allTileVariations {
			if tile.topUnique && !tile.rightUnique && !tile.bottomUnique && !tile.leftUnique && tile.leftBorder() == placement[0][i-1].rightBorder() {
				placement[0][i] = tile
				removeTileByID(allTileVariations, tile)
				break
			}
		}
	}

	// fill top right corner
	for _, tile := range allTileVariations {
		if tile.topUnique && tile.rightUnique && !tile.bottomUnique && !tile.leftUnique && tile.leftBorder() == placement[0][rootNumTiles-2].rightBorder() {
			placement[0][rootNumTiles-1] = tile
			removeTileByID(allTileVariations, tile)
			break
		}
	}

	// fill left column
	for i := 1; i < rootNumTiles-1; i++ {
		for _, tile := range allTileVariations {
			if !tile.topUnique && !tile.rightUnique && !tile.bottomUnique && tile.leftUnique && tile.topBorder() == placement[i-1][0].bottomBorder() {
				placement[i][0] = tile
				removeTileByID(allTileVariations, tile)
				break
			}
		}
	}

	// fill bottom left corner
	for _, tile := range allTileVariations {
		if !tile.topUnique && !tile.rightUnique && tile.bottomUnique && tile.leftUnique && tile.topBorder() == placement[rootNumTiles-2][0].bottomBorder() {
			placement[rootNumTiles-1][0] = tile
			removeTileByID(allTileVariations, tile)
			break
		}
	}

	// fill bottom row
	for i := 1; i < rootNumTiles-1; i++ {
		for _, tile := range allTileVariations {
			if !tile.topUnique && !tile.rightUnique && tile.bottomUnique && !tile.leftUnique && tile.leftBorder() == placement[rootNumTiles-1][i-1].rightBorder() {
				placement[rootNumTiles-1][i] = tile
				removeTileByID(allTileVariations, tile)
				break
			}
		}
	}

	// fill bottom right corner
	for _, tile := range allTileVariations {
		if !tile.topUnique && tile.rightUnique && tile.bottomUnique && !tile.leftUnique && tile.leftBorder() == placement[rootNumTiles-1][rootNumTiles-2].rightBorder() {
			// tile.print()
			placement[rootNumTiles-1][rootNumTiles-1] = tile
			removeTileByID(allTileVariations, tile)
			break
		}
	}

	// fill right column
	for i := 1; i < rootNumTiles-1; i++ {
		for _, tile := range allTileVariations {
			if !tile.topUnique && tile.rightUnique && !tile.bottomUnique && !tile.leftUnique && tile.topBorder() == placement[i-1][rootNumTiles-1].bottomBorder() {
				placement[i][rootNumTiles-1] = tile
				removeTileByID(allTileVariations, tile)
				break
			}
		}
	}

	// fill centre
	for i := 1; i < rootNumTiles-1; i++ {
		for j := 1; j < rootNumTiles-1; j++ {

			for _, tile := range allTileVariations {
				if !tile.topUnique && !tile.rightUnique && !tile.bottomUnique && !tile.leftUnique && tile.leftBorder() == placement[i][j-1].rightBorder() && tile.topBorder() == placement[i-1][j].bottomBorder() {
					placement[i][j] = tile
					removeTileByID(allTileVariations, tile)
					break
				}
			}
		}
	}

	fmt.Println(placement[0][0].id * placement[rootNumTiles-1][0].id * placement[0][rootNumTiles-1].id * placement[rootNumTiles-1][rootNumTiles-1].id)

	finalImageSize := (rootNumTiles * 8)
	fmt.Println(finalImageSize)

	finalImage := make([][]bool, finalImageSize)
	for i := range finalImage {
		finalImage[i] = make([]bool, finalImageSize)
	}

	for puzzleRowIndex, puzzleLine := range placement {
		for puzzleColIndex, tile := range puzzleLine {
			for tileRowIndex, line := range tile.image {
				for tileColIndex, char := range line {
					if tileRowIndex == 0 || tileRowIndex == 9 || tileColIndex == 0 || tileColIndex == 9 {
						continue
					}
					finalImage[tileRowIndex-1+(puzzleRowIndex*8)][tileColIndex-1+(puzzleColIndex*8)] = char
				}
			}
		}
	}

	seaMonster := make([][]bool, 3)
	seaMonster[0] = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false}
	seaMonster[1] = []bool{true, false, false, false, false, true, true, false, false, false, false, true, true, false, false, false, false, true, true, true}
	seaMonster[2] = []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, false}
	seaMonsterRows, seaMonsterCols := len(seaMonster), len(seaMonster[0])

	printImage(seaMonster)
	fmt.Println(seaMonsterRows, seaMonsterCols)
	printImage(finalImage)

	// find a sea monster
	// Found the orientation manually because its late and i've spent all day on this wtf

	for i := 0; i < 8; i++ {
		if i == 4 {
			flip(finalImage)
		}
		rotate(finalImage, i%4)
		if containsSeaMonster(finalImage, seaMonster) > 0 {
			break
		}
	}

	fmt.Println(containsSeaMonster(finalImage, seaMonster))

	monsterMap := makeImage(finalImageSize, finalImageSize)
	createSeaMonsterMap(finalImage, monsterMap, seaMonster)

	printImage(monsterMap)

	part2 := 0

	for row := 0; row < finalImageSize; row++ {
		for col := 0; col < finalImageSize; col++ {
			if finalImage[row][col] && !monsterMap[row][col] {
				part2++
			}
		}
	}
	fmt.Println(part2)

	// // Now need to make a placement
	// // Array of array of pointers to tiles. Size of placement is (2*squaresize)-1
	// // set middle piece as one tile
	// placement[rootNumTiles][rootNumTiles] = allTileVariations[0]
	// // Remove that tile id from all variations.
	// removeTileByID(allTileVariations, allTileVariations[0])
	// // Go and place the remaining tiles around that one.
	// for safetyDance := 0; safetyDance < 100 && allTileVariations[0] != nil; safetyDance++ {
	// 	// Loop through variations
	// 	// loop through placements
	// 	for _, testingTile := range allTileVariations {
	// 		gotPlaced := placeTileInPlacement(testingTile, placement)
	// 		if gotPlaced {
	// 			// fmt.Println("Tile got placed!")
	// 			// testingTile.print()
	// 			removeTileByID(allTileVariations, testingTile)
	// 			break
	// 		}
	// 	}
	// }

	// tightPlacement := [][]*Tile{}
	// for _, line := range placement {
	// 	newRow := []*Tile{}
	// 	for _, tileP := range line {
	// 		if tileP != nil {
	// 			newRow = append(newRow, tileP)
	// 		}
	// 	}
	// 	if len(newRow) > 0 {
	// 		tightPlacement = append(tightPlacement, newRow)
	// 	}
	// }

	// fmt.Println("Tight Placement", tightPlacement)
	// for _, line := range tightPlacement {
	// 	fmt.Println(line)
	// }

	// for _, line := range placement {
	// 	fmt.Println(line)
	// }
	// End Part 2 attempt ------------------------------------------------------------------------
	// fmt.Println("Part1 ans", tightPlacement[0][0].id*tightPlacement[rootNumTiles-1][rootNumTiles-1].id*tightPlacement[rootNumTiles-1][0].id*tightPlacement[0][rootNumTiles-1].id)

	// fmt.Printf("%v %v %v \n", allTileVariations[0].id, allTileVariations[0].flipped, allTileVariations[0].rotation)
	// fmt.Println(len(allTileVariations))
	// fmt.Println(len(allTileVariations))
	// for _, tile := range allTileVariations {
	// 	// tile.print()
	// 	fmt.Printf("%v %v %v : ", tile.id, tile.flipped, tile.rotation)
	// }
	// fmt.Printf("\n")
	// // tileCopy := tile.copy()
	// // tile.print()
	// // tileCopy.print()

}

func createSeaMonsterMap(finalImage [][]bool, monsterMap [][]bool, seaMonster [][]bool) {
	seaMonsterRows, seaMonsterCols := len(seaMonster), len(seaMonster[0])
	finalImageSize := len(finalImage)
	// find a sea monster
	for row := 0; row < finalImageSize-seaMonsterRows+1; row++ {
		for col := 0; col < finalImageSize-seaMonsterCols+1; col++ {
			testMonsterWindow := makeImage(seaMonsterRows, seaMonsterCols)
			for testMonsterRowIndex, line := range testMonsterWindow {
				for testMonsterColIndex := range line {
					testMonsterWindow[testMonsterRowIndex][testMonsterColIndex] = finalImage[row+testMonsterRowIndex][col+testMonsterColIndex]
				}
			}
			if isSeaMonster(testMonsterWindow, seaMonster) {
				// copy seamonster onto monster map starting at these coords
				for seaMonsterRow, line := range seaMonster {
					for seaMonsterCol, char := range line {
						if char {
							monsterMap[row+seaMonsterRow][col+seaMonsterCol] = true
						}
					}
				}
			}
		}
	}
	return
}

func containsSeaMonster(finalImage [][]bool, seaMonster [][]bool) int {
	seaMonsterRows, seaMonsterCols := len(seaMonster), len(seaMonster[0])
	finalImageSize := len(finalImage)
	numSeaMonsters := 0
	// find a sea monster
	for row := 0; row < finalImageSize-seaMonsterRows+1; row++ {
		for col := 0; col < finalImageSize-seaMonsterCols+1; col++ {
			testMonsterWindow := makeImage(seaMonsterRows, seaMonsterCols)
			for testMonsterRowIndex, line := range testMonsterWindow {
				for testMonsterColIndex := range line {
					testMonsterWindow[testMonsterRowIndex][testMonsterColIndex] = finalImage[row+testMonsterRowIndex][col+testMonsterColIndex]
				}
			}
			if isSeaMonster(testMonsterWindow, seaMonster) {
				// fmt.Println("FoundSea", row, col)
				// printImage(testMonsterWindow)
				numSeaMonsters++
			}
		}
	}
	return numSeaMonsters
}

func isSeaMonster(image [][]bool, seaMonster [][]bool) bool {
	for row, seaLine := range seaMonster {
		for col, seaChar := range seaLine {
			if seaChar == true && image[row][col] == false {
				return false
			}
		}
	}
	return true
}
func printImage(image [][]bool) {
	for _, line := range image {
		for _, char := range line {
			if char {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func makeImage(rows int, cols int) [][]bool {
	newImage := make([][]bool, rows)
	for i := range newImage {
		newImage[i] = make([]bool, cols)
	}
	return newImage
}

func flip(finalImage [][]bool) {
	size := len(finalImage[0])
	newImage := makeImage(size, size)

	for i, line := range finalImage {
		for i2, char := range line {
			newImage[i][size-1-i2] = char
		}
	}
	for rowIndex, line := range newImage {
		for colIndex, char := range line {
			finalImage[rowIndex][colIndex] = char
		}
	}
	return
}

func rotate(finalImage [][]bool, times int) {
	size := len(finalImage[0])

	for rotationAmount := 1; rotationAmount <= times; rotationAmount++ {
		newImage := makeImage(size, size)
		for rowIndex, line := range finalImage {
			for colIndex := range line {
				newImage[rowIndex][colIndex] = finalImage[size-colIndex-1][rowIndex]
			}
		}
		for rowIndex, line := range newImage {
			for colIndex, char := range line {
				finalImage[rowIndex][colIndex] = char
			}
		}

	}
	return
}

func placeTileInPlacement(tile *Tile, placement [][]*Tile) bool {
	placementSize := len(placement[0])
	for row := 1; row < placementSize-2; row++ {
		for col := 1; col < placementSize-2; col++ {
			//up
			if placement[row-1][col] != nil && placement[row-1][col].verticalMatch(tile) {
				placement[row][col] = tile

				// fmt.Println("Up Match")
				// placement[row-1][col].print()
				// tile.print()

				return true
			}
			//down
			if placement[row+1][col] != nil && tile.verticalMatch(placement[row+1][col]) {
				placement[row][col] = tile

				// fmt.Println("Down Match")
				// tile.print()
				// placement[row+1][col].print()

				return true
			}
			//left
			if placement[row][col-1] != nil && placement[row][col-1].horizontalMatch(tile) {
				placement[row][col] = tile

				// fmt.Println("Left Match")
				// placement[row][col-1].print()
				// tile.print()

				return true
			}
			//right
			if placement[row][col+1] != nil && tile.horizontalMatch(placement[row][col+1]) {
				placement[row][col] = tile

				// fmt.Println("Right Match")
				// tile.print()
				// placement[row][col+1].print()

				return true
			}
		}
	}
	return false
}

// Compare tiles in a line
func (tile *Tile) horizontalMatch(rightTile *Tile) bool {
	if tile.rightBorder() == rightTile.leftBorder() {
		return true
	}
	return false
}

func (tile *Tile) verticalMatch(bottomTile *Tile) bool {
	if tile.bottomBorder() == bottomTile.topBorder() {
		return true
	}
	return false
}

// Creates a slice representing a border. Heading
func (tile *Tile) rightBorder() [10]bool {
	border := [10]bool{}
	for i, v := range tile.image {
		border[i] = v[9]
	}
	return border
}

func (tile *Tile) leftBorder() [10]bool {
	border := [10]bool{}
	for i, v := range tile.image {
		// border[9-i] = v[0]
		border[i] = v[0]
	}
	return border
}

func (tile *Tile) topBorder() [10]bool {
	return tile.image[0]
}

func (tile *Tile) bottomBorder() [10]bool {
	return tile.image[9]
}

func removeTileByID(variations []*Tile, tile *Tile) {
	i := 0
	for i < len(variations) {
		if variations[i] != nil && tile.id == variations[i].id {
			// variations = append(variations[:i], variations[i+1:]...)
			copy(variations[i:], variations[i+1:])
			variations[len(variations)-1] = nil // or the zero value of T
			variations = variations[:len(variations)-1]

		} else {
			i++
		}
	}
	return
}

func initPlacement(puzzleSize int) [][]*Tile {
	placement, placementSize := [][]*Tile{}, puzzleSize //(2*puzzleSize)+1
	for i := 0; i < placementSize; i++ {
		newNilRow := []*Tile{}
		for j := 0; j < placementSize; j++ {
			newNilRow = append(newNilRow, nil)
		}
		placement = append(placement, newNilRow)
	}
	return placement
}

func (tile *Tile) genTileVariations() []*Tile {
	// Take tile,
	variations := []*Tile{}
	for i := 0; i < 8; i++ {
		newTile := tile.copy()
		if i >= 4 {
			newTile.flip()
		}
		newTile.rotate(i % 4)
		variations = append(variations, newTile)
	}
	return variations
}

func (tile *Tile) flip() {
	newImage := [10][10]bool{}
	for rowIndex, line := range tile.image {
		for colIndex, char := range line {
			newImage[rowIndex][9-colIndex] = char
		}
	}
	for rowIndex, line := range newImage {
		for colIndex, char := range line {
			tile.image[rowIndex][colIndex] = char
		}
	}
	newLeft := tile.rightUnique
	tile.rightUnique = tile.leftUnique
	tile.leftUnique = newLeft
	tile.flipped = true
	return
}

func (tile *Tile) rotate(times int) {
	for rotationAmount := 1; rotationAmount <= times; rotationAmount++ {
		newImage := [10][10]bool{}
		for rowIndex, line := range tile.image {
			for colIndex := range line {
				newImage[rowIndex][colIndex] = tile.image[10-colIndex-1][rowIndex]
			}
		}
		for rowIndex, line := range newImage {
			for colIndex, char := range line {
				tile.image[rowIndex][colIndex] = char
			}
		}
		newTopUnique := tile.leftUnique
		tile.leftUnique = tile.bottomUnique
		tile.bottomUnique = tile.rightUnique
		tile.rightUnique = tile.topUnique
		tile.topUnique = newTopUnique
		tile.rotation++
	}
	return
}

// Tile ...
type Tile struct {
	id           int
	flipped      bool
	rotation     int
	image        [10][10]bool
	topUnique    bool
	bottomUnique bool
	leftUnique   bool
	rightUnique  bool
}

func (tile *Tile) copy() *Tile {
	newTile := new(Tile)
	newTile.id = tile.id
	newTile.flipped = tile.flipped
	newTile.rotation = tile.rotation
	newTile.image = [10][10]bool{}
	for rowIndex, line := range tile.image {
		for colIndex, char := range line {
			newTile.image[rowIndex][colIndex] = char
		}
	}
	newTile.topUnique = tile.topUnique
	newTile.leftUnique = tile.leftUnique
	newTile.rightUnique = tile.rightUnique
	newTile.bottomUnique = tile.bottomUnique
	return newTile
}

func (tile *Tile) print() {
	fmt.Printf("Tile %v, Flipped:%v, Rotation:%v\n", tile.id, tile.flipped, tile.rotation)
	fmt.Println(tile.topUnique, tile.rightUnique, tile.bottomUnique, tile.leftUnique)
	for _, line := range tile.image {
		for _, char := range line {
			if char {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
}

//
func parseInput(input string) []*Tile {
	tilesRaw := strings.Split(input, "\n\n")
	output := []*Tile{}
	for _, tileRaw := range tilesRaw {
		output = append(output, parseTile(tileRaw))
	}
	return output
}

func parseTile(input string) *Tile {
	newTile, inputSplit := new(Tile), strings.Split(input, "\n")

	// Tile Id
	re := regexp.MustCompile(`^Tile (\d+):$`)
	tileID, _ := strconv.Atoi(re.FindStringSubmatch(inputSplit[0])[1])
	newTile.id = tileID

	// Tile Image
	for rowIndex, line := range inputSplit[1:] {
		for colIndex, char := range strings.Split(line, "") {
			if string(char) == "#" {
				newTile.image[rowIndex][colIndex] = true
			} else {
				newTile.image[rowIndex][colIndex] = false
			}
		}
	}

	return newTile
}

var inputTest = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

var input = `Tile 3011:
#####..#..
#..##.#...
..#....#..
##...#.#.#
.#...#.#.#
..##....#.
##..#..#.#
..##......
..#..#....
####..###.

Tile 1451:
####.#...#
#.....#.#.
..##..#..#
#........#
....##....
..#..#....
.##...##.#
#.....##..
.#.##..###
#.##.#.###

Tile 2003:
...##.##.#
.##...##..
##.#..#...
......##..
..#.#.....
.#####...#
....##.#.#
.#...#.##.
#.........
#....##.#.

Tile 1117:
.#.##.#...
..###..#..
#..#.....#
....#.#...
...#....#.
#..###...#
#.#....#..
..#...#...
#.##..#.#.
####.##...

Tile 2671:
.###.#####
...###..#.
.#.......#
#..##.....
#.........
#.......#.
#.##...#.#
.....##.##
#.........
...#..##.#

Tile 2713:
...#..##..
..#.....#.
.....#.###
#...#.....
..........
......##..
..#.##...#
..#.##...#
..##.##..#
.#...#.#..

Tile 1301:
.##.#.#..#
...###..#.
#........#
#.......#.
.#.......#
###......#
.#.#..###.
..#..#.#.#
##..#...##
..#..#..#.

Tile 1103:
.#...#.#.#
.....##...
..#..#.#..
###.#...##
......#..#
......#..#
.#..##.##.
.##....###
.#.#.#.#.#
#...#.#.##

Tile 1129:
###.###...
...#.....#
#...###...
.###....#.
#..#...#..
..###.....
#...##..##
#.........
###.......
#..#......

Tile 1559:
#...###.#.
.......###
##........
#...#.#..#
......#..#
#.#.#.#..#
.....##...
...#....#.
#........#
.#.#......

Tile 3343:
####....#.
..###.#..#
..........
........##
#..#...###
#......#.#
..........
#.....####
...#....#.
##..#..##.

Tile 2207:
..#####.##
#.##..#.#.
.###...#..
#..#.....#
...#.#.#.#
##........
#.........
.#.##.....
#.##..#.#.
##..#..#.#

Tile 2161:
#.###...##
.......#..
.#..######
#.##....##
#.........
.......#..
##.......#
#.........
........#.
##.####.#.

Tile 2969:
#.#.#.####
###..#.#..
.##...#...
.##.##...#
#.####.###
...###.#.#
...#...##.
#......###
#.....#...
#.#....##.

Tile 3361:
......#...
.##..#...#
....#.....
#....##...
.#..##.#.#
......##.#
..#......#
#.....#...
.##.#...#.
#...###.#.

Tile 3709:
###....#.#
....#.#...
#..#......
.........#
#...#.##..
#.#.#.#.##
#..#..#..#
.##.#.#..#
..#..#...#
..##..###.

Tile 1423:
#.##..##..
....###..#
###.......
......#...
...#.....#
..#.......
####..#.##
..#...#...
.#......#.
.#.###.##.

Tile 3049:
.#.##.##.#
#......#..
#.......#.
......#..#
..........
#.#....#.#
....#...#.
#.........
..##.##..#
#..#...##.

Tile 3907:
.###..##.#
#.#..##...
.#..#....#
.#.##..##.
#.#..#..##
#..#..#..#
##.......#
.#....##.#
...###.#.#
.#.#####..

Tile 2711:
..###.#.##
#..#......
#.........
#.......##
#...#....#
.....#.#.#
##.......#
#........#
#..#...#.#
#.....####

Tile 1307:
#.#..##..#
#..#......
#..#..##..
#.....###.
#.#...##.#
###....#.#
.##.#..#.#
........##
.......#..
##.###.#..

Tile 2351:
#....#....
#..#...#.#
##.##..###
.....##...
.........#
##..##...#
.#......#.
#..#.#.#.#
#.....#...
.#.#.#....

Tile 2647:
#..##.#...
#...#.#...
##..#.....
#...##...#
#.#.....#.
.#..#.....
#.#..#....
....#....#
.#..#..#..
.#..#.....

Tile 2543:
.#..#####.
#..##...#.
#...#...##
##...#...#
###..#..#.
#....#..#.
#.....####
#.#...#..#
......#.#.
.#..#.##.#

Tile 3947:
.##.......
#.#...#...
....##....
#.#...##.#
##.#.....#
#.......#.
#...###.#.
#..####...
#...#.#..#
##.###.#..

Tile 1291:
##.##..#..
..#.......
#.#.#..#.#
..#...##..
..##...#..
...#..#...
#...##....
.#.#..#.##
...#.....#
##.#####.#

Tile 3001:
.##.##..##
........#.
#........#
.#..#.#..#
#..###....
##.#......
#.#.#.....
#...#..#..
#..#....##
.###.##.##

Tile 2243:
..##...###
#...#..#..
##.....#..
......#..#
.#.....#..
.........#
..#......#
##.....#..
....#.##.#
##.###.###

Tile 1741:
###..#####
#.#..#...#
#.......##
#....#..##
#.##.##...
..#.#....#
##..#...##
..#..#....
#..#.#..#.
###.#..###

Tile 1913:
#....####.
#...#.###.
#........#
#..#..#.##
.##....###
#.#......#
#.........
##........
.....#..##
.####..#..

Tile 1069:
..#.#.....
.......#.#
#..####...
.....#....
.#..##.#.#
...#.##..#
#...#....#
.....#...#
.#..##.###
#.#.####..

Tile 3461:
..##.#....
...#...##.
#...#...##
#...##.##.
##......#.
#..##.#..#
#....#.#.#
#.#..#...#
#...#.#..#
.##...##..

Tile 1787:
.##.......
.......#..
#........#
...#......
##...#.#..
#..#..#...
##.#..#..#
#..##..###
#.......##
..#.##.##.

Tile 1901:
.##.#..##.
#......#.#
.......#.#
..........
##..#..#.#
.##.......
.#.#......
..####....
..........
###.##...#

Tile 2267:
.##.####.#
.....#....
#..#.#...#
.......#..
.#........
.#..#.....
..........
##....#...
#..#####.#
..#.#.####

Tile 3917:
....#...#.
#.##....##
##..##.#.#
.#........
......##..
##..#..#.#
#.#.#..#..
..........
#...#....#
#..#....##

Tile 1801:
.#####...#
##..#...##
..#..#.###
.........#
#.......#.
......#..#
........#.
#..#....##
##...#...#
..##.##..#

Tile 1213:
#.#.####..
.....###.#
.#...#....
...#.#....
#........#
....#.....
#......#.#
..........
##.##..##.
#.##.##...

Tile 1733:
#.#####.#.
........##
#.#.....##
##.#..#.##
#.....##..
..#.#...##
#..##.....
.#.#......
...#.###.#
#..#...##.

Tile 2473:
..#...####
.......##.
#..#.#....
.###..##..
#..######.
#..##.....
#.#...#..#
....#.....
...#.#...#
#.###.#...

Tile 2657:
###..#..##
...#..#...
###.#..###
..#.....##
..#..#..#.
.#...#...#
..........
.#......#.
##.#...#..
.##.#.##..

Tile 2957:
#..####.##
#..#.##..#
##....#..#
....####..
.#......#.
.#...#....
...#......
.....##..#
###....#..
..#..#.#..

Tile 3463:
#..##....#
...#..#..#
#........#
##........
..##...#.#
.#.......#
....#.....
#.#.......
..###.....
..#.####.#

Tile 3299:
##.#..#...
..#..##..#
...#..#.#.
...#....#.
.....#....
..........
#...#.....
#.........
........##
.####..###

Tile 3389:
..#..##...
##.#...#.#
.###..####
#...###.#.
..#.#.....
#..#......
##....###.
..........
.....##...
####.....#

Tile 2663:
#....#####
.####.#.##
.#......#.
#.........
#..##.#...
...###....
..#..#...#
..##.##..#
.....###..
...#.####.

Tile 3061:
#...#.##.#
.#......##
..........
...#..#.##
..#..#...#
..##.....#
..#....#.#
...#...#.#
#....####.
#.#.##.##.

Tile 2591:
###..##.#.
..##....#.
..#....#.#
####.##..#
..##.#....
##..#..#.#
.....##...
#.........
...##...##
.##.#.###.

Tile 2333:
..#...#..#
#.#..#....
.#...#.##.
.#..#...#.
....#..#..
#.##.##.#.
#...#....#
#...##.##.
#..#####..
##..####..

Tile 1483:
.#...##.##
.#.#....##
#.#.##.#.#
#....##..#
#.#...#...
.#..##....
#..#..#...
...#...#.#
.#..#.....
#..#...#.#

Tile 1543:
.####..#.#
........#.
###.......
.#....#...
...#......
..........
##...##.##
#.#....#..
.#..#...##
#.##.#...#

Tile 2879:
####...###
.##.#.##..
#...#.#...
..#....###
.....#.#.#
..#.......
#..#.....#
.##...#..#
#.##..#...
##...##...

Tile 3449:
..##.#...#
#.##.##.##
......#...
.........#
####.##...
.#..#...#.
.#.......#
.##..##..#
........#.
....######

Tile 1373:
#.#.#.#...
##....##..
....#.....
...##.##..
#.#...###.
..#......#
#..##.#..#
#...##.#..
.#....##..
..##.#.#..

Tile 3391:
#####.#.#.
##...#...#
#.....#.##
##..#.....
.##....#..
.#...#...#
.####.#...
##...#...#
#...#.....
#...###.##

Tile 1453:
##..#.....
..#..#....
#...####..
#...#..##.
#.#.#...##
..##.###..
#.##.#.#..
.....##...
#..#......
#..#..#.#.

Tile 1621:
.#.##...##
#..#....##
#.#.#.....
##..#.#.#.
...#.#....
..#.#...#.
..##.#...#
...#.##.##
....#..#..
..##..#.##

Tile 1493:
.#...#.#..
##..#.....
..........
##..##..#.
#.###....#
#....#....
..#.###...
#...#.....
#.##......
..#.....##

Tile 2137:
.#.#..##.#
#.........
#...###...
#..#...##.
#..#...#..
.....##...
#.##....#.
#.#...#.##
##.#.#.##.
#..##.#..#

Tile 2837:
..##.#....
#.##......
..##......
#...#....#
........##
....#...#.
....##....
..##.....#
...#.##..#
..#.#.#..#

Tile 1237:
#.##.####.
#.......##
..#......#
##..#....#
#...#.....
#......#.#
#.#.##...#
##.....#.#
.#...#...#
##...#....

Tile 1811:
##...#..#.
....##.#..
.........#
###......#
#......#..
.......###
........##
#...#...#.
#...#...##
.##.##....

Tile 1979:
.#.....#.#
.....#..##
.##.##....
##.#.....#
#....#..##
#......#.#
.####.##..
.......#.#
#...#.#..#
##...##.##

Tile 3733:
.###.####.
##....#.#.
##.#.#.#..
#......#..
.....###..
.#...#...#
#.#.....##
....#.#.#.
#.##....#.
..###.#.#.

Tile 1201:
####.##...
....#..#..
#..#....##
...##...#.
..#......#
..##..#..#
##..#.....
#......#.#
#....#...#
#.##...###

Tile 1637:
...###.###
.....#..#.
#...###.##
#........#
#......#..
...#...###
#...####..
#....#....
....#.....
##..#.#...

Tile 3911:
#...#.##..
...###..#.
#.........
##..#.....
...#.....#
###.#.##..
.###...##.
#...#...##
...#.....#
.##.#.....

Tile 1723:
.#.##....#
....#....#
...#.#.#.#
#.###.###.
.....#.#..
#....#.###
#.#..#....
#......#.#
....##...#
#..#......

Tile 1297:
.###.#####
.....##..#
...#..#...
###....##.
.#...#...#
..#.#....#
##...#.#.#
.#........
##...##.#.
#.##.....#

Tile 3631:
##....####
.#..##.##.
.#...#...#
..#.#.#...
#.....##..
...##....#
.........#
........#.
#........#
.........#

Tile 3691:
##....###.
#.#.#.....
.#....#..#
..#....#..
.....#...#
.....#..#.
.........#
.##.#...##
#.#......#
...####.#.

Tile 1579:
.#..##.##.
#.......##
..#.......
#..###...#
#.#.......
##.......#
#....#...#
##..#....#
.#...##..#
##.#.#....

Tile 1409:
..#.#.##.#
#.#.###..#
..#...#.##
#..###..#.
...#..#..#
...#...#..
....#.####
#...######
....#.#.#.
.#.#.#####

Tile 1873:
#.####..#.
#.#...#...
...#..##.#
..#.......
#.#.#....#
#..##.....
......#.##
..........
...#...###
##.##...##

Tile 3671:
...#.#.##.
...##.....
#.#..#.###
....#....#
##........
....##...#
....#.....
.##.#....#
..#..#....
##..#.###.

Tile 3697:
....##..##
##.....#.#
..#.....##
.....#....
...#.#..##
#....#...#
.#.##..##.
......#.##
...#...##.
...#.#.#..

Tile 2777:
##........
.##.#.....
#........#
.#..##.###
...#..#..#
######....
##.###..##
##.....#.#
.........#
.####.#..#

Tile 3323:
#####.#...
.##...##..
.#.###..#.
...#......
..#.#....#
#.#...#...
#.#....##.
#.....#.##
#.......##
####..###.

Tile 3259:
.###.#..##
#...#.....
...##..#.#
##.##...##
...###..##
##..##..#.
#.#.#....#
##......##
....#...#.
..#.##..##

Tile 2063:
...#####.#
#.#...#..#
.#.....###
.#....##.#
.##....#.#
##......#.
#...#..#.#
#.#......#
###.....#.
###.###.#.

Tile 2707:
###..#..#.
.....#...#
#..#..##..
#........#
...#..#..#
#.#.#.#..#
.....#..##
....##....
#..##.#...
##..#....#

Tile 3089:
#..#.###..
.....##...
##.......#
###......#
....#....#
##.#.....#
#........#
#.#....#..
#...#.##..
.#.##.###.

Tile 2203:
#....####.
.......#.#
#........#
...#....##
......###.
.......#..
....##.#.#
##......#.
#.#.......
#####..#..

Tile 2803:
##.#.#....
##........
...#.....#
#........#
..###....#
......#...
#.......#.
#.......#.
###.......
###.#.#...

Tile 3607:
####.##..#
......#...
#...#..#.#
#........#
#.....#.#.
.....#.#.#
#........#
#...#.....
#....##..#
.#.#...#..

Tile 3517:
.##...###.
..#....#..
#.....#.#.
#...#.#.##
...###...#
.##.#.##.#
...#....#.
..#......#
#....#.#..
#..#..#.##

Tile 3613:
.###..##..
..#..#...#
#.#......#
#....#..##
..........
..#..#..#.
#........#
.#........
#....#....
.#...##.##

Tile 3823:
.##.#..##.
##.#.#...#
..#..#....
.#####...#
##....#..#
..#..#....
...#......
#.#.##....
#......#..
##...#..#.

Tile 1471:
####..#.##
#.#.#.....
#..#.....#
#...#...##
..#..#...#
##.#..#.#.
.#.#.....#
#..#...#..
#......#..
#...###.##

Tile 3779:
#.##...#.#
...##....#
.#..##..##
....##....
##.#.....#
.......#..
..##...#..
#.###..#..
#..###.##.
...#....##

Tile 3623:
..#..#..#.
#.##.#..##
######..#.
.........#
.##...#.#.
.........#
#....#...#
...#..#.#.
.#..##...#
......####

Tile 3919:
###.#.....
#.#.#....#
.#...#.#.#
#.#.......
#.#.#....#
......#...
#.#....###
..#..#....
#...#.....
##.##.##.#

Tile 2477:
#..###.#.#
.....#.###
...#...###
.........#
.#..##.#.#
.#......##
..##.#..#.
..#.......
##....#..#
#.##.##.##

Tile 1091:
.#.#.#...#
.##.#..#..
#....#....
........##
#...#....#
...##.#.##
#........#
#..##..#.#
##...##..#
#.#.#...#.

Tile 3929:
##.#...#..
#.#....###
.......#.#
##..#.#..#
#.#.......
..........
#...#.#.#.
#.#.#..###
....#.....
....##...#

Tile 1789:
#...#####.
#....#....
#..#...#.#
.....##.##
######.###
.........#
#..#......
...#....##
.......###
..####.#..

Tile 1667:
#.####.##.
#.....#...
#....#.#..
##..#..###
..##....#.
......#...
.##.#.....
##..#.....
.#.......#
..##...###

Tile 1847:
#.#..#....
..#.#.#..#
#####..##.
#.........
#...##....
.#..#.....
..#.#....#
.......#..
...##...##
..##..#.##

Tile 3119:
##..##..#.
##..#.....
.....#....
#..#..##.#
#......#..
.#...##..#
..#.##.##.
...###.#.#
..........
..#..#....

Tile 1607:
.##..####.
..###....#
###.......
.##..####.
...#.#...#
#.....#..#
.....#.#..
#..##.#..#
..#.#.#...
.#.##.#...

Tile 2089:
...#.#...#
##.#..#..#
#..#.....#
#..#..##.#
...#.....#
#........#
.#.#..#...
..........
##.......#
.##.#...##

Tile 1951:
###.#..###
#.#.....#.
.##.##....
#.#.#..###
#.#...#...
.........#
#..#...##.
.....#....
#.#.....##
##...#.##.

Tile 3373:
##.#.....#
..#...##.#
.#..#.....
##..#...##
#.##.....#
#.#.......
..#.......
#...#....#
.#.......#
#.###.####

Tile 1619:
###...###.
......#..#
#........#
..........
....#.#..#
#.#.#..##.
....##..##
......##..
......#.#.
##.##.#...

Tile 2069:
##..#.#.##
##.......#
#........#
#....#....
.#.#..#..#
.##..#..#.
...#.#.#..
..........
.#....##.#
.#####.#.#

Tile 2801:
##...#....
##....#...
.....#....
#..#.#..##
#.#....##.
#.#.##...#
#......#..
..##.#...#
..........
##.##..#.#

Tile 1753:
##.####.#.
.#.....#.#
###.#.#..#
......#..#
##....##..
....#....#
.#.....##.
.....##..#
#..#....#.
#..##..#..

Tile 3659:
###.##.#.#
.......#..
...#....#.
###.#...#.
#....#.#..
....##...#
##.#..#.#.
##......##
..........
#......#..

Tile 1097:
#.#.##..#.
#......#..
.........#
...##....#
....###..#
#...#.#..#
.#.#.#....
##...#.#..
#....#.#..
#....#..##

Tile 1583:
#####..##.
.#......#.
#...#....#
..#.....#.
....#...##
#.#...#...
.#####..##
....##.#..
.........#
..#...##..

Tile 1759:
.###..#.##
...#.#...#
....#...##
#.##...#.#
.#..##....
#....###..
#..#.##...
.#..#..#..
.......#.#
#..##.....

Tile 3203:
.##...###.
#..#.....#
..#..#...#
###......#
#..#......
..........
#..##.#..#
.#.#..#...
....#...#.
###...##.#

Tile 3877:
#.#.#.#.#.
#.....##..
....#....#
...#..#.#.
##...#.###
##.....#.#
#..#......
##........
.......#.#
.######..#

Tile 3319:
.#..#....#
#..#......
.#.......#
##.....#..
#..#..#...
##.###...#
#.........
..#....##.
.......#.#
##.#.#..##

Tile 1999:
#####.#..#
#........#
#....##...
......#.#.
...#......
..........
#....#..##
#.##.....#
##...#...#
..##.#..##

Tile 1381:
###.#.....
##..#.#.##
........#.
#......#.#
..#..##..#
##.#....#.
#......#..
#....###.#
#..#..#.##
###.##..##

Tile 3137:
###.#.#.##
#...###..#
...#....#.
...##...##
....##...#
#..##.#...
..#...#.##
#..#..#...
...#....#.
#......#..

Tile 2843:
##..##....
.##....#..
.........#
#......#..
..#..#.#.#
#....#.#.#
##....#..#
...#....#.
#........#
.####.#...

Tile 3413:
.#.##..##.
...#.....#
#.#...#...
...#.#...#
#..####.#.
...#.#....
##....##.#
..#..#..#.
#....#....
.###...#.#

Tile 1163:
..#...##..
##......#.
#..###...#
..#.##..#.
#.........
...##.##..
.....##..#
#....#....
........##
#...######

Tile 2381:
#.##.....#
####......
....#...##
#........#
#.#....#.#
#......#.#
##....#...
..........
.#.....#.#
..#..###..

Tile 1231:
.##.##..#.
#..##.##..
#.#.##....
#..#..##.#
#..#....#.
#...##...#
#.........
##...#....
.....#..#.
####.###.#

Tile 1721:
.#.#..###.
##.#..####
.##.#..#..
###...#.##
.##..##.##
#.........
....#.#.##
#.##...#.#
.#.....#.#
###.###.#.

Tile 2797:
#.###.##.#
......#...
#.........
.......#..
#....#...#
#.##......
....####.#
#....####.
#.#.#....#
##.#..#.#.

Tile 2437:
.##.#.###.
#.....####
..#####..#
.#.#.#....
##...##...
#.........
#..##....#
..##.#..#.
##....###.
.#.#####..

Tile 3739:
##.##.###.
#....#.#..
.....#..#.
#..#..##..
..##.#.#.#
#.......#.
##..#.#..#
#....#..#.
#....#...#
.#....#...

Tile 3559:
#.....#.#.
..........
#.#....#..
.........#
.#...#...#
........#.
...##...#.
.#.#.###..
#...##....
#..###.###

Tile 1181:
###.#....#
#...##.#..
.#..#####.
....####.#
#.........
##.#....##
#...#..#.#
......#...
..#....#.#
##..##...#

Tile 2357:
#...#..##.
#....#..#.
.#..#....#
..........
.#.....#..
#.........
#..#..#...
#........#
...###...#
##.#..#.#.

Tile 1657:
.#..##.#..
#..#.#..#.
#....###..
.####..#..
..#.....#.
.........#
....#....#
...#......
...#....##
#..##.#..#

Tile 3433:
.###.#.###
...##.....
.........#
##..##...#
#.#.....##
.#...##...
####.#..##
.#....#..#
...##..#.#
.....#.##.

Tile 2383:
##........
..##......
##.#...#.#
#.......##
##....#...
#.......#.
#..#...#..
...#.....#
....##...#
..#.#.....

Tile 2917:
..###..#.#
.........#
#....##..#
#........#
......#.##
.#....#...
#........#
......#.##
##.....#..
..#####...

Tile 2887:
##.#.###.#
..#......#
........##
......#.#.
#..#..#.#.
#...#..#..
....##....
...#..##..
..#.#.....
.##.##....

Tile 1217:
#..##..#.#
#......#.#
.##.....#.
...#.....#
#..##..#.#
..#.......
#......#.#
#.#.#....#
..#.......
#..##..##.

Tile 3191:
.........#
.....#.###
##....####
#.......#.
..##......
##..##....
#.#..#.#..
##.#..##..
..#......#
##..#####.

Tile 1061:
#..##..###
#..#......
.##.......
#.......#.
#.........
....#.....
..........
#..#..#..#
#...###.#.
.#..##..##

Tile 2903:
.##.##..##
#...#....#
...#..###.
#.....#..#
#..##..#..
#......##.
#......##.
#.##.##...
##.#.#...#
#.#..#.##.

Tile 2417:
#.....##..
#..#......
..##.....#
#.#.#.....
##.##..#..
##....#..#
#.........
.##.#.#...
.#..#.##.#
..#..#.###

Tile 1039:
#.#.##..#.
.....#...#
#...#.....
#.........
...#....#.
...###...#
#...#....#
#....#....
#.#.......
...#..##..

Tile 1109:
.##..####.
#.#....#..
##.#...#..
#.#......#
#####..#.#
#......#..
........#.
........##
#..#.#....
#.####....

Tile 2819:
##...##.#.
#.....##.#
.#.#.#....
#..#.....#
.#........
#....##...
##....#...
##.##.#..#
#..#.#...#
..###.##.#

Tile 2729:
###..#####
..###.#..#
#..##....#
.....#....
....#..#.#
.#........
#....##..#
..#.#....#
#.###....#
.###..#...

Tile 3533:
.#.#...#.#
.........#
##..#....#
###.###...
#....#...#
#......#.#
.###......
.##...#...
##.##.##.#
...##..###`
