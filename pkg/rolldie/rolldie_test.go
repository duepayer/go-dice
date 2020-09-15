package rolldie

import "testing"

func TestRollDiceCount(t *testing.T) {
	res := RollDice(6)
	if len(res) != 6 {
		t.Errorf("Roll Dice slice length (%v) does not match expected output (6)", len(res))
	}
}