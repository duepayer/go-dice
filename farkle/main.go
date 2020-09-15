package farkle

import (
	"fmt"

	"github.com/duepayer/pkg/rolldie"
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

*/

type Farkle struct {
}

func main() {
	fmt.Println("Start Game")
	// dice := make([]int, 6)
	// fmt.Println("dice:", dice)

	fmt.Println(rolldie.RollDice(6))

}
