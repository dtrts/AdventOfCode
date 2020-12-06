package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// input = inputTest
	inputLines := []string{}
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n\n") {
		inputLines = append(inputLines, strings.Trim(strings.Replace(line, "\n", "", -1), " "))
	}
	fmt.Println(inputLines)

	sum := 0

	for _, line := range inputLines {
		// fmt.Println(strings.Join(unique(line), ""))
		sum += len(unique(line))
	}

	fmt.Println(sum)

	// Part 2
	inputGroups := [][]string{}
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n\n") {
		inputGroups = append(inputGroups, strings.Split(line, "\n"))
	}

	inputMaps := make([]map[string]bool, 0, len(inputGroups))

	totalAnswers := 0

	for _, group := range inputGroups {
		m := make(map[string]bool)

		for _, s := range group[0] {
			m[string(s)] = true

			for _, next := range group[1:] {
				if !strings.Contains(next, string(s)) {
					delete(m, string(s))
				}
			}

		}
		totalAnswers += len(m)
		inputMaps = append(inputMaps, m)
	}

	fmt.Println(inputGroups, inputMaps, totalAnswers)

}

func unique(s string) []string {
	u := []string{}
	m := make(map[string]bool)

	for _, val := range s {
		m[string(val)] = true
	}

	for key := range m {
		u = append(u, key)
	}

	sort.Strings(u)

	return u
}
