/*
--- Day 15: Science for Hungry People ---
Today, you set out on the task of perfecting your milk-dunking cookie recipe.
All you have to do is find the right balance of ingredients.

Your recipe leaves room for exactly 100 teaspoons of ingredients. You make a
list of the remaining ingredients you could use to finish the recipe (your
puzzle input) and their properties per teaspoon:

capacity (how well it helps the cookie absorb milk)
durability (how well it keeps the cookie intact when full of milk)
flavor (how tasty it makes the cookie)
texture (how it improves the feel of the cookie)
calories (how many calories it adds to the cookie)
You can only measure ingredients in whole-teaspoon amounts accurately, and you
have to be accurate so you can reproduce your results in the future. The total
score of a cookie can be found by adding up each of the properties (negative
totals become 0) and then multiplying together everything except calories.

For instance, suppose you have these two ingredients:

Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
Then, choosing to use 44 teaspoons of butterscotch and 56 teaspoons of cinnamon
(because the amounts of each ingredient must add up to 100) would result in a
cookie with the following properties:

A capacity of 44*-1 + 56*2 = 68
A durability of 44*-2 + 56*3 = 80
A flavor of 44*6 + 56*-2  = 152
A texture of 44*3 + 56*-1 = 76
Multiplying these together (68 * 80 * 152 * 76, ignoring calories for now)
results in a total score of 62842880, which happens to be the best score
possible given these ingredients. If any properties had produced a negative
total, it would have instead become zero, causing the whole score to multiply to
zero.

Given the ingredients in your kitchen and their properties, what is the total
score of the highest-scoring cookie you can make?

Sprinkles: capacity 2, durability 0, flavor -2, texture 0, calories 3
Butterscotch: capacity 0, durability 5, flavor -3, texture 0, calories 3
Chocolate: capacity 0, durability 0, flavor 5, texture -1, calories 8
Candy: capacity 0, durability -1, flavor 0, texture 5, calories 8


--- Part Two ---
Your cookie recipe becomes wildly popular! Someone asks if you can make another
recipe that has exactly 500 calories per cookie (so they can use it as a meal
replacement). Keep the rest of your award-winning process the same (100
	teaspoons, same ingredients, same scoring system).

For example, given the ingredients above, if you had instead selected 40
teaspoons of butterscotch and 60 teaspoons of cinnamon (which still adds to
100), the total calorie count would be 40*8 + 60*3 = 500. The total score would
go down, though: only 57600000, the best you can do in such trying
circumstances.

Given the ingredients in your kitchen and their properties, what is the total
score of the highest-scoring cookie you can make with a calorie total of 500?

*/

package main

import "fmt"

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

//  pass downa growing slice of ratios
//  pass up the max score

func calcMaxScore(ingredients []ingredient, ratios []int, maxProportion int, calorieGoal int) (score int) {
	maxScore := 0
	ratioSum := 0
	for _, v := range ratios {
		ratioSum += v
	}

	// calc values if last ingredient or all proportions have been assigned
	if ratioSum == maxProportion || len(ratios) == len(ingredients)-1 {
		//add on remaining proportion if last ingredient
		if len(ratios) == len(ingredients)-1 {
			ratios = append(ratios, maxProportion-ratioSum)
		}
		capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0
		for i, v := range ratios {
			capacity += ingredients[i].capacity * v
			durability += ingredients[i].durability * v
			flavor += ingredients[i].flavor * v
			texture += ingredients[i].texture * v
			calories += ingredients[i].calories * v
		}
		if capacity < 0 {
			capacity = 0
		}
		if durability < 0 {
			durability = 0
		}
		if flavor < 0 {
			flavor = 0
		}
		if texture < 0 {
			texture = 0
		}

		//for part 2
		if calorieGoal > 0 && calories != calorieGoal {
			return 0
		}

		return capacity * durability * flavor * texture
		// calc value
	}

	for i := 0; i <= maxProportion-ratioSum; i++ {

		newScore := calcMaxScore(ingredients, append(ratios, i), maxProportion, calorieGoal)
		if newScore > maxScore {
			// fmt.Println(ingredients, ratios, newScore)
			maxScore = newScore
		}

	}
	//  get sum of ratios
	// loop from that to max proportion
	// add the loop value to the ratio and

	return maxScore
}

func main() {
	fmt.Println("Hello World!")
	/*
		butterscotch := ingredient{"Butterscotch", -1, -2, 6, 3, 8}
		cinnamon := ingredient{"Cinnamon", 2, 3, -2, -1, 3}
		ingredients := []ingredient{butterscotch, cinnamon}
	*/
	sprinkles := ingredient{"Sprinkles", 2, 0, -2, 0, 3}
	butterscotch := ingredient{"Butterscotch", 0, 5, -3, 0, 3}
	chocolate := ingredient{"Chocolate", 0, 0, 5, -1, 8}
	candy := ingredient{"Candy", 0, -1, 0, 5, 8}

	ingredients := []ingredient{sprinkles, butterscotch, chocolate, candy}

	maxScore := calcMaxScore(ingredients, []int{}, 100, 0)
	fmt.Println("Part 1:", maxScore)

	maxScore = calcMaxScore(ingredients, []int{}, 100, 500)
	fmt.Println("Part 2:", maxScore)
}
