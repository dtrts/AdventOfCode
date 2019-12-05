package main

import (
	"fmt"
)

func main() {

	pwdMinArr := []int{2, 0, 6, 9, 3, 8}
	// pwdMaxArr := [6]int{6, 7, 9, 1, 2, 8}

	pwdCombinations := 0
	pwd2Combinations := 0
	for i := 0; i <= 472190; i++ {
		if isAscending(pwdMinArr) && hasDouble(pwdMinArr) {
			pwdCombinations++
		}
		if isAscending(pwdMinArr) && hasProperDouble(pwdMinArr) {
			pwd2Combinations++
		}

		addOne(&pwdMinArr)
	}
	fmt.Println(pwdMinArr, pwdCombinations, pwd2Combinations)

	// for pwd := pwdMin ; pwd <= pwdMax ; pwd++ {

	// }
	// fmt.Println("asd")
}

func isAscending(arr []int) bool {
	for i := 0; i < 5; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func hasDouble(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}

func hasProperDouble(arr []int) bool {
	count := 1
	digit := arr[0]
	for i := 1; i < len(arr); i++ {

		if arr[i] == digit {
			count++
		} else {
			count = 1
			digit = arr[i]
		}

		if count == 2 {
			if i == len(arr)-1 {
				return true
			}
			if arr[i+1] != digit {
				return true
			}
		}

	}
	return false
}

func addOne(arr *[]int) {
	addOneR(arr, len(*arr)-1)

}

func addOneR(arr *[]int, i int) {
	if i == 0 {
		if (*arr)[i] < 9 {
			(*arr)[i]++
		}
		return
	}

	if (*arr)[i] == 9 {
		(*arr)[i] = 0
		addOneR(arr, i-1)
	} else {
		(*arr)[i]++
	}

	return
}
