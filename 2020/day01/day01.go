package main

import (
	"fmt"
	"sort"
)

func main() {
	// input := []int{1234, 123, 345, 5654, 2000, 20}
	// input = inputTest

	sort.Ints(input)
	fmt.Println(input)

	var ans1 int
	var ans2 int

	for i, v := range input {
		for i2, v2 := range input[i+1:] {
			fmt.Println(i, v, i2, v2, v+v2)

			if v+v2 == 2020 {
				ans1 = v
				ans2 = v2
				fmt.Println("Values summing to 2020 found:", ans1, ans2)
				break
			}

			if v+v2 > 2020 {
				fmt.Println("values over 2020. Breaking out", i, i2+1, v, v2)
				break
			}
		}
		if ans1+ans2 == 2020 {
			fmt.Println("Answer found in outer loop, breaking out")
			break
		}
	}

	fmt.Println("Part1 solution = ", ans1*ans2)

	var answer1 int
	var answer2 int
	var answer3 int

	for i, v := range input {
		for i2, v2 := range input[i+1:] {
			for i3, v3 := range input[i2+1:] {
				// fmt.Println(i, v, i2, v2, v+v2)

				if v+v2+v3 == 2020 {
					answer1 = v
					answer2 = v2
					answer3 = v3
					fmt.Println("Values summing to 2020 found:", answer1, answer2, answer3)
					break
				}

				if v+v2+v3 > 2020 {
					fmt.Println("values over 2020. Breaking out", i, i2+i+1, i3+i2+1, v, v2, v3)
					break
				}
			}
			if answer1+answer2+answer3 == 2020 {
				fmt.Println("Answer found in middle loop, breaking out")
				break
			}
		}

		if answer1+answer2+answer3 == 2020 {
			fmt.Println("Answer found in outer loop, breaking out")
			break
		}

	}

	fmt.Println("Part2 solution = ", answer1*answer2*answer3)

}
