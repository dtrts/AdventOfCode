package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

	inputString := string(inputBytes)
	// BOILER PLATE --------------------------------------------------------------------

	rootDirectory := &Directory{
		name:        "",
		path:        "",
		files:       map[string]int{},
		totalSize:   -1,
		directories: make(map[string](*Directory), 0),
		parent:      nil,
	}

	currentDirectory := rootDirectory
	fileLineRegExp := regexp.MustCompile(`^\d+ \w+`)

	for _, line := range strings.Split(inputString, "\n") {

		if strings.HasPrefix(line, "$ cd ") {
			cdCommand := ""
			fmt.Sscanf(line, "$ cd %v", &cdCommand)

			if cdCommand == ".." {

				currentDirectory = currentDirectory.parent

			} else if cdCommand == "/" {

				currentDirectory = rootDirectory

			} else {

				currentDirectory = currentDirectory.directories[cdCommand]

			}

		} else if strings.HasPrefix(line, "dir ") {
			name := ""
			fmt.Sscanf(line, "dir %v", &name)
			newDirectory := &Directory{
				name:        name,
				path:        "",
				files:       map[string]int{},
				totalSize:   -1,
				directories: make(map[string](*Directory), 0),
				parent:      currentDirectory,
			}
			newDirectory.path = getParentPath(newDirectory)

			currentDirectory.directories[name] = newDirectory

		} else if fileLineRegExp.MatchString(line) {
			size, file := 0, ""
			fmt.Sscanf(line, "%d %s", &size, &file)

			currentDirectory.files[file] = size
		}
	}

	updateTotalSize(rootDirectory)

	part1 := part1(rootDirectory)

	totalSize, goalFreeSize := 70000000, 30000000

	currentFreeSize := totalSize - rootDirectory.totalSize

	wantToRemove := goalFreeSize - currentFreeSize

	p("want to remove", wantToRemove)

	part2 := part2(rootDirectory, rootDirectory.totalSize, wantToRemove)

	// BOILER PLATE --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

	fmt.Println("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	elapsed = time.Since(start)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// BOILER PLATE --------------------------------------------------------------------

}

type Directory struct {
	name        string
	path        string
	files       map[string]int
	totalSize   int
	directories map[string](*Directory)
	parent      *Directory
}

func part2(dir *Directory, currentBestSize int, goalSize int) int {
	if dir.totalSize < goalSize {
		return currentBestSize
	}

	newBest := currentBestSize

	for _, childDir := range dir.directories {

		if childDir.totalSize < goalSize {
			continue
		}

		childDirBest := part2(childDir, childDir.totalSize, goalSize)

		if childDirBest < newBest {
			newBest = childDirBest
		}
	}

	return newBest
}

func part1(rootDirectory *Directory) int {

	sum := 0

	if rootDirectory.totalSize <= 100000 {
		sum += rootDirectory.totalSize
	}

	for _, dir := range rootDirectory.directories {
		sum += part1(dir)
	}

	return sum

}

func updateTotalSize(rootDir *Directory) {

	sumOfFiles := 0
	for _, size := range rootDir.files {
		sumOfFiles += size
	}

	sumOfDirs := 0
	for _, dir := range rootDir.directories {
		if dir.totalSize == -1 {
			updateTotalSize(dir)
		}

		sumOfDirs += dir.totalSize
	}

	rootDir.totalSize = sumOfDirs + sumOfFiles
}

func getParentPath(dir *Directory) string {

	dirNames := []string{}

	for dir.parent != nil {
		dir = dir.parent
		dirNames = append(dirNames, dir.name)
	}

	sort.Slice(dirNames, func(i, j int) bool {
		return j < i
	})

	return strings.Join(dirNames, "/")

}

func p(s ...interface{}) {
	fmt.Println(s...)
}
