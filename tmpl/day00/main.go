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

	p("Calculating Part 1....")

	p("Calculating Part 2....")

	// BOILER PLATE --------------------------------------------------------------------
	log.Printf("Duration: %s", time.Since(start))
	p("Part1:")
	p("Part2:")
	// BOILER PLATE --------------------------------------------------------------------
}

type GameState struct {
	bossHitPoints   int
	bossDamage      int
	playerHitPoints int
	playerMana      int
	shieldTurns     int
	PoisonTurns     int
	rechargeTurns   int
}

/*
Magic Missile costs 53 mana. It instantly does 4 damage.
Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
*/

/*
Do a turn.
Apply spells

boss damage if applicable.
- recursively call next turn with updated game state
reduce all spells by 1

player casts spell if applicable
- recursively call next turn with updated game state for each spell.


if player or bass is dead can end the game.


key for cache:
spellCounts
hitPoints + mana.

Value for cache:
min mana spent for player to win from that point.

Prune branches:
- keep track of bestMin Mana spent, if crossing that value then prune the branch.
-


*/
