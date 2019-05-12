/*
--- Day 19: Medicine for Rudolph ---
Rudolph the Red-Nosed Reindeer is sick! His nose isn't shining very brightly,
and he needs medicine.

Red-Nosed Reindeer biology isn't similar to regular reindeer biology; Rudolph is
going to need custom-made medicine. Unfortunately, Red-Nosed Reindeer chemistry
isn't similar to regular reindeer chemistry, either.

The North Pole is equipped with a Red-Nosed Reindeer nuclear fusion/fission
plant, capable of constructing any Red-Nosed Reindeer molecule you need. It
works by starting with some input molecule and then doing a series of
replacements, one per step, until it has the right molecule.

However, the machine has to be calibrated before it can be used. Calibration
involves determining the number of molecules that can be generated in one step
from a given starting point.

For example, imagine a simpler machine that supports only the following
replacements:

H => HO
H => OH
O => HH
Given the replacements above and starting with HOH, the following molecules
could be generated:

HOOH (via H => HO on the first H).
HOHO (via H => HO on the second H).
OHOH (via H => OH on the first H).
HOOH (via H => OH on the second H).
HHHH (via O => HH).
So, in the example above, there are 4 distinct molecules (not five, because HOOH
appears twice) after one replacement from HOH. Santa's favorite molecule,
HOHOHO, can become 7 distinct molecules (over nine replacements: six from H, and
three from O).

The machine replaces without regard for the surrounding characters. For example,
given the string H2O, the transition H => OO would result in OO2O.

Your puzzle input describes all of the possible replacements and, at the bottom,
the medicine molecule for which you need to calibrate the machine. How many
distinct molecules can be created after all the different ways you can do one
replacement on the medicine molecule?

Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg

CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl

--- Part Two ---
Now that the machine is calibrated, you're ready to begin molecule fabrication.

Molecule fabrication always begins with just a single electron, e, and applying
replacements one at a time, just like the ones during calibration.

For example, suppose you have the following replacements:

e => H
e => O
H => HO
H => OH
O => HH
If you'd like to make HOH, you start with e, and then make the following
replacements:

e => O to get O
O => HH to get HH
H => OH (on the second H) to get HOH
So, you could make HOH after 3 steps. Santa's favorite molecule, HOHOHO, can be
made in 6 steps.

How long will it take to make the medicine? Given the available replacements and
the medicine molecule in your puzzle input, what is the fewest number of steps
to go from e to the medicine molecule?

*/
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	inputElement := "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
	// inputElement := "HOH"

	// bring in and process replacements
	inputReplacements, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("File could not be read.")
	}

	inputReplacementsSplit := strings.Split(string(inputReplacements), "\n")

	elementOriginal, elementReplacement := "", ""
	elementReplacements := make([][2]string, len(inputReplacementsSplit))

	for i, inputLine := range inputReplacementsSplit {

		fmt.Sscanf(inputLine, "%s => %s", &elementOriginal, &elementReplacement)
		// elementReplacements = append(elementReplacements, [2]string{elementOriginal, elementReplacement})
		elementReplacements[i] = [2]string{elementOriginal, elementReplacement}

	}

	fmt.Println(elementReplacements) // a slice of all replacements

	// Split inputElement into slice
	// Count Elements
	numElementsInput := 0
	for _, char := range inputElement {
		if char < 'a' {
			numElementsInput++
		}
	}

	//  Generate inputElements
	inputMoleculeSplit := make([]string, numElementsInput)

	elementAppendCount := 0

	for i, char := range inputElement {

		newInputElement := ""

		if char < 'a' {

			if i < len(inputElement)-1 {
				if inputElement[i+1] >= 'a' {
					newInputElement = string(char) + string(inputElement[i+1])
				} else {
					newInputElement = string(char)
				}
			} else {
				newInputElement = string(char)
			}

		}

		if newInputElement != "" {
			inputMoleculeSplit[elementAppendCount] = newInputElement
			elementAppendCount++
		}

	}

	outputMolecules := []string{}
	// Do some matching and add the new generated elements to outputMoleculeSet
	for i, inputElement := range inputMoleculeSplit {
		for i2, replacement := range elementReplacements {

			if inputElement == replacement[0] {

				newMolecule := reactionAndCombine(inputMoleculeSplit, i, elementReplacements, i2)

				moleculeSeen := false

				for _, molecule := range outputMolecules {
					if molecule == newMolecule {
						moleculeSeen = true
					}
				}

				if !moleculeSeen {
					outputMolecules = append(outputMolecules, newMolecule)
				}

				// fmt.Println(newMolecule, outputMolecules)

			}
		}
	}

	fmt.Println("Part 1:", len(outputMolecules))
	// fmt.Println(inputElements)

	// fmt.Println(elementReplacements, "\n", inputElement)

	// fmt.Println("Hello World!")

	// stolen from the subreddit...
	count0 := len(inputMoleculeSplit)
	count1 := 0
	count2 := 0
	for _, v := range inputMoleculeSplit {
		if v == "Rn" || v == "Ar" {
			count1++
		}
		if v == "Y" {
			count2++
		}
	}

	steps := count0 - count1 - (2 * count2) - 1
	fmt.Println(inputMoleculeSplit)
	fmt.Println(count0, count1, count2)
	fmt.Println("Part2:", steps)

}

func reactionAndCombine(inputElements []string, inputIndex int, elementReplacements [][2]string, replacementIndex int) (newMolecule string) {
	for i, inputElement := range inputElements {
		if i == inputIndex {
			newMolecule += elementReplacements[replacementIndex][1]
		} else {
			newMolecule += inputElement
		}
	}
	return
}
