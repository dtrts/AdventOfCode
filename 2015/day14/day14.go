/*
--- Day 14: Reindeer Olympics ---
This year is the Reindeer Olympics! Reindeer can fly at high speeds, but must
rest occasionally to recover their energy. Santa would like to know which of his
reindeer is fastest, and so he has them race.

Reindeer can only either be flying (always at their top speed) or resting (not
	moving at all), and always spend whole seconds in either state.

For example, suppose you have the following Reindeer:

Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
After one second, Comet has gone 14 km, while Dancer has gone 16 km. After ten
seconds, Comet has gone 140 km, while Dancer has gone 160 km. On the eleventh
second, Comet begins resting (staying at 140 km), and Dancer continues on for a
total distance of 176 km. On the 12th second, both reindeer are resting. They
continue to rest until the 138th second, when Comet flies for another ten
seconds. On the 174th second, Dancer flies for another 11 seconds.

In this example, after the 1000th second, both reindeer are resting, and Comet
is in the lead at 1120 km (poor Dancer has only gotten 1056 km by that point).
So, in this situation, Comet would win (if the race ended at 1000 seconds).

Given the descriptions of each reindeer (in your puzzle input), after exactly
2503 seconds, what distance has the winning reindeer traveled?

--- Part Two ---
Seeing how reindeer move in bursts, Santa decides he's not pleased with the old
scoring system.

Instead, at the end of each second, he awards one point to the reindeer
currently in the lead. (If there are multiple reindeer tied for the lead, they
each get one point.) He keeps the traditional 2503 second time limit, of course,
as doing otherwise would be entirely ridiculous.

Given the example reindeer from above, after the first second, Dancer is in the
lead and gets one point. He stays in the lead until several seconds into Comet's
second burst: after the 140th second, Comet pulls into the lead and gets his
first point. Of course, since Dancer had been in the lead for the 139 seconds
before that, he has accumulated 139 points by the 140th second.

After the 1000th second, Dancer has accumulated 689 points, while poor Comet,
our old champion, only has 312. So, with the new scoring system, Dancer would
win (if the race ended at 1000 seconds).

Again given the descriptions of each reindeer (in your puzzle input), after
exactly 2503 seconds, how many points does the winning reindeer have?


*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type reindeer struct {
	name     string
	flySpeed int
	flyTime  int
	flyRest  int
}

func main() {
	fmt.Println("Hello world!")
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("File unable to be read.")
	}

	inputSplit := strings.Split(string(input), "\n")
	name, flySpeed, flyTime, flyRest := "", 0, 0, 0
	reindeers := []reindeer{}

	fmt.Println(inputSplit)
	for _, singleInput := range inputSplit {
		fmt.Sscanf(singleInput, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &flySpeed, &flyTime, &flyRest)
		reindeers = append(reindeers, reindeer{name, flySpeed, flyTime, flyRest})
	}
	fmt.Println(reindeers)

	//part1
	//loop through reindeers. do entire race and find distance.
	distance := 0
	maxDistance := 0

	for _, oneReindeer := range reindeers {
		distance = 0
		for time := 0; time < 2503; time++ {
			cycleTime := oneReindeer.flyTime + oneReindeer.flyRest
			if time%cycleTime < oneReindeer.flyTime {
				distance += oneReindeer.flySpeed
			}

			// fmt.Println(time, oneReindeer.name, distance)
		}

		if distance > maxDistance {
			maxDistance = distance
		}
		fmt.Println(oneReindeer.name, distance, maxDistance)

	}

	//part2 need to find the winning reineer on each second.
	//max map of names distacne and points.
	race := map[reindeer][2]int{}
	for _, oneReindeer := range reindeers {
		race[oneReindeer] = [2]int{0, 0}
	}

	for time := 0; time < 2503; time++ {
		//  loop through map
		maxDistance := 0

		for oneReindeer, onePosition := range race {
			//  add distance
			cycleTime := oneReindeer.flyTime + oneReindeer.flyRest
			if time%cycleTime < oneReindeer.flyTime {

				onePosition[0] += oneReindeer.flySpeed
				race[oneReindeer] = onePosition

				//  find winning distance

			}
			if onePosition[0] > maxDistance {
				maxDistance = onePosition[0]
			}

		}

		//  add point to all
		for oneReindeer, onePosition := range race {
			if onePosition[0] == maxDistance {
				onePosition[1]++
			}
			race[oneReindeer] = onePosition
		}

	}

	//get max points
	fmt.Println(race)
	maxPoints := 0
	for _, onePosition := range race {
		if onePosition[1] > maxPoints {
			maxPoints = onePosition[1]
		}
	}
	fmt.Println(maxPoints)
}
