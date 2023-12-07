package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
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

	inputString := strings.TrimSpace(string(inputBytes))
	fmt.Println("Input:", inputString)
	// BOILER PLATE --------------------------------------------------------------------

	hands := parseInput(inputString)

	fmt.Println("-----------------------------------")

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handStrength != hands[j].handStrength {
			return hands[i].handStrength < hands[j].handStrength
		}
		for idx := 0; idx < 5; idx++ {
			cardI := hands[i].cardsStrength[idx]
			cardJ := hands[j].cardsStrength[idx]
			if cardI != cardJ {
				return cardI < cardJ
			}
		}
		return false
	})

	part1 := 0
	for rank, hand := range hands {
		part1 += (rank + 1) * hand.bid
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].jokerHandStrength != hands[j].jokerHandStrength {
			return hands[i].jokerHandStrength < hands[j].jokerHandStrength
		}
		for idx := 0; idx < 5; idx++ {
			cardI := hands[i].jokerCardsStrength[idx]
			cardJ := hands[j].jokerCardsStrength[idx]
			if cardI != cardJ {
				return cardI < cardJ
			}
		}
		return false
	})

	part2 := 0
	for rank, hand := range hands {
		part2 += (rank + 1) * hand.bid
	}

	printHands(hands)

	// ANS --------------------------------------------------------------------
	elapsed := time.Since(start)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	log.Printf("Duration: %s", elapsed)
	// ANS --------------------------------------------------------------------
}

func printHands(hands []*Hand) {
	for _, h := range hands {

		fmt.Println(h)
	}
}

func parseInput(input string) []*Hand {
	lines := strings.Split(input, "\n")

	hands := make([]*Hand, 0, len(lines))

	for _, line := range strings.Split(input, "\n") {
		cards, bidString, _ := strings.Cut(line, " ")

		bidInt, _ := strconv.Atoi(bidString)
		hand := &Hand{
			cards: cards,
			bid:   bidInt,
		}

		hand.processCardStrength()
		hand.processType()
		hand.processJokerType()

		hands = append(hands, hand)
	}

	return hands
}

type Hand struct {
	bid int

	cards         string
	cardsStrength [5]int
	handType      string
	handStrength  int

	jokerCards         string
	jokerHandType      string
	jokerHandStrength  int
	jokerCardsStrength [5]int
}

func (hand *Hand) processJokerType() {
	cardMap := make(map[string]int)

	// Count cards
	for _, c := range hand.cards {
		if _, ok := cardMap[string(c)]; !ok {
			cardMap[string(c)] = 0
		}
		cardMap[string(c)]++
	}

	// Play with cardMap, add joker cards to the most numerous non-joker card
	jokerCount, _ := cardMap["J"]
	if jokerCount != 5 {
		nonJokerMaxCard, nonJokerMaxCardCount := "", 0
		for card, count := range cardMap {
			if card == "J" {
				continue
			}
			if count > nonJokerMaxCardCount {
				nonJokerMaxCard, nonJokerMaxCardCount = card, count
			}
		}

		cardMap[nonJokerMaxCard] += jokerCount
		delete(cardMap, "J")
		hand.jokerCards = strings.Replace(hand.cards, "J", nonJokerMaxCard, -1)
	} else {
		hand.jokerCards = hand.cards
	}

	hand.jokerHandType, hand.jokerHandStrength = whatIsHand(cardMap)
}

func (hand *Hand) processType() {
	cardMap := make(map[string]int)

	for _, c := range hand.cards {
		if _, ok := cardMap[string(c)]; !ok {
			cardMap[string(c)] = 0
		}
		cardMap[string(c)]++
	}

	hand.handType, hand.handStrength = whatIsHand(cardMap)
}

func whatIsHand(cardMap map[string]int) (string, int) {
	switch len(cardMap) {
	case 1:
		return "Five of a kind", 6
	case 2:
		switch maxCount(cardMap) {
		case 4:
			return "Four of a kind", 5
		case 3:
			return "Full house", 4
		default:
			fmt.Println(cardMap)
			panic("o dear")
		}
	case 3:
		switch maxCount(cardMap) {
		case 3:
			return "Three of a kind", 3
		case 2:
			return "Two pair", 2
		default:
			panic("o no")
		}
	case 4:
		return "One pair", 1
	case 5:
		return "High card", 0
	default:
		panic("uh oh")
	}
}

func maxCount(m map[string]int) int {
	maxCount := 0
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}
	return maxCount
}

func (hand *Hand) processCardStrength() {
	for i, c := range hand.cards {
		idx := slices.Index(cardStrengthLookup, string(c))
		if idx < 0 {
			panic("crikey!")
		}
		hand.cardsStrength[i] = idx
	}
	for i, c := range hand.cards {
		idx := slices.Index(jokerCardStrengthLookup, string(c))
		if idx < 0 {
			panic("crikey!")
		}
		hand.jokerCardsStrength[i] = idx
	}
}

// Five of a kind, where all five cards have the same label: AAAAA 1
// Four of a kind, where four cards have the same label and one card has a different label: AA8AA 2
// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332 2
// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98 3 (3,1,1)
// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432 3 (2,2,1)
// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4 2
// High card, where all cards' labels are distinct: 23456

// var Order []

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
var cardStrengthLookup = []string{
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"T",
	"J",
	"Q",
	"K",
	"A",
}

var jokerCardStrengthLookup = []string{
	"J",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"T",
	"Q",
	"K",
	"A",
}

// 32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
// KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
// T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.

/*
JABCD -> 2 pair
JAABC -> 3 of a kind
JAABB -> full house
JAAAB -> 4 of a kind

JJABC -> 3 of a kind
JJAAB -> 4 of a kind
JJAAA -> 5 of a kind

JJJAB -> 4 of a kind
JJJAA -> 5 of a kind

JJJJA -> 5 of a kind

JJJJJ -> 5 of a kind
*/
