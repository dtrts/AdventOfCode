package main

import (
	"strconv"
	"strings"
)

func topValues(ss [][]string) (tops string) {
	for _, s := range ss {
		tops += lastValue(s)
	}
	return tops
}

func lastValue(s []string) string {
	return s[len(s)-1]
}

func popN(s *[]string, N int) []string {
	if len(*s) == 0 {
		return []string{}
	}

	lastIndex := len(*s) - N
	chars := (*s)[lastIndex:]
	*s = (*s)[:lastIndex]
	return chars
}

func parseDiagram(diagram string) [][]string {
	diagramLines := strings.Split(diagram, "\n")

	columns := strings.Fields(lastValue(diagramLines))
	numColumns, _ := strconv.Atoi(lastValue(columns))

	stacks := make([][]string, numColumns)

	for lineIndex := len(diagramLines) - 2; lineIndex >= 0; lineIndex-- {
		line := diagramLines[lineIndex]

		for stackIndex := 0; stackIndex < numColumns; stackIndex++ {
			// Values found at char 1,5,9,13 etc...
			charIndex := (stackIndex * 4) + 1

			if charIndex > len(line) {
				break
			}

			char := string(line[charIndex])

			if strings.TrimSpace(char) != "" {
				stacks[stackIndex] = append(stacks[stackIndex], char)
			}
		}
	}

	return stacks
}
