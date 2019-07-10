package main

import (
	"fmt"
	"sort"
)

type circle struct {
	activeMarble int
	marbles      []int
	scoreboard   []int
	lastPlayer   int
	init         game
}

type game struct {
	players    int
	lastMarble int
}

func lnrm(c circle) int {
	sortMarbs := sort.IntSlice(c.marbles)
	for i, v := range sortMarbs {
		if i != v {
			return i
		}
	}
	return len(sortMarbs)
}

/*
func (m *[]int) Rmv(position int) {

	modposition := position % len(*m)

	switch modposition {
	case 0:
		*m = m[1:]
	case len(*m):
		*m = m[:len(*m)-1]
	default:
		m = append(m[0:position], m[position+1:]...)
	}

}

func (c *circle) play() circle {
	currentPlayer := (c.lastPlayer + 1) % c.init.players
	lnrm := lnrm(*c)

	if len(marbles) == 0 {
		c.activeMarble = 0
		marbles = []int{0}

	} else if lnrn % 23 {

		scoreboard[currentPlayer] += lnrn

	}

}
*/

func main() {
	//game := circle{}
	//fmt.Println(float64(70833) / float64(23))
	fmt.Println("IYERRR")
}

var wtf = "asdasdasda"
