package day01

import (
	"testing"

	"github.com/wim07101993/advent-of-code-2015/utils"
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
	if a != "232" {
		t.Error("Wrong answer, try again")
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
	if a != "1783" {
		t.Error("Wrong answer, try again")
	}
}
