package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
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

	inputLines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	bitsLength := len(inputLines[0])

	gamma := ""
	epsilon := ""
	for i := 0; i < bitsLength; i++ {
		numZeros := 0
		numOnes := 0
		for _, line := range inputLines {
			if line[i] == 48 {
				numZeros++
			} else {
				numOnes++
			}
		}
		if numZeros == numOnes {
			panic("No most commmon bit")
		}
		if numZeros > numOnes {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}
	log.Println("Part 1:", gammaDec*epsilonDec)

	oxGen, scrubRate := make([]string, len(inputLines)), make([]string, len(inputLines))
	copy(oxGen, inputLines)
	copy(scrubRate, inputLines)

	oxGenPosition := 0
	for len(oxGen) > 1 {
		mostCommon := mostCommon(oxGen, oxGenPosition)
		if mostCommon == 0 || mostCommon == 49 { // No most common or 1 most common
			oxGen = removeElements(oxGen, 48, oxGenPosition)
		} else {
			oxGen = removeElements(oxGen, 49, oxGenPosition)
		}
		oxGenPosition = (oxGenPosition + 1) % bitsLength
	}

	oxGenDec, err := strconv.ParseInt(oxGen[0], 2, 64)
	if err != nil {
		panic(err)
	}

	scrubRatePosition := 0
	for len(scrubRate) > 1 {
		mostCommon := mostCommon(scrubRate, scrubRatePosition)
		if mostCommon == 0 || mostCommon == 49 { // No most common or 1 most common
			scrubRate = removeElements(scrubRate, 49, scrubRatePosition)
		} else {
			scrubRate = removeElements(scrubRate, 48, scrubRatePosition)
		}
		scrubRatePosition = (scrubRatePosition + 1) % bitsLength
	}

	scrubRateDec, err := strconv.ParseInt(scrubRate[0], 2, 64)
	if err != nil {
		panic(err)
	}

	log.Println("Part 2:", oxGenDec*scrubRateDec)

	elapsed := time.Since(start)
	log.Printf("Duration: %s", elapsed)
}

func removeElements(s []string, character byte, position int) []string {
	for i := 0; i < len(s); i++ {
		if s[i][position] == character {
			s = removeBinNum(s, i)
			i--
		}
	}
	return s
}

func removeBinNum(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func mostCommon(s []string, position int) byte {
	numZeros, numOnes := 0, 0
	for _, line := range s {
		if line[position] == 48 {
			numZeros++
		} else {
			numOnes++
		}
	}
	if numZeros == numOnes {
		return 0
	}
	if numZeros > numOnes {
		return 48
	} else {
		return 49
	}
}

var inputTest string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

var input string = `000110010001
101000110000
000110010111
100011100010
111001100001
001010001010
010100100101
011000010000
111111011010
001111011101
011011010010
001100010101
001010101100
000000000000
100010111111
100100110011
111100100001
011110001110
000110100101
011101111001
111101110110
101001001111
010111100010
001110100011
110010111101
110000001101
110110111101
101100000110
101101010110
001011001110
111110000111
011110010110
011001110100
110101101110
101110001100
100111000010
101111010001
111110000101
010100111111
010110111011
000000100101
001000101011
011100101110
010011000100
110011100001
100100101110
111101110011
010100010010
110100010000
000101001101
101010110100
000011100110
000100110100
111001011001
010101001110
110100110100
111010001010
111111101111
111000011001
111110000000
011111111011
010100100001
010110100100
110000011110
001100101101
010011001000
101110110110
001001011011
100100101101
001111000100
101001010000
110101010010
101111100101
001011101010
001110010111
100111111000
010010110110
000100100010
010111100100
001000100001
101100100111
111001000001
110010111100
000110011101
000001010010
001010010111
001101000100
000110111101
111011011000
011110101000
100101000011
001000011100
111101001001
000011110011
011010101001
000111000101
011000011001
011010010100
010110001110
011010010110
111010111110
010101000010
100010011100
110001111100
100101100001
010001110010
001000010011
111010011101
110001111000
101100010011
010000010110
101011001101
000110000101
010000101111
001100010100
100110010111
010000010001
011100100101
110011010000
111101001110
001101110010
000011000110
011000010111
000101010000
101111110110
101001011000
111010001000
001100100010
111100001100
011000001100
111010110100
011111010110
111011100000
110010110110
100111001101
101111111010
000101101110
011011000111
110001101100
001110001110
100100001011
111000001110
110000111111
100110111111
100100000011
111101110001
000100001111
000101000111
001011000111
010011001100
100001101110
111100001010
110100101101
100111100001
011000100111
101010010010
111101000111
000010110101
001100010010
001111111101
001101111011
001011101110
011100000111
110010111111
100110110101
100110000001
001011100010
100010111000
000000111000
111010100100
010101101110
100111011111
111101010100
010000111011
000101100001
101010100000
101110111100
010000011011
000101110110
100110000000
001001001001
110110101100
001100110011
100111001111
110111011101
001110110010
001010110000
001001100101
111100101101
100001111100
011001011110
010010100010
001011100100
000100111010
010110011100
011100000101
001110101010
010011101010
000011010001
011101101100
110100101001
101011110110
001010010101
111110011010
011001101111
001000111111
000011101101
010001001001
111011000010
001011110010
010000011001
110000110001
011001001000
110011011000
110101011001
100000100011
010011001111
100100101000
111111111100
011011010101
101100101010
111110011111
010101001000
011110100000
101001111111
110000101110
101011101110
101111001110
100000101101
010010001011
010100110111
111011110001
000100000010
100000111000
111111111000
110111100111
000001101101
100001110011
011010111110
100111111100
111111001010
001001000100
010010101101
110100100000
100010011111
010000000110
110111010111
101000001000
100100101001
001001000010
001111111010
001011011000
100010001111
110011101110
010111001110
001011010011
111011011011
010100100100
101011000111
000011111111
000110000111
110000110100
100101011011
001001110110
011100001001
001111110100
001111001101
110010011001
000100001001
000110010110
100001001010
011100111000
011101010111
011111000110
001000001001
110011000101
111111011011
111110010011
101111100000
101011000100
111110011101
010100101010
101011110000
000100010100
110110010010
100011001101
110110011010
011100111011
001111100001
110011100101
111001101011
010000110101
100100011100
111000111001
111000001011
010110100011
110011010100
010011101111
011011110111
010101101010
000000100100
100111011101
001101011000
110111000101
001011110100
001100000100
101010110010
011000100000
011001010100
111100000000
000100011000
001000111000
101101110010
000100011001
000100101001
010010010101
100010101101
100010010110
110101110000
100011100110
010010011100
010001011011
011100000110
101000101011
111010010000
011100111001
100001100110
111100000110
100001011100
111001110101
111011100111
110101010001
111100011000
100111101010
011110110000
001011000101
011010011111
111100111110
010100011010
000010011101
000110000110
011001010001
111000011100
000111010110
000000110101
010011101011
110011001110
011011110110
101100010101
000011011001
001100101011
101000010101
101100011010
000111111110
000011001001
111011101110
100111111010
000101000011
111100010011
010001101010
101001011110
110010110000
110101101001
000011101110
101101001010
110100001101
011100000011
100011100101
111010100110
100010110001
001011010000
001001010100
100101110100
100100001111
001000001010
111001000101
110110100000
001010100010
100100001110
100111010010
010110111000
010100001010
000010101100
001010111100
001111101001
101011111101
011011001001
101000100110
001010011101
001011110000
111110100000
101100111010
101111010111
011000000101
101100000010
010101111110
000011011100
111110111010
011110101110
111011100010
111100010001
100000100110
010111010000
100010100101
010111101010
011011111111
110111101000
100000010000
001011000010
001001100011
001000100110
001001001011
011110101101
010111111101
100100111100
000011001101
010101100000
111100011100
010111001010
001101010011
111011101111
000000111101
100000110110
010000001101
011110001011
111100111100
100101111110
001000110010
101011010000
000101001111
100001101101
011011110001
010010110010
111011110111
000001110110
100110011110
101100110011
011011011010
100100101010
011011000110
000100100011
100101101001
100110110000
100100011001
001010010001
000001011101
010011011010
010101100101
111101111011
010100111011
111000001001
101100110101
011111101010
111110110111
111111110000
011001001100
110001100100
001000000011
000110001010
111000010100
010101000000
101001100000
001010010010
010001100000
001000101110
110111100001
001111000010
100111111101
010111101110
000110000011
011100111101
000011001111
011101100110
011110111110
000010001010
111101001100
101101110011
001000101111
001011001000
010110100101
000110110010
101101001000
010100100000
001111011110
111001111111
110010011101
011010110001
101101111010
011000110110
011101011111
000011000111
011010101101
010000100110
111001101110
100001110100
111100110110
000001101011
011111000011
010001011001
101010001000
001100111001
101111111111
100000001101
110011000000
111000101011
110001010110
011100010001
110110100110
000010010100
000001000101
110111101011
101110011011
010010111000
110000101101
110100011101
000111110001
110001001010
011000111110
100111010110
111100111000
000010101001
111100100101
100110111001
001001010000
101110000010
110011110101
100010011101
011000100101
101011001000
100111011000
110101011111
111100010110
111001100110
001110001000
110110101001
111001011010
101011100101
001101100001
000111000110
110001110110
111100101011
100010100100
000101110001
011101001000
001110001100
100100001010
101010101111
100011010011
010001111101
100101001001
111110101111
010011000101
000110111011
101101100011
000110110110
100000110000
000001000001
011111010111
111110010101
111101011011
101011110111
001000011111
010010010000
110110110010
000110100000
010010110101
110110001111
101111101000
110101110101
001100111110
100101101110
101111110101
000111101001
001011011111
111011001011
001001101100
010101111010
111110001100
100010111100
011000110111
111111101100
101001100001
110000110111
100010101000
010010100011
101110011001
011101001001
110000000110
000110000010
000111111101
110001001100
001001101010
101010011010
101100111101
110010100010
000101110101
000000010101
101010110000
101010110001
101001001100
010001101110
010011111001
010010111011
001110110101
101000000110
101110000101
110110001101
101000100000
011001100110
101111100110
111110001010
110010101011
100111000111
111001111101
100001011111
101101101110
010100100010
111100111011
001011100110
000001011010
101111100010
110000000000
111100010101
110100001001
100010011010
111000110010
101000110011
111101011111
100011111011
101001100111
011011111101
111011010111
011101010011
010010000111
110011111101
011101011101
010101100111
010111000000
100111000001
001110101101
000111101110
000100110000
001110110111
010000111010
001010001001
101001000111
111010100010
110001000011
011111010000
000110010000
001001011111
000100011111
110001010011
110000010011
000011010101
100101001110
101110010101
000011110000
011111100111
001111101100
101111101100
101110111101
110111111001
011110111101
111110010111
101110100001
000110100010
011000001000
111000110101
101000011100
100110010100
001111111111
111001111001
111111000110
000110100100
111000110000
110011010110
001010011001
101000011011
100011111111
011001100000
110100101100
001001111000
000100001100
010101000001
101010110111
100110100000
100110000011
111101000011
001101001010
000000011111
000101000110
001110111001
101110001011
001111011010
000100110010
010011011001
011100101001
010110011001
001011001101
011110100010
011000011100
001001110010
100000000010
010111010010
000010100000
011000011110
010100110010
010101010001
011010101100
000100100001
110101001011
111110111101
010100100110
110010000100
000101010001
110100000001
001010100011
011101100101
010101011110
010101110101
010101011100
010011001010
010011110001
111111100000
011110111010
110010101010
001111010001
111110111110
110001010101
001011010001
110111000110
000011010011
000001101001
011010100100
100000010111
000000100001
110000110010
001101111101
101011111111
100100100010
000001110111
001111110101
001111000011
101000111100
100100101111
111000000101
000011100011
110101100010
110001010111
011000111100
010000011111
000001011000
110011100000
111100100111
000010001011
100111101001
011111100000
101101111001
101100101100
001000110100
111111010111
000110001101
011111100001
111000001100
000111100100
001111000000
000101110010
101100011001
000001000100
001001001101
111001010101
000011111011
011111110000
000000110001
100111110011
101001001011
101000110110
111001010100
000011110100
010001100001
001111001110
100001101001
000001101111
101000000001
111101010000
011111011001
111111011110
110110010110
001011101100
100001010111
101011011100
110010101001
001000001100
001101010101
111111111110
001100100011
110000110110
000000000110
101000001010
001111001100
100011110001
111011000001
011010100110
011110100100
001110011111
010101101011
010110010111
101100110110
000101100010
110010010000
010100010000
111011000111
011010111011
100000100111
001111011000
001010111011
000010111111
110001000101
000111010101
010011110111
111010110000
100111110110
101011010001
011001111111
110110100010
110111110110
011010110111
111000100011
100010011000
111001000010
000100000111
011101001010
111001000111
000110010100
011011010110
100100110100
010011000001
101000111011
100100000000
010010101110
011111001101
001100001100
010111100001
100101101010
100110001101
010111000101
010100011011
000010000011
111000111111
100111011100
111110011011
011110110110
001000001111
100110110111
100000001011
001111001011
010010111101
010001111000
110010110101
011001100101
100110001111
110001101010
000100101011
101110111010
011111100011
101010111111
011101100001
100000001010
010000001000
110110001010
110001010010
101001111001
000100001101
111100001111
111101000110
001010111001
000101011010
101001010010
101011110001
011011110101
110101110100
000001111001
111000001111
101110110000
001101110101
011010110010
101000001110
100011111001
000111011101
110011010101
010111101101
000010111011
000100101010
100101100101
100000011001
010110110100
011010100000
010100010001
110010100101
010100111110
011100011101
100011110010
101000000010
011001100001
111111100100
010011000011
001101010010
111100100000
110101101111
101001110001
101100111111
010111110010
100011000100
101101001101
000000101001
110100110000
000110010010
100111001110
110011100110
010110100010
000111010100
001101000001
111010000011
011001010110
101000100010
111010001001
010101100001
011001110010
000111000111
110001010001
000111010001
000001111100
010001000001
100110000010
111111000111
110111001111
010010001001
001001101111
000110101111
110011011111
100010100111
001010000010
000000110110
111111111101
011111101111
100111101111
111111110111
101111011001
111010111011
001001111110
000010001100
111011001111
111111100011
100010010011
101110101011
111010111111
001111000110
110000001100
100110100110
001110001101
001000111100
010011000110
011011011100
110110011001
001000111010
101100000011
110100111101
011100001011
011101000100
011100010011
101001000001
000001101010
110000001010
010100010100
101001100110
000100100101
001000111101
101101011111
011001000010
101000100001
111110100011
111101010010
110111010101
010111110100
010001010001
010110011110
101001010001
100110111100
110011100010
010110001011
110001001101
001100100111
001001001111
011000110100
000010010001
101010101110
111100111111
000010100010
011010111101`
