package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/duepayer/go-dice/internal/dice"
)

/* 	Farkle 5000 Rules
https://www.dicegamedepot.com/farkle-rules/

Scoring
1	100 points
5	50 points
Three 1's	1,000 points
Three 2's	200 points
Three 3's	300 points
Three 4's	400 points
Three 5's	500 points
Three 6's	600 points
Four 1's	2,000 points
Four 2's	400 points
Four 3's	600 points
Four 4's	800 points
Four 5's	1000 points
Four 6's	1200 points
1-2-3-4-5-6 	3000 points
3 Pairs	1500 points (including 4-of-a-kind and a pair)

To_Do
[] Learn how to print dice to screen
[] Create scoring algorithm
[] Support interactive user input
*/

// Farkle - game configuration
type Farkle struct {
	playerCount  int
	players      []Player
	diceNumber   int
	WinningScore int
	Winner       Player
}

// Player - a place to store player meta
type Player struct {
	name        string
	score       Score
	tempScore   Score
	rollHistory []Score
}

// Score - all score related meta
type Score struct {
	Score int
	Dice  []int
}

// RollScore - tally and calculate all possible scoring combinations
// and return an array of mapped scores
func RollScore(count map[int]int) []Score {
	scores := make([]Score, 15)

	// Remove non-scoring die from count map
	for k, i := range count {
		if i == 0 {
			delete(count, k)
		}
	}

	if len(count) == 6 {
		d := []int{1, 2, 3, 4, 5, 6}
		s := Score{
			Score: 3000,
			Dice:  d,
		}
		scores = append(scores, s)
	} else {
		// Separate counts into singles, pairs, triplets and quadruplets
		ones := make([]int, 1)
		fives := make([]int, 1)
		pairs := make([]int, 3)
		triplets := make([]int, 2)
		quadruplets := make([]int, 1)
		quintuplets := make([]int, 1)
		sextuplets := make([]int, 1)

		for k, i := range count {
			switch i {
			case 1:
				if k == 1 {
					ones = append(ones, i)
				} else if k == 5 {
					fives = append(fives, i)
				}
			case 2:
				pairs = append(pairs, k)
			case 3:
				triplets = append(triplets, k)
			case 4:
				quadruplets = append(quadruplets, k)
			case 5:
				quintuplets = append(quintuplets, k)
			case 6:
				sextuplets = append(sextuplets, k)
			}
		}
	}
	return scores
}

func main() {
	fmt.Println("Start Game")

	numOfDice, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("User Arg Parsing Error")
	}
	fmt.Printf("Number of Dice: %v\n", numOfDice)
	roll := dice.RollDice(numOfDice)
	rc := dice.RollCount(roll)
	fmt.Println(roll)
	fmt.Println(RollScore(rc))
}
