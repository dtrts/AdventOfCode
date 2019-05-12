package utils

//StringSlicePermutate will return a slice of all possible permutations of the
//original slice. Does not check for duplicates
func StringSlicePermutate(s []string) (permutations [][]string) {

	if len(s) == 0 {

		return

	} else if len(s) == 1 {

		permutations = [][]string{[]string{s[0]}}
		return

	} else if len(s) == 2 {

		permutations = [][]string{[]string{s[0], s[1]}, []string{s[1], s[0]}} //  two permutations
		return

	} else {

		for i, v := range s {
			// append onto permutations each element and the permutations it has
			// append sub permutations onto element
			for _, v2 := range StringSlicePermutate(RemoveElementFromSlice(s, i)) {
				//permutationsElement := append([][]string{}, append([]string{v}, v2...))
				permutations = append(permutations, append([][]string{}, append([]string{v}, v2...))...)
			}
		}
		return

	}
}

//RemoveElementFromSlice removes the element from a string slice.
func RemoveElementFromSlice(sliceIncoming []string, i int) (alteredSlice []string) {
	// assuming len(s) > 0 and i < len(s)
	if i == 0 {

		alteredSlice = append(alteredSlice, sliceIncoming[1:]...)
		return alteredSlice

	} else if i == len(sliceIncoming)-1 {

		alteredSlice = append(alteredSlice, sliceIncoming[:len(sliceIncoming)-1]...)
		return alteredSlice

	} else {

		alteredSlice = append(append(alteredSlice, sliceIncoming[:i]...), sliceIncoming[i+1:]...)
		return alteredSlice

	}
}
