package main

import (
	"fmt"

	"github.com/duepayer/go-dice/internal/rolldice"
)

func main() {
	fmt.Println("Start Game")
	// dice := make([]int, 6)
	// fmt.Println("dice:", dice)

	roll := rolldice.RollDice(4)
	rc := rolldice.RollCount(roll)
	fmt.Println(roll)
	fmt.Println(len(rolldice.ScoreRoll(rc)))
}
