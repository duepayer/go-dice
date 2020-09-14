package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/* 	Farkle 5000 Rules
https://www.dicegamedepot.com/farkle-rules/

To_Do
[] Learn how to print dice to screen
[] Create scoring algorithm

*/

type Player struct {
}

func generateRandomNumber() int {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())

	return (rand.Intn(max-min) + min)
}

func rollDice(num int) []int {
	dice := make([]int, num)
	for i := 0; i < num; i++ {
		dice[i] = generateRandomNumber()
	}
	sort.Ints(dice)
	return dice
}

func scoreRoll(roll []int) int {
	return 0
}

func main() {
	// fmt.Println("Start Game")
	// dice := make([]int, 6)
	// fmt.Println("dice:", dice)

	fmt.Println(rollDice(6))

}
