package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("Starting...")

	testInputPtr := flag.Bool("testInput", false, "a bool")
	flag.Parse()

	if *testInputPtr {
		fmt.Println("Using test Input")
		input = inputTest
	}

	pairs, rules, startChar, endChar := parseInput(input)

	part1, part2 := 0, 0

	for i := 1; i <= 40; i++ {
		pairs = insertPairs(pairs, rules)

		if i == 10 {
			part1 = score(countElements(pairs, startChar, endChar))
		}
		if i == 40 {
			part2 = score(countElements(pairs, startChar, endChar))
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	log.Printf("Duration: %s", time.Since(start))
}

func insertPairs(pairs map[string]int, rules map[string][2]string) map[string]int {

	newPairs := map[string]int{}

	for pairs_key, pairs_count := range pairs {

		newPairs[rules[pairs_key][0]] += pairs_count
		newPairs[rules[pairs_key][1]] += pairs_count
	}

	return newPairs
}

func score(elementCount map[string]int) int {

	max, min := 0, 0
	for _, v := range elementCount {
		max, min = v, v
		break
	}

	for _, v := range elementCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min

}

func countElements(pairs map[string]int, startChar string, endChar string) map[string]int {

	counter := map[string]int{}

	for pair, count := range pairs {

		for _, element := range strings.Split(pair, "") {
			counter[element] += count
		}
	}

	counter[startChar] += 1
	counter[endChar] += 1

	for k, _ := range counter {
		counter[k] /= 2
	}
	return counter
}

func parseInput(input string) (map[string]int, map[string][2]string, string, string) {

	inputs := strings.Split(input, "\n\n")
	startChar, endChar := "", ""
	pairs := map[string]int{}
	chars := strings.Split(inputs[0], "")
	for i, c := range chars {
		if i == 0 {
			startChar = c
			continue
		}
		if i == len(chars)-1 {
			endChar = c
		}

		pairs[strings.Join([]string{chars[i-1], c}, "")] += 1
	}

	rules := strings.Split(inputs[1], "\n")

	rulesOut := map[string][2]string{}

	for _, rule := range rules {
		ruleSplit := strings.Split(rule, " -> ")
		ruleInputSplit := strings.Split(ruleSplit[0], "")
		rulesOut[ruleSplit[0]] = [2]string{ruleInputSplit[0] + ruleSplit[1], ruleSplit[1] + ruleInputSplit[1]}
	}

	return pairs, rulesOut, startChar, endChar
}

var inputTest string = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

var input string = `BCHCKFFHSKPBSNVVKVSK

OV -> V
CO -> V
CS -> O
NP -> H
HH -> P
KO -> F
VO -> B
SP -> O
CB -> N
SB -> F
CF -> S
KS -> P
OH -> H
NN -> O
SF -> K
FH -> F
VV -> B
VH -> O
BV -> V
KF -> K
CC -> F
NF -> H
VS -> O
SK -> K
HV -> O
CK -> K
VP -> F
HP -> S
CN -> K
OB -> H
NS -> F
PS -> S
KB -> S
VF -> S
FP -> H
BB -> N
HF -> V
CH -> N
BH -> F
KK -> B
OO -> N
NO -> K
BP -> K
KH -> P
KN -> P
OF -> B
VC -> F
NK -> F
ON -> O
OC -> P
VK -> O
SH -> C
NH -> C
FB -> B
FC -> K
OP -> O
PV -> V
BN -> V
PC -> K
PK -> S
FF -> C
SV -> S
HK -> H
NB -> C
OK -> C
PH -> B
SO -> O
PP -> F
KV -> V
FO -> B
FN -> H
HN -> C
VB -> K
CV -> O
BC -> C
CP -> S
FS -> S
KP -> V
BS -> V
BK -> B
PN -> C
PF -> S
HO -> V
NC -> N
SS -> N
BO -> P
BF -> N
NV -> P
PB -> K
HB -> H
VN -> H
FV -> B
FK -> K
PO -> S
SC -> S
HS -> S
KC -> F
HC -> S
OS -> K
SN -> N`
