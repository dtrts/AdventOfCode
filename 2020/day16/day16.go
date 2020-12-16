package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// input = inputTest

	// Parse Input

	fmt.Printf("Part1 Start\n")
	rules, yourTicket, nearbyTickets := parseInput(input)

	// fmt.Printf("rules:%v yourTicket:%v nearbyTickets:%v\n", rules, yourTicket, nearbyTickets)

	Part1 := 0
	// Loop through All near by tickets
	for _, nearbyTicket := range nearbyTickets {
		// Check all fields to find one which don't match any rule
		for _, fieldValue := range nearbyTicket {
			if !validFieldAtLeastOneRule(fieldValue, rules) {
				Part1 += fieldValue
			}
		}
	}

	fmt.Printf("Part1: %v\n\n", Part1)

	// Part 2 --------------------------------------------------------------------
	fmt.Println("Part 2")
	rules, yourTicket, nearbyTickets = parseInput(input)

	// fmt.Printf("rules:%v yourTicket:%v nearbyTickets:%v\n\n", rules, yourTicket, nearbyTickets)

	// Get validTickets
	validTickets := filterValidTickets(nearbyTickets, rules)

	// fmt.Println(validTickets)

	// I need to assign a ruleID to every fieldID
	// I could go through each ticket, and each field, and get a map of which
	// rules are valid. []map[rule]bool. If the map only has one in then that is
	// its place, go through all maps and remove that one from the rest.

	matcher := initMatcher(validTickets, rules)
	lenValidTickets := len(validTickets)

	catch := 0
	for isValidMatcher(matcher) == false && catch < 10000 {

		matcherPassThrough(matcher, lenValidTickets)
		catch++
	}
	fmt.Println(matcher, isValidMatcher(matcher), catch)

	Part2 := 1

	for i, m := range matcher {
		for k := range m {
			if strings.HasPrefix(k, "departure") {
				Part2 *= yourTicket[i]
			}
		}
	}

	fmt.Println("Part2", Part2) //2843534243843
}

func matcherPassThrough(matcher []map[string]int, lenValidTickets int) {

	for i, m := range matcher {
		if len(m) == 1 {
			continue
		}

		validRule := theOnlyValidRuleAtIndex(matcher, i, lenValidTickets)

		// fmt.Println("ValidRule", i, validRule)

		if validRule != "" {
			removeRuleFromMatcher(matcher, validRule, i)
			cleanMatcherAtIndex(matcher, validRule, i)
			// return // If we return this will it run faster? // From 19 pass
			// throughs to 10. Because slices are pointers the slice is altered as we
			// go, which is fine :).
		}
	}
	return

}

func theOnlyValidRuleAtIndex(matcher []map[string]int, index int, lenValidTickets int) string {

	validRuleNames := []string{}
	for k, v := range matcher[index] {
		if v == lenValidTickets {
			validRuleNames = append(validRuleNames, k)
			if len(validRuleNames) > 1 {
				return ""
			}
		}
	}

	// fmt.Println("\n\n", matcher[index], index, lenValidTickets, validRuleNames, "\n")

	if len(validRuleNames) == 1 {
		return validRuleNames[0]
	} else {
		panic("asdwqndjks")
	}
	return ""
}

func cleanMatcherAtIndex(matcher []map[string]int, ruleName string, index int) {
	for k := range matcher[index] {
		if k == ruleName {
			continue
		}
		delete(matcher[index], k)
	}
}

func removeRuleFromMatcher(matcher []map[string]int, ruleName string, protectIndex int) {
	for i, m := range matcher {
		if i == protectIndex {
			continue
		}
		delete(m, ruleName)
	}
}

func isValidMatcher(matcher []map[string]int) bool {
	for _, validRules := range matcher {
		if len(validRules) > 1 {
			return false
		}
	}
	return true
}

func initMatcher(tickets [][]int, rules []ruleType) []map[string]int {

	matcher := make([]map[string]int, len(rules))
	for i := range matcher {
		matcher[i] = make(map[string]int)
	}

	for _, ticket := range tickets {
		for i, fieldValue := range ticket {
			validRules := validRules(fieldValue, rules)
			for rule := range validRules {
				matcher[i][rule.name]++
			}
		}
	}

	return matcher
}

func validRules(fieldValue int, rules []ruleType) (validRules map[ruleType]bool) {
	validRules = make(map[ruleType]bool)
	for _, rule := range rules {
		if validField(fieldValue, rule.ranges) {
			validRules[rule] = true
		}
	}
	return validRules
}

type ruleType struct {
	name   string
	ranges [4]int
}

func filterValidTickets(tickets [][]int, rules []ruleType) (validTickets [][]int) {
	for _, ticket := range tickets {
		if validTicket(ticket, rules) {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func validTicket(ticket []int, rules []ruleType) bool {
	// Find one field which doesn't match all rules
	for _, fieldValue := range ticket {

		if !validFieldAtLeastOneRule(fieldValue, rules) {
			return false
		}

	}
	return true
}

func validFieldAtLeastOneRule(fieldValue int, rules []ruleType) bool {
	// Compare field against all rules
	// If it succeeds one return false
	for _, rule := range rules {
		if validField(fieldValue, rule.ranges) {
			return true
		}
	}
	return false

}

func validField(fieldValue int, ruleRanges [4]int) bool {
	if (ruleRanges[0] <= fieldValue && fieldValue <= ruleRanges[1]) || (ruleRanges[2] <= fieldValue && fieldValue <= ruleRanges[3]) {
		return true
	}
	return false
}

func invalidField(fieldValue int, rule [4]int) bool {
	return !validField(fieldValue, rule)
}

func parseInput(input string) (rules []ruleType, yourTicket []int, nearbyTickets [][]int) {

	// Split input into sections, rules, yourTicket, nearbyTickets
	sections := strings.Split(input, "\n\n")

	// RULES ---------------------------------------------------------------------
	// rules = make(map[string][4]int)
	rulesString := strings.Split(sections[0], "\n")
	for _, ruleString := range rulesString {

		re := regexp.MustCompile(`([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)`) //SOME OF THE RULE NAME HAVE SPACES IN THEM AIDNBLJAKSJNDCKHBFNIEKHBDJHDBNS
		ruleStringParse := re.FindStringSubmatch(ruleString)

		ruleName := ruleStringParse[1]
		ruleLow1, _ := strconv.Atoi(ruleStringParse[2])
		ruleHigh1, _ := strconv.Atoi(ruleStringParse[3])
		ruleLow2, _ := strconv.Atoi(ruleStringParse[4])
		ruleHigh2, _ := strconv.Atoi(ruleStringParse[5])
		ranges := [4]int{ruleLow1, ruleHigh1, ruleLow2, ruleHigh2}

		rules = append(rules, ruleType{ruleName, ranges})
	}

	// YOURTICKET ----------------------------------------------------------------
	// Only one ticket here, take it, split by comma, create int slice

	yourTicketString := strings.Split(sections[1], "\n")
	yourTicketStringNums := strings.Split(yourTicketString[1], ",")

	for _, numString := range yourTicketStringNums {
		num, _ := strconv.Atoi(numString)
		yourTicket = append(yourTicket, num)
	}

	// NEARBYTICKETS -------------------------------------------------------------
	nearbyTicketStrings := strings.Split(sections[2], "\n")

	// Loop through all near by tickets
	for _, nearbyTicketString := range nearbyTicketStrings[1:] {

		nearbyTicketStringNums := strings.Split(nearbyTicketString, ",")
		nextTicket := []int{}

		// loop through each string ticket after aplitting by commas
		for _, numString := range nearbyTicketStringNums {
			num, _ := strconv.Atoi(numString)
			nextTicket = append(nextTicket, num)
		}

		nearbyTickets = append(nearbyTickets, nextTicket)

	}

	return rules, yourTicket, nearbyTickets
}
