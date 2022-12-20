package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func p(s ...interface{}) {
	fmt.Println(s...)
}

func main() {
	// BOILER PLATE --------------------------------------------------------------------
	start := time.Now()
	log.Printf("Starting... %s", start.Format("Jan 2 15:04:05 2006 MST"))

	var inputFileName string
	flag.StringVar(&inputFileName, "i", "input.txt", "Name of the input file")
	flag.Parse()

	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic("Input file unable to be read.")
	}

	inputString := strings.TrimSpace(string(inputBytes))
	p("Input string", inputString)

	// BOILER PLATE --------------------------------------------------------------------
	blueprints := []RobotRecipe{}
	for _, line := range strings.Split(inputString, "\n") {
		blueprints = append(blueprints, generateRobotRecipe(line))
	}

	p(blueprints)

	part1 := 0

	for i, blueprint := range blueprints {
		cache := make(map[State]int)
		maxCache := 0
		maxFromStart := getMaxGeodes(24, State{oreRobots: 1}, &blueprint, cache, &maxCache)

		part1 += maxFromStart * blueprint.blueprint

		p("Completed blueprint:", i, blueprint.blueprint, maxFromStart, "Adding to score", maxFromStart*blueprint.blueprint)

	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:", part1)
	// BOILER PLATE --------------------------------------------------------------------

	part2 := 1
	for i := 0; i < 3; i++ {
		cache := make(map[State]int)
		maxCache := 0
		maxFromStart := getMaxGeodes(32, State{oreRobots: 1}, &blueprints[i], cache, &maxCache)
		p("Completed blueprint:", i, blueprints[i].blueprint, maxFromStart)
		part2 *= maxFromStart
		log.Printf("Duration: %s", time.Since(start))
	}

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part2:", part2)
	// BOILER PLATE --------------------------------------------------------------------
}

// Start a round.
// Options are:
// - build nothing
// - build ore if possible
// - build clay if possible
// - build obsidian if possible
// - build geode if possible

// Other things to consider:
// Don't build a ore/clay/obsidian robot if we're already producing the same as the max to spend in one go.
// If making more than the max you can spend,

// Let a minute mass, get the new resources

func getMaxGeodes(minutes int, state State, blueprint *RobotRecipe, cache map[State]int, maxCache *int) int {

	// p("Sarting get max Geodes", state)

	currentMax := state.geode
	if state.time == minutes {
		return currentMax
	}

	if val, ok := cache[state]; ok {
		// p("Found State in cache")
		return val
	}

	remainingTime := minutes - state.time
	// return early if even making pure geode smashers doesn't pass the current max.
	bestPossibleMax := state.geode + ((remainingTime) * state.geodeRobots) + ((remainingTime * (remainingTime - 1)) / 2)
	if bestPossibleMax < *maxCache {
		return currentMax
	}

	maxOre := max(blueprint.ore.ore, max(blueprint.clay.ore, max(blueprint.obsidian.ore, blueprint.geode.ore)))

	canBuildOre := state.ore >= blueprint.ore.ore
	canBuildClay := state.ore >= blueprint.clay.ore
	canBuildObsidian := state.ore >= blueprint.obsidian.ore && state.clay >= blueprint.obsidian.clay
	canBuildGeode := state.ore >= blueprint.geode.ore && state.obsidian >= blueprint.geode.obsidian

	shouldBuildOre := state.oreRobots < maxOre
	shouldBuildClay := state.clayRobots < blueprint.obsidian.clay
	shouldBuildObsidian := state.obsidianRobots < blueprint.geode.obsidian

	if canBuildOre && shouldBuildOre {

		newState := copyState(state)
		newState.ore -= blueprint.ore.ore

		updateState(&newState, blueprint, maxOre, minutes)

		newState.oreRobots++

		currentMax = max(currentMax, getMaxGeodes(minutes, newState, blueprint, cache, maxCache))

	}

	if canBuildClay && shouldBuildClay {

		newState := copyState(state)

		newState.ore -= blueprint.clay.ore

		updateState(&newState, blueprint, maxOre, minutes)

		newState.clayRobots++

		currentMax = max(currentMax, getMaxGeodes(minutes, newState, blueprint, cache, maxCache))

	}

	if canBuildObsidian && shouldBuildObsidian {
		newState := copyState(state)
		newState.ore -= blueprint.obsidian.ore
		newState.clay -= blueprint.obsidian.clay

		updateState(&newState, blueprint, maxOre, minutes)

		newState.obsidianRobots++

		currentMax = max(currentMax, getMaxGeodes(minutes, newState, blueprint, cache, maxCache))
	}

	if canBuildGeode {

		newState := copyState(state)
		newState.ore -= blueprint.geode.ore
		newState.obsidian -= blueprint.geode.obsidian

		updateState(&newState, blueprint, maxOre, minutes)

		newState.geodeRobots++

		currentMax = max(currentMax, getMaxGeodes(minutes, newState, blueprint, cache, maxCache))
	}

	if true {

		newState := copyState(state)

		updateState(&newState, blueprint, maxOre, minutes)

		currentMax = max(currentMax, getMaxGeodes(minutes, newState, blueprint, cache, maxCache))
	}

	cache[state] = currentMax

	if currentMax > *maxCache {
		*maxCache = currentMax
	}

	return currentMax
}

func updateState(state *State, blueprint *RobotRecipe, maxOre int, minutes int) {

	state.time++
	// state.ore = min(state.ore+state.oreRobots, maxOre)
	// state.clay = min(state.clay+state.clayRobots, blueprint.obsidian.clay)
	// state.obsidian = min(state.obsidian+state.obsidianRobots, blueprint.geode.obsidian)
	if state.oreRobots >= maxOre {
		state.ore = maxOre
	} else {
		state.ore += state.oreRobots
	}
	if state.clayRobots >= blueprint.obsidian.clay {
		state.clay = blueprint.obsidian.clay
	} else {
		state.clay += state.clayRobots
	}

	if state.obsidianRobots >= blueprint.geode.obsidian {
		state.obsidian = blueprint.geode.obsidian
	} else {
		state.obsidian += state.obsidianRobots
	}

	state.geode = state.geode + state.geodeRobots

}

func copyState(state State) State {
	return State{
		ore:            state.ore,
		clay:           state.clay,
		obsidian:       state.obsidian,
		geode:          state.geode,
		oreRobots:      state.oreRobots,
		clayRobots:     state.clayRobots,
		obsidianRobots: state.obsidianRobots,
		geodeRobots:    state.geodeRobots,
		time:           state.time,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateRobotRecipe(line string) RobotRecipe {

	var blueprint, oreCostOre, clayCostOre, obsidianCostOre, obsidianCostClay, geodeCostOre, geodeCostObsidian int
	fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &blueprint, &oreCostOre, &clayCostOre, &obsidianCostOre, &obsidianCostClay, &geodeCostOre, &geodeCostObsidian)
	return RobotRecipe{
		blueprint: blueprint,
		ore: struct{ ore int }{
			ore: oreCostOre,
		},
		clay: struct{ ore int }{
			ore: clayCostOre,
		},
		obsidian: struct{ ore, clay int }{
			ore:  obsidianCostOre,
			clay: obsidianCostClay,
		},
		geode: struct{ ore, obsidian int }{
			ore:      geodeCostOre,
			obsidian: geodeCostObsidian,
		},
	}
}

type RobotRecipe struct {
	blueprint int
	ore       struct {
		ore int
	}
	clay struct {
		ore int
	}
	obsidian struct {
		ore  int
		clay int
	}
	geode struct {
		ore      int
		obsidian int
	}
}

type State struct {
	ore            int
	clay           int
	obsidian       int
	geode          int
	oreRobots      int
	clayRobots     int
	obsidianRobots int
	geodeRobots    int
	time           int
}
