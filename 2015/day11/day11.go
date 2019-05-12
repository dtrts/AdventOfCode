/*
--- Day 11: Corporate Policy ---
Santa's previous password expired, and he needs help choosing a new one.

To help him remember his new password after the old one expires, Santa has
devised a method of coming up with a password based on the previous one.
Corporate policy dictates that passwords must be exactly eight lowercase letters
(for security reasons), so he finds his new password by incrementing his old
password string repeatedly until it is valid.

Incrementing is just like counting with numbers: xx, xy, xz, ya, yb, and so on.
Increase the rightmost letter one step; if it was z, it wraps around to a, and
repeat with the next letter to the left until one doesn't wrap around.

Unfortunately for Santa, a new Security-Elf recently started, and he has imposed
some additional password requirements:

Passwords must include one increasing straight of at least three letters, like
abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't
count.
Passwords may not contain the letters i, o, or l, as these letters can be
mistaken for other characters and are therefore confusing.
Passwords must contain at least two different, non-overlapping pairs of letters,
like aa, bb, or zz.
For example:

hijklmmn meets the first requirement (because it contains the straight hij) but
fails the second requirement requirement (because it contains i and l).
abbceffg meets the third requirement (because it repeats bb and ff) but fails
the first requirement.
abbcegjk fails the third requirement, because it only has one double letter
(bb).
The next password after abcdefgh is abcdffaa.
The next password after ghijklmn is ghjaabcc, because you eventually skip all
the passwords that start with ghi..., since i is not allowed.
Given Santa's current password (your puzzle input), what should his next
password be?

Your puzzle input is cqjxjnds.

--- Part Two ---
Santa's password expired again. What's the next one?

*/
package main

import (
	"fmt"
	"strings"
)

func incrementString(s string) string {
	outS := []byte(s)
	for i := 7; i >= 0; i-- {
		if outS[i] == 122 {
			outS[i] = 97
		} else {
			outS[i]++
			break
		}
	}
	return string(outS)
}

func checkString(s string) (valid bool) {
	//three characters in a row
	// no i o l
	//two different non over lapping pairs
	threeCheck := false
	twoDoublesCheck := false

	for i, c := range s {
		if string(c) == "i" || string(c) == "o" || string(c) == "l" {

			// fmt.Println("--  invalid character")
			return false

		}

		// three characters in a row
		if i < len(s)-2 {

			// fmt.Println(string(c), string(s[i+1]), string(s[i+2]))
			if s[i+1] == byte(c)+1 && s[i+2] == byte(c)+2 {
				// fmt.Println("Three check true")
				threeCheck = true
			}

		}

		//two doubles
		if i < len(s)-3 {

			// fmt.Println("double start", string(c), string(s[i+1]))

			if byte(c) == s[i+1] {

				// fmt.Println("double one true")

				for i2 := i + 2; i2 < len(s)-1; i2++ {
					if s[i2] == s[i2+1] {
						twoDoublesCheck = true
					}
				}

			}

		}

	}
	if threeCheck && twoDoublesCheck {
		valid = true
	}
	return valid
}

func main() {
	fmt.Println(strings.Repeat("-", 80))
	s := "cqjxxxyz"
	i := 0
	fmt.Println(i, s, checkString(s))

	validOutputs := 0
	for validOutputs < 2 {
		s = incrementString(s)
		if checkString(s) {
			fmt.Println(i, s, checkString(s))
			validOutputs++
		}
		i++
	}

}
