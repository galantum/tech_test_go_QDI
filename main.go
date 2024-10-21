package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	id           int
	point        int
	dice         []int
	diceTransfer []int
}

// method for rolling
func roll(number int) []int {
	dice := make([]int, number)
	for i := range dice {
		dice[i] = rand.Intn(6) + 1 // Generate number 1-6
	}
	return dice
}

// Evaluate after rolling
func evaluate(player *Player, nextPlayer *Player, playerQty int) {

	newDice := []int{}
	transferDice := []int{}

	for _, num := range player.dice {
		if num == 6 { // if number = 6 add point
			player.point++
		} else if num == 1 { // if number = 1 transfer dice to next player
			transferDice = append(transferDice, 1)
		} else {
			newDice = append(newDice, num)
		}
	}

	// combine dice with dice 1 from prev player
	player.dice = append(newDice, player.diceTransfer...)

	if player.id == playerQty { //if turn in last player combine dice player #1 with dice 1 last player
		nextPlayer.dice = append(nextPlayer.dice, transferDice...)
	} else {
		nextPlayer.diceTransfer = append(nextPlayer.diceTransfer, transferDice...) // transfer dice 1 to next player
	}

	// set value diceTransfer to empty
	player.diceTransfer = nil
}

// method for playing
func play(playerQty int, dice int) {
	// Init player
	players := make([]*Player, playerQty)
	for x := 0; x < playerQty; x++ {
		players[x] = &Player{
			id:    x + 1,
			dice:  roll(dice),
			point: 0,
		}
	}

	round := 1
	leftPlayer := 0

	for {
		fmt.Printf("==================\n Turn %d rolling dice:\n", round)

		// Every player roll dice
		for _, player := range players {
			if len(player.dice) > 0 {
				player.dice = roll(len(player.dice))
			}
			if len(player.dice) > 0 {
				fmt.Printf("Player #%d (%d): %v\n", player.id, player.point, player.dice)
			} else {
				fmt.Printf("Player #%d (%d): nil\n", player.id, player.point)
			}
		}

		// evaluate
		for x, player := range players {
			if len(player.dice) > 0 {
				nextPlayer := players[(x+1)%playerQty]
				evaluate(player, nextPlayer, playerQty)
			}
		}

		fmt.Println("Evaluate:")
		for _, player := range players {
			if len(player.dice) > 0 {
				fmt.Printf("Player %d (%d): %v \n", player.id, player.point, player.dice)
			} else {
				fmt.Printf("Player %d (%d): nil\n", player.id, player.point)
			}
		}

		// stop playing if only one player have a dice
		activePlayer := 0
		for _, player := range players {
			if len(player.dice) > 0 {
				activePlayer++
			}
		}

		if activePlayer <= 1 {
			leftPlayer = activePlayer
			break
		}

		round++
	}

	if leftPlayer == 1 {
		fmt.Println("==================")
		fmt.Println("the game is over because only 1 player has dice.")
	} else {
		fmt.Println("==================")
		fmt.Println("the game is over because no one player has dice.")
	}

	// define winning player
	winners := []*Player{players[0]}
	maxPoints := 0

	for _, player := range players {
		if player.point > maxPoints {
			winners = []*Player{player}
			maxPoints = player.point
		} else if player.point == maxPoints {
			winners = append(winners, player)
		}
	}

	if len(winners) <= 1 { // we have 1 winner
		fmt.Printf("The game was won by Player #%d with %d points!\n", winners[0].id, maxPoints)
	} else { // we have more than one winner
		fmt.Printf("The game has ended in a draw with a maximum of %d points each obtained by: \n", maxPoints)
		for _, winner := range winners {
			fmt.Printf("Player #%d\n", winner.id)
		}
	}
}

func main() {
	play(3, 4) // 3 player with 4 dice
}
