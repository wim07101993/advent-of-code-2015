package day02

import (
	"testing"

	"github.com/wim07101993/advent-of-code-2015/utils"
)

func TestPart1(t *testing.T) {
	s := New(func() []string {
		ls, err := utils.ReadFileLines("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart1()
	if a != "1606483" {
		t.Error("Wrong answer, try again")
	}
}

func TestPart2(t *testing.T) {
	s := New(func() []string {
		ls, err := utils.ReadFileLines("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart2()
	if a != "3842356" {
		t.Error("Wrong answer, try again")
	}
}
