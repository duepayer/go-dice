package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/duepayer/go-dice/internal/rolldice"
)

func main() {
	fmt.Println("Start Game")
	// dice := make([]int, 6)
	// fmt.Println("dice:", dice)
	numOfDice, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("User Arg Parsing Error")
	}
	fmt.Printf("Number of Dice: %v\n", numOfDice)
	roll := rolldice.RollDice(numOfDice)
	rc := rolldice.RollCount(roll)
	fmt.Println(roll)
	fmt.Println(len(rolldice.RollScore(rc)))
}
