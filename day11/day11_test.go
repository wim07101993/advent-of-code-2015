package day11

import "testing"

func TestCheckPwd(t *testing.T) {
	b := CheckPwd("hijklmmn")
	if b {
		t.Error("hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).")
	}
	b = CheckPwd("abbceffg")
	if b {
		t.Error("abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.")
	}
	b = CheckPwd("abbcegjk")
	if b {
		t.Errorf("abbcegjk fails the third requirement, because it only has one double letter (bb).")
	}
}
