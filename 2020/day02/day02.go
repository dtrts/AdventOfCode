package main

import (
	"fmt"
	"strings"
)

func main() {
	// input = inputTest

	var inputLines []string

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	minOccurrences, maxOccurrences, validPasswords, valid2Passwords := 0, 0, 0, 0
	password, letter := "", ""

	for _, line := range inputLines {

		fmt.Sscanf(line, "%d-%d %1s: %s", &minOccurrences, &maxOccurrences, &letter, &password)

		// Part 1
		occurrences := strings.Count(password, letter)
		valid := minOccurrences <= occurrences && occurrences <= maxOccurrences
		if valid {
			validPasswords++
		}

		// Part 2
		valid2 := (string(password[minOccurrences-1]) == letter && string(password[maxOccurrences-1]) != letter) || (string(password[minOccurrences-1]) != letter && string(password[maxOccurrences-1]) == letter)
		if valid2 {
			valid2Passwords++
		}
		// fmt.Println("Part2debug", string(password[minOccurrences-1]), string(password[maxOccurrences-1]))
		fmt.Println(line, minOccurrences, maxOccurrences, letter, password, valid, valid2)
	}

	fmt.Println("Valid Passwords", validPasswords, valid2Passwords)
}
