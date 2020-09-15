package rolldie

import (
	"math/rand"
	"sort"
	"time"
)

type scoreRoll interface {
	ScoreRoll(roll []int) int
}

func generateRandomNumber() int {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())

	return (rand.Intn(max-min) + min)
}

// RollDice - Rolls a given number of dice and returns an array of random ints
func RollDice(num int) []int {
	dice := make([]int, num)
	for i := 0; i < num; i++ {
		dice[i] = generateRandomNumber()
	}
	sort.Ints(dice)
	return dice
}
