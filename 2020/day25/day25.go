package main

import (
	"fmt"
)

func main() {

	card, door := 14012298, 74241

	loop := 0
	for k := 1; k != card; loop++ {
		k = k * 7 % 20201227
	}

	key := 1
	for l := 0; l < loop; l++ {
		key = key * door % 20201227
	}
	fmt.Println(key)
}
