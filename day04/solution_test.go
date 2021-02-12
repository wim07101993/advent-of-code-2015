package day04

import (
	"testing"

	"advent-of-code-2015/utils"
)

func TestPart1(t *testing.T) {
	s := New(func() []byte {
		ls, err := utils.ReadFileAsBytes("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart1()
	if a != "346386" {
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}

func TestPart2(t *testing.T) {
	s := New(func() []byte {
		ls, err := utils.ReadFileAsBytes("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart2()
	if a != "9958218" {
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}
