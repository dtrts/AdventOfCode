/*
--- Day 21: RPG Simulator 20XX ---
Little Henry Case got a new video game for Christmas. It's an RPG, and he's
stuck on a boss. He needs to know what equipment to buy at the shop. He hands
you the controller.

In this game, the player (you) and the enemy (the boss) take turns attacking.
The player always goes first. Each attack reduces the opponent's hit points by
at least 1. The first character at or below 0 hit points loses.

Damage dealt by an attacker each turn is equal to the attacker's damage score
minus the defender's armor score. An attacker always does at least 1 damage. So,
if the attacker has a damage score of 8, and the defender has an armor score of
3, the defender loses 5 hit points. If the defender had an armor score of 300,
the defender would still lose 1 hit point.

Your damage score and armor score both start at zero. They can be increased by
buying items in exchange for gold. You start with no items and have as much gold
as you need. Your total damage or armor is equal to the sum of those stats from
all of your items. You have 100 hit points.

Here is what the item shop is selling:

Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3
You must buy exactly one weapon; no dual-wielding. Armor is optional, but you
can't use more than one. You can buy 0-2 rings (at most one for each hand). You
must use any items you buy. The shop only has one of each item, so you can't
buy, for example, two rings of Damage +3.

For example, suppose you have 8 hit points, 5 damage, and 5 armor, and that the
boss has 12 hit points, 7 damage, and 2 armor:

The player deals 5-2 = 3 damage; the boss goes down to 9 hit points.
The boss deals 7-5 = 2 damage; the player goes down to 6 hit points.
The player deals 5-2 = 3 damage; the boss goes down to 6 hit points.
The boss deals 7-5 = 2 damage; the player goes down to 4 hit points.
The player deals 5-2 = 3 damage; the boss goes down to 3 hit points.
The boss deals 7-5 = 2 damage; the player goes down to 2 hit points.
The player deals 5-2 = 3 damage; the boss goes down to 0 hit points.
In this scenario, the player wins! (Barely.)

You have 100 hit points. The boss's actual stats are in your puzzle input. What
is the least amount of gold you can spend and still win the fight?

PUZZLE INPUT ::

Hit Points: 103
Damage: 9
Armor: 2

--- Part Two ---
Turns out the shopkeeper is working with the boss, and can persuade you to buy
whatever items he wants. The other rules still apply, and he still only has one
of each item.

What is the most amount of gold you can spend and still lose the fight?

*/

package main

import "fmt"

type character struct {
	hitPoints int
	damage    int
	armor     int
}

type item struct {
	cost   int
	damage int
	armor  int
}

func main() {
	// Initialise fighters
	boss := character{hitPoints: 103, damage: 9, armor: 2}
	playerBase := character{hitPoints: 100, damage: 0, armor: 0}

	// Weapons
	dagger := item{cost: 8, damage: 4, armor: 0}
	shortsword := item{cost: 10, damage: 5, armor: 0}
	warhammer := item{cost: 25, damage: 6, armor: 0}
	longsword := item{cost: 40, damage: 7, armor: 0}
	greataxe := item{cost: 74, damage: 8, armor: 0}

	weapons := []item{dagger, shortsword, warhammer, longsword, greataxe}

	// armor
	leather := item{cost: 13, damage: 0, armor: 1}
	chainmail := item{cost: 31, damage: 0, armor: 2}
	splintmail := item{cost: 53, damage: 0, armor: 3}
	bandedmail := item{cost: 75, damage: 0, armor: 4}
	platemail := item{cost: 102, damage: 0, armor: 5}
	noArmor := item{cost: 0, damage: 0, armor: 0}

	armors := []item{leather, chainmail, splintmail, bandedmail, platemail, noArmor}

	// rings
	damage1 := item{cost: 25, damage: 1, armor: 0}
	damage2 := item{cost: 50, damage: 2, armor: 0}
	damage3 := item{cost: 100, damage: 3, armor: 0}
	defense1 := item{cost: 20, damage: 0, armor: 1}
	defense2 := item{cost: 40, damage: 0, armor: 2}
	defense3 := item{cost: 80, damage: 0, armor: 3}
	noRing1 := item{cost: 0, damage: 0, armor: 0}
	noRing2 := item{cost: 0, damage: 0, armor: 0}

	rings := []item{damage1, damage2, damage3, defense1, defense2, defense3, noRing1, noRing2}

	//  just loop through all combinations of equipable items.

	minCost := -1
	maxCost := -1

	for _, weapon := range weapons {
		for _, armor := range armors {
			for ring1Index, ring1 := range rings {
				for ring2Index, ring2 := range rings {

					if ring1Index == ring2Index { // can't have two of the same ring (which is why there are two no ring rings)
						break
					}

					itemsLoadout := []item{weapon, armor, ring1, ring2}
					player := playerBase
					goldSpent := 0

					for _, item := range itemsLoadout {
						goldSpent += item.cost
						player.damage += item.damage
						player.armor += item.armor
					}

					// Calculate attack per move for each fighter
					playerAttack := player.damage - boss.armor
					if playerAttack <= 0 {
						playerAttack = 1
					}

					bossAttack := boss.damage - player.armor
					if bossAttack <= 0 {
						bossAttack = 1
					}

					// Calc number of moves needed to kill opponent
					playerMoves := 0
					bossMoves := 0

					playerMoves = (boss.hitPoints / playerAttack)
					if boss.hitPoints%playerAttack != 0 {
						playerMoves++
					}

					bossMoves = (player.hitPoints / bossAttack)
					if player.hitPoints%bossAttack != 0 {
						bossMoves++
					}

					//Calc Winner
					winner := ""
					if playerMoves <= bossMoves {
						winner = "Player"
					} else {
						winner = "Boss"
					}

					// Check is min/max cost has changed and print info
					if (maxCost == -1 || goldSpent > maxCost) && winner == "Boss" {
						maxCost = goldSpent
						// fmt.Printf("-----------\nMinCost: %d. MaxCost: %d.\nWinner: %s. Player Stats: %v. Player Moves: %d. Boss Moves: %d.\nCurrent Cost: %v. Loadout: %v\n", minCost, maxCost, winner, player, playerMoves, bossMoves, goldSpent, itemsLoadout)
					}

					// Set min cost.
					if (minCost == -1 || goldSpent < minCost) && winner == "Player" {
						minCost = goldSpent
						// fmt.Printf("-----------\nMinCost: %d. MaxCost: %d.\nWinner: %s. Player Stats: %v. Player Moves: %d. Boss Moves: %d.\nCurrent Cost: %v. Loadout: %v\n", minCost, maxCost, winner, player, playerMoves, bossMoves, goldSpent, itemsLoadout)
					}

					// Show results of battle

				}
			}
		}
	}

	// Print results
	fmt.Printf("Part 1: %d\nPart 2: %d\n", minCost, maxCost)

}
