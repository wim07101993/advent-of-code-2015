package day06

import (
	"log"
	"testing"

	"advent-of-code-2015/utils"
)

func TestExample1(t *testing.T) {
	s := New(func() []string {
		return []string{
			"turn on 0,0 through 999,999",
			"toggle 0,0 through 999,0",
			"turn off 499,499 through 500,500",
		}
	})

	a := s.SolvePart1()
	if a != "998996" {
		log.Println(s.grid)
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}

func TestSolverDay6_SolvePart1(t *testing.T) {
	s := New(func() []string {
		ls, err := utils.ReadFileLines("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart1()
	if a != "400410" {
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}

func TestExample2(t *testing.T) {
	s := New(func() []string {
		return []string{
			"turn on 0,0 through 0,0",
			"toggle 0,0 through 999,999",
		}
	})

	a := s.SolvePart2()
	if a != "2000001" {
		log.Println(s.grid)
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}

func TestSolverDay6_SolvePart2(t *testing.T) {
	s := New(func() []string {
		ls, err := utils.ReadFileLines("./input")
		if err != nil {
			t.Error(err)
		}
		return ls
	})

	a := s.SolvePart2()
	if a != "15343601" {
		t.Errorf("Wrong answer, try again (%s)", a)
	}
}
