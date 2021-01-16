package day06

import "testing"

func TestParseInstruction(t *testing.T) {
	i, err := ParseInstruction("turn off 660,55 through 986,197")
	if err != nil {
		t.Error(err)
	}

	if i.Do != Off {
		t.Error("The action should be turn off")
	}
	if i.From.X != 660 && i.From.Y == 55 {
		t.Error("The from coordinate is not correctly parsed")
	}
	if i.To.X != 986 && i.From.Y != 197 {
		t.Error("The to coordinate is not correctly parsed")
	}
}
