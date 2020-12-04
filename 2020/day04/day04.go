package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// input = valid

	inputLines := [][]string{}
	var passportMaps []map[string]string
	validPassports := 0
	validPassports2 := 0

	// Split into assports
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n\n") {
		inputLines = append(inputLines, strings.Split(strings.Trim(strings.Replace(line, "\n", " ", -1), " "), " "))
	}
	// convert into maps
	for _, line := range inputLines {
		passportMap := make(map[string]string)
		for _, keyValue := range line {
			keyValueSplit := strings.Split(keyValue, ":")
			passportMap[keyValueSplit[0]] = keyValueSplit[1]
		}
		passportMaps = append(passportMaps, passportMap)
	}

	for _, passport := range passportMaps {
		fmt.Println(passport)
		if isValidPassport(passport) {
			validPassports++
			if dataValidation(passport) {
				validPassports2++
			}

		}
	}

	fmt.Println(validPassports, validPassports2)

	// fmt.Println(len(passportMaps[0]))

}

func dataValidation(passport map[string]string) bool {

	bry, err := strconv.Atoi(passport["byr"])
	if err != nil {
		fmt.Println("Invalid Bry err")
		return false
	}
	if bry < 1920 || 2002 < bry {
		fmt.Println("Invalid Bry")
		return false
	}

	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		fmt.Println("Invalid iyr err")
		return false
	}
	if iyr < 2010 || 2020 < iyr {
		fmt.Println("Invalid iyr")
		return false
	}

	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		fmt.Println("Invalid eyr err")
		return false
	}
	if eyr < 2020 || 2030 < eyr {
		fmt.Println("Invalid eyr")
		return false
	}

	// HGT in: 59 76, cm: 150 193
	if len(passport["hgt"]) <= 2 {
		fmt.Println("Invalid hgt len")
		return false
	}
	hgtType := passport["hgt"][len(passport["hgt"])-2:]
	hgtNum, err := strconv.Atoi(passport["hgt"][:len(passport["hgt"])-2])
	if err != nil {
		fmt.Println("Invalid hgt err")
		return false
	}
	if hgtType != "in" && hgtType != "cm" {
		fmt.Println("Invalid hgt type")
		return false
	}
	if hgtType == "in" && (hgtNum < 59 || 76 < hgtNum) {
		fmt.Println("Invalid hgt type in")
		return false
	}
	if hgtType == "cm" && (hgtNum < 150 || 193 < hgtNum) {
		fmt.Println("Invalid hgt type cm")
		return false
	}

	// hcl
	matched, err := regexp.MatchString(`^#[0-9a-f]{6}$`, passport["hcl"])
	if err != nil {
		fmt.Println("Invalid hcl err")
		return false
	}
	if matched == false {
		fmt.Println("Invalid hcl match")
		return false
	}

	//ecl
	validEcl := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	if !contains(validEcl, passport["ecl"]) {
		fmt.Println("Invalid ecl match")
		return false
	}

	//pid
	matchedPid, err := regexp.MatchString(`^[0-9]{9}$`, passport["pid"])
	if err != nil {
		fmt.Println("Invalid pid err")
		return false
	}
	if matchedPid == false {
		fmt.Println("Invalid pid match")
		return false
	}

	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ```
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in: If cm, the number must
// be at least 150 and at most 193. If in, the number must be at least 59 and at
// most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
// ```

func isValidPassport(passport map[string]string) bool {
	// required := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid"}

	if len(passport) == 8 {
		return true
	}

	if len(passport) == 7 {
		_, ok := passport["cid"]
		if ok {
			return false
		}
		return true
	}

	fmt.Println("Invalid keys")
	return false
}
