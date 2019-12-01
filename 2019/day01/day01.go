package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read file and creat int array
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	massesStr := strings.Split(string(file), "\n")
	masses := make([]int, len(massesStr), len(massesStr))

	for i, massStr := range massesStr {
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			panic("")
		}
		masses[i] = mass
	}

	// Calculate fuel requirement (Part 1)
	fuelRequirement := 0

	for _, mass := range masses {
		fuelRequirement += (mass / 3) - 2
	}

	fmt.Printf("Part 1: Fuel requirement = %d\n", fuelRequirement)

	// Calculate fuel requirement for fuel requirement (Part 2)
	fuelFuelRequirement := 0

	for _, mass := range masses[:] {
		mass = mass/3 - 2
		for mass > 0 {
			fuelFuelRequirement += mass
			mass = mass/3 - 2
		}
	}

	fmt.Printf("Part 2: Fuel Fuel requirement = %d\n", fuelFuelRequirement)
}
