/*
--- Day 5: Doesn't He Have Intern-Elves For This? ---
Santa needs help figuring out which strings in his text file are naughty or
nice.

A nice string is one with all of the following properties:

It contains at least three vowels (aeiou only), like aei, xazegov, or
aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde
(dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one
of the other requirements.
For example:

ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a
double letter (...dd...), and none of the disallowed substrings.
aaa is nice because it has at least three vowels and a double letter, even
though the letters used by different rules overlap.
jchzalrnumimnmhp is naughty because it has no double letter.
haegwjzuvuyypxyu is naughty because it contains the string xy.
dvszwmarrgswjxmb is naughty because it contains only one vowel.
How many strings are nice?

--- Part Two ---
Realizing the error of his ways, Santa has switched to a better model of
determining whether a string is naughty or nice. None of the old rules apply,
as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

It contains a pair of any two letters that appears at least twice in the string
without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa,
but it overlaps).
It contains at least one letter which repeats with exactly one letter between
them, like xyx, abcdefeghi (efe), or even aaa.
For example:

qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a
letter that repeats with exactly one letter between them (zxz).
xxyxx is nice because it has a pair that appears twice and a letter that
repeats with one between, even though the letters used by each rule overlap.
uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a
single letter between them.
ieodomkazucvgmuy is naughty because it has a repeating letter with one
between (odo), but no pair that appears twice.

How many strings are nice under these new rules?
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func threeVowelsCheck(s string) bool {
	numVowels := 0
	for _, c := range s {
		switch string(c) {
		case "a":
			numVowels++
		case "e":
			numVowels++
		case "i":
			numVowels++
		case "o":
			numVowels++
		case "u":
			numVowels++
		}
		if numVowels >= 3 {
			return true
		}
	}
	return false
}

func doubleLetterCheck(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func noBadPairCheck(s string) bool {

	for i := 0; i < len(s)-1; i++ {
		switch string(s[i : i+2]) {
		case "ab":
			return false
		case "cd":
			return false
		case "pq":
			return false
		case "xy":
			return false
		}
	}
	return true
}

func twoPairCheck(s string) bool {
	for i := 0; i < len(s)-1-2; i++ {
		for i2 := i + 2; i2 < len(s)-1; i2++ {
			if s[i] == s[i2] && s[i+1] == s[i2+1] {
				return true
			}
		}
	}
	return false
}

func oneLetterGapCheck(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func main() {

	//  Pull in input file
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("file not read")
	}
	strings := strings.Split(string(input), "\n")

	//part1
	niceStrings := 0
	for _, s := range strings {
		if threeVowelsCheck(s) && doubleLetterCheck(s) && noBadPairCheck(s) {
			niceStrings++
		}
	}
	fmt.Println(niceStrings)
	//part2
	niceStrings = 0
	for _, s := range strings {
		if twoPairCheck(s) && oneLetterGapCheck(s) {
			niceStrings++
		}
	}
	fmt.Println(niceStrings)

}
