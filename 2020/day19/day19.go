package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	rules, messages := parseInput(input)

	regexRules := make(map[string]string)

	genRegexRule("0", rules, regexRules, -1, -1)

	testRegex := regexp.MustCompile("^" + strings.ReplaceAll(regexRules["0"], " ", "") + "$")

	part1 := 0
	for _, message := range messages {
		if testRegex.MatchString(message) {
			part1++
		}
	}

	fmt.Println("Part1:", part1)

	regexRules2 := make(map[string]string)

	genRegexRule("0", rules, regexRules2, 0, 0)

	testRegex = regexp.MustCompile("^" + strings.ReplaceAll(regexRules2["0"], " ", "") + "$")

	part2 := 0
	for _, message := range messages {
		if testRegex.MatchString(message) {
			part2++
		}
	}
	fmt.Println("Part2:", part2)
}

func genRegexRule(ruleIndex string, rules map[string]string, regexRules map[string]string, eightCount int, elevenCount int) {
	if _, exists := regexRules[ruleIndex]; exists {
		return
	}

	ruleValue := rules[ruleIndex]

	if ruleIndex == "8" && 0 <= eightCount && eightCount <= 10 {
		eightCount++
		ruleValue = "(42 | 42 8)"
	}
	if ruleIndex == "11" && 0 <= elevenCount && elevenCount <= 10 {
		elevenCount++
		ruleValue = "(42 31 | 42 11 31)"
	}

	re := regexp.MustCompile(`\d+`)

	ruleValue = re.ReplaceAllStringFunc(ruleValue, func(ruleIndex string) string {

		genRegexRule(ruleIndex, rules, regexRules, eightCount, elevenCount)

		return regexRules[ruleIndex]
	})

	regexRules[ruleIndex] = ruleValue

	return
}

func parseInput(input string) (map[string]string, []string) {
	bigSplit := strings.Split(input, "\n\n")
	return parseRules(bigSplit[0]), parseMessages(bigSplit[1])
}

func parseRules(input string) map[string]string {
	output := make(map[string]string)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lineSplit := strings.Split(line, ": ")
		output[lineSplit[0]] = strings.Trim(lineSplit[1], "\" ")

		if strings.ContainsAny(output[lineSplit[0]], "0123456789") {
			output[lineSplit[0]] = "(" + output[lineSplit[0]] + ")"
		} else {
			output[lineSplit[0]] = strings.ReplaceAll(output[lineSplit[0]], "\"", "")
		}

	}
	return output
}

func parseMessages(input string) []string {
	return strings.Split(input, "\n")
}

var inputTest1 = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

var inputTest2 = `0: 8 11
1: "a"
2: 1 24 | 14 4
3: 5 14 | 16 1
4: 1 1
5: 1 14 | 15 1
6: 14 14 | 1 14
7: 14 5 | 1 21
8: 42
9: 14 27 | 1 26
10: 23 14 | 28 1
11: 42 31
12: 24 14 | 19 1
13: 14 3 | 1 12
14: "b"
15: 1 | 14
16: 15 1 | 14 14
17: 14 2 | 1 7
18: 15 15
19: 14 1 | 14 14
20: 14 14 | 1 15
21: 14 1 | 1 14
22: 14 14
23: 25 1 | 22 14
24: 14 1
25: 1 1 | 1 14
26: 14 22 | 1 20
27: 1 6 | 14 18
28: 16 1
31: 14 17 | 1 13
42: 9 14 | 10 1

aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
ababaaaaaabaaab
ababaaaaabbbaba
abbbbabbbbaaaababbbbbbaaaababb
abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
baabbaaaabbaaaababbaababb
babaaabbbaaabaababbaabababaaab
babbbbaabbbbbabbbbbbaabaaabaaa
bbabbbbaabaabba
bbbababbbbaaaaaaaabbababaaababaabab
bbbbbbbaaaabbbbaaabbabaaa`

var input = `0: 8 11
1: 41 41 | 48 48
2: 48 85 | 41 128
3: 48 118 | 41 25
4: 131 48 | 70 41
5: 41 110
6: 48 100 | 41 132
7: 41 41
8: 42
9: 89 41 | 100 48
10: 108 48 | 53 41
11: 42 31
12: 41 18 | 48 98
13: 43 41 | 69 48
14: 9 48 | 93 41
15: 41 20 | 48 63
16: 48 5 | 41 45
17: 62 48 | 26 41
18: 48 48
19: 48 72 | 41 18
20: 61 48 | 57 41
21: 48 83 | 41 51
22: 103 48 | 18 41
23: 84 41 | 27 48
24: 48 109 | 41 38
25: 48 81 | 41 7
26: 41 72 | 48 81
27: 100 41 | 7 48
28: 41 132 | 48 120
29: 116 48 | 3 41
30: 89 41 | 18 48
31: 48 133 | 41 127
32: 41 96 | 48 95
33: 120 41
34: 48 41 | 48 48
35: 123 48 | 91 41
36: 48 81 | 41 100
37: 41 7 | 48 120
38: 2 48 | 77 41
39: 41 67 | 48 41
40: 81 48 | 7 41
41: "a"
42: 41 68 | 48 105
43: 52 48 | 112 41
44: 101 48 | 114 41
45: 7 48
46: 48 29 | 41 82
47: 120 41 | 81 48
48: "b"
49: 81 41 | 7 48
50: 7 41 | 7 48
51: 41 1 | 48 81
52: 41 73 | 48 104
53: 54 48 | 119 41
54: 41 36 | 48 55
55: 48 1 | 41 18
56: 34 48 | 39 41
57: 41 92 | 48 80
58: 41 22 | 48 111
59: 41 47 | 48 125
60: 30 41 | 86 48
61: 41 37 | 48 6
62: 100 48 | 98 41
63: 41 59 | 48 17
64: 48 44 | 41 74
65: 16 48 | 58 41
66: 41 126 | 48 121
67: 48 | 41
68: 64 41 | 24 48
69: 106 41 | 32 48
70: 87 48 | 115 41
71: 67 34
72: 48 48 | 41 67
73: 1 48 | 103 41
74: 41 107 | 48 124
75: 89 41 | 39 48
76: 48 14 | 41 35
77: 40 48 | 47 41
78: 41 65 | 48 66
79: 67 1
80: 120 48 | 98 41
81: 48 48 | 67 41
82: 23 48 | 21 41
83: 67 89
84: 18 48
85: 120 41 | 89 48
86: 48 89 | 41 39
87: 103 48 | 72 41
88: 7 48 | 89 41
89: 41 41 | 41 48
90: 81 41 | 34 48
91: 34 48
92: 48 103 | 41 34
93: 110 41 | 34 48
94: 41 4 | 48 76
95: 34 41 | 103 48
96: 100 48 | 103 41
97: 41 120 | 48 34
98: 41 48
99: 97 48 | 92 41
100: 48 41 | 67 48
101: 129 48 | 93 41
102: 48 62 | 41 50
103: 67 67
104: 48 1 | 41 81
105: 46 41 | 78 48
106: 88 41
107: 79 41 | 33 48
108: 102 41 | 60 48
109: 48 113 | 41 99
110: 48 41 | 41 41
111: 103 67
112: 48 71 | 41 123
113: 49 41 | 56 48
114: 41 117 | 48 73
115: 67 132
116: 41 93 | 48 122
117: 48 120 | 41 39
118: 89 41 | 132 48
119: 48 12 | 41 19
120: 48 41
121: 48 49 | 41 47
122: 48 120 | 41 89
123: 110 48 | 98 41
124: 83 48 | 130 41
125: 41 120 | 48 132
126: 75 48 | 27 41
127: 41 94 | 48 10
128: 48 72 | 41 1
129: 48 110 | 41 120
130: 103 48 | 89 41
131: 90 41 | 28 48
132: 41 48 | 48 41
133: 48 15 | 41 13

aaaaaaaaabaabaababbaabbb
aaaaaaaaabbaabaaaabaabbbabbabbab
aaaaaaabbbbbaaabbabbaabaababbababababbbaaabbabbbbaaaabaa
aaaaaabbababbbabbaababbb
aaaaaabbbabaaaabbabbaaba
aaaabaababaaabaabbaaaaab
aaaabaabababbabbaaabaabbabaababbbaaaabbbbbbababa
aaaabaabbabbbbaabbbbbaabbaaaaaabbbaaaaabaabaabbbbabbbabbabbaabba
aaaabaabbbabaaaabbabababbbaabbbbbabbaaaabaaaaaab
aaaabaabbbababbbbbbbbbbb
aaaababbabaaabaaabbbbbaaaabaaabaaababbbbbaabaabababbbbbbbaabaaaa
aaaabbaaaababaaabaaabbbabbabaabababaabbb
aaaabbaaaabbbbaaaaababba
aaaabbaaabaaabaabaabbbabaaaabbbbabbaaabbbbabbbbbbbaabbabbbbbababababaabb
aaaabbaaabaababbbaaabbbb
aaaabbabaabaaaaabaabaaaaaaabaababbbaabaa
aaaabbabaabaabbaabaaabba
aaaabbabaabbaaabbaaabbbabbabbabaababaaabaaabbaabbabbbbbabaabbabb
aaaabbababbbaaaabbabbaabbbaaabbbabbbbababbbaaaab
aaaabbababbbaaabababbbabbbabbaba
aaaabbabbaaaababaabbbabb
aaaabbabbaaabbaabaabbaba
aaaabbbbabbaabbaaabbbbbb
aaaabbbbbabaabaaababaaabbbbbaaabbbbbabbbaaaabaaabababaaa
aaaabbbbbabbabaabaaaaaab
aaaabbbbbbbbaabaaabababaababbbaa
aaabaaaaaabaaaababaaaabaaaaabaabaababbabaabababa
aaabaaaabbbbbaababbbbababbabbbaaaabbabbaababaabbbaaabaabababaababbabaabb
aaabaabbaaabaabbbbaaababaaababbbaaabaaabbababaaa
aaabaabbaabbbbaaabaababa
aaabaabbbbaaabbbabababbb
aaababbbaaabbbaaaaabbbab
aaababbbaaabbbbbbbabaaabaaabbbaabbaabbbaaaaabbbabbababaa
aaabbaabaaaabbbababbbbbbbabaaabaababaabb
aaabbbaabababaaaaabaabab
aaabbbaababbaaaabaaaabbbbaaabaaa
aaabbbaabbbaabbababaabbabbbbabaabbaabbaaaabababbabbbbbba
aaabbbaabbbbabaaabaabbabbabbaaaaababaaabaaabaabbbaaabaab
aaabbbbbabaabbababaaaabb
aaabbbbbbbaaabbbbaaaaaba
aabaaaaaaaababbbabbbaabaabbbbaabbbbbbbabaababababaababaa
aabaaaaaababbbbbaababaabbbbabaabaabbbbaa
aabaaaaabbbabaababaaaaab
aabaaaababaaabaababaaaabaaaaaaba
aabaaaababbaaabbbbbbaaaa
aabaaaababbabaaaabaabbabaaaaaababbbbbbbb
aabaaaabbbabbbaaaaaaabaa
aabaaabaabababababaaaababbbbbbababaabbbbbbbbbbbbbaaabbbb
aabaaabaabbabababbaaaabbababaabbbbaabbba
aabaaabbbbabbbbbbaabbbaababbaaabbaaaaabaaaaaabbbbbaaaabbbaabababbabbaabaabbabaaaaabbabab
aabaabaaaabaaaaaaaaabbabaaababbbaaabbbbaabaaaaab
aabaabaaabbaaaababbabbab
aabaabaaabbbbaaaaaaaaaaaaababaaabbabaaabababaaaaababbaba
aabaabbabababbbbabbabbabaaabaaab
aabaabbbaabbabbababbaaabaabbabbbabbabaab
aabaabbbabbaaaabaabbbaba
aabaabbbbaaaabbbaabaabbbbbabaaabbbabbbbaaabbbababababbbbbaabbbbb
aabaabbbbabaaabbaababbaabbbbbbabbbbaabababbabbaa
aabaabbbbabaababbabbbabaabababbb
aabaabbbbbaabaabbbabaaaabababbba
aababaaaaabaaaababbbaaabbbabbaabababbabababbabbb
aababaaaaabbbaabaaaaabbb
aababaaabaaabbababaaabba
aababaabaaaaaaabaabbaabaaaabbaaaaababbbbabbbabbbbaaabbab
aababaabaabaabaaababbaba
aababaabaabbbaabaabbaaaa
aababaabaabbbbaaaabaaabb
aababaababbaaaaaabaaaabb
aababaabbabaabbabbbabbba
aababaabbbaaaaaabaabbbaabbaabbbbabbbbaaaaaaababb
aababaabbbaaaabbabbaabaababbbbbbbbbbbbba
aabababbbbaaabbaaaababababababbb
aabababbbbbbaaabababaababbbaabbabbbababbabbabaabaabbbabbabbbabbb
aababbaaabababbaaabbaaaa
aababbaaababbbbbbaaabbbaabbbabab
aababbaaabbbbababbaababa
aababbababaaaaaaabbbbbbbbaabaaab
aababbabbabaababbaaaabbbbaaabaaa
aababbabbabbaaabbbbabbaaaaabbbba
aababbabbbabbaaaaabbaaaa
aababbbbabaababbbbaababaaabbaaab
aababbbbababaaaababbabba
aababbbbbaabbbaaaaaabbabaababbba
aababbbbbaabbbabbaaaaaba
aababbbbbababaababbaababababbbbbbaaabbbbaaabaaab
aababbbbbbbabaaaaaabaabbbbbabbaabbabbababbbbaaba
aabbaabaaababaaabbababab
aabbaabaabbaaabaaabaaaab
aabbaabaabbabaaaaabbbbaabaaabbababbabbba
aabbaabababaababaaababbbbabaabbb
aabbaababababbabbaabbaba
aabbaabbaabaaabababbbbabaababbabbbaababbbabbbbaa
aabbabaabaabbaabbbbabbbbaaabbaababbabbbb
aabbabaabbababbabbbaaaab
aabbabbabbbbbbabaaaababa
aabbabbbbabababbaaabbbbaaababaabbabbbbbbbbaabbbaabbbabaababbbaabbbaabbaabbbaabbb
aabbbaabaaaaabababaaaaab
aabbbaababbababababbabaabaabababbbaaaaab
aabbbaababbbabbbaababaaaabbbbababbbaabaaabaaabbbbbaabbaa
aabbbaabbaaabbaabaabbaaaabbabaaa
aabbbaabbabbabaaabbaaaaaababbaaaababaabaabababbbaaabbbbaaabbabab
aabbbaabbbabbbaaaababbba
aabbbbaaabaabaabababbabbbaabbbaabaaabaabaaaaaabaaababbba
aabbbbaabaaaaaaaaabbbaabbbbabaababbaaabbbabbabbbaaabababbaabbaba
aabbbbaababbbabaabaabbba
aabbbbaabbabbbbaabbbbbaabbbbbababaabbbabbaabaaabbbaabaaa
aabbbbabababbbbbbaabbaba
aabbbbababbbbaabababbbabbaabaabb
aabbbbbaaabaaaaabbaaaaaabbabaabbabbaaaabbbbbbbbbbbababbbbbababab
aabbbbbabbbbabbbbbbbbababbbabababaaaababbbaaaabaaaabaaaaabbaabbaaababaaaabaaabab
abaaaaaaaaababbbaaaabaaa
abaaaaaaaabaaaabbabaababbbababbaaabbbaabaaaababbbbbabbba
abaaaaaaabbaababbbaababbbbbaabbbbbaabbaa
abaaaaaabbaaaabaaabbaaab
abaaaaaabbaaaababababbabababbabbbbaaabaabbbaabaa
abaaaaaabbbababbbaaaabba
abaaaabaaaababbbaaabbbbbabbbabbaaaababaa
abaaaabaabbaaabababababa
abaaaabbaaaaaaabbbbaababbabbbbaabbababaaaabbabbb
abaaabaaaababbabbaabaaba
abaaabaaabaaabaaaababaaaaaaaabbb
abaaabbbbaabbaababababbaaabbbaaabaababab
abaabaaaaaaaaabbabbabbbb
abaabaaaaaabbbbbabbabaab
abaabaaaababababbaabaabb
abaabaaababaabbaabababbaababbaaabbbaabbaabaaaabbbaaaaabbbbbaabbbaaaabaaa
abaabaaababbbababbabbaabaababaaaaaabbbbabbbaabaaaaabbaaaababbababbbbbbaa
abaabaabbababaabbababbabbbabbaaabbaabaaa
abaabaabbabbaabbbabaabaabbbabbbbabbbaaabaabbabaaabbabbaabbaaabba
abaabaabbbaabaabaabaabab
abaabaabbbaabaabbaaabbabaabbbbaabaaabaab
abaabbabababbbbbabbbabbabaabbaababababaa
abaabbabbabbabbababaababbabbbabbbbbbbbbb
abaabbbbaabbabaabaaabbaaababaabaabbabaab
ababaaabaabaabbbbaaaabba
ababaaabaabbbbabbbbbaabbaabbabbababbbbbbaabbbbbb
ababaaababaaaaaaaabbabaabbabbbbbbabbbaaa
ababaaabbabaaaaaabbaaaaabbaabaaa
ababaaabbbbaaabbababaaaaababbaababbbbaaaaababaaabbabbbaababababbabaabaabbabbabbbbaabaabb
ababaabbbbbbbaaaababbaba
ababababaabbbbbbbaaabaaabababababaaababbaaaaabbababaabaaaaaaaabaaababaab
abababababbbaaaaabbababbaabaaaaabbaaababaabbbababbabbbabaaaaabaa
ababababbbaaaaaaababbbbbbbbabbab
ababababbbaabaababbbbbba
abababbabbaaaaaaaabbabbabaabaaba
abababbbaaaaabbbbbbbbabaabaababbabbbbbbbbbabbabbaaabaaba
abababbbbababbbbbaaaabbbaabaaaaaababbabababbbbbbbbaabbba
ababbaaaaabaabbbabababbb
ababbaaaabbaaaaabbbbaabbaaabaaaabaaaaaabaaaabaaaaabbaaaababbbbaa
ababbaaabababaabbaababbaaaabbbab
ababbaaabbbaabbababbbbab
ababbabbaaabaababaababaabaababaaababbababaaaabaa
ababbabbaaabaabababaabbaaabbabababaabbaa
ababbabbbbaaabbbababbabbaaaaaaba
ababbabbbbbaaabbbbbbaaaa
ababbabbbbbabbbbaababaaa
ababbbabaababbbbbabbbbba
ababbbabbabbbaabbbbbbbbb
ababbbbbbbabaaaaabaabbabbbbaaaabbabbaaba
ababbbbbbbababbbbbaaabbbaaababba
abbaaaaaaabaaaaaaaababaa
abbaaaaaaabbabbaaaababab
abbaaaaaabbbabbababbbbbb
abbaaaaabaabbaaababbbbabbbbbaaababaaaabbaabbabab
abbaaaaababababbabaaaabaaabbbabaababbbba
abbaaaaabbaabaabbbabbbab
abbaaaabaaaaaabbaabbabbb
abbaaaabaaaaababbaaaabbbababaaabbaaabaaa
abbaaaababababbaababaaabbaabbbabbabbabaabbaabbabbbbbabbb
abbaaaababbaaaabbabaaabbbbabbaaabaaaabababbbbbbaaaababaaabaabbaabababaaa
abbaaaabbaabbbababaaaabb
abbaaaabbbabbbaaabbbbbab
abbaaabaabbaaaabaabaaabaababbaba
abbaaabababaababaabbbbaaaabbbbaaababaaaa
abbaaababababaabaaaababa
abbaaabababbbbabbbbaabbaabaabaabaaaaaaabbaaabbbb
abbaaababbbabbbbabbbbababaabababbbbbbbabababbababbbbabbb
abbaaabbbbaaababbbbaaabbabbbbbbbbbbbbaababbbbbbbabbbbabbabbabbbb
abbaabaaababbaaaaaabaabbbbaaaabaabbbaabaaabaabab
abbaabaabaaaabbbbbbbaabbaaababbbababaaba
abbaababababaaababaaabaaabaabaabbbababbabbbbbaaaaaaabbba
abbaabababbbabbbbbaaabbbaabaaaabbbbaaababbbbabba
abbaababbaaababbbbbaaaababbabbbbbbabbaba
abbaababbaaabbbaababababaabbaabbbbaaabaa
abbaabbaaabbbaabbaaabaaa
abbaabbabaaabbabbaabaaab
abbaabbabbabbbaabaaaaabb
abbabaaaababbaaabababbaa
abbabaaaabbababbaaabbaba
abbabaaabaabbbababbbaabbbabaaaaaaaabaabbbbbbbabb
abbabaaabbaaaaaabaaaabaa
abbababaabbbbaababaaabba
abbababbaabbbbaaaaaaaaaaaaabbbbbbabaabbb
abbababbabaabaabbaaabbabbbbababababaabbb
abbababbbaababababaaabaaababbbabbbbaaaab
abbababbbbbbbaabbaaabbababbbbaabaabbaaaaabbbbbab
abbabbaaaabaaaaaababbaaaabababbbaaaabababaaabbaa
abbbaaabababababbbabbbbb
abbbaaababbabaaababaabaabaaaabba
abbbaaabbaabababbabbbaaa
abbbaabbaabaabbbbaaaaabb
abbbaabbbbbabaaaaaababab
abbbabbaababaaababbbbbab
abbbabbabbabbbbaaabaaaaababbaabbaabaaabb
abbbabbbaabbabaaaaaabaabbbbbbabb
abbbbaaaaabaabaabaababbaababbbbabbabbbab
abbbbaaaaababbabbbabaaabbbaaabba
abbbbaaaabaabbabbaabbabbaabbbabb
abbbbaaaababaaabaaabaabaabbbbbbbbabbbaaa
abbbbaaaababbabbabbaabbaabaabbba
abbbbaaabaaabbbaababbbaa
abbbbaaabbbaabababaababa
abbbbabaabbbabbbabbaaaaaaaaaabba
abbbbababaabaababbbbabbbabaaabbbaaaaabbbabababaabbbaabaa
abbbbababbaababbbaababbaabaabbabbabaabbabaaaaaab
abbbbbaaaabaaaaaabababbb
abbbbbaaababbaaaabbaabaabbabbaaabbababbbaaaaabaaababaabb
abbbbbaabaabbbaaaaaaababbaaaaabb
abbbbbaababbbbabaabaaabaaabbbbaaaabbaabababbaaaaabaaaabbbbbababa
abbbbbbbbbbabaabbabbabbb
baaaaaaaaaabaabbaabbbbba
baaaaaaabbaababbbbabaabaababbaabbbaaaabaaaababab
baaaaaaabbaabbbbbaabaaba
baaaaaaabbbbbababbbbbbba
baaaaaabaabbbbbbabaababa
baaaaaababababbaaaabbbbaaabaabaa
baaaabaaaabaabaabaaabbaabaababbaabbbbbabaabaaababbbabaabaaaababb
baaaababaabaaaaabaabaaba
baaaababaabbaabbabbaaabbbaababbabbbbabab
baaaabababbaaabbbaabbaba
baaaababbbaaababbbabbaabaaabbaaa
baaaababbbababbaaaabaaab
baaaabbbbaabbaaaabaaabaabaababaabbbaabaa
baaaabbbbbaababbbbaabbab
baaaabbbbbabaaaabbbabaabbabbbaaa
baaaabbbbbbbaabbbabbbababaaaabba
baaabbaaabaaaaaababaabbb
baaabbaaabbaaaaababababa
baaabbaabababbabbaababbb
baaabbaabbabbbbaaabaaababbbbaaabbabbabababaaabab
baaabbababbbaaaabaabbabbaabbabbb
baaabbabbaaaaaaababaaaabbaabbbba
baaabbabbaababbaaabbbaaa
baaabbabbabbababaaaabbabaabaaaabbbbaabbbabaaaabb
baaabbabbbbaaababbaaabbbababbaba
baaabbbaaaabbbaabbaababa
baaabbbaabbaaababbbbbabb
baaabbbaabbbbaabbbabbbbababaaaaaabaabaabbbaaabbaabbabbba
baaabbbababaaaabbbbbabab
baaabbbababbaaabaabaaabb
baaabbbabbaabaabababbbba
baaabbbbaaababaabababbba
baabaaaaababbbabbbbaaaaabbbabababaababaa
baabaaaabaaaabbbabbbbbbbabbabbaabbbaabab
baabaaaabaaabbbaabbbaabbbbbaaaab
baabaaaababbabaaaaabaabbbaaaaababbbbabba
baabaaaababbabaaababaaabbbbaaabbbbbbbbbaaabbaaaaaaabbbba
baabababbaaaaaaabbbbabab
baabababbbaaaabbbbabbaaaabbabbaabbabbbbb
baababbaaabbaababaaaabbbbaabbbba
baababbabababbabaababbab
baababbababbbaababbbaabbbbabbaababaabababbbabbba
baabbaaabaaaababbbabbaabbbabaababbbbbbaaaaaaabaa
baabbaaabbaaaabbaaabbbbbabaabbabbbbabbba
baabbaaabbbbabaabaaaaabb
baabbaabaaabaabaababbaaaabaaaabaaaababbbbbbabbbbbbbbabab
baabbaabababbabbaaaaabba
baabbaabbbabbaaaabababbb
baabbabababaabbbabbabbbaabaabaabaabbbbaabbabbabbabbbaaaabaaaaabaabbbaaba
baabbbaababaaaabbbaababa
baabbbaababaabaabaaabbaaaabaabbababaabaaabbbabaa
baabbbaabbaaaabaaaaababa
baabbbaabbabbabaabbabbbb
baabbbabaaaaaaaabbbabaaaababbabbaaababab
baabbbabaaaabaabbabbababbabaaaabaabbbabaaaabaaabbbbbbabb
baabbbabbabbaaabbbbaabab
babaaaaaaabaaabaabbbabab
babaaaaabaaaababbbabbabb
babaaaabaaaabbbbbaabbbabababbaba
babaaaababbbbbbbaaaaaababaabbbba
babaaaabbbbabbaaaaaabaabaaaababa
babaaabbbaabbbabbbabaaaabaaaaaaaabbbabaa
babaaabbbbbabbaaaaababbbabbbabaaabbabbaa
babaabaaaaaaabababbaaabbbbbaabaabbababab
babaababaabbbbabbabbbbabbababbaa
babaabbaaaabbabbbaababbbaaababab
babaabbaaababbaaabaabbaa
babaabbaabaaaaaabababbaa
babaabbaabbbbaabbbbaabbb
babaabbabaaabbabbbaaaaaaabaaaababaaabbba
babaabbababaaabbbababbbb
babaabbabbaaaabbbaabaaba
bababaababbbaaabbabbaaba
bababaabbaabbaaaaababbbbaaababababababbb
bababaabbabbbbbabaabbbbabbbaaaab
bababaabbbbabaabbababbba
babababbbaaaabbbbabaababbbbaaaaaaabbaababbaababbaaabaaababaaaaabbbabbbab
babababbbbaabaaababbaaababaababa
bababbaaabaabbabaaababaaaabbaababaabbaabaabbabbababaabbabaaabaaaaababbabaabaaabb
bababbabaabaaababbbaabab
bababbabbabbbbabaababbbbabaabbabaababaabbbaaaabaaaabbaabaaabbaaa
babbaaaabbbbabaabaababbbababaaabaaaabaabbababbaabbbabbaa
babbaaababababbabbbbbbba
babbaaabababbaaabbbaabab
babbaaabbaaabbababaaaaab
babbaaabbbabbbbbbbbabbbbbbbaabbaaabaabbabbababaaaaaabbaa
babbaaabbbbaaaabbaababbbabbbabaababbbbbabbbabbbaaabbbbaaaaabaaaabaaabbbb
babbabaaaaabaabbbbbaaaab
babbabaaaaabbbbbbbbabbba
babbabaaaabaaaabbbbbabab
babbabaabaaaabbbbbbaaaab
babbabaabbaabbabbbabbabaabaaabab
babbabababbaabaaabaaabbb
babbababbabaabbabababaaa
babbababbbbabbbbaabaaabaabbaabbaabbabaab
babbababbbbabbbbbabaaaaabaabbbaaaaaabaabaabbbbabbbaaabaa
babbabbaaabbbbbbabaaabbb
babbabbabaababbaabbbabbbaaababaa
babbbaaabaabaaabaabaaabb
babbbaababbaabaaaabbbbbb
babbbaabbbababbbbbbabaababbbbabbbbbbabbabbaaaabaabababbabbbbbbbbabbbaaabaaabbbab
babbbaabbbabbaaabbabbbbabaaabbbb
babbbaabbbbababbbbabbaabaababaaa
babbbabaaababaaaaabbbabb
babbbabababaaabbbbbaabab
babbbababababbabbabbbabb
babbbbababaaaabaabbbbbaaababbaabaabbbbbabbaabaaaaabababb
babbbbabababbbaabbaabbbaabbaababaabbbbbaabbbbbbabbaaaaaaaaaaaabbabababbb
babbbbababbbbbbbaaaaabaa
babbbbabbabbbbabbbbbaabbbbbbaabbaaaaaabbbbbaabaa
babbbbabbbbbbbabaabbbaaa
bbaaaaaaaabaaaaaababaaaa
bbaaaaaaabbbbaaabaaaaabb
bbaaaaaabaabbbababbaabbabbabaaaabbbaababaaaabbbaabababaa
bbaaaaaabbbbaabbbbbbaaaabbaabbaa
bbaaaabaabbaaabbbbabaaabbabbbbbbbbbabbabbbaabaaa
bbaaaabaabbbbaabaaababaababbaabbbabbbaaa
bbaaaabaabbbbaabbbababbbaaaabbabbbaabbab
bbaaaabababbabaaaaaabbbb
bbaaaabbaaaabbbbbbbbbbbbbabababa
bbaaaabbaabbaabbbbbabbba
bbaaaabbbaaabaaaabababbbbaabaaabbbbbbbbabaaaaabbababaaaaaaaabbabbbaaabaababbbaab
bbaaabaaaabbabbbaabaaaaababbababbbabbbaa
bbaaabaaabababbaaabbababaaaabaaabaabbaaabbaababaaaaababababababbbaaaabbabaaabaaaabbaabba
bbaaababaaaaaaaabbbaabbb
bbaaabbbbaabbaababaababbbabbbaaa
bbaaabbbbaabbbabbaababbababaaaabbaaaaababbaaabaabababbba
bbaaabbbbabaabbaaabbbaba
bbaabaabaabbabbababbabaaababaababaaaaaba
bbaabaabababbbabbabbabbababbabababaaabab
bbaabababbbaabaaaaaaaaaababbaaabbbbababbabbaaaabababbaaababaaabbaababaababaaaabaaaaabbaa
bbaababbbabbaaaaababbbabbbaaabba
bbaababbbbaaababbabaaaaaaabababbaaabbabb
bbaababbbbababbaaaababaa
bbaababbbbbbbbbabbabaabbbbaababa
bbaabbaababbaabaabaaaabaaaabbaaabbababaabbaabbbbbabbaaaabbbbbbababbbababaabbaaaabbaaabbb
bbaabbbbaabaabbaabbbbbab
bbaabbbbaabbaabbbaaababa
bbaabbbbbbabababbbbabbabbaabbbbaaabaabbbbbabbbbbbaabaaaabababbbaaabbbaaa
bbabaaaaabbabaaababaaaabaaabbaab
bbabaaabababbbabbaabaaaababbaabb
bbabaaabbaaaaaaaaaaabbbbaaaaabaa
bbabaaabbaabbaaaaabbbaba
bbabaabaabaabaaaabbababaaabaabaa
bbabaabababbbbabbababaabaabaabbaaabaabab
bbabaababbbbbabaabaaaabaabaababa
bbabababbbaabbababbababaabbaaaaaaaabbaba
bbababbaaaabaaaaaaaabbbbbbbaaaabbbaabbab
bbababbaaaabaaaaabbabbaa
bbababbababababbbbaaabaa
bbababbabbbabbaabbaabbbbabbaabbbbbaaabba
bbababbabbbbaaabbabbbbabbbabaaabaaababaa
bbababbbababbbababbabbbb
bbababbbbabaabaabbbbabaaaaaaaabb
bbababbbbababbabbbaababbbabaabbaabbaaabaabababaa
bbabbaaaaaaabaaaaabbabbb
bbabbaaabbabbbbabaabbaaa
bbabbaababbaabbabaabbaaaaabbbbbbbbbabababbabbbababababaa
bbabbaabbbaaaabaabbaababbbbbabba
bbabbaabbbbabaaabbbaaababbbaaaabbbababab
bbabbabbbbbbbabaabbaaaababbabbababaabbaabbbbbbaaabbbbbbaabbabababbabaababbaaaaba
bbabbbaaaabaabbbaababaabaababaabaaaaabba
bbabbbaabaaaabbbbbabaaabababaababaaababa
bbabbbaabaaabbabaaabaabbbbabbbaabbbaababbaaaaabbaabbbbba
bbabbbaabbaaaabbaaaababb
bbabbbbaaababbbbaaababbbbbbaabbaababbabbaaaabbbaaaaabbbaaabbabbbaaababab
bbabbbbaababaaabbaabbbaabbbaabbbbbbaabaa
bbabbbbaabbbaabbbabbaaabbababaababbbbbabbbabbbbb
bbabbbbaabbbabbbabaabaabababbbbbabbabbbb
bbabbbbaabbbbababaaaaaaaabbaabbb
bbabbbbaabbbbbbbabababbabbbababa
bbabbbbabbabbbbaabbbbabb
bbabbbbbabaaabbabbbaaabbbbbaaababbbaababaabaabbaaabbaaabbabaaabbbbbbabbaaabbbaab
bbbaaaaaabbbbbaabbabaaabbbbaaaab
bbbaaaaabbbabaabaabaaaaaabaabbbaaabbaaab
bbbaaabaaaababbbbbaabaabbaabaaba
bbbaaabababaaaabbabaabaabbaabaabbbbbaaaa
bbbaaababbbabaaaababbbbbbaababaa
bbbaaabbaababaabaabbabbb
bbbaaabbbababbabbaabaabb
bbbaaabbbbaabaabbaaaaaab
bbbaaabbbbbaaabbabbbabaa
bbbaabaaaabbbaaabbabaabb
bbbaabaabaabbabbaabaaabbbbbaabaabaabaabbbbabbbaabaabbbbbbbbbbaaaabbbabab
bbbaabbaaababbbbabbbaabb
bbbaabbaaabbabaababbababaabbabbbbabbaabaaaabbaabaabaabaabaaaabbb
bbbaabbabaaabbabbababbba
bbbaabbababaababbaabbbba
bbbaabbabababbabbbbbbbbb
bbbaabbababbabbaababbaabbbbbabaaabbaaabbbaabababaabbababbbbabbbabbaabaaabbbbaaba
bbbaabbbabaaaaabaaaababb
bbbababaaabababbabaabaaabbabbbab
bbbababbabbbabbbabbaabbababaabbb
bbbababbabbbbaabaaaaaabbaaaabbbbaaaaaababbaabbab
bbbababbabbbbaabbabaabababbbaabaaabbaaab
bbbababbbaaabbaaabaabbba
bbbababbbaabbbaaabbbbabb
bbbababbbabababbaaaabaabababaabaaabbabbb
bbbababbbabbaaaabbabbaaaaabbbbaababababbaaaababb
bbbababbbbbbbaababbbaabbbababbbabbabbabb
bbbabbaaaabbaababaabaaba
bbbabbaababbbaabaabbaaaa
bbbabbaabbaababbbabbbbabaaaabbbbababbaaababaabbb
bbbabbaabbbbbabbaababababaabaabb
bbbabbbbbabaabaaabbbabab
bbbbaaabbabababbbbbabbaabbbbbabbababaaaa
bbbbaaabbbabbaaabbbaaababbabbaaabaabbabbabbbabab
bbbbaabbaaabaabbbbbbbabb
bbbbaabbaabbbbaabbbababa
bbbbaabbababbabbababababbbbaaababbbaaaba
bbbbaabbabbbbaabbbaaaaaaaabbbbaababaabaaaabbaaaabbbaabbb
bbbbaabbbabbabaaabbbbbaaabaaabab
bbbbaabbbbbaaabaaabbabbabbababbaababaabb
bbbbaabbbbbbaabbbaabbbaaaabbbaabbaabbbababbbabababbaabbb
bbbbaabbbbbbbbababababbabababbaa
bbbbabaaaaaaabbbababbbba
bbbbabaaabaabaabbbaaaabaaaababbbbbaababababababa
bbbbabaababaababaaaaababbbbbabaababaababbbbaaaab
bbbbabaababbbabaabaabbbb
bbbbbaabaaaaaabbaabbbaababaabaaabbabaaabbbabbaaaabaaaabb
bbbbbaabbabaaaaaabbabbab
bbbbbaabbabaaaaabbaaaabbabaababbbbabbbbababaabaabaaaaaabbbaaaaab
bbbbbaabbabaaabbbbaaabba
bbbbbaabbabbbbabaabbbbaaaaaaabba
bbbbbababaaababbabbbbaabbabbabbbaabbaaaabaaaabaaaaaabababbbababa
bbbbbababaabbaaaaaaaaaaaabaaaabaabaabaabbaabaaab
bbbbbababbabbaaabbbabaabaabbbbbb
bbbbbababbbabbbbabbbbbaaabbaabaaabbbabaaabababbb
bbbbbbabaaaaabababbbbabb
bbbbbbabaaaabbbbbbbaabbabbaabbba
bbbbbbabababbabbbbaaaaaabaabaaaaaababaaabbbbbbaabbbabbbaaaaaabaa
bbbbbbababbbbbbbbbababab
bbbbbbabbabababbabbbaabbaababababbaabbababbbbbab
bbbbbbabbbaaabbbbabaabaa`
