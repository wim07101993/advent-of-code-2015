package day11

import (
	"testing"

	"advent-of-code-2015/utils"
)

func TestPart1(t *testing.T) {
	s := New(func() string {
		ls, err := utils.ReadFileAsString("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart1()
	if a != "hxbxxyzz" {
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}

func TestPart2(t *testing.T) {
	s := New(func() string {
		ls, err := utils.ReadFileAsString("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart2()
	if a != "hxcaabcc" {
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}
