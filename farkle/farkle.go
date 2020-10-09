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
	winningScore int
	winner       Player
}

// Player - a place to store player meta
type Player struct {
	name               string
	score              Score
	tempScore          Score
	scoringRollHistory []Score
}

// RollTally - Tally roll
type RollTally struct {
	ones        int
	fives       int
	pairs       []int
	triplets    []int
	quadruplets []int
	quintuplets []int
	sextuplets  []int
}

// Score - all score related meta
type Score struct {
	Score int
	Dice  []int
}

// CountRoll - compile each die number count
func CountRoll(roll []int) map[int]int {
	rc := map[int]int{ // Map literal
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}
	for _, die := range roll {
		switch die {
		case 1:
			rc[1] = rc[1] + 1
		case 2:
			rc[2] = rc[2] + 1
		case 3:
			rc[3] = rc[3] + 1
		case 4:
			rc[4] = rc[4] + 1
		case 5:
			rc[5] = rc[5] + 1
		case 6:
			rc[6] = rc[6] + 1
		}
	}
	return rc
}

// TallyScoringCombinations - tally all possible scoring combinations
func TallyScoringCombinations(count map[int]int) RollTally {
	var rt RollTally

	// Remove non-scoring die from count map
	for k, i := range count {
		if i == 0 {
			delete(count, k)
		}
	}
	fmt.Printf("Dice count after removing zero count dice:\n %v \n\n", count)

	for k, i := range count {
		switch i {
		case 1:
			if k == 1 {
				rt.ones = i
			} else if k == 5 {
				rt.fives = i
			}
		case 2:
			if k == 1 {
				rt.ones = i
			} else if k == 5 {
				rt.fives = i
			}
			rt.pairs = append(rt.pairs, k)
		case 3:
			if k == 1 {
				rt.ones = i
			} else if k == 5 {
				rt.fives = i
			}
			rt.triplets = append(rt.triplets, k)
		case 4:
			if k == 1 {
				rt.ones = i
			} else if k == 5 {
				rt.fives = i
			}
			rt.quadruplets = append(rt.quadruplets, k)
		case 5:
			if k == 1 {
				rt.ones = i
			} else if k == 5 {
				rt.fives = i
			}
			rt.quintuplets = append(rt.quintuplets, k)
		case 6:
			if k == 1 {
				rt.ones = i
			} else if k == 5 {
				rt.fives = i
			}
			rt.sextuplets = append(rt.sextuplets, k)
		}
	}
	return rt
}

// ScoreRoll - tally and calculate all possible scoring combinations
// and return an array of mapped scores
func ScoreRoll(rt RollTally) Score {
	var score Score

	return score
}

func main() {
	fmt.Println("Start Game")

	numOfDice, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("User Arg Parsing Error")
	}
	fmt.Printf("Number of Dice: %v\n", numOfDice)
	roll := dice.RollDice(numOfDice)
	fmt.Println(roll)
	cr := CountRoll(roll)
	fmt.Printf("\n%+v", TallyScoringCombinations(cr))
}
