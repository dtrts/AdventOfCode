package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numGames int
var debug = false

func main() {
	fmt.Println("hw")
	// 	// Start game
	// 	// Track hands at each round
	// 	// Play a round
	// 	// To find winner:
	// 	// check if hands have been played in that game, if they have the game's winner is player 1.
	// 	// Else compare the first card for each player
	// 	// if number on each card <= number of cards remaining start sub game with new hands, which are the number of cards
	// 	// shown on the top card from each players deck
	// 	// the winner is the winner of that sub game
	// 	// else compare the top, the winner is the one

	// input = inputTest
	// input = inputTest2

	player1, player2, winner := []int{}, []int{}, new([]int)

	player1, player2 = parseInput(input)
	winner = combat(&player1, &player2, false)
	fmt.Printf("Part 1: %v\n%v %v\n", score(winner), player1, player2) //34324

	player1, player2 = parseInput(input)
	winner = combat(&player1, &player2, true)
	fmt.Printf("Part 2: %v\n%v %v\n", score(winner), player1, player2) //33259
}

func combat(player1 *[]int, player2 *[]int, recursive bool) (gameWinner *[]int) {
	numGames++
	if debug {
		fmt.Printf("=== GAME %v ===\n\n", numGames)
	}
	roundNumber := 0

	// Create an array to hold a history of game states
	gameStateHistory := [][2][]int{}

	// Keep playing rounds whilst both players have cards
	for len(*player1) > 0 && len(*player2) > 0 {
		// Print out round number and game number
		// Game number doesn't return to original so not sure about the depth
		roundNumber++
		if debug {
			fmt.Printf("-- Round %v (Game %v) --\n", roundNumber, numGames)
		}

		// Check current game state and end game early if we find a loop
		currentGameState := newGameState(player1, player2)
		if gameStateExists(gameStateHistory, currentGameState) {
			if debug {
				fmt.Printf("Player 1 wins due to repeated positions\n\n")
			}
			return player1
		}
		gameStateHistory = append(gameStateHistory, currentGameState)

		// Play a round to find the winner
		roundWinner := round(player1, player2, recursive)

		// Print out round success
		if debug {
			fmt.Printf("Player 1's deck: %v\nPlayer 2's deck: %v\nPlayer 1 plays: %v\nPlayer 2 plays: %v\n", player1, player2, (*player1)[0], (*player2)[0])
			if roundWinner == player1 {
				fmt.Printf("Player 1 wins round %v of game %v\n\n", roundNumber, numGames)
			} else {
				fmt.Printf("Player 2 wins round %v of game %v\n\n", roundNumber, numGames)
			}
		}

		// Process the winner of round 2
		processWinner(player1, player2, roundWinner)
	}

	// Return the winner
	if len(*player1) > 0 {
		return player1
	}
	return player2
}

func round(player1 *[]int, player2 *[]int, recursive bool) (roundWinner *[]int) {
	// Get top card, and the number of cards left in hand
	card1, lenRemaining1 := (*player1)[0], len(*player1)-1
	card2, lenRemaining2 := (*player2)[0], len(*player2)-1

	// If recursive AND the number or remaining cards for each player is less than the number on the card
	if recursive == true && lenRemaining1 >= card1 && lenRemaining2 >= card2 {
		subPlayer1 := make([]int, card1)
		subPlayer2 := make([]int, card2)
		copy(subPlayer1, (*player1)[1:card1+1])
		copy(subPlayer2, (*player2)[1:card2+1])

		if debug {
			fmt.Printf("Player 1's deck: %v\nPlayer 2's deck: %v\nPlayer 1 plays: %v\nPlayer 2 plays: %v\n", player1, player2, (*player1)[0], (*player2)[0])
			fmt.Printf("Playing a sub-game to determine the winner...\n\n")
		}

		subGameWinner := combat(&subPlayer1, &subPlayer2, true)

		if subGameWinner == &subPlayer1 {
			if debug {
				fmt.Printf("Player 1 wins the subgame\n")
			}
			return player1
		}

		if debug {
			fmt.Printf("Player 2 wins the subgame\n")
		}
		return player2
	}

	if card1 > card2 {
		return player1
	}
	return player2

}

func gameStateExists(gameStateHistory [][2][]int, currentGameState [2][]int) bool {
	for _, oldGameState := range gameStateHistory {
		if gameStatesMatch(oldGameState, currentGameState) {
			return true
		}
	}
	return false
}

func gameStatesMatch(gameStateA [2][]int, gameStateB [2][]int) bool {
	if len(gameStateA[0]) != len(gameStateB[0]) || len(gameStateA[1]) != len(gameStateB[1]) {
		return false
	}
	for playerIndex, player := range gameStateA {
		for cardIndex, card := range player {
			if card != gameStateB[playerIndex][cardIndex] {
				return false
			}
		}
	}
	return true
}

func newGameState(player1 *[]int, player2 *[]int) (gameState [2][]int) {
	gameState[0] = make([]int, len(*player1))
	gameState[1] = make([]int, len(*player2))
	copy(gameState[0], *player1)
	copy(gameState[1], *player2)
	return gameState
}

func processWinner(player1 *[]int, player2 *[]int, winner *[]int) {
	if player1 == winner {
		*player1 = append(*player1, popFront(player1), popFront(player2))
	} else {
		*player2 = append(*player2, popFront(player2), popFront(player1))
	}
}

func popFront(hand *[]int) int {
	front := (*hand)[0]
	*hand = (*hand)[1:]
	return front
}

func score(deck *[]int) int {
	score, numCards := 0, len(*deck)
	for i, card := range *deck {
		score += (numCards - i) * card
	}
	return score
}

func parseInput(input string) ([]int, []int) {
	players := strings.Split(input, "\n\n")
	player1 := []int{}
	player2 := []int{}
	player1 = parsePlayer(players[0], player1)
	player2 = parsePlayer(players[1], player2)
	return player1, player2

}

func parsePlayer(input string, player []int) []int {
	for i, card := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		cardInt, err := strconv.Atoi(card)
		if err != nil {
			panic("NO INT")
		}
		player = append(player, cardInt)
	}
	return player
}

//start game.
//

// func recursiveCombat(player1 *[]int, player2 *[]int) (winner *[]int) {

// 	fmt.Println("\n\nRECURSIVE COMBAT SUB GAME", player1, player2)

// 	// Create a new game history for each game
// 	gameHistory := []GameState{}

// 	// Whilst both players have cards
// 	for len(*player1) > 0 && len(*player2) > 0 {
// 		// Check if hand has been played before in this
// 		fmt.Println(player1, player2)
// 		if gamePlayed(gameHistory, player1, player2) {

// 			// fmt.Printf("Player 1 wins on history!\n\n")
// 			// fmt.Printf("Player 1's deck:%v\nPlayer 2's deck:%v\n", player1, player2)

// 			card1 := popFront(player1)
// 			card2 := popFront(player2)

// 			// fmt.Printf("Player 1 plays: %v\nPlayer 2 plays: %v\n", card1, card2)
// 			// fmt.Printf("len1: %v len2: %v\n", len(*player1), len(*player2))
// 			// fmt.Printf("Player 1 wins on history!\n\n")

// 			*player1 = append(*player1, card1, card2)
// 			// fmt.Printf("Player 1's deck:%v\nPlayer 2's deck:%v\n", player1, player2)

// 			return player1
// 		}
// 		gameHistory = append(gameHistory, newGameState(player1, player2))
// 		// Check history for existing hand and return early
// 		// Check if game has been played before

// 		winner = playRecursiveHand(player1, player2)
// 		// Finish once one player runs out of cards
// 		for _, player := range []*[]int{player1, player2} {
// 			if len(*player) == len(*player1)+len(*player2) {
// 				return player // This is the winner if they have all the cards
// 			}
// 		}
// 		// add hand to history
// 	}
// 	return winner
// }

// func playRecursiveHand(player1 *[]int, player2 *[]int) *[]int {
// 	// fmt.Printf("Player 1's deck:%v\nPlayer 2's deck:%v\n", player1, player2)

// 	card1 := popFront(player1)
// 	card2 := popFront(player2)

// 	// fmt.Printf("Player 1 plays: %v\nPlayer 2 plays: %v\n", card1, card2)
// 	// fmt.Printf("len1: %v len2: %v\n", len(*player1), len(*player2))

// 	// length of hand 1 and hand 2 less than the value on the respective cards then the winner is a game of recursive
// 	// combat.
// 	if len(*player1) >= card1 && len(*player2) >= card2 {

// 		// add Game State to game history

// 		// Make duplicates
// 		player1Sub := make([]int, card1)
// 		copy(player1Sub, (*player1)[:card1])
// 		player2Sub := make([]int, card2)
// 		copy(player2Sub, (*player2)[:card2])

// 		subWinner := recursiveCombat(&player1Sub, &player2Sub)
// 		if subWinner == player1 {
// 			// fmt.Printf("Player 1 wins!\n\n")
// 			*player1 = append(*player1, card1, card2)
// 		} else {
// 			// fmt.Printf("Player 2 wins!\n\n")
// 			*player2 = append(*player2, card2, card1)
// 		}
// 		return subWinner
// 	}

// 	if card1 > card2 {
// 		// fmt.Printf("Player 1 wins!\n\n")
// 		*player1 = append(*player1, card1, card2)
// 		return player1
// 	}

// 	// fmt.Printf("Player 2 wins!\n\n")
// 	*player2 = append(*player2, card2, card1)
// 	return player2

// }

// func gamePlayed(gameHistory []GameState, player1 *[]int, player2 *[]int) bool {
// 	for _, gameState := range gameHistory {
// 		gsP1 := gameState.player1
// 		gsP2 := gameState.player2
// 		if equalHands(&gsP1, player1) && equalHands(&gsP2, player2) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func newGameState(player1 *[]int, player2 *[]int) GameState {
// 	gameState := GameState{}
// 	gameState.player1 = make([]int, len(*player1))
// 	gameState.player2 = make([]int, len(*player2))
// 	copy(gameState.player1, *player1)
// 	copy(gameState.player2, *player2)
// 	return gameState
// }

// // GameState ...
// type GameState struct {
// 	player1 []int
// 	player2 []int
// }

// func combat(player1 *[]int, player2 *[]int) {
// 	for len(*player1) > 0 && len(*player2) > 0 {
// 		playHand(player1, player2)
// 	}
// }

// func equalHands(hand1 *[]int, hand2 *[]int) bool {
// 	if len(*hand1) != len(*hand2) {
// 		return false
// 	}
// 	for i, card := range *hand1 {
// 		if card != (*hand2)[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// // func compareGameStates(gameState1 GameState, gameState2 GameState) bool {

// // }

// func score(deck []int) int {
// 	score, numCards := 0, len(deck)
// 	for i, card := range deck {
// 		score += (numCards - i) * card
// 	}
// 	return score
// }

// func playHand(player1 *[]int, player2 *[]int) {
// 	card1 := popFront(player1)
// 	card2 := popFront(player2)
// 	if card1 > card2 {
// 		*player1 = append(*player1, card1, card2)
// 	} else {
// 		*player2 = append(*player2, card2, card1)
// 	}
// 	return
// }

// func popFront(deck *[]int) int {
// 	front := (*deck)[0]
// 	*deck = (*deck)[1:]
// 	return front
// }

var inputTest = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

var inputTest2 = `Player 1:
43
19

Player 2:
2
29
14`

var input = `Player 1:
14
29
25
17
13
50
33
32
7
37
26
34
46
24
3
28
18
20
11
1
21
8
44
10
22

Player 2:
5
38
27
15
45
40
43
30
35
9
48
12
16
47
42
4
2
31
41
39
23
19
36
49
6`
