/*
--- Day 10: Elves Look, Elves Say ---
Today, the Elves are playing a game called look-and-say. They take turns making
sequences by reading aloud the previous sequence and using that reading as the
next sequence. For example, 211 is read as "one two, two ones", which becomes
1221 (1 2, 2 1s).

Look-and-say sequences are generated iteratively, using the previous value as
input for the next step. For each step, take the previous value, and replace
each run of digits (like 111) with the number of digits (3) followed by the
digit itself (1).

For example:

1  becomes 11 (1 copy of digit 1).
11 becomes 21 (2 copies of digit 1).
21 becomes 1211 (one 2 followed by one 1).
1211 becomes 111221 (one 1, one 2, and two 1s).
111221 becomes 312211 (three 1s, two 2s, and one 1).
Starting with the digits in your puzzle input, apply this process 40 times. What is the length of the result?

Your puzzle input is 1321131112.

--- Part Two  ---
Neat, right? You might also enjoy hearing John Conway talking about this
sequence (that's Conway of Conway's Game of Life fame).

Now, starting again with the digits in your puzzle input, apply this process
50 times. What is the length of the new result?


*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(look string) (say string) {
	i := 0

	for i < len(look) {

		currentChar := string(look[i])
		numChars := 1

		//  Dela with being the last character
		if i == len(look)-1 {
			say += strconv.Itoa(numChars) + currentChar
			break
		}

		j := i + 1

		for j < len(look) {
			if string(look[j]) == currentChar {
				numChars++
				j++
			} else {
				break
			}
		}

		say += strconv.Itoa(numChars) + currentChar
		i += numChars
	}
	return
}

func main() {
	input := "1321131112"
	/*
		   1      becomes 11 (1 copy of digit 1).
		   11     becomes 21 (2 copies of digit 1).
		   21     becomes 1211 (one 2 followed by one 1).
		   1211   becomes 111221 (one 1, one 2, and two 1s).
		   111221 becomes 312211 (three 1s, two 2s, and one 1).
			 fmt.Println(lookAndSay("1"))
			 fmt.Println(lookAndSay("11"))
			 fmt.Println(lookAndSay("21"))
			 fmt.Println(lookAndSay("1211"))
			 fmt.Println(lookAndSay("111221"))
	*/
	fmt.Println(strings.Repeat("_", 80))

	for i := 1; i <= 50; i++ {
		input = lookAndSay(input)
		fmt.Println(i, len(input))
	}
	/*
		1 14
		2 16
		3 22
		4 34
		5 42
		6 56
		7 76
		8 96
		9 126
		10 172
		11 226
		12 286
		13 382
		14 500
		15 638
		16 860
		17 1118
		18 1420
		19 1890
		20 2466
		21 3182
		22 4178
		23 5460
		24 7066
		25 9224
		26 12100
		27 15684
		28 20474
		29 26744
		30 34772
		31 45370
		32 59214
		33 77146
		34 100396
		35 131124
		36 170930
		37 222442
		38 290526
		39 378526
		40 492982
		41 643280
		42 838402
		43 1092528
		44 1424438
		45 1857356
		46 2420082
		47 3155270
		48 4114514
		49 5361162
		50 6989950
		That took a long ass time.
	*/

}
