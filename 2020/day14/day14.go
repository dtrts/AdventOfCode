package main

import (
	"fmt"
	"strconv"
	"strings"
)

//14954914379452

func main() {
	// input = inputTest2

	inputLines := strings.Split(input, "\n")

	// fmt.Println(inputLines)

	stringMask := ""

	var zeroMask, oneMask, addr, memVal, part1, part2 uint64
	mem := make(map[uint64]uint64)

	for _, line := range inputLines {

		if strings.HasPrefix(line, "mask") {
			stringMask = strings.TrimPrefix(line, "mask = ")

			// All the 1s will be part of an OR mask, with the rest as 0s
			zeroMask, _ = strconv.ParseUint(strings.ReplaceAll(stringMask, "X", "1"), 2, 64)
			// All the 0s will be part of an AND mask, with the rest as 1s
			oneMask, _ = strconv.ParseUint(strings.ReplaceAll(stringMask, "X", "0"), 2, 64)

		} else if strings.HasPrefix(line, "mem") {

			fmt.Sscanf(line, "mem[%d] = %d", &addr, &memVal)
			mem[addr] = (memVal & zeroMask) | oneMask

		} else {

			fmt.Println("Should not be here")

		}

	}

	for _, value := range mem {
		part1 += value
	}
	fmt.Println(part1)

	// Start in the same way, however apply the bit mask to the address.
	mem = make(map[uint64]uint64)

	for _, line := range inputLines {
		if strings.HasPrefix(line, "mask") {
			stringMask = strings.TrimPrefix(line, "mask = ")

			// All the 1s will be part of an OR mask, with the rest as 0s
			zeroMask, _ = strconv.ParseUint(strings.ReplaceAll(stringMask, "X", "1"), 2, 64)
			// All the 0s will be part of an AND mask, with the rest as 1s
			oneMask, _ = strconv.ParseUint(strings.ReplaceAll(stringMask, "X", "0"), 2, 64)

		} else if strings.HasPrefix(line, "mem") {

			fmt.Sscanf(line, "mem[%d] = %d", &addr, &memVal)
			newAddr := (addr) | oneMask

			// Go through mask, for every X append onto map the two variations.
			baseAddrs := make(map[uint64]bool)
			baseAddrs[newAddr] = true

			for revPos, maskBit := range stringMask {
				if string(maskBit) == "X" {
					for baseAddr := range baseAddrs {
						baseAddrs[setBit(baseAddr, uint64(35-revPos))] = true
						baseAddrs[clearBit(baseAddr, uint64(35-revPos))] = true
					}
				}
			}

			// Go through baseAddrs and set val
			for baseAddr := range baseAddrs {
				mem[baseAddr] = memVal
			}
		} else {

			fmt.Println("Should not be here")

		}
	}

	for _, value := range mem {
		part2 += value
	}
	fmt.Println(part2)

}

// Sets the bit at pos in the integer n.
func setBit(n uint64, pos uint64) uint64 {
	n |= (uint64(1) << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n uint64, pos uint64) uint64 {
	mask := ^(uint64(1) << pos)
	n &= mask
	return n
}
